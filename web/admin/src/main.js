import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import './plugin/ant-design-vue-ui'
import './assets/css/style.css'

// config axios and detech to vue
axios.defaults.baseURL = 'http://127.0.0.1:8081/api/v1'
axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})
Vue.prototype.$http = axios
Vue.config.productionTip = false
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
