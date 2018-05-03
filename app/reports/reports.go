package reports

import (
	"github.com/qor/admin"
	"github.com/qor/metas/daterange"
	"github.com/qor/metas/frequency"
	"github.com/qor/notification"
	"github.com/qor/qor-example/config/application"
	"github.com/qor/worker"
)

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
	kpiWorker := worker.New()

	// KPI Report
	type kpiReport struct {
		Template string
		daterange.DateRange
		frequency.Frequency
		notification.Notification
	}

	kpiWorker.RegisterJob(&worker.Job{
		Name: "KPI Report",
		Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
			return nil
		},
		Resource: application.Admin.NewResource(&kpiReport{}),
	})

	application.Admin.AddResource(kpiWorker, &admin.Config{Menu: []string{"Report Management", "KPI Report"}, Priority: 1})

	productWorker := worker.New()

	// Product Report
	type productReport struct {
		Template string
		daterange.DateRange
		frequency.Frequency
		notification.Notification
	}

	productWorker.RegisterJob(&worker.Job{
		Name: "Product Report",
		Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
			return nil
		},
		Resource: application.Admin.NewResource(&productReport{}),
	})

	application.Admin.AddResource(productWorker, &admin.Config{Menu: []string{"Report Management", "Product Report"}})

	orderWorker := worker.New()

	// Order & Campaign Report
	type orderReport struct {
		Template string
		daterange.DateRange
		frequency.Frequency
		notification.Notification
	}

	orderWorker.RegisterJob(&worker.Job{
		Name: "Order & Campaign Report",
		Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
			return nil
		},
		Resource: application.Admin.NewResource(&orderReport{}),
	})

	application.Admin.AddResource(orderWorker, &admin.Config{Menu: []string{"Report Management", "Order & Campaign Report"}})

	userWorker := worker.New()

	// User Report
	type userReport struct {
		Template string
		daterange.DateRange
		frequency.Frequency
		notification.Notification
	}

	userWorker.RegisterJob(&worker.Job{
		Name: "User Report",
		Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
			return nil
		},
		Resource: application.Admin.NewResource(&userReport{}),
	})

	application.Admin.AddResource(userWorker, &admin.Config{Menu: []string{"Report Management", "User Report"}})
}
