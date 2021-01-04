
import Vue from 'vue'
import {
  Button,
  FormModel,
  Icon,
  Input,
  message,
  Layout,
  Menu,
  Card,
  Table,
  Row,
  Col,
  ConfigProvider
} from 'ant-design-vue'

Vue.prototype.$message = message

Vue.use(Button)
Vue.use(FormModel)
Vue.use(Input)
Vue.use(Icon)
Vue.use(Layout)
Vue.use(Menu)
Vue.use(Card)
Vue.use(Table)
Vue.use(Row)
Vue.use(Col)
Vue.use(ConfigProvider)
