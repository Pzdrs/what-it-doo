<script lang="ts">
	import type { Snippet } from 'svelte';

	const ALERT_TYPE_CLASSES = {
		error: 'alert-error',
		info: 'alert-info',
		success: 'alert-success',
		warning: 'alert-warning'
	};

	interface Props {
		message: string;
		children?: Snippet;
		type?: 'error' | 'info' | 'success' | 'warning';
	}

	let { children, message, type }: Props = $props();

	let alertClass = $derived(() => {
		if (!type || !(type in ALERT_TYPE_CLASSES)) {
			return 'alert';
		}
		return `alert ${ALERT_TYPE_CLASSES[type]}`;
	});
</script>

{#if message}
	<div role="alert" class={alertClass()}>
		<svg
			xmlns="http://www.w3.org/2000/svg"
			class="h-6 w-6 shrink-0 stroke-current"
			fill="none"
			viewBox="0 0 24 24"
		>
			{#if type === 'error'}
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
				/>
			{:else if type === 'info'}
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
				></path>
			{:else if type === 'success'}
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
				></path>
			{:else if type === 'warning'}
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
				/>
			{/if}
		</svg>
		<span>{message}</span>
		{@render children?.()}
	</div>
{/if}
