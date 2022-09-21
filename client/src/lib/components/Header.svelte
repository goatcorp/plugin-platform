<script lang="ts">
	import { connectBackend } from '$lib/backend';
	import type { User } from 'pocketbase';
	import { onMount } from 'svelte';
	import UserMenu from './UserMenu.svelte';
	import { SearchIcon } from 'svelte-feather-icons';

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

<nav class="flex justify-between">
	<a href="/" class="btn btn-primary">Dalamud Plugin Presets</a>
	<div class="flex justify-end">
		<form method="get" action="/search" class="mr-1">
			<label class="input-group">
				<input
					type="search"
					placeholder="Search for presets..."
					name="q"
					class="input input-sm input-bordered input-primary w-full max-w-xs"
				/>
				<button class="btn btn-sm"><SearchIcon /></button>
			</label>
		</form>
		{#if user?.profile == null}
			<a href="/login" class="btn btn-sm">Log in</a>
		{:else}
			<a href="/create" class="btn btn-sm btn-accent mr-1">Create</a>
			<UserMenu id={user.profile.id} />
		{/if}
	</div>
</nav>
