import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import './plugin/ant-design-vue-ui'
import './assets/css/style.css'

// config axios and detech to vue
axios.defaults.baseURL = 'http://localhost:8081/api/v1'
Vue.prototype.$http = axios

Vue.config.productionTip = false
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
