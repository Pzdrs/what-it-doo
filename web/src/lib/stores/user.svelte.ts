import type { User } from "$lib/types";

let user = $state<User>();

export const getUser = (): User | undefined => {
    return user;
};

export const setUser = (_user: User) => {
    user = _user;
    return user;
};

export const resetSavedUser = () => {
    user = undefined;
};
