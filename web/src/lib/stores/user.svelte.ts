import type { ModelUser } from '$lib/api/client';

class UserStore {
	user = $state<ModelUser | null>(null);
}

export const userStore = new UserStore();
