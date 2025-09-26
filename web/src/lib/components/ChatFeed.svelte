<script lang="ts">
	import { getMessages } from '$lib/stores/chats.svelte';
	import type { Chat } from '$lib/types';
	import { mdiHandWaveOutline, mdiSendOutline } from '@mdi/js';
	import ChatMessage from './ChatMessage.svelte';
	import Icon from './Icon.svelte';

	interface Props {
		chat: Chat;
	}

	const { chat }: Props = $props();

	$effect(() => {
		// TODO: send typing indicator to server
		console.log('Actively typing changed:', activelyTyping());
	});

	$effect(() => {
		console.log('Loaded chat ID:', chat.id);
	});
	let messages = $derived(() => {
		return getMessages(chat.id);
	});

	let message = $state('');
	let focus = $state(false);
	let typing = $derived(message.length > 0);
	let activelyTyping = $derived(() => typing && focus);

	function sendMessage() {
		console.log('Sending message:', message);
		message = '';
	}
	function dapUp() {
		console.log('Dapping up!');
	}
</script>

<div class="flex">
	<div class="flex items-center rounded-lg p-2 transition-colors duration-200">
		<div class="avatar avatar-online">
			<div class="w-9 rounded-full">
				<img
					src={chat.getAvatarUrl()}
					alt="Chat Avatar"
				/>
			</div>
		</div>
		<div class="ml-4">
			<h3 class="font-bold">
				{chat.getTitle()}
			</h3>
			<p class="text-base-content/75 text-sm">Active now</p>
		</div>
	</div>
	<div></div>
</div>
<div class="flex-grow overflow-auto">
	{#if messages().length > 0}
		{#each messages() as msg (msg.id)}
			<ChatMessage message={msg} />
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
		class="input w-full"
	/>
	<span>
		<div class="tooltip" data-tip={typing ? 'Send message' : 'Dap a homie up'}>
			<button
				class="btn btn-ghost"
				onclick={typing ? sendMessage : dapUp}
				aria-label={typing ? 'Send Message' : 'Dap Up'}
			>
				<Icon icon={typing ? mdiSendOutline : mdiHandWaveOutline} size="30" />
			</button>
		</div>
	</span>
</div>
