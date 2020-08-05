import {getJWTKey} from "../auth/auth"


// TODO maybe replace with a reverse proxy
const baseURL = "http://127.0.0.1:4000/api/"
const baseHeaders = {
    'Content-Type': 'application/json',
}


// Will only make json type of requests
export async function post(endpoint, data) {
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

    console.log(headers)
    const response = await fetch(baseURL + endpoint, {
        method: 'POST',
        body: JSON.stringify(data),
        headers: new Headers(headers),
    })

    return await response.json()
}

export async function get(endpoint, data) {
    // TODO implement
}