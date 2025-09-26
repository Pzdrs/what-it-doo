import { redirect } from '@sveltejs/kit';

export const load = async ({ params, url }) => {
    redirect(302, '/chat/1');
};