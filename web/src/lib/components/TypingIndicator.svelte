<script lang="ts">
	import type { User } from '$lib/types';

	interface Props {
		user: User;
	}

	let { user }: Props = $props();
</script>

<div class="chat chat-start">
	<!-- Avatar -->
	<div class="chat-image avatar">
		<div class="w-10 rounded-full">
			<img src={user.avatarUrl} alt={user.name} />
		</div>
	</div>

	<!-- Typing indicator bubble -->
	<div class="chat-bubble bg-neutral py-3">
		<div class="flex space-x-1">
			<span class="dot animate-jump h-2 w-2 rounded-full bg-gray-400"></span>
			<span class="dot animate-jump h-2 w-2 rounded-full bg-gray-400 delay-150"></span>
			<span class="dot animate-jump h-2 w-2 rounded-full bg-gray-400 delay-300"></span>
		</div>
	</div>
</div>

<style>
	.dot {
		display: inline-block;
	}

	.animate-jump {
		animation: jump 1s infinite;
	}

	.delay-150 {
		animation-delay: 0.15s;
	}

	.delay-300 {
		animation-delay: 0.3s;
	}

	@keyframes jump {
		0%,
		100% {
			transform: translateY(0);
		}
		50% {
			transform: translateY(-6px);
		}
	}
</style>
