<template>
  <div>
    <h3>用户列表</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search placeholder="请输入用户名查找" enter-button />
        </a-col>
        <a-col :span="4">
          <a-button type="primary">新增</a-button>
        </a-col>
      </a-row>
      <a-table
        rowKey="username"
        :columns="columns"
        :pagination="pageOptions"
        :dataSource="userlist"
      >
        <span slot="role" slot-scope="role">{{
          role == 1 ? "管理员" : "订阅者"
        }}</span>

        <template slot="action">
          <div class="actionSlot">
            <a-button type="primary" style="margin-right: 15px">编辑</a-button>
            <a-button type="danger">删除</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
const columns = [
  { title: 'ID', dataIndex: 'ID', width: '10%', key: 'id', align: 'center' },
  { title: '用户名', dataIndex: 'username', width: '20%', key: 'username', align: 'center' },
  { title: '角色', dataIndex: 'role', width: '20%', key: 'role', align: 'center', scopedSlots: { customRender: 'role' } },
  { title: '操作', width: '30%', key: 'action', align: 'center', scopedSlots: { customRender: 'action' } }
]
export default {
  created() {
    this.getUserList()
  },
  data() {
    return {
      userlist: [],
      columns,
      pageOptions: {
        pageSizeOptions: ['2', '5', '10'],
        defaultPageSize: 5,
        defaultCurrent: 1,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
        onChange: (page, pageSize) => {
          this.pageOptions.defaultCurrent = page
          this.pageOptions.defaultPageSize = pageSize
          this.getUserList()
        },
        onshowSizeChange: (current, size) => {
          this.pageOptions.defaultCurrent = current
          this.pageOptions.defaultPageSize = size
          this.getUserList()
        }
      }
    }
  },
  methods: {
    async getUserList() {
      const { data: res } = await this.$http.get('users', {
        params: {
          pageNum: this.pageOptions.defaultCurrent,
          pageSize: this.pageOptions.defaultPageSize
        }
      })
      console.log(res)
      this.userlist = res.data.list
      this.pageOptions.total = res.data.total
    }
  }
}
</script>
