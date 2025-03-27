export default class {
    constructor(baseURL, headers) {
        this._baseURL = baseURL
        this._headers = headers
    }

    get(route) {
        return fetch(this._baseURL + route, { headers: this._headers })
    }

    post(route, contentType, body) {
        const headers = {}
        if (contentType) {
            headers["Content-Type"] = contentType
        }
        if (this._headers) {
            for (let key in this._headers) {
                headers[key] = this._headers[key]
            }
        }
        return fetch(this._baseURL + route, { method: "POST", headers, body })
    }

    postJson(route, data) {
        return this.post(route, "application/json", JSON.stringify(data))
    }

    postForm(route, data) {
        return this.post(route, "", data)
    }
}

export function getBlob(url, callback) {
    const request = new XMLHttpRequest()
    request.open("GET", url, true)
    request.responseType = "blob"
    request.onload = function (e) {
        callback(e.target.response)
    }
    request.send()
}