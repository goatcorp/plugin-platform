<script lang="ts">
	import PocketBase from 'pocketbase';
	import { onMount } from 'svelte';
	import type { Preset } from './preset';
	import type { Profile } from './profile';

	export let preset: Preset;
	const url = `/presets/${preset.id}`;

	let author: Profile | null = null;

	const connect = () => {
		return new PocketBase('http://127.0.0.1:8090');
	};

	const fetchAuthor = async () => {
		const client = connect();
		const authorRecord = await client.records.getOne('profiles', preset.author);
		author = {
			name: authorRecord.name,
			avatar: authorRecord.avatar
		};
	};

	onMount(async () => {
		await fetchAuthor();
	});
</script>

<div class="preset-card">
	<a href={url}>
		<div class="thumbnail" />
	</a>
	<div class="info">
		<a href={url}><p class="title">{preset.title}</p></a>
		<p class="author">By <a href={`/users/${preset.author}`}>{author?.name}</a></p>
	</div>
</div>

<style lang="scss">
	p,
	a {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}

	.preset-card {
		height: 200px;
		width: 300px;

		border: 1px solid #ccc;
		border-radius: 10px 10px 15px 15px; // Hides the top corners behind the thumbnail

		.info {
			padding-left: 10px;
			padding-right: 10px;
		}

		.thumbnail {
			height: 60%;
			width: 100%;
			background-color: blue;

			border-radius: 10px 10px 0 0;
		}
	}
</style>
