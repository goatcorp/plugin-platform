package routes

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/tools/search"

	"github.com/karashiiro/plugin-platform/server/pkg/entity"
)

// https://gist.github.com/heri16/077282d46ae95d48d430a90fb6accdff?permalink_comment_id=4219415#gistcomment-4219415
func interfaceSlice[P any](slice []P) []any {
	ret := make([]any, len(slice))
	for i := 0; i < len(slice); i++ {
		v := slice[i]
		ret[i] = any(v)
	}
	return ret
}

func BuildSearchPresetsHandler(app *pocketbase.PocketBase) func(echo.Context) error {
	return func(c echo.Context) error {
		var presets []*entity.Preset

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
	}
}
