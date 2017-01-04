package admin

import (
	"github.com/qor/admin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config/seo"
	qor_seo "github.com/qor/seo"
)

func initSeo() {
	seo.SeoCollection = qor_seo.New("Common SEO")
	seo.SeoCollection.RegisterGlobalVaribles(&seo.SeoGlobalSetting{SiteName: "Qor Shop"})
	seo.SeoCollection.SettingResource = Admin.AddResource(&seo.MySeoSetting{}, &admin.Config{Invisible: true})
	seo.SeoCollection.RegisterSeo(&qor_seo.SEO{
		Name: "Default Page",
	})
	seo.SeoCollection.RegisterSeo(&qor_seo.SEO{
		Name:     "Product Page",
		Varibles: []string{"Name", "Code"},
		Context: func(objects ...interface{}) map[string]string {
			product := objects[0].(models.Product)
			context := make(map[string]string)
			context["Name"] = product.Name
			context["Code"] = product.Code
			return context
		},
	})
	Admin.AddResource(seo.SeoCollection, &admin.Config{Name: "SEO Setting", Menu: []string{"Site Management"}, Singleton: true, Priority: 2})
}
