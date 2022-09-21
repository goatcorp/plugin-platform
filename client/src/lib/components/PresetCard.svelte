<script lang="ts">
	import type { Preset, PresetStats } from '../preset';

	export let preset: Preset;
	export let stats: PresetStats;
	export let authorName: string;
	export let showSpoilers: boolean;

	const url = `/preset/${preset.id}`;
	const thumbnail =
		preset.thumbnail != ''
			? `${
					typeof window !== 'undefined'
						? import.meta.env.VITE_PUBLIC_API_ADDRESS
						: import.meta.env.VITE_API_ADDRESS
			  }/api/files/presets/${preset.id}/${preset.thumbnail}`
			: '';

	let spoilerEnabled = preset.spoiler;
</script>

<div class="card w-72 bg-base-100 shadow-xl">
	<a href={url}>
		<div class="relative">
			{#if preset.spoiler}
				<div class="absolute badge bottom-8 right-2">Spoiler</div>
			{/if}
			<div class="absolute badge bottom-1 right-2">{stats.views} views</div>
			<figure>
				{#if thumbnail}
					{#if !showSpoilers && spoilerEnabled}
						<div class="h-32 w-full bg-secondary" />
					{:else}
						<img src={thumbnail} alt="Preset thumbnail" class="h-32 w-full" />
					{/if}
				{:else}
					<div class="h-32 w-full bg-primary" />
				{/if}
			</figure>
		</div>
	</a>
	<div class="card-body">
		<h2 class="card-title">
			<a href={url}>{preset.title}</a>
		</h2>
		{#if authorName}
			<p>By <a href={`/user/${preset.author}`} class="link">{authorName}</a></p>
		{/if}
	</div>
</div>
