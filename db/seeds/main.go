// +build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/jinzhu/configor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"github.com/qor/slug"
)

var Seeds = struct {
	Categories []struct {
		Name string
	}
	Colors []struct {
		Name string
		Code string
	}
	Sizes []struct {
		Name string
		Code string
	}
	Products []struct {
		CategoryName    string
		Name            string
		NameWithSlug    string
		Code            string
		Price           float32
		Description     string
		MadeCountry     string
		ColorVariations []struct {
			ColorName string
			Images    []struct {
				URL string
			}
		}
		SizeVariations []struct {
			SizeName string
		}
	}
	Stores []struct {
		Name      string
		Phone     string
		Email     string
		Country   string
		Zip       string
		City      string
		Region    string
		Address   string
		Latitude  float64
		Longitude float64
	}
}{}

var Tables = []interface{}{
	&models.Category{}, &models.Color{}, &models.Size{},
	&models.Product{}, &models.ColorVariation{}, &models.ColorVariationImage{}, &models.SizeVariation{},
	&models.Store{},
}

func main() {
	filepaths, _ := filepath.Glob("db/seeds/data/*.yml")
	if err := configor.Load(&Seeds, filepaths...); err != nil {
		panic(err)
	}

	truncateTables()
	createRecords()
}

func truncateTables() {
	for _, table := range Tables {
		if err := db.DB.DropTableIfExists(table).Error; err != nil {
			panic(err)
		}
		if err := db.Publish.DraftDB().DropTableIfExists(table).Error; err != nil {
			panic(err)
		}
		db.DB.AutoMigrate(table)
		db.Publish.AutoMigrate(table)
	}
}

func createRecords() {
	// fmt.Println("Start create sample data...")
	// createCategories()
	// fmt.Println("--> Created categories.")
	// createColors()
	// fmt.Println("--> Created colors.")
	// createSizes()
	// fmt.Println("--> Created sizes.")
	// createProducts()
	// fmt.Println("--> Created products.")
	createStores()
	fmt.Println("--> Created stores.")
	fmt.Println("--> Done!")
}

func createCategories() {
	for _, c := range Seeds.Categories {
		category := models.Category{}
		category.Name = c.Name
		if err := db.DB.Create(&category).Error; err != nil {
			log.Fatalf("create category (%v) failure, got err %v", category, err)
		}
	}
}

func createColors() {
	for _, c := range Seeds.Colors {
		color := models.Color{}
		color.Name = c.Name
		color.Code = c.Code
		if err := db.DB.Create(&color).Error; err != nil {
			log.Fatalf("create color (%v) failure, got err %v", color, err)
		}
	}
}

func createSizes() {
	for _, s := range Seeds.Sizes {
		size := models.Size{}
		size.Name = s.Name
		size.Code = s.Code
		if err := db.DB.Create(&size).Error; err != nil {
			log.Fatalf("create size (%v) failure, got err %v", size, err)
		}
	}
}

func createProducts() {
	for _, p := range Seeds.Products {
		category := findCategoryByName(p.CategoryName)

		product := models.Product{}
		product.CategoryID = category.ID
		product.Name = p.Name
		product.NameWithSlug = slug.Slug{p.NameWithSlug}
		product.Code = p.Code
		product.Price = p.Price
		product.Description = p.Description
		product.MadeCountry = p.MadeCountry

		if err := db.DB.Create(&product).Error; err != nil {
			log.Fatalf("create product (%v) failure, got err %v", product, err)
		}

		for _, cv := range p.ColorVariations {
			color := findColorByName(cv.ColorName)

			colorVariation := models.ColorVariation{}
			colorVariation.ProductID = product.ID
			colorVariation.ColorID = color.ID
			if err := db.DB.Create(&colorVariation).Error; err != nil {
				log.Fatalf("create color_variation (%v) failure, got err %v", colorVariation, err)
			}

			for _, i := range cv.Images {
				image := models.ColorVariationImage{}
				if file, err := openFileByURL(i.URL); err != nil {
					fmt.Printf("open file (%q) failure, got err %v", i.URL, err)
				} else {
					defer file.Close()
					image.Image.Scan(file)
				}
				image.ColorVariationID = colorVariation.ID
				if err := db.DB.Create(&image).Error; err != nil {
					log.Fatalf("create color_variation_image (%v) failure, got err %v", image, err)
				}
			}

			for _, sv := range p.SizeVariations {
				size := findSizeByName(sv.SizeName)

				sizeVariation := models.SizeVariation{}
				sizeVariation.ColorVariationID = colorVariation.ID
				sizeVariation.SizeID = size.ID
				sizeVariation.AvailableQuantity = 20
				if err := db.DB.Create(&sizeVariation).Error; err != nil {
					log.Fatalf("create size_variation (%v) failure, got err %v", sizeVariation, err)
				}
			}
		}
	}
}

func createStores() {
	for _, s := range Seeds.Stores {
		store := models.Store{}
		store.Name = s.Name
		store.Phone = s.Phone
		store.Email = s.Email
		store.Country = s.Country
		store.City = s.City
		store.Region = s.Region
		store.Address = s.Address
		store.Zip = s.Zip
		store.Latitude = s.Latitude
		store.Longitude = s.Longitude
		if err := db.DB.Create(&store).Error; err != nil {
			log.Fatalf("create store (%v) failure, got err %v", store, err)
		}
	}
}

func findCategoryByName(name string) *models.Category {
	category := &models.Category{}
	if err := db.DB.Where(&models.Category{Name: name}).First(category).Error; err != nil {
		log.Fatalf("can't find category with name = %q, got err %v", name, err)
	}
	return category
}

func findColorByName(name string) *models.Color {
	color := &models.Color{}
	if err := db.DB.Where(&models.Color{Name: name}).First(color).Error; err != nil {
		log.Fatalf("can't find color with name = %q, got err %v", name, err)
	}
	return color
}

func findSizeByName(name string) *models.Size {
	size := &models.Size{}
	if err := db.DB.Where(&models.Size{Name: name}).First(size).Error; err != nil {
		log.Fatalf("can't find size with name = %q, got err %v", name, err)
	}
	return size
}

func openFileByURL(rawURL string) (*os.File, error) {
	if fileURL, err := url.Parse(rawURL); err != nil {
		return nil, err
	} else {
		path := fileURL.Path
		segments := strings.Split(path, "/")
		fileName := segments[len(segments)-1]

		basePath, _ := filepath.Abs(".")
		filePath := fmt.Sprintf("%s/tmp/%s", basePath, fileName)

		if _, err := os.Stat(filePath); err == nil {
			return os.Open(filePath)
		}

		file, err := os.Create(filePath)
		if err != nil {
			return file, err
		}

		check := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}
		resp, err := check.Get(rawURL) // add a filter to check redirect
		if err != nil {
			return file, err
		}
		defer resp.Body.Close()
		fmt.Printf("----> Downloaded %v\n", rawURL)

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return file, err
		}
		return file, nil
	}
}
