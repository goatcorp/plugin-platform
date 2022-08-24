<script lang="ts">
	import type { User } from 'pocketbase';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import type { Tag } from '$lib/tags';
	import type { Plugin } from '$lib/plugins';
	import { connectBackend } from '$lib/backend';
	import TagSelector from '$lib/components/TagSelector.svelte';
	import PluginSelector from '$lib/components/PluginSelector.svelte';

	export let data: PageData;

	let user: User | null = null;
	let isFavorite = data.isFavoriteInitial;
	let tagOptions: Tag[] = [];
	let tags: Tag[] = data.presetTags;
	let plugins = data.plugins;
	let presetPlugin: Plugin | null | undefined = data.presetPlugin;

	const searchTags = async (q: string) => {
		const backend = connectBackend();
		try {
			tagOptions = (await backend.searchTags(q)).filter(
				(tag) => !data.presetTags.map((t) => t.id).includes(tag.id)
			);
		} catch (err) {
			console.error(err);
		}
	};

	const unsetPlugin = async (pluginId: string) => {
		const plugin = plugins.find((p) => p.id === pluginId);
		if (plugin == null) {
			return;
		}

		const backend = connectBackend();
		try {
			const relation = await backend.app.records.getList('preset_plugins', 1, 1, {
				filter: `preset='${data.preset.id}' && plugin='${pluginId}'`
			});
			if (relation.items.length === 0) {
				console.error('No such plugin found.');
				return;
			}

			await backend.app.records.delete('preset_plugins', relation.items[0].id);
		} catch (err) {
			console.error(err);
		}
	};

	const setPlugin = async (plugin: Plugin | null | undefined) => {
		if (presetPlugin != null) {
			await unsetPlugin(presetPlugin.id);
		}

		const backend = connectBackend();
		try {
			if (plugin != null) {
				await backend.app.records.create('preset_plugins', {
					preset: data.preset.id,
					plugin: plugin.id
				});
			} else {
				const relations = await backend.app.records.getFullList('preset_plugins', undefined, {
					filter: `preset='${data.preset.id}'`
				});
				for (const relation of relations) {
					await backend.app.records.delete('preset_plugins', relation.id);
				}
			}

			presetPlugin = plugin;
		} catch (err) {
			console.error(err);
		}
	};

	const addTagById = async (tagId: string) => {
		const tag = tagOptions.find((t) => t.id === tagId);
		if (tag == null) {
			return;
		}

		const backend = connectBackend();
		try {
			await backend.addPresetTagById(data.preset.id, tag.id);
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

	const removeTagByLabel = async (tagLabel: string) => {
		const tag = tags.find((t) => t.label === tagLabel);
		if (tag == null) {
			return;
		}

		const backend = connectBackend();
		try {
			await backend.removePresetTagByLabel(data.preset.id, tagLabel);
			tags.splice(
				tags.findIndex((t) => t.id === tag.id),
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

		const backend = connectBackend();

		// This should be done on the server later, it's just simpler to do this here for now
		try {
			const favorites = await backend.app.records.getList('profile_preset_favorites', 1, 1, {
				filter: `profile='${user.profile?.id}'`
			});

			if (favorites.totalItems === 0) {
				try {
					await backend.app.records.create('profile_preset_favorites', {
						profile: user.profile?.id,
						preset: data.preset.id
					});
					isFavorite = true;
				} catch (err) {
					console.error(err);
				}
			} else {
				try {
					await backend.app.records.delete('profile_preset_favorites', favorites.items[0].id);
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
		const backend = connectBackend();
		if (backend.isAuthenticated()) {
			user = backend.getCurrentUser();
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
	{/if}

	<div>
		<h2>Tags</h2>
		<TagSelector
			readOnly={user?.profile?.id !== data.preset.author}
			tags={tags.map((tag) => tag.label)}
			{tagOptions}
			onSearch={searchTags}
			onAdd={addTagById}
			onRemove={removeTagByLabel}
		/>

		<h2>Plugin</h2>
		<PluginSelector
			{plugins}
			selectedPlugin={presetPlugin}
			readOnly={user?.profile?.id !== data.preset.author}
			onSet={setPlugin}
		/>
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
