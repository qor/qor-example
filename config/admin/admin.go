package admin

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/qor/action_bar"
	"github.com/qor/admin"
	"github.com/qor/i18n/exchange_actions"
	"github.com/qor/page_builder"
	"github.com/qor/qor"
	"github.com/qor/qor-example/config/db"
	"github.com/qor/qor-example/config/i18n"
	"github.com/qor/qor-example/models/blogs"
	"github.com/qor/qor-example/models/products"
	"github.com/qor/qor-example/models/settings"
	"github.com/qor/qor-example/models/stores"
	"github.com/qor/qor-example/models/users"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/validations"
	"github.com/qor/widget"
)

var ActionBar *action_bar.ActionBar
var Genders = []string{"Men", "Women", "Kids"}

func SetupAdmin(Admin *admin.Admin) {
	// Add User
	user := Admin.AddResource(&users.User{}, &admin.Config{Menu: []string{"User Management"}})
	user.Meta(&admin.Meta{Name: "Gender", Config: &admin.SelectOneConfig{Collection: []string{"Male", "Female", "Unknown"}}})
	user.Meta(&admin.Meta{Name: "Birthday", Type: "date"})
	user.Meta(&admin.Meta{Name: "Role", Config: &admin.SelectOneConfig{Collection: []string{"Admin", "Maintainer", "Member"}}})
	user.Meta(&admin.Meta{Name: "Password",
		Type:   "password",
		Valuer: func(interface{}, *qor.Context) interface{} { return "" },
		Setter: func(resource interface{}, metaValue *resource.MetaValue, context *qor.Context) {
			if newPassword := utils.ToString(metaValue.Value); newPassword != "" {
				bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
				if err != nil {
					context.DB.AddError(validations.NewError(user, "Password", "Can't encrpt password"))
					return
				}
				u := resource.(*users.User)
				u.Password = string(bcryptPassword)
			}
		},
	})
	user.Meta(&admin.Meta{Name: "Confirmed", Valuer: func(user interface{}, ctx *qor.Context) interface{} {
		if user.(*users.User).ID == 0 {
			return true
		}
		return user.(*users.User).Confirmed
	}})
	user.Meta(&admin.Meta{Name: "DefaultBillingAddress", Config: &admin.SelectOneConfig{Collection: userAddressesCollection}})
	user.Meta(&admin.Meta{Name: "DefaultShippingAddress", Config: &admin.SelectOneConfig{Collection: userAddressesCollection}})

	user.Filter(&admin.Filter{
		Name: "Role",
		Config: &admin.SelectOneConfig{
			Collection: []string{"Admin", "Maintainer", "Member"},
		},
	})

	user.IndexAttrs("ID", "Email", "Name", "Gender", "Role", "Balance")
	user.ShowAttrs(
		&admin.Section{
			Title: "Basic Information",
			Rows: [][]string{
				{"Name"},
				{"Email", "Password"},
				{"Avatar"},
				{"Gender", "Role", "Birthday"},
				{"Confirmed"},
			},
		},
		&admin.Section{
			Title: "Credit Information",
			Rows: [][]string{
				{"Balance"},
			},
		},
		&admin.Section{
			Title: "Accepts",
			Rows: [][]string{
				{"AcceptPrivate", "AcceptLicense", "AcceptNews"},
			},
		},
		&admin.Section{
			Title: "Default Addresses",
			Rows: [][]string{
				{"DefaultBillingAddress"},
				{"DefaultShippingAddress"},
			},
		},
		"Addresses",
	)
	user.EditAttrs(user.ShowAttrs())

	// Add Store
	store := Admin.AddResource(&stores.Store{}, &admin.Config{Menu: []string{"Store Management"}})
	store.Meta(&admin.Meta{Name: "Owner", Type: "single_edit"})
	store.AddValidator(&resource.Validator{
		Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if meta := metaValues.Get("Name"); meta != nil {
				if name := utils.ToString(meta.Value); strings.TrimSpace(name) == "" {
					return validations.NewError(record, "Name", "Name can't be blank")
				}
			}
			return nil
		},
	})

	// Blog Management
	article := Admin.AddResource(&blogs.Article{}, &admin.Config{Menu: []string{"Blog Management"}})
	article.IndexAttrs("ID", "VersionName", "ScheduledStartAt", "ScheduledEndAt", "Author", "Title")

	// Add Translations
	Admin.AddResource(i18n.I18n, &admin.Config{Menu: []string{"Site Management"}, Priority: 1})

	// Add Worker
	Worker := getWorker()
	exchange_actions.RegisterExchangeJobs(i18n.I18n, Worker)
	Admin.AddResource(Worker, &admin.Config{Menu: []string{"Site Management"}})

	// Add Setting
	Admin.AddResource(&settings.Setting{}, &admin.Config{Name: "Shop Setting", Singleton: true})

	// Add Search Center Resources
	// Admin.AddSearchResource(product, user, order)

	// Add ActionBar
	initWidgets()

	PageBuilderWidgets := widget.New(&widget.Config{DB: db.DB})
	PageBuilderWidgets.WidgetSettingResource = Admin.NewResource(&QorWidgetSetting{}, &admin.Config{Name: "PageBuilderWidgets"})
	PageBuilderWidgets.WidgetSettingResource.NewAttrs(
		&admin.Section{
			Rows: [][]string{{"Kind"}, {"SerializableMeta"}},
		},
	)
	PageBuilderWidgets.WidgetSettingResource.AddProcessor(&resource.Processor{
		Handler: func(value interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if widgetSetting, ok := value.(*QorWidgetSetting); ok {
				if widgetSetting.Name == "" {
					var count int
					context.GetDB().Set(admin.DisableCompositePrimaryKeyMode, "off").Model(&QorWidgetSetting{}).Count(&count)
					widgetSetting.Name = fmt.Sprintf("%v %v", utils.ToString(metaValues.Get("Kind").Value), count)
				}
			}
			return nil
		},
	})
	Admin.AddResource(PageBuilderWidgets)

	page := page_builder.New(&page_builder.Config{
		Admin:      Admin,
		PageModel:  &blogs.Page{},
		Containers: PageBuilderWidgets,
		// AdminConfig: &admin.Config{Name: "Campaign Pages or Builder", Menu: []string{"Sites & Campaign Pages"}, Priority: 2},
	})
	page.IndexAttrs("ID", "Title", "PublishLiveNow")

	// page := Admin.AddResource(&models.Page{})
	// page.Meta(&admin.Meta{
	// 	Name: "QorWidgetSettings",
	// 	Valuer: func(value interface{}, context *qor.Context) interface{} {
	// 		scope := context.GetDB().NewScope(value)
	// 		field, _ := scope.FieldByName("QorWidgetSettings")
	// 		context.GetDB().Model(value).Where("scope = ?", "default").Related(field.Field.Addr().Interface(), "QorWidgetSettings")
	// 		return field.Field.Interface()
	// 	},
	// 	Config: &admin.SelectManyConfig{
	// 		SelectionTemplate:  "metas/form/sortable_widgets.tmpl",
	// 		SelectMode:         "bottom_sheet",
	// 		DefaultCreating:    true,
	// 		RemoteDataResource: PageBuilderWidgets.WidgetSettingResource,
	// 	}})
	// page.Meta(&admin.Meta{Name: "QorWidgetSettingsSorter"})

	initSeo()
	initFuncMap()
	initRouter()
}

func sizeVariationCollection(resource interface{}, context *qor.Context) (results [][]string) {
	for _, sizeVariation := range products.SizeVariations() {
		results = append(results, []string{strconv.Itoa(int(sizeVariation.ID)), sizeVariation.Stringify()})
	}
	return
}

func userAddressesCollection(resource interface{}, context *qor.Context) (results [][]string) {
	var (
		user users.User
		DB   = context.DB
	)

	DB.Preload("Addresses").Where(context.ResourceID).First(&user)

	for _, address := range user.Addresses {
		results = append(results, []string{strconv.Itoa(int(address.ID)), address.Stringify()})
	}
	return
}
