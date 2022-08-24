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

	const dismissSpoiler = async (target: Element) => {
		target.setAttribute('style', 'blur: none;');
		await new Promise((resolve) => setTimeout(resolve, 100));
		spoilerEnabled = false;
	};

	const dismissSpoilerText = async (target: Element) => {
		target.setAttribute('style', 'blur: none;');
		await new Promise((resolve) => setTimeout(resolve, 100));
		spoilerEnabled = false;
	};
</script>

<div class="result-entry">
	<div class="thumbnail-wrapper" style="position: relative;">
		{#if !showSpoilers && spoilerEnabled}
			<div
				class="spoiler-overlay"
				style={!showSpoilers && spoilerEnabled && thumbnail !== ''
					? 'background-color: #ccc;'
					: undefined}
				on:click={(e) => dismissSpoiler(e.currentTarget)}
			>
				<div
					class="thumbnail"
					style={thumbnail ? `background-image: url(${thumbnail});` : undefined}
				/>
			</div>
		{/if}
		<a href={url}>
			<div
				class={!showSpoilers && spoilerEnabled ? '' : 'thumbnail'}
				style={thumbnail ? `background-image: url(${thumbnail});` : undefined}
			/>
		</a>
	</div>
	<div class="info">
		<div class="basic">
			<div style="position: relative;">
				{#if !showSpoilers && spoilerEnabled}
					<span
						class="title spoiler-overlay-text"
						on:click={(e) => dismissSpoilerText(e.currentTarget)}
					>
						<a href={url} disabled>{preset.title}</a>
					</span>
				{:else}
					<span class="title"><a href={url}>{preset.title}</a></span>
				{/if}
			</div>
			{#if preset.spoiler}
				<small>Spoiler</small>
			{/if}
			{#if authorName}
				<span>By <a href={`/user/${preset.author}`}>{authorName}</a></span>
			{/if}
		</div>
		<div class="extra">
			<small>Last updated: {preset.updated.toLocaleDateString()}</small>
			<small>{stats.views} views</small>
		</div>
	</div>
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	small {
		color: #444;
	}

	a[disabled] {
		pointer-events: none;
	}

	.result-entry {
		display: flex;

		margin-bottom: 15px;

		border: 1px solid #ccc;
		border-radius: 15px 10px 10px 15px; // Hides the left corners behind the thumbnail

		.thumbnail-wrapper {
			height: 100px;
			width: 150px;
		}

		.thumbnail {
			height: 100px;
			width: 150px;

			background-color: blue;
			background-size: 150px;

			border-radius: 10px 0 0 10px;
		}

		.spoiler-overlay {
			position: absolute;
			width: 100%;
			height: 100%;
			filter: blur(5px);
			cursor: pointer;

			animation-duration: 100ms;
		}

		.spoiler-overlay-text {
			filter: blur(4px);
			cursor: pointer;

			animation-duration: 100ms;
		}

		.info {
			display: flex;
			justify-content: space-between;
			width: 100%;

			padding: 10px;

			.basic > * {
				display: block;
			}

			.extra > * {
				display: block;
				text-align: right;
			}

			.title {
				margin-bottom: 10px;
			}
		}
	}
</style>
