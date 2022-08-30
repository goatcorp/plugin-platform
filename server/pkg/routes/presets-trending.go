package routes

import (
	"fmt"

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

		// The search provider replaces the SELECT field with a COUNT(*) when using pagination to determine
		// the total number of records in the result set. However, when using GROUP BY, this actually counts
		// the number of subrows in each group, returning each count as a separate row. When there are no
		// groups, no rows are returned, which causes the query to fail. To work around this, we need to use
		// a wrapper query that maintains the aggregation of the inner query.
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
