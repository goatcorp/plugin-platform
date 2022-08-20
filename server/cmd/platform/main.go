package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
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
					Spoiler   bool   `json:"spoiler"`
					Created   string `json:"created"`
					Updated   string `json:"updated"`
					Views     int64  `json:"views"`
				}

				err := app.Dao().DB().
					NewQuery("SELECT presets.*, SUM(preset_stats.views) AS views FROM presets JOIN preset_stats ON presets.id = preset_stats.preset GROUP BY preset_stats.preset ORDER BY views DESC LIMIT 20").
					All(&results)
				if err != nil {
					return err
				}

				return c.JSON(200, results)
			},
		})

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/presets/trending",
			Handler: func(c echo.Context) error {
				var results []struct {
					Id        string `json:"id"`
					Thumbnail string `json:"thumbnail"`
					Title     string `json:"title"`
					Author    string `json:"author"`
					Spoiler   bool   `json:"spoiler"`
					Created   string `json:"created"`
					Updated   string `json:"updated"`
				}

				err := app.Dao().DB().
					NewQuery("SELECT presets.*, MAX(preset_stats.date) FROM presets JOIN preset_stats ON presets.id = preset_stats.preset GROUP BY preset_stats.preset ORDER BY preset_stats.views DESC LIMIT 20").
					All(&results)
				if err != nil {
					return err
				}

				return c.JSON(200, results)
			},
		})

		return nil
	})

	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		if e.Record.Collection().Name == "presets" {
			// Get the preset stats collection
			coll, err := app.Dao().FindCollectionByNameOrId("preset_stats")
			if err != nil {
				return err
			}

			presetId := e.Record.Id

			// Create a stats record for the new preset
			record := models.NewRecord(coll)
			record.SetDataValue("views", 0)
			record.SetDataValue("date", time.Now().UTC())
			record.SetDataValue("preset", presetId)

			err = app.Dao().SaveRecord(record)
			if err != nil {
				return err
			}
		}

		return nil
	})

	app.OnRecordViewRequest().Add(func(e *core.RecordViewEvent) error {
		if e.Record.Collection().Name == "presets" {
			// Get the preset stats collection
			coll, err := app.Dao().FindCollectionByNameOrId("preset_stats")
			if err != nil {
				return err
			}

			presetId := e.Record.Id

			// Get all of the records from today
			y, m, d := time.Now().UTC().Date()
			loc := time.Now().UTC().Location()
			from := time.Date(y, m, d, 0, 0, 0, 0, loc)
			to := time.Date(y, m, d+1, 0, 0, 0, 0, loc)
			records, err := app.Dao().
				FindRecordsByExpr(coll, dbx.And(dbx.Between("date", from, to), dbx.HashExp{"preset": presetId}))
			if err != nil {
				return err
			}

			if len(records) == 0 {
				// Create a new stats record for today (this may create a duplicate record due to race conditions; that's fine)
				next := models.NewRecord(coll)
				next.SetDataValue("views", 1)
				next.SetDataValue("preset", presetId)
				next.SetDataValue("date", time.Now().UTC())

				err := app.Dao().SaveRecord(next)
				if err != nil {
					return err
				}
			} else {
				// Update the record (this is not atomic)
				stats := records[0]
				_, err = app.Dao().DB().
					Update("preset_stats", dbx.Params{"views": stats.GetIntDataValue("views") + 1}, dbx.HashExp{"id": stats.Id}).
					Execute()
				if err != nil {
					return err
				}
			}

			return nil
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
