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
    baseUrl: "/",
};
const oazapfts = Oazapfts.runtime(defaults);
export const servers = {};
export type DtoLoginRequest = {
    email: string;
    password: string;
    remember_me?: boolean;
};
export type DtoUserDetails = {
    avatar_url?: string;
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
export type ModelUser = {
    avatar_url?: string;
    bio?: string;
    created_at?: string;
    email?: string;
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
