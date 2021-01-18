/* eslint-disable indent */
import Vue from 'vue'
import axios from 'axios'
import router from '../router'

const URL = 'http://127.0.0.1:8081/api/v1'
// config axios and detech to vue
axios.defaults.baseURL = URL
axios.interceptors.request.use(config => {
    console.log(config)
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
    return config
})

function onFulfilled(response) {
    console.log('interceptor response: ', response)
    if (response.data.status === 1005) {
        Vue.prototype.$message.error(response.data.msg)
        router.replace('/login')
    }
    return response
}

function onRejected(error) {
    return Promise.reject(error)
}

// msg: "TOKEN已过期"
// status: 1005

axios.interceptors.response.use(onFulfilled, onRejected)

Vue.prototype.$http = axios

export { URL }
