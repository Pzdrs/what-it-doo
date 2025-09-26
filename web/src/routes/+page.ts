import { AppRoute } from "$lib/constants";
import { loadUser } from "$lib/utils/auth";
import { redirect } from "@sveltejs/kit";

export const ssr = false;

export const load = async () => {
    const authenticated = await loadUser();

    if (authenticated) redirect(302, AppRoute.CHAT);
};