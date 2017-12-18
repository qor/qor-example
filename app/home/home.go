package home

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

func (App) ConfigureApplication(application *Application) {
}
