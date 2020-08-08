import {post} from "../api/request";
import {onAuthenticationStateChange} from "./auth-store";

export async function register(username, password) {
    // Returns a bool whether it failed or not
    const result = await post('auth/register', {
        'username': username,
        'password': password
    })

    if (result.code === 0) {
        return await login(username, password)
    }
    else {
        alert("Failed to register this user")
    }
}

export async function login(username, password) {
    // Returns a bool whether it failed or not
    const result = await post('auth/login',  {
        'username': username,
        'password': password
    })
    if (result.code === 0) {
        saveJWTKey(result.content.token)
        saveUserId(result.content.userId)
        saveUsername(username)
        onAuthenticationStateChange.set(true)
    } else {
        return false
    }
}

export function isAuthenticated() {
    // TODO add expiration
    return getJWTKey() !== null
}

export function saveJWTKey(jwtToken) {
    // TODO add expiration
    localStorage.setItem("jwtKey", jwtToken)
}

export function getJWTKey() {
    return localStorage.getItem("jwtKey")
}

export function saveUsername(username) {
    localStorage.setItem("username", username)
}

export function getUsername() {
    return localStorage.getItem("username")
}

export function saveUserId(id) {
    localStorage.setItem("id", id)
}

export function getUserId() {
    return localStorage.getItem("id")
}


export async function renewKey() {
    // TODO implement
}