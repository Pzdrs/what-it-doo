import { authenticate } from '$lib/utils/auth';
import { messagingStore } from '$stores/chats.svelte';
import type { PageLoad } from './$types';

export const load: PageLoad = (async ({ params, url }) => {
	await authenticate(url);

	const chatId = parseInt(params.chat_id, 10);

	// Await, so we can set the meta title correctly
	await messagingStore.setCurrentChat(chatId);
	await messagingStore.initLoadMessages(chatId);

	return {
		meta: {
			title: messagingStore.currentChat?.title
		}
	};
}) satisfies PageLoad;
