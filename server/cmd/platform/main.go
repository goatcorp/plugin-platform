package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/search"

	_ "github.com/karashiiro/plugin-platform/server/pkg/migrations"
)

type preset struct {
	Id        string `json:"id"`
	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Spoiler   bool   `json:"spoiler"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}

type plogon struct {
	InternalName string `json:"InternalName"`
	Name         string `json:"Name"`
}

// https://gist.github.com/heri16/077282d46ae95d48d430a90fb6accdff?permalink_comment_id=4219415#gistcomment-4219415
func interfaceSlice[P any](slice []P) []any {
	ret := make([]any, len(slice))
	for i := 0; i < len(slice); i++ {
		v := slice[i]
		ret[i] = any(v)
	}
	return ret
}

func main() {
	app := pocketbase.New()

	checkPlugins := func() error {
		// Get plogons
		res, err := http.Get("https://kamori.goats.dev/Plugin/PluginMaster")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		var plogons []plogon
		err = json.Unmarshal(data, &plogons)
		if err != nil {
			return err
		}

		// Reformat as a map
		reformatted := map[string]plogon{}
		for _, plogon := range plogons {
			reformatted[plogon.InternalName] = plogon
		}

		// Get known plogons
		result, err := app.Dao().DB().
			Select("internal_name").
			From("plugins").
			Build().
			Rows()
		if err != nil {
			return err
		}

		// Reformat as a map
		reformattedKnown := map[string]bool{}
		for result.Next() {
			var internalName string

			err := result.Scan(&internalName)
			if err != nil {
				return err
			}

			reformattedKnown[internalName] = true
		}

		// Store new plugins
		for _, p := range reformatted {
			if _, ok := reformattedKnown[p.InternalName]; !ok {
				coll, err := app.Dao().FindCollectionByNameOrId("plugins")
				if err != nil {
					return err
				}

				record := models.NewRecord(coll)

				record.SetDataValue("internal_name", p.InternalName)
				record.SetDataValue("name", p.Name)
				err = app.Dao().SaveRecord(record)
				if err != nil {
					return err
				}

				log.Printf("Added %s (%s) to known plugins", p.Name, p.InternalName)
			}
		}

		return nil
	}

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
			Method: http.MethodGet,
			Path:   "/api/presets/popular",
			Handler: func(c echo.Context) error {
				var presets []struct {
					preset
					Views int64 `json:"views"`
				}

				// Return the presets ordered by their views
				query := app.Dao().DB().
					Select("presets.*", "SUM(preset_stats.views) AS views").
					From("presets").
					InnerJoin("preset_stats", dbx.NewExp("presets.id = preset_stats.preset")).
					GroupBy("preset_stats.preset").
					OrderBy("views DESC")
				fieldResolver := search.NewSimpleFieldResolver("*")

				result, err := search.NewProvider(fieldResolver).Query(query).Exec(&presets)
				if err != nil {
					return err
				}

				return c.JSON(200, result)
			},
		})

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/presets/trending",
			Handler: func(c echo.Context) error {
				var presets []*preset

				// Return the presets ordered only by today's views
				query := app.Dao().DB().
					Select("presets.*", "SUM(views)").
					From("presets").
					InnerJoin("(SELECT * FROM preset_stats WHERE date >= date('now', '-1 day')) AS recent", dbx.NewExp("presets.id = recent.preset")).
					GroupBy("recent.preset").
					OrderBy("recent.views DESC")
				fieldResolver := search.NewSimpleFieldResolver("*")

				result, err := search.NewProvider(fieldResolver).Query(query).Exec(&presets)
				if err != nil {
					return err
				}

				return c.JSON(200, result)
			},
		})

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/presets/search",
			Handler: func(c echo.Context) error {
				var presets []*preset

				q := c.QueryParam("q")

				rawPage := c.QueryParam("page")
				page, err := strconv.ParseInt(rawPage, 10, 32)
				if err != nil {
					page = 1
				}

				var tags []string
				rawTags := c.QueryParam("tags")
				if rawTags != "" {
					tags = strings.Split(rawTags, ",")
				}

				plugin := c.QueryParam("plugin")

				var whereExpr dbx.Expression
				whereExpr = dbx.Like("presets.title", q)
				if len(tags) > 0 {
					whereExpr = dbx.And(dbx.In("tags.label", interfaceSlice(tags)...), whereExpr)
				}

				if plugin != "" {
					whereExpr = dbx.And(dbx.NewExp("plugins.internal_name = {:internal_name}", dbx.Params{"internal_name": plugin}), whereExpr)
				}

				query := app.Dao().DB().
					Select("presets.*").
					From("presets").
					LeftJoin("preset_tags", dbx.NewExp("presets.id = preset_tags.preset")).
					LeftJoin("tags", dbx.NewExp("preset_tags.tag = tags.id")).
					LeftJoin("preset_plugins", dbx.NewExp("presets.id = preset_plugins.preset")).
					LeftJoin("plugins", dbx.NewExp("preset_plugins.plugin = plugins.id")).
					Where(whereExpr)
				fieldResolver := search.NewSimpleFieldResolver("*")

				result, err := search.NewProvider(fieldResolver).Query(query).Page(int(page)).Exec(&presets)
				if err != nil {
					return err
				}

				return c.JSON(200, result)
			},
		})

		return nil
	})

	app.OnUserAfterOauth2Register().Add(func(e *core.UserOauth2RegisterEvent) error {
		query := app.Dao().DB().
			Update("profiles", dbx.Params{"name": e.AuthData.Name}, dbx.HashExp{"userid": e.User.Id})
		res, err := query.Execute()
		if err != nil {
			return err
		}

		n, err := res.RowsAffected()
		if err != nil {
			return err
		}

		if n == 0 {
			return errors.New("app: user not found")
		}

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
