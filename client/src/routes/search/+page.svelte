<script lang="ts">
	import PocketBase from 'pocketbase';
	import PresetRow from '$lib/PresetRow.svelte';
	import { getSettings } from '$lib/settings';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import type { Tag } from '$lib/tags';

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
			console.log(selectedTags);
			tagOptions = (await getTags(q)).filter((tag) => !selectedTags.includes(tag.label));
		} catch (err) {
			console.error(err);
		}
	};

	const addTag = (label: string) => {
		if (!selectedTags.includes(label)) {
			selectedTags.push(label);
			selectedTags = selectedTags;
		}
	};

	const removeTag = (label: string) => {
		if (selectedTags.includes(label)) {
			selectedTags.splice(selectedTags.indexOf(label), 1);
			selectedTags = selectedTags;
		}
	};

	onMount(async () => {
		const client = connect();
		const user = client.authStore.model;
		const settings = await getSettings('id' in user ? user.id : undefined);
		showSpoilers = settings.show_spoilers;
	});
</script>

<h1>Results for "{data.query}":</h1>

<div>
	<h2>Tags</h2>
	{#each selectedTags as tag}
		<div>
			<button on:click={() => removeTag(tag)}>Remove</button>
			<span>{tag}</span>
		</div>
	{/each}
	<input type="text" style="display: block;" on:keyup={(e) => searchTags(e.currentTarget.value)} />
	<select style="display: block;">
		{#each tagOptions as tag}
			<option value={tag.id} on:click={() => addTag(tag.label)}>{tag.label}</option>
		{/each}
	</select>
</div>

<div class="search">
	<div>
		<form
			method="get"
			on:submit={(e) => {
				e.preventDefault();

				const form = new FormData(e.currentTarget);
				const query = form.get('q');

				const params = new URL(location.toString()).searchParams;
				params.set('q', typeof query === 'string' ? query : '');
				params.set('tags', selectedTags.join(','));
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
