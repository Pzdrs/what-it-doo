<script lang="ts">
	import type { ChatMessage, User } from '$lib/types';
	import { formatDateOrTime } from '$lib/utils';
	import { getUser } from '$lib/stores/user.svelte';

	type ChatMessageOrigin = 'us' | 'them';

	interface Props {
		message: ChatMessage;
	}

	let { message }: Props = $props();

	const user = getUser();

	const origin = user.id === message.sender.id ? 'us' : 'them';
</script>

<div class="chat" class:chat-start={origin === 'them'} class:chat-end={origin === 'us'}>
	<div class="chat-image avatar">
		<div class="w-10 rounded-full">
			<img alt="{message.sender.name} Avatar" src={message.sender.avatarUrl} />
		</div>
	</div>
	<div class="chat-header">
		{message.sender.name}
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
	{#if message.sender.id === user.id}
		<div class="chat-footer opacity-50">
			{#if message.readAt}
				<span class="text-xs">Seen at {formatDateOrTime(message.readAt)}</span>
			{:else if message.deliveredAt}
				<span class="text-xs">Delivered</span>
			{:else if message.timestamp}
				<span class="text-xs">Sent</span>
			{:else}
				<span class="text-xs">Sending</span>
			{/if}
		</div>
	{/if}
</div>
