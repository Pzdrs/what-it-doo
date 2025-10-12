import { AppRoute } from "$lib/constants";
import { init } from "$lib/utils/server";
import { redirect } from "@sveltejs/kit";
import type { LayoutLoad } from "./$types";

export const ssr = false;
export const csr = true;

export const load = (async ({fetch}) => {
    await init(fetch);

    redirect(302, AppRoute.CHAT);
}) satisfies LayoutLoad;