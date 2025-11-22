<script lang="ts">
	import type { DtoUserDetails } from '$lib/api/client';
	import ChatMessage from '$lib/components/ChatMessage.svelte';
	import Icon from '$lib/components/Icon.svelte';
	import TypingIndicator from '$lib/components/TypingIndicator.svelte';
	import { getGroupChatTitle, getOtherChatParticipants, getTheOtherParticipant } from '$lib/utils/chat';
	import { messagingStore } from '$stores/chats.svelte';
	import { userStore } from '$stores/user.svelte';
	import { sendWebSocketMessage } from '$stores/websocket.svelte';
	import { mdiHandWaveOutline, mdiSendOutline } from '@mdi/js';
	import { t } from 'svelte-i18n';

	let scrollEl: HTMLDivElement | null = null;
	let initialized = false;
	const chat = $derived(messagingStore.currentChat);

	let loadingOlder = $state(false);
	let message = $state('');
	let focus = $state(false);
	let typing = $derived(message.length > 0);
	let activelyTyping = $derived(() => typing && focus);
	let autoScroll = $state(true);

	$effect(() => {
		[messagingStore.currentMessages, messagingStore.currentTypingUsers];

		if (scrollEl && autoScroll) {
			requestAnimationFrame(() => {
				scrollEl.scrollTop = scrollEl.scrollHeight;
			});
		}
	});

	$effect(() => {
		const typing = activelyTyping();
		if (!initialized) {
			initialized = true;
			return;
		}
		sendWebSocketMessage('typing', { typing, chat_id: chat?.id });
	});

	function sendMessage() {
		// negative so we dont conflict with server ids
		const localId = -(Date.now() * 1000 + Math.floor(Math.random() * 1000));
		sendWebSocketMessage('message', { message, chat_id: chat?.id, temp_id: localId });
		messagingStore.sendMessageOptimistic(localId, chat?.id, userStore.user?.id, message);
		message = '';
	}

	function dapUp() {
		sendWebSocketMessage('dap_up', {});
	}

	async function handleScroll() {
		if (!scrollEl) return;
		autoScroll = scrollEl.scrollTop + scrollEl.clientHeight >= scrollEl.scrollHeight - 50;
		if (scrollEl.scrollTop === 0 && !loadingOlder) {
			loadingOlder = true;
			const newMessages = await messagingStore.loadOlderMessages();
			loadingOlder = false;

			if (!newMessages) return;
			// scroll down a little so we dont have to scroll down and up again to continue loading
			requestAnimationFrame(() => {
				scrollEl!.scrollTop = 275;
			});
		}
	}
</script>

{#snippet avatarGroupMember(participant: DtoUserDetails)}
	<div class="avatar" title={participant.name}>
		<div class="w-9">
			<img alt={participant.name} src={participant.avatar_url} />
		</div>
	</div>
{/snippet}

<div class="flex">
	<div class="flex items-center rounded-lg p-2 transition-colors duration-200">
		{#if chat?.participants.length == 2}
			<div class="avatar" class:avatar-online={true} class:avatar-offline={false}>
				<div class="w-9 rounded-full">
					<img src={getTheOtherParticipant(chat, userStore.user!)?.avatar_url} alt="Chat Avatar" />
				</div>
			</div>
		{:else if chat?.participants.length == 3}
			<div class="avatar-group -space-x-6">
				{#each getOtherChatParticipants(chat, userStore.user!).slice(0, 2) as participant}
					{@render avatarGroupMember(participant)}
				{/each}
			</div>
		{:else if chat?.participants.length! > 3}
			<div class="avatar-group -space-x-6">
				{#each getOtherChatParticipants(chat!, userStore.user!).slice(0, 2) as participant}
					{@render avatarGroupMember(participant)}
				{/each}

				<div class="avatar avatar-placeholder">
					<div class="bg-neutral text-neutral-content w-9">
						<span>+{chat!.participants.length - 3}</span>
					</div>
				</div>
			</div>
		{/if}

		<div class="ml-4">
			<h3 class="font-bold">
				{#if chat?.title}
					{chat.title}
				{:else if chat?.participants?.length === 2}
					{getTheOtherParticipant(chat, userStore.user!)?.name || 'Unknown User'}
				{:else}
					{getGroupChatTitle(chat!, userStore.user!, 50)}
				{/if}
			</h3>
			<p class="text-base-content/75 text-sm">Active now</p>
		</div>
	</div>
</div>

<div class="flex-grow overflow-auto" bind:this={scrollEl} onscroll={handleScroll}>
	{#if loadingOlder}
		<div class="text-center">
			<span class="loading loading-spinner loading-xl"></span>
		</div>
	{/if}

	{#if messagingStore.currentMessages.length > 0}
		{#each messagingStore.currentMessages as message (message.id)}
			<ChatMessage {message} />
		{/each}

		{#each messagingStore.currentTypingUsers as user (user.id)}
			<TypingIndicator {user} />
		{/each}
	{:else}
		<p class="text-base-content/50 mt-20 text-center">No messages yet. Start the conversation!</p>
	{/if}
</div>

<div class="mt-2 flex justify-between gap-2">
	<input
		type="text"
		placeholder="Aa"
		bind:value={message}
		onfocus={() => (focus = true)}
		onblur={() => (focus = false)}
		onkeydown={(e) => {
			if (e.key === 'Enter' && !e.shiftKey) {
				e.preventDefault();
				if (typing) sendMessage();
			}
		}}
		class="input w-full"
	/>
	<div class="tooltip" data-tip={typing ? 'Send message' : 'Dap a homie up'}>
		<button
			class="btn btn-ghost"
			onclick={typing ? sendMessage : dapUp}
			aria-label={typing ? 'Send Message' : 'Dap Up'}
		>
			<Icon icon={typing ? mdiSendOutline : mdiHandWaveOutline} size="30" />
		</button>
	</div>
</div>
