import { authenticate } from "$lib/utils/auth";

export const load = async ({url}) => {
	await authenticate(url)
	return {
		meta: {
			title: 'Chats'
		}
	};
};
