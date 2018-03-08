package products

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/media"
	"github.com/qor/media/media_library"
	"github.com/qor/qor"
	"github.com/qor/qor-example/config/application"
	"github.com/qor/qor-example/models/products"
	"github.com/qor/qor-example/utils/funcmapmaker"
	"github.com/qor/render"
)

var Genders = []string{"Men", "Women", "Kids"}

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	controller := &Controller{View: render.New(&render.Config{AssetFileSystem: application.AssetFS.NameSpace("products")}, "app/products/views")}

	funcmapmaker.AddFuncMapMaker(controller.View)
	app.ConfigureAdmin(application.Admin)

	application.Router.Get("/products", controller.Index)
	application.Router.Get("/products/{code}", controller.Show)
	application.Router.Get("/{gender:^(men|women|kids)$}", controller.Gender)
	application.Router.Get("/category/{code}", controller.Category)
}

// ConfigureAdmin configure admin interface
func (App) ConfigureAdmin(Admin *admin.Admin) {
	// Produc Management
	Admin.AddMenu(&admin.Menu{Name: "Product Management", Priority: 1})
	color := Admin.AddResource(&products.Color{}, &admin.Config{Menu: []string{"Product Management"}, Priority: -5})
	Admin.AddResource(&products.Size{}, &admin.Config{Menu: []string{"Product Management"}, Priority: -4})
	Admin.AddResource(&products.Material{}, &admin.Config{Menu: []string{"Product Management"}, Priority: -4})

	category := Admin.AddResource(&products.Category{}, &admin.Config{Menu: []string{"Product Management"}, Priority: -3})
	category.Meta(&admin.Meta{Name: "Categories", Type: "select_many"})

	collection := Admin.AddResource(&products.Collection{}, &admin.Config{Menu: []string{"Product Management"}, Priority: -2})

	// Add ProductImage as Media Libraray
	ProductImagesResource := Admin.AddResource(&products.ProductImage{}, &admin.Config{Menu: []string{"Product Management"}, Priority: -1})

	ProductImagesResource.Filter(&admin.Filter{
		Name:       "SelectedType",
		Label:      "Media Type",
		Operations: []string{"contains"},
		Config:     &admin.SelectOneConfig{Collection: [][]string{{"video", "Video"}, {"image", "Image"}, {"file", "File"}, {"video_link", "Video Link"}}},
	})
	ProductImagesResource.Filter(&admin.Filter{
		Name:   "Color",
		Config: &admin.SelectOneConfig{RemoteDataResource: color},
	})
	ProductImagesResource.Filter(&admin.Filter{
		Name:   "Category",
		Config: &admin.SelectOneConfig{RemoteDataResource: category},
	})
	ProductImagesResource.IndexAttrs("File", "Title")

	// Add Product
	product := Admin.AddResource(&products.Product{}, &admin.Config{Menu: []string{"Product Management"}})
	product.Meta(&admin.Meta{Name: "Gender", Config: &admin.SelectOneConfig{Collection: Genders, AllowBlank: true}})

	productPropertiesRes := product.Meta(&admin.Meta{Name: "ProductProperties"}).Resource
	productPropertiesRes.NewAttrs(&admin.Section{
		Rows: [][]string{{"Name", "Value"}},
	})
	productPropertiesRes.EditAttrs(&admin.Section{
		Rows: [][]string{{"Name", "Value"}},
	})

	product.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{Plugins: []admin.RedactorPlugin{
		{Name: "medialibrary", Source: "/admin/assets/javascripts/qor_redactor_medialibrary.js"},
		{Name: "table", Source: "/vendors/redactor_table.js"},
	},
		Settings: map[string]interface{}{
			"medialibraryUrl": "/admin/product_images",
		},
	}})
	product.Meta(&admin.Meta{Name: "Category", Config: &admin.SelectOneConfig{AllowBlank: true}})
	product.Meta(&admin.Meta{Name: "Collections", Config: &admin.SelectManyConfig{SelectMode: "bottom_sheet"}})

	product.Meta(&admin.Meta{Name: "MainImage", Config: &media_library.MediaBoxConfig{
		RemoteDataResource: ProductImagesResource,
		Max:                1,
		Sizes: map[string]*media.Size{
			"main": {Width: 560, Height: 700},
		},
	}})
	product.Meta(&admin.Meta{Name: "MainImageURL", Valuer: func(record interface{}, context *qor.Context) interface{} {
		if p, ok := record.(*products.Product); ok {
			result := bytes.NewBufferString("")
			tmpl, _ := template.New("").Parse("<img src='{{.image}}'></img>")
			tmpl.Execute(result, map[string]string{"image": p.MainImageURL()})
			return template.HTML(result.String())
		}
		return ""
	}})

	product.Filter(&admin.Filter{
		Name:   "Collections",
		Config: &admin.SelectOneConfig{RemoteDataResource: collection},
	})

	product.Filter(&admin.Filter{
		Name: "Featured",
	})

	product.Filter(&admin.Filter{
		Name: "Name",
		Type: "string",
	})

	product.Filter(&admin.Filter{
		Name: "Code",
	})

	product.Filter(&admin.Filter{
		Name: "Price",
		Type: "number",
	})

	product.Filter(&admin.Filter{
		Name: "CreatedAt",
	})

	product.Action(&admin.Action{
		Name:        "Import Product",
		URLOpenType: "slideout",
		URL: func(record interface{}, context *admin.Context) string {
			return "/admin/workers/new?job=Import Products"
		},
		Modes: []string{"collection"},
	})

	type updateInfo struct {
		CategoryID  uint
		Category    *products.Category
		MadeCountry string
		Gender      string
	}

	updateInfoRes := Admin.NewResource(&updateInfo{})
	product.Action(&admin.Action{
		Name:     "Update Info",
		Resource: updateInfoRes,
		Handler: func(argument *admin.ActionArgument) error {
			newProductInfo := argument.Argument.(*updateInfo)
			for _, record := range argument.FindSelectedRecords() {
				fmt.Printf("%#v\n", record)
				if product, ok := record.(*products.Product); ok {
					if newProductInfo.Category != nil {
						product.Category = *newProductInfo.Category
					}
					if newProductInfo.MadeCountry != "" {
						product.MadeCountry = newProductInfo.MadeCountry
					}
					if newProductInfo.Gender != "" {
						product.Gender = newProductInfo.Gender
					}
					argument.Context.GetDB().Save(product)
				}
			}
			return nil
		},
		Modes: []string{"batch"},
	})

	product.UseTheme("grid")

	// variationsResource := product.Meta(&admin.Meta{Name: "Variations", Config: &variations.VariationsConfig{}}).Resource
	// if imagesMeta := variationsResource.GetMeta("Images"); imagesMeta != nil {
	// 	imagesMeta.Config = &media_library.MediaBoxConfig{
	// 		RemoteDataResource: ProductImagesResource,
	// 		Sizes: map[string]*media.Size{
	// 			"icon":    {Width: 50, Height: 50},
	// 			"thumb":   {Width: 100, Height: 100},
	// 			"display": {Width: 300, Height: 300},
	// 		},
	// 	}
	// }

	// variationsResource.EditAttrs("-ID", "-Product")
	// oldSearchHandler := product.SearchHandler
	// product.SearchHandler = func(keyword string, context *qor.Context) *gorm.DB {
	// 	context.SetDB(context.GetDB().Preload("Variations.Color").Preload("Variations.Size").Preload("Variations.Material"))
	// 	return oldSearchHandler(keyword, context)
	// }
	colorVariationMeta := product.Meta(&admin.Meta{Name: "ColorVariations"})
	colorVariation := colorVariationMeta.Resource
	colorVariation.Meta(&admin.Meta{Name: "Images", Config: &media_library.MediaBoxConfig{
		RemoteDataResource: ProductImagesResource,
		Sizes: map[string]*media.Size{
			"icon":    {Width: 50, Height: 50},
			"preview": {Width: 300, Height: 300},
			"listing": {Width: 640, Height: 640},
		},
	}})

	colorVariation.NewAttrs("-Product", "-ColorCode")
	colorVariation.EditAttrs("-Product", "-ColorCode")

	sizeVariationMeta := colorVariation.Meta(&admin.Meta{Name: "SizeVariations"})
	sizeVariation := sizeVariationMeta.Resource
	sizeVariation.EditAttrs(
		&admin.Section{
			Rows: [][]string{
				{"Size", "AvailableQuantity"},
				{"ShareableVersion"},
			},
		},
	)
	sizeVariation.NewAttrs(sizeVariation.EditAttrs())

	product.SearchAttrs("Name", "Code", "Category.Name", "Brand.Name")
	product.IndexAttrs("MainImageURL", "Name", "Featured", "Price", "VersionName", "PublishLiveNow")
	product.EditAttrs(
		&admin.Section{
			Title: "Seo Meta",
			Rows: [][]string{
				{"Seo"},
			}},
		&admin.Section{
			Title: "Basic Information",
			Rows: [][]string{
				{"Name", "Featured"},
				{"Code", "Price"},
				{"MainImage"},
			}},
		&admin.Section{
			Title: "Organization",
			Rows: [][]string{
				{"Category", "Gender"},
				{"Collections"},
			}},
		"ProductProperties",
		"Description",
		"ColorVariations",
		"PublishReady",
	)
	// product.ShowAttrs(product.EditAttrs())
	product.NewAttrs(product.EditAttrs())

	for _, gender := range Genders {
		var gender = gender
		product.Scope(&admin.Scope{Name: gender, Group: "Gender", Handler: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
			return db.Where("gender = ?", gender)
		}})
	}

	product.Action(&admin.Action{
		Name: "View On Site",
		URL: func(record interface{}, context *admin.Context) string {
			if product, ok := record.(*products.Product); ok {
				return fmt.Sprintf("/products/%v", product.Code)
			}
			return "#"
		},
		Modes: []string{"menu_item", "edit"},
	})

}
