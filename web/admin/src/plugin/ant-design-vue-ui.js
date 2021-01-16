
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
  ConfigProvider,
  Modal,
  Select,
  Upload,
  Spin
} from 'ant-design-vue'

Vue.prototype.$message = message // 语法糖挂载 提示消息toast 到prototype
Vue.prototype.$confirm = Modal.confirm // 语法糖挂载 确认对话框 到prototype

Vue.use(Button) // 按钮
Vue.use(FormModel) // 表单
Vue.use(Input) // 输入框
Vue.use(Icon) // 图标
Vue.use(Layout) // 布局用到的
Vue.use(Menu) // 导航用到的菜单
Vue.use(Card) // 卡片
Vue.use(Table) // 表格
Vue.use(Row) // 栅格布局 行
Vue.use(Col) // 栅格布局 列
Vue.use(ConfigProvider) // 多语言
Vue.use(Modal) // 对话框
Vue.use(Select) // 选择下拉框
Vue.use(Upload) // 上传按钮
Vue.use(Spin) // 加载中
