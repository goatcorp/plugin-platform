package events

import (
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func BuildPresetViewUpdateStatsHandler(app *pocketbase.PocketBase) func(*core.RecordViewEvent) error {
	return func(e *core.RecordViewEvent) error {
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
	}
}
