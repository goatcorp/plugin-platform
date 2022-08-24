package events

import (
	"errors"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func BuildOauth2AddUsernameHandler(app *pocketbase.PocketBase) func(*core.UserOauth2RegisterEvent) error {
	return func(e *core.UserOauth2RegisterEvent) error {
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
	}
}
