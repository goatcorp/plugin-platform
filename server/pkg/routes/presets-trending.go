package routes

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/tools/search"

	"github.com/karashiiro/plugin-platform/server/pkg/entity"
)

func BuildTrendingPresetsHandler(app *pocketbase.PocketBase) func(echo.Context) error {
	return func(c echo.Context) error {
		var presets []*entity.Preset

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
	}
}
