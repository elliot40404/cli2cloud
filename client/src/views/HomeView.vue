<script setup>
	import { ref } from 'vue';
	import { useRouter, useRoute } from 'vue-router';
	import { useIdStore } from '../stores/main';
	const store = useIdStore();
	const router = useRouter();
	const route = useRoute();
	const id = ref("");
	const submit = () => {
		store.setId(id.value);
		router.push({
			name: 'cli',
			params: {
				id: id.value,
			},
		});
	};
</script>

<template>
	<div class="container">
		<h1>CLI 2 CLOUD</h1>
		<form @submit="submit">
			<label for="cli_id">
				CLI-ID:
				<input
					autocomplete="off"
					required
					minlength="8"
					maxlength="8"
					type="text"
					id="cli_id"
					placeholder="elliot27"
					v-model="id"
				/>
			</label>
			<button type="submit">Connect !</button>
		</form>
	</div>
</template>

<style scoped>
	.container {
		display: flex;
		flex-direction: column;
		align-items: center;
		height: 100vh;
	}
	h1 {
		font-size: 3em;
		margin: 0;
		padding: 2rem;
		background: -webkit-linear-gradient(
			45deg,
			rgb(223, 255, 118),
			rgb(89, 165, 251)
		);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	form {
		padding: 4rem;
	}
	label {
		font-weight: bold;
	}
	input {
		font-family: inherit;
		padding: 0.5rem;
		text-align: center;
		letter-spacing: 2px;
		border-radius: 0.25rem;
		color: aquamarine;
		font-weight: bold;
		border: none;
		background: #505050;
	}
	button {
		font-family: inherit;
		padding: 0.5rem;
		border-radius: 0.25rem;
		color: aquamarine;
		font-weight: bold;
		border: none;
		background: #505050;
		margin: 1rem;
		cursor: pointer;
	}
</style>
