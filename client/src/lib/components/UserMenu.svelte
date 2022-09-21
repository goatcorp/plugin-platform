<script lang="ts">
	import { UserIcon } from 'svelte-feather-icons';

	export let id: string;
</script>

<div class="dropdown dropdown-end">
	<span tabindex="0" class="btn btn-sm"><UserIcon /></span>
	<ul tabindex="0" class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-52">
		<li><a href={`/user/${id}`}>My account</a></li>
		<li><a href="/settings">Settings</a></li>
		<li><a href="/logout">Log out</a></li>
	</ul>
</div>
