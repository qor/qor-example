package admin

import (
	"github.com/qor/admin"
	"github.com/qor/qor-example/models/products"
	"github.com/qor/qor-example/models/seo"
	qor_seo "github.com/qor/seo"
)

// SetupSEO add seo
func SetupSEO(Admin *admin.Admin) {
	openGraphConfig := &qor_seo.OpenGraphConfig{ImageResource: Admin.GetResource("MediaLibrary")}

	seo.SEOCollection = qor_seo.New("Common SEO")
	seo.SEOCollection.RegisterGlobalVaribles(&seo.SEOGlobalSetting{SiteName: "Qor Shop"})
	seo.SEOCollection.SettingResource = Admin.AddResource(&seo.MySEOSetting{}, &admin.Config{Invisible: true})

	seo.SEOCollection.RegisterSEO(&qor_seo.SEO{
		Name:      "Default Page",
		OpenGraph: openGraphConfig,
	})

	seo.SEOCollection.RegisterSEO(&qor_seo.SEO{
		Name:      "Product Page",
		OpenGraph: openGraphConfig,
		Varibles:  []string{"Name", "Code", "CategoryName", "VariantURL", "VariantImage"},
		Context: func(objects ...interface{}) map[string]string {
			context := make(map[string]string)

			var product products.Product
			for _, obj := range objects {
				var ok bool
				product, ok = obj.(products.Product)
				if ok {
					context["Name"] = product.Name
					context["Code"] = product.Code
					context["CategoryName"] = product.Category.Name
				}

				if cv, ok := obj.(products.ColorVariation); ok {
					context["VariantURL"] = cv.URL(product)
					context["VariantImage"] = cv.MainImageURL()
				}
			}

			return context
		},
	})

	seo.SEOCollection.RegisterSEO(&qor_seo.SEO{
		Name:      "Category Page",
		OpenGraph: openGraphConfig,
		Varibles:  []string{"Name", "Code"},
		Context: func(objects ...interface{}) map[string]string {
			category := objects[0].(products.Category)
			context := make(map[string]string)
			context["Name"] = category.Name
			context["Code"] = category.Code
			return context
		},
	})

	Admin.AddResource(seo.SEOCollection, &admin.Config{Name: "SEO Setting", Menu: []string{"Site Management"}, Singleton: true, Priority: 2})
}
