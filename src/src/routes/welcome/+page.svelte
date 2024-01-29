<script>
	async function welcome() {
		const response = await fetch('http://localhost:8080/welcome', {
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			},
		});

		if (response.ok) {
			console.log('Login successful');
			// Handle successful login
		} else {
			console.error('Login failed');
			// Handle login failure
		}
	}

	welcome()
</script>

<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div class="text-column">
	<h1>Welcome</h1>


</div>
