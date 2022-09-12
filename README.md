# plugin-platform
(WIP) Dalamud plugin platform website for sharing presets and possibly also [this stuff](https://github.com/goatcorp/DIPs/issues/34).

## Contributing
Dependencies:
* Node v16.14.0+
* Yarn v1.22.18+
* Go v1.19+

Copy `/client/.env.example` into `/client/.env.local`.

Shell 1:
```sh
cd server
go run ./cmd/platform/main.go migrate up
go run ./cmd/platform/main.go serve
```

Navigate to the admin UI and create an admin account. Create an application in the Discord Developer Portal and copy the application's client ID and client secret into the authentication settings in the admin UI. Make sure to add `http://YOUR_LOCALHOST_URL/oauth2/redirect` to the OAuth2 redirects.
Please use `localhost` for local development - do not use `0.0.0.0` since Discord doesn't support that.

Shell 2:
```sh
cd client
yarn
yarn dev -- --open
```

### OAuth2
Open up the admin UI at `http://127.0.0.1:8090/_/` and go to the **Settings** page. Under **Auth Providers**, select an
option other than **Email/Password** and add your application credentials from the provider. After the provider is enabled
and the changes are saved, the login page should show the provider as a login option.

To add new login options, PR changes to [PocketBase](https://github.com/pocketbase/pocketbase) upstream.

### Adding data
Open up the admin UI and go to the **Collections** page. Select a collection and add a record.
Users can be added manually under the **Users** page, or by going through the standard registration process.

### Frontend development
For most changes, refer to the [SvelteKit documentation](https://kit.svelte.dev/docs/introduction).

### Backend development
For most changes, refer to the [PocketBase documentation](https://pocketbase.io/docs).

### Schema modifications
Open up the admin UI and make your changes, then run:
```sh
cd server/pkg
go run ../cmd/platform/main.go --dir="../pb_data" migrate collections
```

Note that PocketBase uses SQLite (though I suspect this can be customized at the expense of having to configure another database),
so certain features like schema-level permissions are not possible. These should instead be implemented through denormalization and
application-level permissions.

Migrations can be applied by running:
```sh
cd server
go run ./cmd/platform/main.go migrate up
```