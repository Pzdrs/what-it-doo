<script lang="ts">
	import { type DtoUserDetails, getMyChats } from '$lib/api/client';
	import { messagingStore } from '$lib/stores/chats.svelte';
	import { getGroupChatTitle, getOtherChatParticipants, getTheOtherParticipant } from '$lib/utils/chat';
	import { userStore } from '$stores/user.svelte';
	import { t } from 'svelte-i18n';
	import { format } from 'timeago.js';

	interface Props {}

	let {}: Props = $props();

	async function loadChats() {
		const fetchedChats = await getMyChats();
		messagingStore.chats = fetchedChats;
		return fetchedChats;
	}
</script>

{#snippet avatarGroupMember(participant: DtoUserDetails)}
	<div class="avatar" title={participant.name}>
		<div class="w-12">
			<img alt={participant.name} src={participant.avatar_url} />
		</div>
	</div>
{/snippet}

<ul class="menu flex w-full flex-1 flex-col gap-4 overflow-y-auto overflow-x-hidden">
	{#await loadChats()}
		<div class="mt-10 text-center">
			<span class="loading loading-spinner loading-xl"></span>
		</div>
	{:then chats}
		{#each chats as chat}
			<li>
				<a
					data-sveltekit-preload-data="tap"
					href="/chat/{chat.id}"
					class:menu-active={messagingStore.currentChat &&
						messagingStore.currentChat.id === chat.id}
				>
					<div class="rounded-box flex items-center p-2 transition-colors duration-200">
						{#if chat.participants.length == 2}
							<div class="avatar" class:avatar-online={true} class:avatar-offline={false}>
								<div class="w-12 rounded-full">
									<img
										src={getTheOtherParticipant(chat, userStore.user!)?.avatar_url}
										alt="Chat Avatar"
									/>
								</div>
							</div>
						{:else if chat.participants.length > 2}
							<div class="avatar avatar-placeholder">
								<div class="bg-neutral text-neutral-content w-12 rounded-full">
									<span class="text-3xl">{chat.participants.length}</span>
								</div>
							</div>
						{/if}

						<div class="ml-4">
							<h3 class="font-bold">
								{#if chat.title}
									{chat.title}
								{:else if chat.participants?.length == 2}
									{getTheOtherParticipant(chat, userStore.user!)?.name || 'Unknown User'}
								{:else}
									{getGroupChatTitle(chat!, userStore.user!, 20)}
								{/if}
							</h3>
							<p class="text-sm">
								<span>{$t('last_message')}</span> &middot;
								<span>{format(Date.now())}</span>
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
