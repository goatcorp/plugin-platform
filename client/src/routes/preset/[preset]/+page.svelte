<script lang="ts">
	import PocketBase, { type User } from 'pocketbase';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import type { Tag } from '$lib/tags';

	export let data: PageData;

	let user: User | null = null;
	let isFavorite = data.isFavoriteInitial;
	let tagOptions: Tag[] = [];
	let tags: Tag[] = data.presetTags;

	const connect = () => {
		return new PocketBase('http://127.0.0.1:8090');
	};

	const getTags = async (q: string) => {
		const client = connect();
		const tags = await client.records.getList('tags', 1, 10, {
			filter: `label~'${q}'`
		});
		return tags.items.map((record) => ({ id: record.id, label: record.label }));
	};

	const searchTags = async (q: string) => {
		try {
			tagOptions = (await getTags(q)).filter(
				(tag) => !data.presetTags.map((t) => t.id).includes(tag.id)
			);
		} catch (err) {
			console.error(err);
		}
	};

	const addTag = async (tagId: string) => {
		const tag = tagOptions.find((t) => t.id === tagId);
		if (tag == null) {
			return;
		}

		const client = connect();
		try {
			await client.records.create('preset_tags', {
				preset: data.preset.id,
				tag: tag.id
			});
			tags.push(tag);
			tags = tags; // Force reactivity

			const optionIndex = tagOptions.findIndex((tag) => tag.id === tagId);
			if (optionIndex !== -1) {
				tagOptions.splice(optionIndex, 1);
				tagOptions = tagOptions; // Force reactivity
			}
		} catch (err) {
			console.error(err);
		}
	};

	const removeTag = async (tagId: string) => {
		const tag = tags.find((t) => t.id === tagId);
		if (tag == null) {
			return;
		}

		const client = connect();
		try {
			const relation = await client.records.getList('preset_tags', 1, 1, {
				filter: `preset='${data.preset.id}' && tag='${tagId}'`
			});
			if (relation.items.length === 0) {
				console.error('No such tag found.');
				return;
			}

			await client.records.delete('preset_tags', relation.items[0].id);
			tags.splice(
				tags.findIndex((tag) => tag.id === tagId),
				1
			);
			tags = tags; // Force reactivity
		} catch (err) {
			console.error(err);
		}
	};

	const toggleFavorite = async () => {
		if (user == null) {
			return;
		}

		const client = connect();

		// This should be done on the server later, it's just simpler to do this here for now
		try {
			const favorites = await client.records.getList('profile_preset_favorites', 1, 1, {
				filter: `profile='${user.profile?.id}'`
			});

			if (favorites.totalItems === 0) {
				try {
					await client.records.create('profile_preset_favorites', {
						profile: user.profile?.id,
						preset: data.preset.id
					});
					isFavorite = true;
				} catch (err) {
					console.error(err);
				}
			} else {
				try {
					await client.records.delete('profile_preset_favorites', favorites.items[0].id);
					isFavorite = false;
				} catch (err) {
					console.error(err);
				}
			}
		} catch (err) {
			console.error(err);
		}
	};

	const loadCurrentUser = async () => {
		const client = connect();
		const model = client.authStore.model;
		if ('id' in model) {
			user = await client.users.getOne(model.id);
			return;
		}
	};

	onMount(async () => {
		await loadCurrentUser();
	});
</script>

<h1>{data.preset.title}</h1>
<div class="info">
	<span>By <a href={`/user/${data.preset.author}`}>{data.authorName}</a></span>
	<span>Created {data.preset.created.toLocaleDateString()}</span>
	<span>Last updated {data.preset.updated.toLocaleDateString()}</span>
	<span>{data.presetStats.views} views</span>
	{#if isFavorite != null}
		<button on:click={toggleFavorite}
			>{isFavorite ? 'Remove from favorites' : 'Add to favorites'}</button
		>
	{/if}

	{#if user?.profile?.id === data.preset.author}
		<a href={`/preset/${data.preset.id}/edit`}>Edit preset</a>
		<div>
			<input
				type="text"
				style="display: block;"
				on:keyup={(e) => searchTags(e.currentTarget.value)}
			/>
			<select style="display: block;">
				{#each tagOptions as tag}
					<option value={tag.id} on:click={() => addTag(tag.id)}>{tag.label}</option>
				{/each}
			</select>
		</div>
	{/if}

	<div>
		<h2>Tags</h2>
		{#each tags as tag}
			<div style="display: flex;">
				{#if user?.profile?.id === data.preset.author}
					<button on:click={() => removeTag(tag.id)}>Remove</button>
				{/if}
				<div class="tag">{tag.label}</div>
			</div>
		{/each}
		{#if tags.length === 0}
			<span>No tags.</span>
		{/if}
	</div>
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
