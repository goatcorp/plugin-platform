package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

// Auto generated migration with the most recent collections configuration.
func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "systemprofiles0",
				"created": "2022-08-18 22:13:23.202",
				"updated": "2022-08-23 14:07:44.254",
				"name": "profiles",
				"system": true,
				"schema": [
					{
						"system": true,
						"id": "pbfielduser",
						"name": "userId",
						"type": "user",
						"required": true,
						"unique": true,
						"options": {
							"maxSelect": 1,
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "pbfieldname",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "pbfieldavatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif"
							],
							"thumbs": null
						}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "userId = @request.user.id",
				"updateRule": "userId = @request.user.id",
				"deleteRule": null
			},
			{
				"id": "bqZxJ0Bxhp2JrfF",
				"created": "2022-08-19 03:38:22.597",
				"updated": "2022-08-23 14:07:44.257",
				"name": "presets",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "vrnuccti",
						"name": "title",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 3,
							"max": 32,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "fb4jzllu",
						"name": "author",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "systemprofiles0",
							"cascadeDelete": false
						}
					},
					{
						"system": false,
						"id": "bh1seu3l",
						"name": "thumbnail",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif"
							],
							"thumbs": []
						}
					},
					{
						"system": false,
						"id": "s93or6zf",
						"name": "spoiler",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "@request.user.id != \"\"",
				"updateRule": "@request.user.profile.id = author",
				"deleteRule": "@request.user.profile.id = author"
			},
			{
				"id": "glYbMLATGZuGFA9",
				"created": "2022-08-19 17:11:11.206",
				"updated": "2022-08-23 14:07:44.259",
				"name": "preset_data",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "zr1dnhnd",
						"name": "preset",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "bqZxJ0Bxhp2JrfF",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "kjjikcy2",
						"name": "data",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 0,
							"max": 1000,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "2pzyu2fo",
						"name": "title",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 32,
							"pattern": ""
						}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "@request.user.id != \"\"",
				"updateRule": "@request.user.id = @collection.preset_data.preset.author.userId",
				"deleteRule": "@request.user.id = @collection.preset_data.preset.author.userId"
			},
			{
				"id": "49sRfFrkmlK2nMN",
				"created": "2022-08-19 21:48:28.538",
				"updated": "2022-08-23 14:07:44.261",
				"name": "preset_stats",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "rybn4gbr",
						"name": "views",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					},
					{
						"system": false,
						"id": "byiwlgts",
						"name": "preset",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "bqZxJ0Bxhp2JrfF",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "jzdb81kd",
						"name": "date",
						"type": "date",
						"required": true,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null
			},
			{
				"id": "GSqUBLsYVFGAAkg",
				"created": "2022-08-20 14:07:18.235",
				"updated": "2022-08-23 14:07:44.262",
				"name": "profile_preset_favorites",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "syzicbp8",
						"name": "preset",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "bqZxJ0Bxhp2JrfF",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "tgkqt7rc",
						"name": "profile",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "systemprofiles0",
							"cascadeDelete": true
						}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "@request.user.profile.id = profile.id",
				"updateRule": null,
				"deleteRule": "@request.user.profile.id = profile.id"
			},
			{
				"id": "NTu2C51ttK80ICw",
				"created": "2022-08-20 17:08:18.551",
				"updated": "2022-08-23 14:07:44.265",
				"name": "tags",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "6afu0qkp",
						"name": "label",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": 32,
							"pattern": "(\\w|\\d|-)+"
						}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null
			},
			{
				"id": "1LIRgKfmr6KGiH9",
				"created": "2022-08-20 17:12:07.841",
				"updated": "2022-08-24 16:33:44.803",
				"name": "preset_tags",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "5ewr8ogp",
						"name": "preset",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "bqZxJ0Bxhp2JrfF",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "q5gipge2",
						"name": "tag",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "NTu2C51ttK80ICw",
							"cascadeDelete": true
						}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "@request.user.id = preset.author.userId",
				"updateRule": null,
				"deleteRule": "@request.user.id = preset.author.userId"
			},
			{
				"id": "Go3oFqSnlKsvTSB",
				"created": "2022-08-20 19:00:56.707",
				"updated": "2022-08-23 14:07:44.268",
				"name": "user_settings",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "hryivhzo",
						"name": "user",
						"type": "user",
						"required": true,
						"unique": true,
						"options": {
							"maxSelect": 1,
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "txm6cr6p",
						"name": "show_spoilers",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					}
				],
				"listRule": "@request.user.id = user",
				"viewRule": "@request.user.id = user",
				"createRule": "@request.user.id = user",
				"updateRule": "@request.user.id = user",
				"deleteRule": null
			},
			{
				"id": "g2iirc5mkp7omjh",
				"created": "2022-08-23 02:20:02.684",
				"updated": "2022-08-24 16:16:07.680",
				"name": "plugins",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "7ye4yx81",
						"name": "internal_name",
						"type": "text",
						"required": true,
						"unique": true,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "x6erycco",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null
			},
			{
				"id": "bey7jv8p7f4a9a4",
				"created": "2022-08-23 16:40:51.066",
				"updated": "2022-08-24 16:33:35.669",
				"name": "preset_plugins",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "d97b0jmw",
						"name": "preset",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "bqZxJ0Bxhp2JrfF",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "zgqxaqc8",
						"name": "plugin",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "g2iirc5mkp7omjh",
							"cascadeDelete": true
						}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "@request.user.profile.id = preset.author.id",
				"updateRule": null,
				"deleteRule": "@request.user.profile.id = preset.author.id"
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		// no revert since the configuration on the environment, on which
		// the migration was executed, could have changed via the UI/API
		return nil
	})
}
