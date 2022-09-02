package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"

	"github.com/karashiiro/plugin-platform/server/pkg/events"
	"github.com/karashiiro/plugin-platform/server/pkg/kamori"
	_ "github.com/karashiiro/plugin-platform/server/pkg/migrations"
	"github.com/karashiiro/plugin-platform/server/pkg/routes"
)

func main() {
	app := pocketbase.New()

	checkPlugins := kamori.BuildCheckPlugins(app)
	ticker := time.NewTicker(5 * time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				err := checkPlugins()
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()
	defer func() {
		ticker.Stop()
		done <- true
	}()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodGet,
			Path:    "/api/presets/popular",
			Handler: routes.BuildPopularPresetsHandler(app),
		})

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodGet,
			Path:    "/api/presets/trending",
			Handler: routes.BuildTrendingPresetsHandler(app),
		})

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodGet,
			Path:    "/api/presets/search",
			Handler: routes.BuildSearchPresetsHandler(app),
		})

		return nil
	})

	app.OnRecordAfterCreateRequest().Add(events.BuildPresetCreateAddStatsHandler(app))

	app.OnRecordViewRequest().Add(events.BuildPresetViewUpdateStatsHandler(app))

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
