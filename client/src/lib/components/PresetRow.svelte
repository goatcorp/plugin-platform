<script lang="ts">
	import type { Preset, PresetStats } from '../preset';

	export let preset: Preset;
	export let stats: PresetStats;
	export let authorName: string;
	export let showSpoilers: boolean;

	const url = `/preset/${preset.id}`;
	const thumbnail =
		preset.thumbnail != ''
			? `http://127.0.0.1:8090/api/files/presets/${preset.id}/${preset.thumbnail}`
			: '';

	let spoilerEnabled = preset.spoiler;
</script>

<div class="card card-side bg-base-100 shadow-xl mb-2">
	<figure>
		<a href={url}>
			{#if thumbnail}
				{#if !showSpoilers && spoilerEnabled}
					<div class="h-full w-36 bg-secondary" />
				{:else}
					<img src={thumbnail} alt="Preset thumbnail" class="h-full w-36" />
				{/if}
			{:else}
				<div class="h-full w-36 bg-primary" />
			{/if}
		</a>
	</figure>
	<div class="card-body">
		<div class="flex justify-between">
			<div>
				<a href={url}><h2 class="card-title">{preset.title}</h2></a>
				{#if preset.spoiler}
					<p>Spoiler</p>
				{/if}
				{#if authorName}
					<p>By <a href={`/user/${preset.author}`} class="link">{authorName}</a></p>
				{/if}
			</div>
			<div class="text-right">
				<p>Last updated: {preset.updated.toLocaleDateString()}</p>
				<p>{stats.views} views</p>
			</div>
		</div>
	</div>
</div>
