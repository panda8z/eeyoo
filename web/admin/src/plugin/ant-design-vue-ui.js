
import Vue from 'vue'
import {
  Button,
  FormModel,
  Icon,
  Input,
  message
} from 'ant-design-vue'

Vue.prototype.$message = message

Vue.use(Button)
Vue.use(FormModel)
Vue.use(Input)
Vue.use(Icon)
