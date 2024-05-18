package monolith_config

import (
	course_config "github.com/TesyarRAz/penggerak/internal/app/course/config"
	user_config "github.com/TesyarRAz/penggerak/internal/app/user/config"
	"github.com/TesyarRAz/penggerak/internal/pkg/config"
)

type App struct {
	userModule   *user_config.App
	courseModule *course_config.App
}

var _ config.App = &App{}

func NewApp(cfg *config.BootstrapConfig) *App {
	userModule := user_config.NewApp(cfg)
	courseModule := course_config.NewApp(cfg)

	return &App{
		userModule:   userModule,
		courseModule: courseModule,
	}
}

func (a *App) Provider() config.Provider {
	return config.CombineProvider(
		a.userModule.Provider(),
		a.courseModule.Provider(),
	)
}

func (a *App) Service(providers config.Provider) {
	a.userModule.Service(providers)
	a.courseModule.Service(providers)
}
