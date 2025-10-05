import type { ModelUser } from "$lib/api/client";

let user = $state<ModelUser>();

export const getUser = (): ModelUser | undefined => {
    return user;
};

export const setUser = (_user: ModelUser) => {
    user = _user;
    return user;
};

export const resetSavedUser = () => {
    user = undefined;
};
