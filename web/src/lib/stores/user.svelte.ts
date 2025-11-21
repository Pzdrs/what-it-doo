import type { DtoUserDetails } from '$lib/api/client';

class UserStore {
	user = $state<DtoUserDetails | null>(null);

	setUser(user: DtoUserDetails) {
		this.user = user;
	}
}

export const userStore = new UserStore();
