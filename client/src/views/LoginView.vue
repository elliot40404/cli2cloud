<script setup>
	import { ref, onMounted } from 'vue';
	import { useRouter, useRoute } from 'vue-router';
	import { io } from 'socket.io-client';
	import { useIdStore } from '../stores/main';
	import stripAnsi from 'strip-ansi';
	const router = useRouter();
	const route = useRoute();
	const store = useIdStore();
	const text = ref('');
	const scroll = ref(null);
	let socket;
	const room = route.query.id || route.params.id || store.id;
	store.id = room;
	onMounted(() => {
		const ws = import.meta.env.VITE_SOCKET || import.meta.env.BASE_URL;
		socket = io(`${ws}`);
		socket.on('connect', () => {
			console.log('connected');
			socket.emit('join', { room: room });
		});
		socket.on('message', (data) => {
			const parsed = stripAnsi(data.data);
			text.value +=
				`\n$ root@${room} ${new Date().toLocaleString()}\n` + parsed;
			scroll.value.scrollIntoView(scroll.value, {
				behavior: 'smooth',
				block: 'end',
			});
		});
	});
	const close = () => {
		store.$reset();
		router.push('/');
	};
</script>

<template>
	<div class="container">
		<header>
			<h3 class="logo">CLI 2 CLOUD</h3>
			<h3 class="id">Connected to - {{ room }}</h3>
			<h3 @click="close" class="exit">close &#10060;</h3>
		</header>
		<pre class="output"><code ref="output" id="output">{{text}}</code></pre>
		<button @click="text = ''" class="clear">Clear</button>
		<div class="scroll" ref="scroll"></div>
	</div>
</template>

<style scoped>
	header {
		display: flex;
		justify-content: space-between;
		padding: 0.5rem;
		align-items: center;
		height: 50px;
	}
	.logo {
		background: -webkit-linear-gradient(
			45deg,
			rgb(223, 255, 118),
			rgb(89, 165, 251)
		);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.id {
		color: rgb(119, 251, 141);
	}
	.exit {
		color: rgb(255, 119, 119);
		cursor: pointer;
	}
	.output {
		width: calc(100% - 2rem);
		margin: 1rem auto;
	}
	code {
		color: white;
		font-family: 'FiraCode NF', monospace;
		font-weight: 400;
	}
	.clear {
		position: fixed;
		bottom: 10px;
		right: 10px;
		font-family: inherit;
		padding: 0.5rem;
		font-size: 1rem;
		line-height: 1.5;
		background: rgb(228, 56, 56);
		font-weight: bold;
		border-radius: 5px;
		border: none;
		color: white;
		cursor: pointer;
	}
</style>
