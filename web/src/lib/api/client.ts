/**
 * What-it-doo API
 * 1.0
 * DO NOT MODIFY - This file has been generated using oazapfts.
 * See https://www.npmjs.com/package/oazapfts
 */
import * as Oazapfts from "@oazapfts/runtime";
import * as QS from "@oazapfts/runtime/query";
export const defaults: Oazapfts.Defaults<Oazapfts.CustomHeaders> = {
    headers: {},
    baseUrl: "/api/v1",
};
const oazapfts = Oazapfts.runtime(defaults);
export const servers = {
    server1: "/api/v1"
};
export type DtoLoginRequest = {
    email: string;
    password: string;
    remember_me?: boolean;
};
export type DtoUserDetails = {
    bio?: string;
    email?: string;
    id?: string;
    name?: string;
};
export type DtoLoginResponse = {
    user: DtoUserDetails;
};
export type ProblemDetails = {
    detail?: string;
    instance?: string;
    status?: number;
    title?: string;
    "type"?: string;
};
export type DtoLogoutResponse = {
    redirect_url?: string;
    success?: boolean;
};
export type DtoRegistrationRequest = {
    email: string;
    name: string;
    password: string;
};
export type DtoRegistrationResponse = {
    user: DtoUserDetails;
};
export type DtoChat = {
    created_at?: string;
    id?: number;
    last_message?: string;
    participants?: DtoUserDetails[];
    title?: string;
    updated_at?: string;
};
export type DtoCreateChatRequest = {
    participants: string[];
};
export type DtoChatMessage = {
    content?: string;
    delivered_at?: string;
    id?: number;
    read_at?: string;
    sender_id?: string;
    sent_at?: string;
};
export type DtoChatMessages = {
    has_more?: boolean;
    messages?: DtoChatMessage[];
};
export type DtoServerInfo = {
    version?: string;
};
export type DtoServerConfig = object;
export type ModelUser = {
    avatar_url?: string;
    bio?: string;
    created_at?: string;
    email?: string;
    hashed_password?: string;
    id?: string;
    name?: string;
    updated_at?: string;
};
/**
 * Authenticate user
 */
export function login(dtoLoginRequest: DtoLoginRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: DtoLoginResponse;
    } | {
        status: 400;
        data: ProblemDetails;
    } | {
        status: 401;
        data: ProblemDetails;
    }>("/auth/login", oazapfts.json({
        ...opts,
        method: "POST",
        body: dtoLoginRequest
    })));
}
/**
 * Logout user
 */
export function logout(opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: DtoLogoutResponse;
    } | {
        status: 401;
        data: ProblemDetails;
    }>("/auth/logout", {
        ...opts,
        method: "POST"
    }));
}
/**
 * Register user
 */
export function register(dtoRegistrationRequest: DtoRegistrationRequest, { autoLogin }: {
    autoLogin?: boolean;
} = {}, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 201;
        data: DtoRegistrationResponse;
    } | {
        status: 400;
        data: ProblemDetails;
    } | {
        status: 500;
        data: ProblemDetails;
    }>(`/auth/register${QS.query(QS.explode({
        autoLogin
    }))}`, oazapfts.json({
        ...opts,
        method: "POST",
        body: dtoRegistrationRequest
    })));
}
/**
 * Get my chats
 */
export function getMyChats(opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: DtoChat[];
    }>("/chats/", {
        ...opts
    }));
}
/**
 * Create a new chat
 */
export function createChat(dtoCreateChatRequest: DtoCreateChatRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 201;
        data: DtoChat;
    }>("/chats/", oazapfts.json({
        ...opts,
        method: "POST",
        body: dtoCreateChatRequest
    })));
}
/**
 * Get chat by ID
 */
export function getChatById(chatId: number, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: DtoChat;
    } | {
        status: 404;
    }>(`/chats/${encodeURIComponent(chatId)}`, {
        ...opts
    }));
}
/**
 * Get chat messages
 */
export function getChatMessages(chatId: number, { limit, before }: {
    limit?: number;
    before?: string;
} = {}, opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: DtoChatMessages;
    }>(`/chats/${encodeURIComponent(chatId)}/messages${QS.query(QS.explode({
        limit,
        before
    }))}`, {
        ...opts
    }));
}
/**
 * Get server information
 */
export function getServerInfo(opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: DtoServerInfo;
    }>("/server/about", {
        ...opts
    }));
}
/**
 * Get server configuration
 */
export function getServerConfig(opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: DtoServerConfig;
    }>("/server/config", {
        ...opts
    }));
}
/**
 * Get current user
 */
export function getMyself(opts?: Oazapfts.RequestOpts) {
    return oazapfts.ok(oazapfts.fetchJson<{
        status: 200;
        data: ModelUser;
    }>("/users/me", {
        ...opts
    }));
}
