<script lang="ts">
	import type { Preset, PresetStats } from '../preset';

	export let preset: Preset;
	export let stats: PresetStats;
	export let authorName: string;
	export let showSpoilers: boolean;

	const url = `/preset/${preset.id}`;

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

<div class="preset-card">
	<div class="preview">
		{#if !showSpoilers && spoilerEnabled}
			<div class="spoiler-overlay" on:click={(e) => dismissSpoiler(e.currentTarget)} />
		{/if}
		<a href={url}>
			<div class="thumbnail" />
		</a>
		{#if preset.spoiler}
			<div class="spoiler-notice-wrapper">
				<span class="spoiler-notice">Spoiler</span>
			</div>
		{/if}
		<div class="views-wrapper">
			<span class="views">{stats.views} views</span>
		</div>
	</div>
	<div class="info">
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
		<p class="author">By <a href={`/user/${preset.author}`}>{authorName}</a></p>
	</div>
</div>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	a[disabled] {
		pointer-events: none;
	}

	.preset-card {
		height: 200px;
		width: 300px;

		border: 1px solid #ccc;
		border-radius: 10px 10px 15px 15px; // Hides the top corners behind the thumbnail

		.info {
			padding-left: 10px;
			padding-right: 10px;

			.spoiler-overlay-text {
				filter: blur(4px);
				cursor: pointer;

				animation-duration: 100ms;
			}
		}

		.preview {
			position: relative;
			height: 60%;
			width: 100%;

			background-color: blue;
			border-radius: 10px 10px 0 0;

			.thumbnail {
				height: 100%;
			}

			.spoiler-overlay {
				position: absolute;
				width: 100%;
				height: 100%;
				filter: blur(1.5rem);
				cursor: pointer;

				animation-duration: 100ms;
			}
		}

		.views-wrapper {
			position: absolute;
			bottom: 10px;
			right: 10px;
			padding: 4px;

			cursor: default;

			background-color: rgba(0.7, 0.7, 0.7, 0.6);
			border-radius: 15px;
		}

		.views {
			font-size: smaller;
			color: white;
		}

		.spoiler-notice-wrapper {
			position: absolute;
			bottom: 42px;
			right: 10px;
			padding: 4px;

			cursor: default;

			background-color: rgba(0.7, 0.7, 0.7, 0.6);
			border-radius: 15px;
		}

		.spoiler-notice {
			font-size: smaller;
			color: white;
		}
	}
</style>
