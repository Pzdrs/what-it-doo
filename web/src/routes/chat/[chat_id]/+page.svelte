<script lang="ts">
	import { page } from '$app/stores';
	import ChatListItem from '$lib/components/ChatListItem.svelte';
	import ChatMessage from '$lib/components/ChatMessage.svelte';
	import Icon from '$lib/components/Icon.svelte';
	import type { Chat } from '$lib/types';
	import type { PageProps } from './$types';
	import { mdiSquareEditOutline } from '@mdi/js';

	let { data }: PageProps = $props();

	$effect(() => {
		console.log('Load messages for chat:', $page.params.chat_id);
	});

	const currentChat = $derived($page.params.chat_id);

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
</script>

<div class="mt-3 flex h-full w-full px-5">
	<section class="flex-1 rounded-xl bg-gray-300 p-4">
		<div class="flex justify-between">
			<p class="p-2 text-4xl font-bold">Chats</p>
			<span class="text-white my-auto">
				<a href="#new-chat" aria-label="New Chat">
					<Icon icon={mdiSquareEditOutline} size="30" />
				</a>
			</span>
		</div>
		<div class="my-4">
			<input type="text" placeholder="Search What It Doo" />
		</div>
		<div>
			{#each chats as chat}
				<ChatListItem {chat} active={parseInt(currentChat) === chat.id} />
			{/each}
		</div>
	</section>
	<section class="flex-3 ml-3 rounded-xl bg-gray-100 p-4">
		<div class="flex">
			<div class="flex items-center rounded-lg p-2 transition-colors duration-200">
				<img
					class="h-9 w-9 rounded-full"
					src="https://scontent-prg1-1.xx.fbcdn.net/v/t39.30808-1/385763548_6948274105206399_3560856272069698376_n.jpg?stp=dst-jpg_p200x200_tt6&_nc_cat=103&ccb=1-7&_nc_sid=e99d92&_nc_ohc=bW-Nk_SfhrEQ7kNvwEYj1ab&_nc_oc=AdmnK9O53ElfigOxXct-Vi8G0jm10Q64AR71Rb62wFxLKOt4gJJsq8UQPuRQm9IpmMI&_nc_ad=z-m&_nc_cid=1097&_nc_zt=24&_nc_ht=scontent-prg1-1.xx&_nc_gid=iTU3jBr-XOTG0tPi5IzqBA&oh=00_AfZm3SJPNquCb3o8xQWx0bJc7fl10-3xZwVbUlw_WVim9g&oe=68D61548"
					alt="Chat Icon"
				/>
				<div class="ml-4">
					<h3 class="font-bold">Chat Title</h3>
					<p class="text-sm text-gray-600">Active now</p>
				</div>
			</div>
			<div></div>
		</div>
		<div class="">
			<ChatMessage message="Hello! How can I help you today?" origin="them" />
			<ChatMessage message="I have a question about my order." origin="us" />
		</div>
		<div class="flex">
			<div class="">
				<input name="message" type="text" placeholder="Aa" />
			</div>
			<span>
				<button>Send</button>
			</span>
		</div>
	</section>
</div>
