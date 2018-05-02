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
	Worker := worker.New()

	// KPI Report
	type kpiReport struct {
		Template string
		daterange.DateRange
		frequency.Frequency
		notification.Notification
	}

	Worker.RegisterJob(&worker.Job{
		Name: "KPI Report",
		Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
			return nil
		},
		Resource: application.Admin.NewResource(&kpiReport{}),
	})

	application.Admin.AddResource(Worker, &admin.Config{Menu: []string{"Report Management", "KPI Report"}})

	// Product Report
	type productReport struct {
		Template string
		daterange.DateRange
		frequency.Frequency
		notification.Notification
	}

	Worker.RegisterJob(&worker.Job{
		Name: "Product Report",
		Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
			return nil
		},
		Resource: application.Admin.NewResource(&productReport{}),
	})

	application.Admin.AddResource(Worker, &admin.Config{Menu: []string{"Report Management", "Product Report"}})

	// Order & Campaign Report
	type orderReport struct {
		Template string
		daterange.DateRange
		frequency.Frequency
		notification.Notification
	}

	Worker.RegisterJob(&worker.Job{
		Name: "Order & Campaign Report",
		Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
			return nil
		},
		Resource: application.Admin.NewResource(&orderReport{}),
	})

	application.Admin.AddResource(Worker, &admin.Config{Menu: []string{"Report Management", "Order & Campaign Report"}})

	// User Report
	type userReport struct {
		Template string
		daterange.DateRange
		frequency.Frequency
		notification.Notification
	}

	Worker.RegisterJob(&worker.Job{
		Name: "User Report",
		Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
			return nil
		},
		Resource: application.Admin.NewResource(&userReport{}),
	})

	application.Admin.AddResource(Worker, &admin.Config{Menu: []string{"Report Management", "User Report"}})
}
