package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/presets/popular",
			Handler: func(c echo.Context) error {
				var results []struct {
					Id        string `json:"id"`
					Thumbnail string `json:"thumbnail"`
					Title     string `json:"title"`
					Author    string `json:"author"`
					Created   string `json:"created"`
					Updated   string `json:"updated"`
					Views     int64  `json:"views"`
				}

				err := app.Dao().DB().
					NewQuery("SELECT presets.*, preset_stats.views FROM presets JOIN preset_stats ON presets.id = preset_stats.preset LIMIT 20").
					All(&results)
				if err != nil {
					return err
				}

				return c.JSON(200, results)
			},
		})

		return nil
	})

	app.OnRecordViewRequest().Add(func(e *core.RecordViewEvent) error {
		if e.Record.Collection().Name == "presets" {
			coll, err := app.Dao().FindCollectionByNameOrId("preset_stats")
			if err != nil {
				return err
			}

			presetId := e.Record.Id
			log.Println(presetId)
			stats, err := app.Dao().FindFirstRecordByData(coll, "preset", presetId)
			if err != nil {
				return err
			}

			_, err = app.Dao().DB().
				Update("preset_stats", dbx.Params{"views": stats.GetIntDataValue("views") + 1}, dbx.HashExp{"id": stats.Id}).
				Execute()
			if err != nil {
				return err
			}

			return nil
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
