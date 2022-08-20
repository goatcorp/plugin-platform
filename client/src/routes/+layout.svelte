<script lang="ts">
	import 'normalize.css';
	import PocketBase, { type User } from 'pocketbase';
	import { onDestroy, onMount } from 'svelte';

	import { id } from '$lib/session';

	let user: User | null = null;
	let invalidate: (() => void) | null = null;

	const connect = () => {
		return new PocketBase('http://127.0.0.1:8090');
	};

	const loadCurrentUser = async () => {
		const client = connect();
		if (!client.authStore.isValid || $id == null) {
			user = null;
			return;
		}

		user = await client.users.getOne($id);
	};

	onMount(async () => {
		invalidate = id.subscribe(async () => {
			await loadCurrentUser();
		});

		await loadCurrentUser();
	});

	onDestroy(() => {
		if (invalidate != null) {
			invalidate();
		}
	});
</script>

<header>
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
				<a href={`/user/${user.profile?.id}`}>{user.profile?.name}</a>
			{/if}
		</div>
	</nav>
</header>

<main><slot /></main>

<style lang="scss">
	header {
		margin: auto;
		padding: 20px;
		width: 70%;

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

	main {
		width: 70%;
		margin: auto;
	}

	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
