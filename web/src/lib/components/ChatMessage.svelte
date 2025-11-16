<script lang="ts">
	import type { DtoChatMessage } from '$lib/api/client';
	import { formatDateOrTime } from '$lib/utils';
	import { messagingStore } from '$stores/chats.svelte';
	import { userStore } from '$stores/user.svelte';

	interface Props {
		message: DtoChatMessage;
	}

	let { message }: Props = $props();

	const origin = $derived(userStore.user?.id === message.sender_id ? 'us' : 'them');
	const sender = $derived.by(() => {
		return messagingStore.allParticipants.find((p) => p.id === message.sender_id);
	});
</script>

<div class="chat" class:chat-start={origin === 'them'} class:chat-end={origin === 'us'}>
	<div class="chat-image avatar">
		<div class="w-10 rounded-full">
			<img alt="{sender?.name} Avatar" src={sender?.avatar_url} />
		</div>
	</div>
	<div class="chat-header">
		{#if origin === 'them'}
			{sender?.name}
		{/if}
		{#if message.timestamp}
			<time class="text-xs opacity-50">{formatDateOrTime(message.timestamp)}</time>
		{/if}
	</div>
	<div
		class="chat-bubble"
		class:bg-primary={origin === 'us'}
		class:bg-neutral={origin === 'them'}
		class:text-primary-content={origin === 'us'}
		class:text-neutral-content={origin === 'them'}
	>
		{message.content}
	</div>

	{#if message.sender_id === userStore.user?.id}
		<div class="chat-footer opacity-50">
			{#if message.read_at}
				<span class="text-xs">Seen at {formatDateOrTime(message.readAt)}</span>
			{:else if message.delivered_at}
				<span class="text-xs">Delivered</span>
			{:else if message.sent_at}
				<span class="text-xs">Sent</span>
			{:else}
				<span class="text-xs">Sending</span>
			{/if}
		</div>
	{/if}
</div>
