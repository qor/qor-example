package admin

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/media_library"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/qor/validations"
	"github.com/qor/transition"
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
	assetManager := Admin.AddResource(&media_library.AssetManager{}, &admin.Config{Invisible: true})

	// Add Product
	product := Admin.AddResource(&models.Product{}, &admin.Config{Menu: []string{"Product Management"}})
	product.Meta(&admin.Meta{Name: "MadeCountry", Type: "select_one", Collection: Countries})
	product.Meta(&admin.Meta{Name: "Description", Type: "rich_editor", Resource: assetManager})

	colorVariationMeta := product.Meta(&admin.Meta{Name: "ColorVariations"})
	colorVariation := colorVariationMeta.Resource
	colorVariation.NewAttrs("-Product")
	colorVariation.EditAttrs("-Product")

	sizeVariationMeta := colorVariation.Meta(&admin.Meta{Name: "SizeVariations"})
	sizeVariation := sizeVariationMeta.Resource
	sizeVariation.NewAttrs("-ColorVariation")
	sizeVariation.EditAttrs(
		&admin.Section{
			Rows: [][]string{
				{"Size", "AvailableQuantity"},
			},
		},
	)

	product.SearchAttrs("Name", "Code", "Category.Name", "Brand.Name")
	product.EditAttrs(
		&admin.Section{
			Title: "Basic Information",
			Rows: [][]string{
				{"Name"},
				{"Code", "Price"},
				{"Disabled"},
			}},
		&admin.Section{
			Title: "Organization",
			Rows: [][]string{
				{"Category", "Collections", "MadeCountry"},
			}},
		"Description",
		"ColorVariations",
	)

	for _, country := range Countries {
		var country = country
		product.Scope(&admin.Scope{Name: country, Group: "Made Country", Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
			return db.Where("made_country = ?", country)
		}})
	}

	product.IndexAttrs("-ColorVariations")
	product.Action(&admin.Action{
		Name: "disable",
		Handle: func(arg *admin.ActionArgument) error {
			for _, record := range arg.FindSelectedRecords() {
				arg.Context.DB.Model(record.(*models.Product)).Update("disabled", true)
			}
			return nil
		},
		Visibles: []string{"index", "edit", "show"},
	})
	product.Action(&admin.Action{
		Name: "enable",
		Handle: func(arg *admin.ActionArgument) error {
			for _, record := range arg.FindSelectedRecords() {
				arg.Context.DB.Model(record.(*models.Product)).Update("disabled", false)
			}
			return nil
		},
		Visibles: []string{"index", "edit", "show"},
	})

	type UpdatePriceActionArgument struct {
		Price float32
	}
	product.Action(&admin.Action{
		Name: "update price",
		Handle: func(arg *admin.ActionArgument) error {
			for _, record := range arg.FindSelectedRecords() {
				price := arg.Argument.(*UpdatePriceActionArgument).Price
				arg.Context.DB.Model(record.(*models.Product)).Update("price", price)
			}
			return nil
		},
		Resource: Admin.NewResource(&UpdatePriceActionArgument{}),
		Visibles: []string{"index"},
	})

	Admin.AddResource(&models.Color{}, &admin.Config{Menu: []string{"Product Management"}})
	Admin.AddResource(&models.Size{}, &admin.Config{Menu: []string{"Product Management"}})
	Admin.AddResource(&models.Category{}, &admin.Config{Menu: []string{"Product Management"}})
	Admin.AddResource(&models.Collection{}, &admin.Config{Menu: []string{"Product Management"}})

	// Add Order
	order := Admin.AddResource(&models.Order{}, &admin.Config{Menu: []string{"Order Management"}})
	order.Meta(&admin.Meta{Name: "ShippingAddress", Type: "single_edit"})
	order.Meta(&admin.Meta{Name: "BillingAddress", Type: "single_edit"})
	order.Meta(&admin.Meta{Name: "ShippedAt", Type: "date"})

	orderItemMeta := order.Meta(&admin.Meta{Name: "OrderItems"})
	orderItemMeta.Resource.Meta(&admin.Meta{Name: "SizeVariation", Type: "select_one", Collection: sizeVariationCollection})

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

	// define actions for Order
	type trackingNumberArgument struct {
		TrackingNumber string
	}

	order.Action(&admin.Action{
		Name: "Ship",
		Handle: func(argument *admin.ActionArgument) error {
			trackingNumberArgument := argument.Argument.(*trackingNumberArgument)
			for _, record := range argument.FindSelectedRecords() {
				argument.Context.GetDB().Model(record).UpdateColumn("tracking_number", trackingNumberArgument.TrackingNumber)
			}
			return nil
		},
		Resource: Admin.NewResource(&trackingNumberArgument{}),
	})

	order.Action(&admin.Action{
		Name: "Cancel",
		Handle: func(argument *admin.ActionArgument) error {
			for _, order := range argument.FindSelectedRecords() {
				db := argument.Context.GetDB()
				if err := models.OrderState.Trigger("cancel", order.(*models.Order), db); err != nil {
					return err
				}
				db.Select("state").Save(order)
			}
			return nil
		},
	})

	order.IndexAttrs("User", "PaymentAmount", "ShippedAt", "CancelledAt", "State", "ShippingAddress")
	order.NewAttrs("-DiscountValue", "-AbandonedReason", "-CancelledAt")
	order.EditAttrs("-DiscountValue", "-AbandonedReason", "-CancelledAt")
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

	// Add User
	user := Admin.AddResource(&models.User{})
	user.IndexAttrs("ID", "Email", "Name", "Gender", "Role")

	// Add Publish
	Admin.AddResource(db.Publish, &admin.Config{Singleton: true})

	// Add Worker
	Admin.AddResource(getWorker())

	initFuncMap()
	initRouter()
}

func sizeVariationCollection(resource interface{}, context *qor.Context) (results [][]string) {
	for _, sizeVariation := range models.SizeVariations() {
		results = append(results, []string{strconv.Itoa(int(sizeVariation.ID)), sizeVariation.Stringify()})
	}
	return
}
