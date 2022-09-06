# API Specification
This is a loose specification of the API format. Database collections and collection permissions are described precisely by [this schema snapshot](https://github.com/goatcorp/plugin-platform/blob/9cc56686d4d6411625ce48986d96e308500c65cc/server/pkg/migrations/1661381271_collections_snapshot.go).

Many REST operations require header-based user authentication, as described [here](https://pocketbase.io/docs/api-authentication/). In general, only data modification requires authentication - exceptions include user settings and other private data.

## Plugins
The list of supported plugins is pulled from the master list, available [here](https://kamori.goats.dev/Plugin/PluginMaster). This endpoint is polled every few minutes to check for updates.

## Route map
```
/api
|-/collections
  |-/presets/records { GET, POST }
    |-/:id { GET, POST, PATCH, DELETE }
  |-/preset_data/records { GET, POST }
    |-/:id { GET, POST, PATCH, DELETE }
  |-/preset_stats/records { GET }
    |-/:id { GET, POST, PATCH, DELETE }
  |-/profile_preset_favorites/records { GET, POST }
    |-/:id { GET, POST, PATCH, DELETE }
  |-/tags/records { GET, POST }
    |-/:id { GET, POST, PATCH, DELETE }
  |-/preset_tags/records { GET, POST }
    |-/:id { GET, POST, PATCH, DELETE }
  |-/user_settings/records { GET, POST }
    |-/:id { GET, POST, PATCH, DELETE }
  |-/plugins/records { GET }
    |-/:id { GET, POST, PATCH, DELETE }
  |-/preset_plugins/records { GET, POST }
    |-/:id { GET, POST, PATCH, DELETE }
|-/presets
  |-/popular { GET }
  |-/trending { GET }
  |-/search { GET }
|-/users
  |-/auth-methods { GET }
  |-/auth-via-oauth2 { POST }
  |-/refresh { POST }
  |-/:id { GET, DELETE }
    |-/external-auths { GET }
```

## `/api/collections`
CRUD endpoints. Each of these is backed by a table that supports basic CRUD operations. They are documented in more detail [here](https://pocketbase.io/docs/api-collections/) (collections) and [here](https://pocketbase.io/docs/api-records/) (records).

### `POST /api/collections/presets/records`
This endpoint additionally adds a new record to the `preset_stats` collection for the current day.

### `GET /api/collections/presets/records/:id`
This endpoint additionally increments the view counter for the current day in the `preset_stats` collection, creating it if it does not yet exist.

## `/api/presets`
A set of custom aggregation endpoints for finding presets.

### `/api/presets/popular`
Returns the list of all presets, ordered in descending order by its number of views. See [**List collections**](https://pocketbase.io/docs/api-collections/) for response format information. Unlike that documentation page notes, this endpoint does not support any query parameters.

### `/api/presets/trending`
Returns the list of all presets, ordered in descending order by its number of views over the previous day. See [**List collections**](https://pocketbase.io/docs/api-collections/) for response format information. Unlike that documentation page notes, this endpoint does not support any query parameters.

### `/api/presets/search`
Returns the list of all presets returned by a query. See [**List collections**](https://pocketbase.io/docs/api-collections/) for response format information.

#### Query parameters
| Param    | Type     | Description                                                           |
| -------- | -------- | --------------------------------------------------------------------- |
| `q`      | `String` | (Optional) The search term.                                           |
| `page`   | `Number` | (Optional) The page number to view.                                   |
| `tags`   | `String` | (Optional) A comma-separated list of tag names to filter for.         |
| `plugin` | `String` | (Optional) The name of the plugin associated with the result presets. |

## `/api/users`
User endpoints for authentication and management. Complete documentation is provided [here](https://pocketbase.io/docs/api-users/), but only the listed endpoints above are known to be used on the website (these are all only used internally by the PocketBase client library).