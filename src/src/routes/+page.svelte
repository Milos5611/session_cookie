<script>
	let username = '';
	let password = '';

	async function login() {
		const response = await fetch('http://localhost:8080/signin', {
			credentials: "include",
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ username, password })
		});


		if (response.ok) {
			console.log('Login successful');
			// Handle successful login
		} else {
			console.error('Login failed');
			// Handle login failure
		}
	}
</script>

<form on:submit|preventDefault={login}>
	<div>
		<label for="username">Username:</label>
		<input id="username" type="text" bind:value={username}>
	</div>
	<div>
		<label for="password">Password:</label>
		<input id="password" type="password" bind:value={password}>
	</div>
	<button type="submit">Login</button>
</form>
