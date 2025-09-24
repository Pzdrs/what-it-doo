<script lang="ts">
	type ChatMessageOrigin = 'us' | 'them';

	interface Props {
		message: string;
		origin: ChatMessageOrigin;
	}

	let { message, origin = 'them' }: Props = $props();
</script>

<div class="my-2 flex" class:justify-end={origin === 'us'} class:justify-start={origin === 'them'}>
	<div
		class="rounded-selector max-w-xs p-3"
		class:bg-primary={origin === 'us'}
		class:bg-neutral={origin === 'them'}
	>
		<span
			class:text-primary-content={origin === 'us'}
			class:text-neutral-content={origin === 'them'}
		>
			{message}
		</span>
	</div>
</div>
