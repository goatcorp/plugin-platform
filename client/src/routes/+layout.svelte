<script lang="ts">
	import 'normalize.css';
	import { connectBackend } from '$lib/backend';
	import type { User } from 'pocketbase';
	import { onMount } from 'svelte';

	let user: User | null = null;

	const loadCurrentUser = async () => {
		const backend = connectBackend();
		const model = backend.getCurrentUser();
		if (model && 'id' in model) {
			// The stored user data doesn't include the profile data
			user = await backend.app.users.getOne(model.id);
			return;
		}
	};

	onMount(async () => {
		await loadCurrentUser();
	});
</script>

<div class="layout">
	<header class="header">
		<nav>
			<div><a href="/">Dalamud Plugin Presets</a></div>
			<div class="controls">
				<div class="search">
					<div>
						<form method="get" action="/search">
							<input type="search" placeholder="Search for presets..." name="q" />
							<button>Search</button>
						</form>
					</div>
				</div>
				{#if user == null}
					<a href="/login">Log in</a>
				{:else}
					<a href="/logout">Log out</a>
					<a href="/create">Create</a>
					<a href="/settings">Settings</a>
					<a href={`/user/${user.profile?.id}`}>{user.profile?.name}</a>
				{/if}
			</div>
		</nav>
	</header>

	<main class="content"><slot /></main>

	<nav class="sidebar">
		<h1>Right text</h1>
	</nav>

	<footer class="footer">
		<p>Bottom text</p>
	</footer>
</div>

<style lang="scss">
	.layout {
		display: grid;
		grid-template: auto / calc(100% - 400px) 400px;

		.header {
			grid-column: 1 / 3;
			padding: 35px;

			nav {
				display: flex;
				justify-content: space-between;

				.controls {
					display: flex;
					justify-content: right;

					> * {
						margin-left: 8px;
					}
				}

				.search {
					display: flex;
				}
			}
		}

		.content {
			grid-column: 1;
			padding-right: 35px;
			padding-left: 35px;
		}

		.sidebar {
			grid-column: 2;
			padding-right: 35px;
			padding-left: 35px;
		}

		.footer {
			grid-column: 1 / 3;
			padding: 35px;
		}
	}

	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
