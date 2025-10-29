import { authenticate } from '$lib/utils/auth';
import { messagingStore } from '$stores/chats.svelte';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, url }) => {
	await authenticate(url);

	// Await, so we can set the meta title correctly
	await messagingStore.setCurrentChat(Number(params.chat_id));

	return {
		meta: {
			title: messagingStore.currentChat?.title
		}
	};
};
