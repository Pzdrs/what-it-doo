<script lang="ts">
	import { getChats } from '$lib/stores/chats.svelte';
	import type { Chat } from '$lib/types';
	import { format } from 'timeago.js';

	interface Props {
		currentChat: Chat;
	}

	let { currentChat }: Props = $props();

	let chatsPromis: Promise<Chat[]> = Promise.resolve(getChats());
</script>

<ul class="rounded-box flex w-full flex-1 flex-col gap-4 overflow-y-auto overflow-x-hidden">
	{#await chatsPromis}
		<div class="mt-10 text-center">
			<span class="loading loading-spinner loading-xl"></span>
		</div>
	{:then chats}
		{#each chats as chat}
			<li>
				<a href="/chat/{chat.id}" class:menu-active={currentChat.id === chat.id}>
					<div class="rounded-box flex items-center p-2 transition-colors duration-200">
						<div
							class="avatar"
							class:avatar-online={chat.getStatus()}
							class:avatar-offline={!chat.getStatus()}
						>
							<div class="w-12 rounded-full">
								<img src={chat.getAvatarUrl()} alt="Chat Avatar" />
							</div>
						</div>
						<div class="ml-4">
							<h3 class="font-bold">{chat.getTitle()}</h3>
							<p class="text-sm">
								<span>{chat.lastMessage.content}</span> &middot;
								<span>{format(chat.lastMessage.timestamp)}</span>
							</p>
						</div>
					</div>
				</a>
			</li>
		{/each}
	{:catch error}
		<div role="alert" class="alert alert-error">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="h-6 w-6 shrink-0 stroke-current"
				fill="none"
				viewBox="0 0 24 24"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
				/>
			</svg>
			<span>Failed to load chats!</span>
		</div>
	{/await}
</ul>
