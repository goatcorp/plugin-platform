package events

import (
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func BuildPresetCreateAddStatsHandler(app *pocketbase.PocketBase) func(*core.RecordCreateEvent) error {
	return func(e *core.RecordCreateEvent) error {
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
	}
}
