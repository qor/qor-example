package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/media"
	"github.com/qor/media/media_library"
	"github.com/qor/media/oss"
	"github.com/qor/publish2"
	qor_seo "github.com/qor/seo"
	"github.com/qor/slug"
	"github.com/qor/sorting"
	"github.com/qor/validations"

	"github.com/qor/qor-example/config/seo"
	"github.com/qor/qor-example/db"
)

type Product struct {
	gorm.Model
	l10n.Locale
	sorting.SortingDESC

	Name                  string
	NameWithSlug          slug.Slug    `l10n:"sync"`
	Code                  string       `l10n:"sync"`
	CategoryID            uint         `l10n:"sync"`
	Category              Category     `l10n:"sync"`
	Collections           []Collection `l10n:"sync" gorm:"many2many:product_collections;"`
	MadeCountry           string       `l10n:"sync"`
	Gender                string       `l10n:"sync"`
	MainImage             media_library.MediaBox
	Price                 float32          `l10n:"sync"`
	Description           string           `sql:"size:2000"`
	ColorVariations       []ColorVariation `l10n:"sync"`
	ColorVariationsSorter sorting.SortableCollection
	ProductProperties     ProductProperties `sql:"type:text"`
	Seo                   qor_seo.Setting

	Variations []ProductVariation

	publish2.Version
	publish2.Schedule
	publish2.Visible
}

type ProductVariation struct {
	gorm.Model
	ProductID *uint
	Product   Product

	Color      Color `variations:"primary"`
	ColorID    *uint
	Size       Size `variations:"primary"`
	SizeID     *uint
	Material   Material `variations:"primary"`
	MaterialID *uint

	SKU               string
	ReceiptName       string
	Featured          bool
	Price             uint
	SellingPrice      uint
	AvailableQuantity uint
	Images            media_library.MediaBox
}

func (product Product) GetSEO() *qor_seo.SEO {
	return seo.SEOCollection.GetSEO("Product Page")
}

func (product Product) DefaultPath() string {
	defaultPath := "/"
	if len(product.ColorVariations) > 0 {
		defaultPath = fmt.Sprintf("/products/%s_%s", product.Code, product.ColorVariations[0].ColorCode)
	}
	return defaultPath
}

func (product Product) MainImageURL(styles ...string) string {
	style := "main"
	if len(styles) > 0 {
		style = styles[0]
	}

	if len(product.MainImage.Files) > 0 {
		return product.MainImage.URL(style)
	}

	for _, cv := range product.ColorVariations {
		return cv.MainImageURL()
	}

	return "/images/default_product.png"
}

func (product Product) Validate(db *gorm.DB) {
	if strings.TrimSpace(product.Name) == "" {
		db.AddError(validations.NewError(product, "Name", "Name can not be empty"))
	}

	if strings.TrimSpace(product.Code) == "" {
		db.AddError(validations.NewError(product, "Code", "Code can not be empty"))
	}
}

type ProductImage struct {
	gorm.Model
	Title        string
	Color        Color
	ColorID      uint
	Category     Category
	CategoryID   uint
	SelectedType string
	File         media_library.MediaLibraryStorage `sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}}"`
}

func (productImage ProductImage) Validate(db *gorm.DB) {
	if strings.TrimSpace(productImage.Title) == "" {
		db.AddError(validations.NewError(productImage, "Title", "Title can not be empty"))
	}
}

func (productImage *ProductImage) SetSelectedType(typ string) {
	productImage.SelectedType = typ
}

func (productImage *ProductImage) GetSelectedType() string {
	return productImage.SelectedType
}

func (productImage *ProductImage) ScanMediaOptions(mediaOption media_library.MediaOption) error {
	if bytes, err := json.Marshal(mediaOption); err == nil {
		return productImage.File.Scan(bytes)
	} else {
		return err
	}
}

func (productImage *ProductImage) GetMediaOption() (mediaOption media_library.MediaOption) {
	mediaOption.Video = productImage.File.Video
	mediaOption.FileName = productImage.File.FileName
	mediaOption.URL = productImage.File.URL()
	mediaOption.OriginalURL = productImage.File.URL("original")
	mediaOption.CropOptions = productImage.File.CropOptions
	mediaOption.Sizes = productImage.File.GetSizes()
	mediaOption.Description = productImage.File.Description
	return
}

type ProductProperties []ProductProperty

type ProductProperty struct {
	Name  string
	Value string
}

func (productProperties *ProductProperties) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, productProperties)
	case string:
		if v != "" {
			return productProperties.Scan([]byte(v))
		}
	default:
		return errors.New("not supported")
	}
	return nil
}

func (productProperties ProductProperties) Value() (driver.Value, error) {
	if len(productProperties) == 0 {
		return nil, nil
	}
	return json.Marshal(productProperties)
}

type ColorVariation struct {
	gorm.Model
	ProductID      uint
	Product        Product
	ColorID        uint
	Color          Color
	ColorCode      string
	Images         media_library.MediaBox
	SizeVariations []SizeVariation
	publish2.SharedVersion
}

type ColorVariationImage struct {
	gorm.Model
	ColorVariationID uint
	Image            ColorVariationImageStorage `sql:"type:varchar(4096)"`
}

type ColorVariationImageStorage struct{ oss.OSS }

func (colorVariation ColorVariation) MainImageURL() string {
	if len(colorVariation.Images.Files) > 0 {
		return colorVariation.Images.URL()
	}
	return "/images/default_product.png"
}

func (ColorVariationImageStorage) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":  {Width: 320, Height: 320},
		"middle": {Width: 640, Height: 640},
		"big":    {Width: 1280, Height: 1280},
	}
}

type SizeVariation struct {
	gorm.Model
	ColorVariationID  uint
	ColorVariation    ColorVariation
	SizeID            uint
	Size              Size
	AvailableQuantity uint
	publish2.SharedVersion
}

func SizeVariations() []SizeVariation {
	sizeVariations := []SizeVariation{}
	if err := db.DB.Preload("ColorVariation.Color").Preload("ColorVariation.Product").Preload("Size").Find(&sizeVariations).Error; err != nil {
		log.Fatalf("query sizeVariations (%v) failure, got err %v", sizeVariations, err)
		return sizeVariations
	}
	return sizeVariations
}

func (sizeVariation SizeVariation) Stringify() string {
	if colorVariation := sizeVariation.ColorVariation; colorVariation.ID != 0 {
		product := colorVariation.Product
		return fmt.Sprintf("%s (%s-%s-%s)", product.Name, product.Code, colorVariation.Color.Code, sizeVariation.Size.Code)
	}
	return fmt.Sprint(sizeVariation.ID)
}
