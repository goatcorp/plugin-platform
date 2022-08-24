package kamori

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

type plogon struct {
	InternalName string `json:"InternalName"`
	Name         string `json:"Name"`
}

func BuildCheckPlugins(app *pocketbase.PocketBase) func() error {
	return func() error {
		// Get plogons
		res, err := http.Get("https://kamori.goats.dev/Plugin/PluginMaster")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
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
}
