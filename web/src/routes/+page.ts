import { AppRoute } from "$lib/constants";
import { loadUser } from "$lib/utils/auth";
import { initLanguage } from "$lib/utils/i18n";
import { redirect } from "@sveltejs/kit";

export const load = async () => {
    await initLanguage();

    const authenticated = await loadUser();

    if (authenticated) redirect(302, AppRoute.CHAT);
};