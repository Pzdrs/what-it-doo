import { browser } from "$app/environment";
import { AppRoute, SESSION_COOKIE_NAME } from "$lib/constants";
import { getUser, setUser } from "$lib/stores/user.svelte";
import type { User } from "$lib/types";
import { redirect } from "@sveltejs/kit";

export interface AuthOptions {
    public?: boolean
}

export function getMyUser(): Promise<User> {
    return new Promise(resolve => {
        resolve({
            id: 'c175238b-88fe-4035-8c0b-c78c49ffee67',
            name: 'Pycrs',
            avatarUrl: 'https://gravatar.com/avatar/76699f36a61375c9dbd2931e0f518be2b1265948d8cd8843072f2ca70db1c6e4?v=1715076597000&size=256&d=initials'
        });
    });
}


export const loadUser = async () => {
    try {
        const user = getUser();
        if (!user && hasAuthCookie()) {
            let [_user] = await Promise.all([getMyUser()]);
            return setUser(_user);
        }
        return user;
    } catch {
        return null;
    }
};

const hasAuthCookie = (): boolean => {
    if (!browser) {
        return false;
    }

    for (const cookie of document.cookie.split('; ')) {
        const [name] = cookie.split('=');
        if (name === SESSION_COOKIE_NAME) {
            return true;
        }
    }

    return false;
};

export const authenticate = async (url: URL, options?: AuthOptions) => {
    const { public: publicRoute } = options || {};
    const user = await loadUser();

    if (publicRoute) {
        return;
    }

    if (!user) {
        redirect(302, `${AppRoute.AUTH_LOGIN}?continue=${encodeURIComponent(url.pathname + url.search)}`);
    }
};