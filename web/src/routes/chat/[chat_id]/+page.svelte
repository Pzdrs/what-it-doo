<script lang="ts">
	import { page } from '$app/state';
	import ChatListItem from '$lib/components/ChatListItem.svelte';
	import ChatMessage from '$lib/components/ChatMessage.svelte';
	import Icon from '$lib/components/Icon.svelte';
	import NewChatModal from '$lib/modals/NewChatModal.svelte';
	import type { Chat } from '$lib/types';
	import type { PageProps } from './$types';
	import { mdiHandWaveOutline, mdiSendOutline, mdiSquareEditOutline } from '@mdi/js';

	let { data }: PageProps = $props();

	$effect(() => {
		// TODO: send typing indicator to server
		console.log('Actively typing changed:', activelyTyping());
	});

	const currentChat = $derived(page.params.chat_id);

	const chats: Chat[] = [
		{
			id: 1,
			title: 'Chat 1',
			lastMessage: {
				id: crypto.randomUUID(),
				chatId: 1,
				content: 'Hello, how are you?',
				timestamp: new Date()
			},
			participants: []
		},
		{
			id: 2,
			title: 'Chat 2',
			lastMessage: {
				id: crypto.randomUUID(),
				chatId: 2,
				content: "What's up?",
				timestamp: new Date('2023-10-01T10:00:00Z')
			},
			participants: []
		},
		{
			id: 3,
			title: 'Chat 3',
			lastMessage: {
				id: crypto.randomUUID(),
				chatId: 3,
				content: "Let's catch up!",
				timestamp: new Date('2025-01-15T15:30:00Z')
			},
			participants: []
		}
	];

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

<div class="bg-base-100 mt-3 flex h-auto w-full px-5">
	<section class="rounded-box bg-base-300 h-max flex-1 p-4">
		<div class="flex justify-between">
			<p class="text-primary p-2 text-4xl font-bold">Chats</p>
			<span class="my-auto text-white">
				<a
					href="#new-chat"
					class="btn tooltip"
					data-tip="Start a new chat"
					aria-label="New Chat"
					onclick={() =>
						(document.getElementById('new-chat-dialog') as HTMLDialogElement).showModal()}
				>
					<Icon icon={mdiSquareEditOutline} class="size-[1.2em]" />
					Compose
				</a>
			</span>
		</div>
		<div class="my-4">
			<label class="input w-full">
				<svg class="h-[1em] opacity-50" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
					<g
						stroke-linejoin="round"
						stroke-linecap="round"
						stroke-width="2.5"
						fill="none"
						stroke="currentColor"
					>
						<circle cx="11" cy="11" r="8"></circle>
						<path d="m21 21-4.3-4.3"></path>
					</g>
				</svg>
				<input type="search" required placeholder="Search conversations" />
			</label>
		</div>
		<ul class="menu rounded-box h-full w-full px-0">
			{#each chats as chat}
				<ChatListItem {chat} active={parseInt(currentChat) === chat.id} />
			{/each}
		</ul>
		<div></div>
	</section>
	<section class="flex-3 rounded-box bg-base-300 ml-3 flex flex-col p-4">
		<div class="flex">
			<div class="flex items-center rounded-lg p-2 transition-colors duration-200">
				<div class="avatar avatar-online">
					<div class="w-9 rounded-full">
						<img
							src="https://scontent-prg1-1.xx.fbcdn.net/v/t39.30808-1/385763548_6948274105206399_3560856272069698376_n.jpg?stp=dst-jpg_p200x200_tt6&_nc_cat=103&ccb=1-7&_nc_sid=e99d92&_nc_ohc=bW-Nk_SfhrEQ7kNvwEYj1ab&_nc_oc=AdmnK9O53ElfigOxXct-Vi8G0jm10Q64AR71Rb62wFxLKOt4gJJsq8UQPuRQm9IpmMI&_nc_ad=z-m&_nc_cid=1097&_nc_zt=24&_nc_ht=scontent-prg1-1.xx&_nc_gid=iTU3jBr-XOTG0tPi5IzqBA&oh=00_AfZm3SJPNquCb3o8xQWx0bJc7fl10-3xZwVbUlw_WVim9g&oe=68D61548"
							alt="Chat Icon"
						/>
					</div>
				</div>
				<div class="ml-4">
					<h3 class="font-bold">Chat Title</h3>
					<p class="text-base-content/75 text-sm">Active now</p>
				</div>
			</div>
			<div></div>
		</div>
		<div class="flex-grow">
			{#if true}
				<ChatMessage message="Hello! How can I help you today?" origin="them" />
				<ChatMessage message="I have a question about my order." origin="us" />
			{:else}
				<p class="text-base-content/50 mt-20 text-center">
					No messages yet. Start the conversation!
				</p>
			{/if}
		</div>
		<div class="flex justify-between gap-2">
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
	</section>
</div>

<NewChatModal />
