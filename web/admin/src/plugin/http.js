/* eslint-disable indent */
import Vue from 'vue'
import axios from 'axios'

const URL = 'http://127.0.0.1:8081/api/v1'
// config axios and detech to vue
axios.defaults.baseURL = URL
axios.interceptors.request.use(config => {
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
    return config
})
Vue.prototype.$http = axios

export { URL }
