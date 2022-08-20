<script lang="ts">
	import PocketBase, { type User } from 'pocketbase';
	import { onDestroy, onMount } from 'svelte';
	import type { PageData } from './$types';
	import { id } from '$lib/session';

	export let data: PageData;

	let user: User | null = null;
	let invalidate: (() => void) | null = null;
	let isFavorite = data.isFavoriteInitial;

	const connect = () => {
		return new PocketBase('http://127.0.0.1:8090');
	};

	const toggleFavorite = async () => {
		const client = connect();

		// This should be done on the server later, it's just simpler to do this here for now
		try {
			const favorites = await client.records.getList('user_preset_favorites', 1, 1, {
				filter: `preset='${data.preset.id}'`
			});

			if (favorites.totalItems === 0) {
				try {
					await client.records.create('user_preset_favorites', {
						user: $id,
						preset: data.preset.id
					});
					isFavorite = true;
				} catch (err) {
					console.error(err);
				}
			} else {
				try {
					await client.records.delete('user_preset_favorites', favorites.items[0].id);
					isFavorite = false;
				} catch (err) {
					console.error(err);
				}
			}
		} catch (err) {
			console.error(err);
		}
	};

	// This should be a store
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

<h1>{data.preset.title}</h1>
<div class="info">
	<span>By <a href={`/user/${data.preset.author}`}>{data.authorName}</a></span>
	<span>Created {data.preset.created.toLocaleDateString()}</span>
	<span>Last updated {data.preset.updated.toLocaleDateString()}</span>
	<span>{data.presetStats.views} views</span>
	<button on:click={toggleFavorite}
		>{isFavorite ? 'Remove from favorites' : 'Add to favorites'}</button
	>
	{#if user?.profile?.id === data.preset.author}
		<a href={`/preset/${data.preset.id}/edit`}>Edit preset</a>
	{/if}
</div>

{#each data.presetData as presetDataEntry, i}
	<h2>{presetDataEntry.title || `File ${i + 1}`}</h2>
	<div class="preset-data">
		<code>{@html presetDataEntry.data.replace(/[\r\n]+/g, '<br />')}</code>
	</div>
{/each}

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	code {
		font-family: 'Courier New', Courier, monospace;
	}

	.info > * {
		display: block;
	}

	.preset-data {
		padding: 10px;
		background-color: #eee;
		overflow-wrap: break-word;
	}
</style>
