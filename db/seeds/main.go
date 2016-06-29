// +build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jinzhu/now"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/media_library"
	"github.com/qor/publish"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/db"
	"github.com/qor/qor-example/db/seeds"
	"github.com/qor/seo"
	"github.com/qor/slug"
	"github.com/qor/sorting"
)

/* How to upload file
 * $ brew install s3cmd
 * $ s3cmd --configure (Refer https://github.com/theplant/qor-example)
 * $ s3cmd put local_file_path s3://qor3/
 */

var (
	fake           = seeds.Fake
	truncateTables = seeds.TruncateTables

	Seeds  = seeds.Seeds
	Tables = []interface{}{
		&models.User{}, &models.Address{},
		&models.Category{}, &models.Color{}, &models.Size{}, &models.Collection{},
		&models.Product{}, &models.ColorVariation{}, &models.ColorVariationImage{}, &models.SizeVariation{},
		&models.Store{},
		&models.Order{}, &models.OrderItem{},
		&models.Setting{},
		&models.SEOSetting{},

		&media_library.AssetManager{},
		&publish.PublishEvent{},
		&database.Translation{},
		&admin.QorWidgetSetting{},
	}
)

func main() {
	truncateTables(Tables...)
	createRecords()
}

func createRecords() {
	fmt.Println("Start create sample data...")

	createSetting()
	fmt.Println("--> Created setting.")

	createSeo()
	fmt.Println("--> Created seo.")

	createAdminUsers()
	fmt.Println("--> Created admin users.")

	createUsers()
	fmt.Println("--> Created users.")
	createAddresses()
	fmt.Println("--> Created addresses.")

	createCategories()
	fmt.Println("--> Created categories.")
	createCollections()
	fmt.Println("--> Created collections.")
	createColors()
	fmt.Println("--> Created colors.")
	createSizes()
	fmt.Println("--> Created sizes.")
	createProducts()
	fmt.Println("--> Created products.")
	createStores()
	fmt.Println("--> Created stores.")

	createOrders()
	fmt.Println("--> Created orders.")

	createWidgets()
	fmt.Println("--> Created widgets.")

	fmt.Println("--> Done!")
}

func createSetting() {
	setting := models.Setting{}

	setting.ShippingFee = Seeds.Setting.ShippingFee
	setting.GiftWrappingFee = Seeds.Setting.GiftWrappingFee
	setting.CODFee = Seeds.Setting.CODFee
	setting.TaxRate = Seeds.Setting.TaxRate
	setting.Address = Seeds.Setting.Address
	setting.Region = Seeds.Setting.Region
	setting.City = Seeds.Setting.City
	setting.Country = Seeds.Setting.Country
	setting.Zip = Seeds.Setting.Zip
	setting.Latitude = Seeds.Setting.Latitude
	setting.Longitude = Seeds.Setting.Longitude

	if err := db.DB.Create(&setting).Error; err != nil {
		log.Fatalf("create setting (%v) failure, got err %v", setting, err)
	}
}

func createSeo() {
	seoSetting := models.SEOSetting{}
	seoSetting.SiteName = Seeds.Seo.SiteName
	seoSetting.DefaultPage = seo.Setting{Title: Seeds.Seo.DefaultPage.Title, Description: Seeds.Seo.DefaultPage.Description, Keywords: Seeds.Seo.DefaultPage.Keywords}
	seoSetting.HomePage = seo.Setting{Title: Seeds.Seo.HomePage.Title, Description: Seeds.Seo.HomePage.Description, Keywords: Seeds.Seo.HomePage.Keywords}
	seoSetting.ProductPage = seo.Setting{Title: Seeds.Seo.ProductPage.Title, Description: Seeds.Seo.ProductPage.Description, Keywords: Seeds.Seo.ProductPage.Keywords}

	if err := db.DB.Create(&seoSetting).Error; err != nil {
		log.Fatalf("create seo (%v) failure, got err %v", seoSetting, err)
	}
}

func createAdminUsers() {
	user := models.User{}
	user.Email = "dev@getqor.com"
	user.Password = "$2a$10$a8AXd1q6J1lL.JQZfzXUY.pznG1tms8o.PK.tYD.Tkdfc3q7UrNX." // Password: testing
	user.Confirmed = true
	user.Name = "QOR Admin"
	user.Role = "admin"
	db.DB.Create(&user)
}

func createUsers() {
	totalCount := 600
	for i := 0; i < totalCount; i++ {
		user := models.User{}
		user.Email = fake.Email()
		user.Name = fake.Name()
		user.Gender = []string{"Female", "Male"}[i%2]
		if err := db.DB.Create(&user).Error; err != nil {
			log.Fatalf("create user (%v) failure, got err %v", user, err)
		}

		day := (-14 + i/45)
		user.CreatedAt = now.EndOfDay().Add(time.Duration(day*rand.Intn(24)) * time.Hour)
		if user.CreatedAt.After(time.Now()) {
			user.CreatedAt = time.Now()
		}
		if err := db.DB.Save(&user).Error; err != nil {
			log.Fatalf("Save user (%v) failure, got err %v", user, err)
		}
	}
}

func createAddresses() {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		log.Fatalf("query users (%v) failure, got err %v", users, err)
	}

	for _, user := range users {
		address := models.Address{}
		address.UserID = user.ID
		address.ContactName = user.Name
		address.Phone = fake.PhoneNumber()
		address.City = fake.City()
		address.Address1 = fake.StreetAddress()
		address.Address2 = fake.SecondaryAddress()
		if err := db.DB.Create(&address).Error; err != nil {
			log.Fatalf("create address (%v) failure, got err %v", address, err)
		}
	}
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

func createCollections() {
	for _, c := range Seeds.Collections {
		collection := models.Collection{}
		collection.Name = c.Name
		if err := db.DB.Create(&collection).Error; err != nil {
			log.Fatalf("create collection (%v) failure, got err %v", collection, err)
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
		for _, c := range p.Collections {
			collection := findCollectionByName(c.Name)
			product.Collections = append(product.Collections, *collection)
		}

		if err := db.DB.Create(&product).Error; err != nil {
			log.Fatalf("create product (%v) failure, got err %v", product, err)
		}

		for _, cv := range p.ColorVariations {
			color := findColorByName(cv.ColorName)

			colorVariation := models.ColorVariation{}
			colorVariation.ProductID = product.ID
			colorVariation.ColorID = color.ID
			colorVariation.ColorCode = cv.ColorCode
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
		product.Name = p.ZhName
		product.Description = p.ZhDescription
		product.MadeCountry = p.ZhMadeCountry
		db.DB.Set("l10n:locale", "zh-CN").Create(&product)
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

func createOrders() {
	var users []models.User
	if err := db.DB.Preload("Addresses").Find(&users).Error; err != nil {
		log.Fatalf("query users (%v) failure, got err %v", users, err)
	}

	var sizeVariations []models.SizeVariation
	if err := db.DB.Find(&sizeVariations).Error; err != nil {
		log.Fatalf("query sizeVariations (%v) failure, got err %v", sizeVariations, err)
	}
	var sizeVariationsCount = len(sizeVariations)

	for i, user := range users {
		order := models.Order{}
		state := []string{"draft", "checkout", "cancelled", "paid", "paid_cancelled", "processing", "shipped", "returned"}[rand.Intn(10)%8]
		abandonedReasons := []string{
			"Unsatisfied with discount",
			"Dropped after check gift wrapping option",
			"Dropped after select expected delivery date",
			"Invalid credit card inputted",
			"Credit card balances insufficient",
			"Created a new order with more products",
			"Created a new order with fewer products",
		}
		abandonedReason := abandonedReasons[rand.Intn(len(abandonedReasons))]

		order.UserID = user.ID
		order.ShippingAddressID = user.Addresses[0].ID
		order.BillingAddressID = user.Addresses[0].ID
		order.State = state
		if rand.Intn(15)%15 == 3 && state == "checkout" || state == "processing" || state == "paid_cancelled" {
			order.AbandonedReason = abandonedReason
		}
		if err := db.DB.Create(&order).Error; err != nil {
			log.Fatalf("create order (%v) failure, got err %v", order, err)
		}

		sizeVariation := sizeVariations[i%sizeVariationsCount]
		product := findProductByColorVariationID(sizeVariation.ColorVariationID)
		quantity := []uint{1, 2, 3, 4, 5}[rand.Intn(10)%5]
		discountRate := []uint{0, 5, 10, 15, 20, 25}[rand.Intn(10)%6]

		orderItem := models.OrderItem{}
		orderItem.OrderID = order.ID
		orderItem.SizeVariationID = sizeVariation.ID
		orderItem.Quantity = quantity
		orderItem.Price = product.Price
		orderItem.DiscountRate = discountRate
		if err := db.DB.Create(&orderItem).Error; err != nil {
			log.Fatalf("create orderItem (%v) failure, got err %v", orderItem, err)
		}

		order.OrderItems = append(order.OrderItems, orderItem)
		order.CreatedAt = user.CreatedAt.Add(1 * time.Hour)
		order.PaymentAmount = order.Amount()
		if err := db.DB.Save(&order).Error; err != nil {
			log.Fatalf("Save order (%v) failure, got err %v", order, err)
		}
	}
}

func createWidgets() {
	// Normal banner
	type ImageStorage struct{ media_library.FileSystem }
	topBannerSetting := admin.QorWidgetSetting{}
	topBannerSetting.Name = "TopBanner"
	topBannerSetting.WidgetType = "NormalBanner"
	topBannerSetting.GroupName = "Banner"
	topBannerSetting.Scope = "from_google"
	topBannerValue := &struct {
		Title           string
		ButtonTitle     string
		Link            string
		BackgroundImage ImageStorage `sql:"type:varchar(4096)"`
		Logo            ImageStorage `sql:"type:varchar(4096)"`
	}{
		Title:       "Welcome Googlistas!",
		ButtonTitle: "LEARN MORE",
		Link:        "http://getqor.com",
	}
	if file, err := openFileByURL("http://qor3.s3.amazonaws.com/google_banner.jpg"); err == nil {
		defer file.Close()
		topBannerValue.BackgroundImage.Scan(file)
	} else {
		fmt.Printf("open file (%q) failure, got err %v", "banner", err)
	}

	if file, err := openFileByURL("http://qor3.s3.amazonaws.com/qor_logo.png"); err == nil {
		defer file.Close()
		topBannerValue.Logo.Scan(file)
	} else {
		fmt.Printf("open file (%q) failure, got err %v", "qor_logo", err)
	}

	topBannerSetting.SetSerializableArgumentValue(topBannerValue)
	if err := db.DB.Create(&topBannerSetting).Error; err != nil {
		log.Fatalf("Save widget (%v) failure, got err %v", topBannerSetting, err)
	}

	// SlideShow
	type slideImage struct {
		Title string
		Image media_library.FileSystem
	}
	slideshowSetting := admin.QorWidgetSetting{}
	slideshowSetting.Name = "TopBanner"
	slideshowSetting.GroupName = "Banner"
	slideshowSetting.WidgetType = "SlideShow"
	slideshowSetting.Scope = "default"
	slideshowValue := &struct {
		SlideImages []slideImage
	}{}
	slideDatas := [][]string{[]string{"Contra legem facit qui id facit quod lex prohibet.", "http://qor3.s3.amazonaws.com/slide1.jpg"},
		[]string{"Fictum, deserunt mollit anim laborum astutumque! Excepteur sint obcaecat cupiditat non proident culpa.", "http://qor3.s3.amazonaws.com/slide2.jpg"},
		[]string{"Excepteur sint obcaecat cupiditat non proident culpa.", "http://qor3.s3.amazonaws.com/slide3.jpg"}}
	for _, data := range slideDatas {
		slide := slideImage{Title: data[0]}
		if file, err := openFileByURL(data[1]); err == nil {
			defer file.Close()
			slide.Image.Scan(file)
		} else {
			fmt.Printf("open file (%q) failure, got err %v", "banner", err)
		}
		slideshowValue.SlideImages = append(slideshowValue.SlideImages, slide)
	}
	slideshowSetting.SetSerializableArgumentValue(slideshowValue)
	if err := db.DB.Create(&slideshowSetting).Error; err != nil {
		log.Fatalf("Save widget (%v) failure, got err %v", slideshowSetting, err)
	}

	// Feature Product
	featureProducts := admin.QorWidgetSetting{}
	featureProducts.Name = "FeatureProducts"
	featureProducts.WidgetType = "Products"
	featureProducts.SetSerializableArgumentValue(&struct {
		Products       []string
		ProductsSorter sorting.SortableCollection
	}{
		Products:       []string{"1", "2", "3", "4", "5", "6"},
		ProductsSorter: sorting.SortableCollection{PrimaryKeys: []string{"1", "2", "3", "4", "5", "6"}},
	})
	if err := db.DB.Create(&featureProducts).Error; err != nil {
		log.Fatalf("Save widget (%v) failure, got err %v", featureProducts, err)
	}
}

func findCategoryByName(name string) *models.Category {
	category := &models.Category{}
	if err := db.DB.Where(&models.Category{Name: name}).First(category).Error; err != nil {
		log.Fatalf("can't find category with name = %q, got err %v", name, err)
	}
	return category
}

func findCollectionByName(name string) *models.Collection {
	collection := &models.Collection{}
	if err := db.DB.Where(&models.Collection{Name: name}).First(collection).Error; err != nil {
		log.Fatalf("can't find collection with name = %q, got err %v", name, err)
	}
	return collection
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

func findProductByColorVariationID(colorVariationID uint) *models.Product {
	colorVariation := models.ColorVariation{}
	product := models.Product{}

	if err := db.DB.Find(&colorVariation, colorVariationID).Error; err != nil {
		log.Fatalf("query colorVariation (%v) failure, got err %v", colorVariation, err)
		return &product
	}
	if err := db.DB.Find(&product, colorVariation.ProductID).Error; err != nil {
		log.Fatalf("query product (%v) failure, got err %v", product, err)
		return &product
	}
	return &product
}

func randTime() time.Time {
	num := rand.Intn(10)
	return time.Now().Add(-time.Duration(num*24) * time.Hour)
}

func openFileByURL(rawURL string) (*os.File, error) {
	if fileURL, err := url.Parse(rawURL); err != nil {
		return nil, err
	} else {
		path := fileURL.Path
		segments := strings.Split(path, "/")
		fileName := segments[len(segments)-1]

		filePath := filepath.Join("/tmp", fileName)

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
