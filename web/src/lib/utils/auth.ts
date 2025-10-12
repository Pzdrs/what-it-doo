import { browser } from "$app/environment";
import { getMyself } from "$lib/api/client";
import { AppRoute, SESSION_COOKIE_NAME } from "$lib/constants";
import { getUser, setUser } from "$lib/stores/user.svelte";
import { redirect } from "@sveltejs/kit";

export interface AuthOptions {
    public?: boolean
}

export const loadUser = async () => {
    try {
        const user = getUser();
        if (!user && hasAuthCookie()) {
            let [_user] = await Promise.all([getMyself()]);
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

export const requireNoAuth = async (url: URL) => {
    const user = await loadUser();
    if (user) {
        redirect(302, AppRoute.HOME);
    }
};
