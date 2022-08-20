<script lang="ts">
	import PocketBase from 'pocketbase';
	import { id } from '$lib/session';
	import { goto } from '$app/navigation';

	let email: string = '';
	let password: string = '';
	let passwordConfirm: string = '';

	const connect = () => {
		return new PocketBase('http://127.0.0.1:8090');
	};

	const login = async (email: string, password: string) => {
		if (email === '' || password === '') {
			return;
		}

		console.log('Logging in');

		const client = connect();
		try {
			const userAuth = await client.users.authViaEmail(email, password);
			$id = userAuth.user.id;

			try {
				await goto('/');
			} catch (err) {
				console.error(err);
			}
		} catch (err) {
			console.error(err);
		}
	};

	const createUser = async (email: string, password: string, passwordConfirm: string) => {
		if (email === '' || password === '' || passwordConfirm !== password) {
			return;
		}

		console.log('Creating user');

		const client = connect();
		try {
			await client.users.create({
				email,
				password,
				passwordConfirm
			});
			await login(email, password);
		} catch (err) {
			console.error(err);
		}
	};
</script>

<form>
	<label for="email">Email</label>
	<input type="text" name="email" bind:value={email} />
	<br />
	<label for="password">Password</label>
	<input type="password" name="password" bind:value={password} />
	<br />
	<input type="button" name="login" value="Log in" on:click={() => login(email, password)} />
</form>

<form>
	<label for="email">Email</label>
	<input type="text" name="email" bind:value={email} />
	<br />
	<label for="password">Password</label>
	<input type="password" name="password" bind:value={password} />
	<br />
	<label for="confirm">Confirm password</label>
	<input type="password" name="confirm" bind:value={passwordConfirm} />
	<br />
	<input
		type="button"
		name="create"
		value="Create account"
		on:click={() => createUser(email, password, passwordConfirm)}
	/>
</form>

<style lang="scss">
	* {
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell,
			'Open Sans', 'Helvetica Neue', sans-serif;
	}
</style>
