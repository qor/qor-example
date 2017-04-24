package admin

import (
	"github.com/qor/admin"
	qor_seo "github.com/qor/seo"

	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config/seo"
)

func initSeo() {
	seo.SEOCollection = qor_seo.New("Common SEO")
	seo.SEOCollection.RegisterGlobalVaribles(&seo.SEOGlobalSetting{SiteName: "Qor Shop"})
	seo.SEOCollection.SettingResource = Admin.AddResource(&seo.MySEOSetting{}, &admin.Config{Invisible: true})
	seo.SEOCollection.RegisterSEO(&qor_seo.SEO{
		Name: "Default Page",
	})
	seo.SEOCollection.RegisterSEO(&qor_seo.SEO{
		Name:     "Product Page",
		Varibles: []string{"Name", "Code", "CategoryName"},
		Context: func(objects ...interface{}) map[string]string {
			product := objects[0].(models.Product)
			context := make(map[string]string)
			context["Name"] = product.Name
			context["Code"] = product.Code
			context["CategoryName"] = product.Category.Name
			return context
		},
	})
	seo.SEOCollection.RegisterSEO(&qor_seo.SEO{
		Name:     "Category Page",
		Varibles: []string{"Name", "Code"},
		Context: func(objects ...interface{}) map[string]string {
			category := objects[0].(models.Category)
			context := make(map[string]string)
			context["Name"] = category.Name
			context["Code"] = category.Code
			return context
		},
	})
	Admin.AddResource(seo.SEOCollection, &admin.Config{Name: "SEO Setting", Menu: []string{"Site Management"}, Singleton: true, Priority: 2})
}
