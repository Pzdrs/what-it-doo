<script lang="ts">
	import { page } from '$app/state';
	import ChatFeed from '$lib/components/ChatFeed.svelte';
	import ChatList from '$lib/components/ChatList.svelte';
	import Icon from '$lib/components/Icon.svelte';
	import NewChatModal from '$lib/modals/NewChatModal.svelte';
	import { getChat } from '$lib/stores/chats.svelte';
	import type { Chat } from '$lib/types';
	import { mdiSquareEditOutline } from '@mdi/js';

	const currentChat: Chat | undefined = $derived(getChat(parseInt(page.params.chat_id || '0')));
</script>

<div
	class="bg-base-100 flex h-full w-full flex-col gap-x-0 gap-y-4 px-4 pb-6 pt-4 md:h-[calc(100vh-80px)] md:flex-row md:gap-x-4 md:gap-y-0"
>
	<section class="rounded-box bg-base-300 flex h-full w-full flex-col p-4 md:w-1/3">
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
		<ChatList {currentChat} />
	</section>
	<section class="flex-3 rounded-box bg-base-300 flex h-full w-full flex-col p-4 md:w-2/3">
		<ChatFeed chat={currentChat} />
	</section>
</div>

<NewChatModal />
