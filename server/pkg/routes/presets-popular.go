package routes

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/tools/search"

	"github.com/karashiiro/plugin-platform/server/pkg/entity"
)

func BuildPopularPresetsHandler(app *pocketbase.PocketBase) func(echo.Context) error {
	return func(c echo.Context) error {
		var presets []struct {
			entity.Preset
			Views int64 `json:"views"`
		}

		// Return the presets ordered by their views
		query := app.Dao().DB().
			Select("presets.*", "SUM(preset_stats.views) AS views").
			From("presets").
			InnerJoin("preset_stats", dbx.NewExp("presets.id = preset_stats.preset")).
			GroupBy("preset_stats.preset").
			OrderBy("views DESC")

		// See BuildTrendingPresetsHandler for an explanation of this hack
		queryWrap := app.Dao().DB().
			Select("*").
			From(fmt.Sprintf("(%s)", query.Build().SQL()))
		fieldResolver := search.NewSimpleFieldResolver("*")

		result, err := search.NewProvider(fieldResolver).Query(queryWrap).Exec(&presets)
		if err != nil {
			return err
		}

		return c.JSON(200, result)
	}
}
