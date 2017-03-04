// +build enterprise

package admin

import (
	"enterprise.getqor.com/microsite"
	"enterprise.getqor.com/microsite/develop/aws_manager"
	"github.com/jinzhu/configor"
	"github.com/qor/admin"
	"github.com/qor/qor-example/config"
)

var MicroSite *microsite.MicroSite

type QorMicroSite struct {
	microsite.QorMicroSite
}

type AWSManagerConfig struct {
	AccessID  string `env:"AWS_ACCESS_KEY_ID" required:"true"`
	AccessKey string `env:"AWS_SECRET_ACCESS_KEY" required:"true"`
	Region    string `env:"AWS_Region" required:"true"`
	Bucket    string `env:"AWS_Bucket" required:"true"`
}

func init() {
	initWidgets()
	awsConfig := AWSManagerConfig{}
	configor.Load(awsConfig)

	MicroSite = microsite.New(&microsite.Config{Dir: config.Root + "/public/microsites", Widgets: Widgets, DevelopManager: aws_manager.New(&aws_manager.Config{
		AccessID:  awsConfig.AccessID,
		AccessKey: awsConfig.AccessKey,
		Region:    awsConfig.Region,
		Bucket:    awsConfig.Bucket,
	})})

	MicroSite.Resource = Admin.AddResource(&QorMicroSite{}, &admin.Config{Name: "MicroSite"})

	Admin.AddResource(MicroSite)
}
