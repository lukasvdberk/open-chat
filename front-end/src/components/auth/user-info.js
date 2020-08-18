import {get} from "../api/request";

const localStorageKey = "userData"
export async function setUserInfoFromApi() {
    const response = await get("user/get-user-info")
    localStorage.setItem(localStorageKey, JSON.stringify(response.content.user))
}

export function getUserInfo() {
    return JSON.parse(localStorage.getItem(localStorageKey))
}