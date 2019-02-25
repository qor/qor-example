// +build enterprise

package enterprise

import (
	"enterprise.getqor.com/microsite"
	"enterprise.getqor.com/microsite/develop/aws_manager"
	"github.com/jinzhu/configor"
	"github.com/qor/admin"
	"github.com/qor/l10n"
	"github.com/qor/publish2"
	adminapp "github.com/qor/qor-example/app/admin"
	"github.com/qor/roles"
)

var MicroSite *microsite.MicroSite

type QorMicroSite struct {
	microsite.QorMicroSite
	publish2.Version
	publish2.Schedule
	publish2.Visible
	l10n.Locale
}

type AWSManagerConfig struct {
	AccessID  string `env:"AWS_ACCESS_KEY_ID" required:"true"`
	AccessKey string `env:"AWS_SECRET_ACCESS_KEY" required:"true"`
	Region    string `env:"AWS_Region" required:"true"`
	Bucket    string `env:"AWS_Bucket" required:"true"`
}

func SetupMicrosite(Admin *admin.Admin) {
	awsConfig := AWSManagerConfig{}
	configor.Load(&awsConfig)

	MicroSite = microsite.New(&microsite.Config{
		Widgets: adminapp.Widgets,
		DevelopManager: aws_manager.New(&aws_manager.Config{
			AccessID:  awsConfig.AccessID,
			AccessKey: awsConfig.AccessKey,
			Region:    awsConfig.Region,
			Bucket:    awsConfig.Bucket,
		})})

	MicroSite.Resource = Admin.AddResource(&QorMicroSite{}, &admin.Config{Name: "MicroSite", Menu: []string{"Pages Management"}})
	MicroSite.Resource.SetPrimaryFields("ID", "VersionName")

	Admin.AddResource(MicroSite, &admin.Config{Menu: []string{"Pages Management"}, Priority: 2})
	MicroSite.Resource.GetAction("VIEW S3 DEV SITE").Permission = roles.Deny("*", "*")
}
