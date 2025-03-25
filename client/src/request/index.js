import { ElLoading, ElMessage } from 'element-plus'
import HTTP from './http'

const http = new HTTP('http://127.0.0.1:3000')

if (window.fetchURL) {
    fetchURL().then(baseURL => {
        http._baseURL = baseURL
    })
}

export function getURL(route) {
    return http._baseURL + route
}

export function post(route, data, callback) {
    const loading = ElLoading.service({ lock: true, background: 'rgba(0,0,0,0.5)' })
    let status = 200
    return http.postJson(route, data).then(res => {
        if (res.status == 200) {
            return res.json()
        }
        status = res.status
        return res.text()
    }).then(result => {
        if (status == 200) {
            if (callback) {
                callback(result)
            }
        } else if (status == 400) {
            ElMessage({ message: "参数错误", type: "warning" })
        } else if (status == 500) {
            ElMessage({ message: result, type: "warning" })
        }
    }).catch(() => {
        ElMessage({ message: "网络错误", type: "error" })
    }).finally(() => {
        loading.close()
    })
}