package admin

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/activity"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/transition"
	"github.com/qor/qor/utils"
	"github.com/qor/qor/validations"
	"github.com/qor/worker"
)

var Admin *admin.Admin
var Countries = []string{"China", "Japan", "USA"}

func init() {
	Admin = admin.New(&qor.Config{DB: db.Publish.DraftDB()})
	Admin.SetSiteName("Qor DEMO")
	Admin.SetAuth(Auth{})

	// Add Dashboard
	Admin.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/admin"})

	// Add Asset Manager, for rich editor
	assetManager := Admin.AddResource(&admin.AssetManager{}, &admin.Config{Invisible: true})

	// Add Product
	product := Admin.AddResource(&models.Product{}, &admin.Config{Menu: []string{"Product Management"}})
	product.Meta(&admin.Meta{Name: "MadeCountry", Type: "select_one", Collection: Countries})
	product.Meta(&admin.Meta{Name: "Description", Type: "rich_editor", Resource: assetManager})
	sizeVariation := Admin.NewResource(&models.SizeVariation{}, &admin.Config{Invisible: true})
	sizeVariation.NewAttrs("-ColorVariation")
	sizeVariation.EditAttrs("-ColorVariation")
	colorVariation := Admin.NewResource(&models.ColorVariation{}, &admin.Config{Invisible: true})
	colorVariation.Meta(&admin.Meta{Name: "SizeVariations", Resource: sizeVariation})
	colorVariation.NewAttrs("-Product")
	colorVariation.EditAttrs("-Product")
	product.Meta(&admin.Meta{Name: "ColorVariations", Resource: colorVariation})
	product.SearchAttrs("Name", "Code", "Category.Name", "Brand.Name")

	for _, country := range Countries {
		var country = country
		product.Scope(&admin.Scope{Name: country, Group: "Made Country", Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
			return db.Where("made_country = ?", country)
		}})
	}

	product.IndexAttrs("-ColorVariations")

	Admin.AddResource(&models.Color{}, &admin.Config{Menu: []string{"Product Management"}})
	Admin.AddResource(&models.Size{}, &admin.Config{Menu: []string{"Product Management"}})
	Admin.AddResource(&models.Category{}, &admin.Config{Menu: []string{"Product Management"}})
	Admin.AddResource(&models.Collection{}, &admin.Config{Menu: []string{"Product Management"}})

	// Add Order
	orderItem := Admin.NewResource(&models.OrderItem{})
	orderItem.Meta(&admin.Meta{Name: "SizeVariation", Type: "select_one", Collection: sizeVariationCollection})

	order := Admin.AddResource(&models.Order{}, &admin.Config{Menu: []string{"Order Management"}})
	order.Meta(&admin.Meta{Name: "ShippingAddress", Type: "single_edit"})
	order.Meta(&admin.Meta{Name: "BillingAddress", Type: "single_edit"})
	order.Meta(&admin.Meta{Name: "OrderItems", Resource: orderItem})
	activity.RegisterActivityMeta(order)

	// define scopes for Order
	for _, state := range []string{"checkout", "cancelled", "paid", "paid_cancelled", "processing", "shipped", "returned"} {
		var state = state
		order.Scope(&admin.Scope{
			Name:  state,
			Label: strings.Title(strings.Replace(state, "_", " ", -1)),
			Group: "Order Status",
			Handle: func(db *gorm.DB, context *qor.Context) *gorm.DB {
				return db.Where(models.Order{Transition: transition.Transition{State: state}})
			},
		})
	}
	order.IndexAttrs("-DiscountValue", "-OrderItems", "-AbandonedReason")
	order.NewAttrs("-DiscountValue", "-AbandonedReason")
	order.EditAttrs("-DiscountValue", "-AbandonedReason")
	order.ShowAttrs("-DiscountValue", "-AbandonedReason")
	order.SearchAttrs("User.Name", "User.Email", "ShippingAddress.ContactName", "ShippingAddress.Address1", "ShippingAddress.Address2")

	// Define another resource for same model
	abandonedOrder := Admin.AddResource(&models.Order{}, &admin.Config{Name: "Abandoned Order", Menu: []string{"Order Management"}})
	abandonedOrder.Meta(&admin.Meta{Name: "ShippingAddress", Type: "single_edit"})
	abandonedOrder.Meta(&admin.Meta{Name: "BillingAddress", Type: "single_edit"})

	// Define default scope for abandoned orders
	abandonedOrder.Scope(&admin.Scope{
		Default: true,
		Handle: func(db *gorm.DB, context *qor.Context) *gorm.DB {
			return db.Where("abandoned_reason IS NOT NULL AND abandoned_reason <> ?", "")
		},
	})

	// Define scopes for abandoned orders
	for _, amount := range []int{5000, 10000, 20000} {
		var amount = amount
		abandonedOrder.Scope(&admin.Scope{
			Name:  fmt.Sprint(amount),
			Group: "Amount Greater Than",
			Handle: func(db *gorm.DB, context *qor.Context) *gorm.DB {
				return db.Where("payment_amount > ?", amount)
			},
		})
	}

	abandonedOrder.IndexAttrs("-ShippingAddress", "-BillingAddress", "-DiscountValue", "-OrderItems")
	abandonedOrder.NewAttrs("-DiscountValue")
	abandonedOrder.EditAttrs("-DiscountValue")
	abandonedOrder.ShowAttrs("-DiscountValue")

	// Add Store
	store := Admin.AddResource(&models.Store{}, &admin.Config{Menu: []string{"Store Management"}})
	store.AddValidator(func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
		if meta := metaValues.Get("Name"); meta != nil {
			if name := utils.ToString(meta.Value); strings.TrimSpace(name) == "" {
				return validations.NewError(record, "Name", "Name can't be blank")
			}
		}
		return nil
	})

	// Add Translations
	Admin.AddResource(config.Config.I18n, &admin.Config{Menu: []string{"Site Management"}})

	// Add Setting
	Admin.AddResource(&models.Setting{}, &admin.Config{Singleton: true})

	// Add Worker
	Worker := worker.New()
	Worker.RegisterJob(worker.Job{
		Name: "send_newsletter",
		Handler: func(interface{}) error {
			fmt.Println("sending newsletter...")
			time.Sleep(5 * time.Second)
			return nil
		},
		Resource: Admin.NewResource(&struct {
			Subject      string
			Content      string `sql:"size:65532"`
			SendPassword string
		}{}),
	})
	Admin.AddResource(Worker)

	// Add User
	user := Admin.AddResource(&models.User{})
	user.IndexAttrs("ID", "Email", "Name", "Gender", "Role")

	// Add Publish
	Admin.AddResource(db.Publish, &admin.Config{Singleton: true})

	// Add Seo
	Admin.AddResource(&models.Seo{}, &admin.Config{Name: "Meta Data", Singleton: true})

	// Add Search Center
	Admin.AddSearchResource(order, user, product)

	initFuncMap()
	initRouter()
}

func sizeVariationCollection(resource interface{}, context *qor.Context) (results [][]string) {
	for _, sizeVariation := range models.SizeVariations() {
		results = append(results, []string{strconv.Itoa(int(sizeVariation.ID)), sizeVariation.Stringify()})
	}
	return
}
