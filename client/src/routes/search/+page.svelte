<script lang="ts">
	import PresetRow from '$lib/components/PresetRow.svelte';
	import { getSettings } from '$lib/settings';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import type { Tag } from '$lib/tags';
	import type { Plugin } from '$lib/plugins';
	import { connectBackend } from '$lib/backend';
	import TagSelector from '$lib/components/TagSelector.svelte';
	import PageSelector from '$lib/components/PageSelector.svelte';
	import PluginSelector from '$lib/components/PluginSelector.svelte';

	export let data: PageData;

	let showSpoilers = false;
	let tagOptions: Tag[] = [];

	let selectedTags: string[] = (() => {
		const tags = data.params.tags;
		return tags != null && tags.length > 0 ? tags.split(',') : [];
	})();
	let selectedPlugin: Plugin | null | undefined = (() => {
		const plugin = data.params.plugin;
		const found = data.plugins.find((p) => p.internal_name === plugin);
		return found || null;
	})();

	const searchTags = async (q: string) => {
		const backend = connectBackend();
		try {
			tagOptions = (await backend.searchTags(q)).filter((tag) => !selectedTags.includes(tag.label));
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

<h1>Results for "{data.params.q}":</h1>

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
	<PluginSelector
		plugins={data.plugins}
		{selectedPlugin}
		onSet={(plugin) => (selectedPlugin = plugin)}
	/>
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
			<input type="search" placeholder="Search for presets..." name="q" value={data.params.q} />
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

<PageSelector page={data.page} totalPages={data.totalPages} />
