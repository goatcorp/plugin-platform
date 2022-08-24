<script lang="ts">
	import PresetRow from '$lib/components/PresetRow.svelte';
	import { getSettings } from '$lib/settings';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import type { Tag } from '$lib/tags';
	import type { Plugin } from '$lib/plugins';
	import { connectBackend } from '$lib/backend';
	import TagSelector from '$lib/components/TagSelector.svelte';

	export let data: PageData;

	let showSpoilers = false;
	let tagOptions: Tag[] = [];

	let selectedTags: string[] =
		typeof location !== 'undefined'
			? (() => {
					const url = new URL(location.toString());
					const tags = url.searchParams.get('tags');
					if (tags != null && tags.length > 0) {
						return tags.split(',');
					}

					return [];
			  })()
			: [];
	let selectedPlugin: Plugin | null =
		typeof location !== 'undefined'
			? (() => {
					const url = new URL(location.toString());
					const plugin = url.searchParams.get('plugin');
					const found = data.plugins.find((p) => p.internal_name === plugin);
					return found || null;
			  })()
			: null;

	const getTags = async (q: string) => {
		const backend = connectBackend();
		const tags = await backend.app.records.getList('tags', 1, 10, {
			filter: `label~'${q}'`
		});
		return tags.items.map((record) => ({ id: record.id, label: record.label }));
	};

	const searchTags = async (q: string) => {
		try {
			console.log(selectedTags);
			tagOptions = (await getTags(q)).filter((tag) => !selectedTags.includes(tag.label));
		} catch (err) {
			console.error(err);
		}
	};

	const addTag = (tagId: string) => {
		const tag = tagOptions.find((t) => t.id === tagId);
		if (tag == null) {
			return;
		}

		if (!selectedTags.includes(tag.label)) {
			selectedTags.push(tag.label);
			selectedTags = selectedTags;
		}
	};

	const removeTag = (tagLabel: string) => {
		if (selectedTags.includes(tagLabel)) {
			selectedTags.splice(selectedTags.indexOf(tagLabel), 1);
			selectedTags = selectedTags;
		}
	};

	onMount(async () => {
		const backend = connectBackend();
		const user = backend.getCurrentUser();
		const settings = await getSettings(user ? user.id : undefined);
		showSpoilers = settings.show_spoilers;
	});
</script>

<h1>Results for "{data.query}":</h1>

<div>
	<h2>Tags</h2>
	<TagSelector
		tags={selectedTags}
		{tagOptions}
		onSearch={searchTags}
		onAdd={addTag}
		onRemove={removeTag}
	/>

	<h2>Plugin</h2>
	<div>
		<select>
			<option value="" disabled selected={selectedPlugin == null}>--Select a plugin--</option>
			{#each data.plugins as plugin}
				<option
					value={plugin.id}
					on:click={() => (selectedPlugin = plugin)}
					selected={selectedPlugin?.id === plugin.id}>{plugin.name}</option
				>
			{/each}
		</select>
	</div>
</div>

<div class="search">
	<div>
		<form
			method="get"
			on:submit={(e) => {
				e.preventDefault();

				const form = new FormData(e.currentTarget);
				const query = form.get('q');

				const params = new URLSearchParams();
				params.set('q', typeof query === 'string' ? query : '');
				if (selectedTags.length !== 0) params.set('tags', selectedTags.join(','));
				if (selectedPlugin != null) params.set('plugin', selectedPlugin.internal_name);
				location.assign('?' + params.toString());
			}}
		>
			<input type="search" placeholder="Search for presets..." name="q" value={data.query} />
			<button>Search</button>
		</form>
	</div>
</div>

<div class="results">
	{#each data.results as preset}
		<PresetRow
			{preset}
			authorName={data.authors[preset.author]}
			stats={data.stats[preset.id]}
			{showSpoilers}
		/>
	{/each}
</div>

<div>
	{#if data.page > 1}
		<a href={`?page=${data.page - 1}`}>Previous</a>
	{:else}
		<a href={`?page=${data.page - 1}`} disabled tabindex="-1">Previous</a>
	{/if}
	<a href={`?page=${data.page}`} disabled tabindex="-1">{data.page}</a>
	{#if data.page < data.totalPages}
		<a href={`?page=${data.page + 1}`}>Next</a>
	{:else}
		<a href={`?page=${data.page + 1}`} disabled tabindex="-1">Next</a>
	{/if}
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	a[disabled] {
		pointer-events: none;
		text-decoration: none;
		color: inherit;
	}
</style>
