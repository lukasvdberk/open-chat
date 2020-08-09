import {getJWTKey} from "../auth/auth"
import {BASE_API_ENDPOINT} from "../../settings";

// TODO maybe replace with a reverse proxy like nginx
// const baseURL = "http://127.0.0.1:4000/api/"

function getHeaders() {
    const baseHeaders = {
        'Content-Type': 'application/json',
    }
    // Needed for authorization
    let jwtToken = getJWTKey()

    let headers = baseHeaders
    if (jwtToken !== undefined) {
        let authHeaders = {
            'Authorization': `Bearer ${jwtToken}`
        }

        // Adds the auth headers
        headers = {
            ...baseHeaders,
            ...authHeaders
        }
    }

    return headers
}


export async function get(endpoint) {
    const response = await fetch(BASE_API_ENDPOINT + endpoint, {
        headers: new Headers(getHeaders()),
    })

    return await response.json()
}

// Will only make json type of requests
export async function post(endpoint, data) {
    const response = await fetch(BASE_API_ENDPOINT + endpoint, {
        method: 'POST',
        body: JSON.stringify(data),
        headers: new Headers(getHeaders()),
    })

    return await response.json()
}
