<template>
  <div>
    <h3>用户列表</h3>
    <a-card>
      <a-row :gutter="20">
        <!-- 查找用户输入框和按钮 -->
        <a-col :span="6">
          <a-input-search
            allowClear
            placeholder="请输入用户名查找"
            v-model="queryParams.username"
            enter-button
            @search="getUserList"
          />
        </a-col>
        <!-- 添加用户按钮 -->
        <a-col :span="4">
          <a-button type="primary" @click="addUser">新增</a-button>
        </a-col>
      </a-row>
      <!-- 用户列表 -->
      <a-table
        rowKey="username"
        :columns="columns"
        :pagination="pageOptions"
        :dataSource="userlist"
        @change="handleTablePageChange"
      >
        <!-- 角色一栏的渲染逻辑 -->
        <span slot="role" slot-scope="role">{{
          role == 1 ? "管理员" : "订阅者"
        }}</span>

        <template slot="action" slot-scope="actionData">
          <div class="actionSlot">
            <a-button
              type="primary"
              style="margin-right: 15px"
              @click="editUser(actionData)"
            >
              编辑
            </a-button>
            <a-button type="danger" @click="deleteUser(actionData.ID)">
              删除
            </a-button>
          </div>
        </template>
      </a-table>
    </a-card>
    <!-- 添加用户弹窗 -->
    <a-modal
      title="新增用户"
      width="25%"
      closable
      destroyOnClose
      :visible="addUserVisible"
      @ok="addUserOk"
      @cancel="addUserCancle"
    >
      <a-form-model ref="addUserModalRef" :model="userInfo" :rules="userRules">
        <!-- <a-form-model-item label="ID">
          <a-input v-model="userInfo.id" />
        </a-form-model-item> -->
        <a-form-model-item label="用户名" prop="uername">
          <a-input v-model="userInfo.username" />
        </a-form-model-item>
        <a-form-model-item has-feedback label="密码" prop="password">
          <a-input-password v-model="userInfo.password" />
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkpwd">
          <a-input-password v-model="userInfo.checkpwd" />
        </a-form-model-item>
        <a-form-model-item label="是否为管理员">
          <a-select
            default-value="2"
            style="width: 120px"
            @change="handleAdminChange"
          >
            <a-select-option key="1" value="1"> 是 </a-select-option>
            <a-select-option key="2" value="2"> 否 </a-select-option>
          </a-select>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑用户弹窗 -->
    <a-modal
      title="编辑用户"
      width="25%"
      closable
      destroyOnClose
      :visible="editUserVisible"
      @ok="editUserOk"
      @cancel="editUserCancle"
    >
      <a-form-model ref="editUserModalRef" :model="userInfo" :rules="userRules">
        <a-form-model-item label="ID">
          <a-input :disabled="true" v-model="userInfo.id" />
        </a-form-model-item>
        <a-form-model-item label="用户名" prop="uername">
          <a-input v-model="userInfo.username" />
        </a-form-model-item>
        <!-- <a-form-model-item has-feedback label="密码" prop="password">
          <a-input-password v-model="userInfo.password" />
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkpwd">
          <a-input-password v-model="userInfo.checkpwd" />
        </a-form-model-item> -->
        <a-form-model-item label="是否为管理员">
          <a-select
            :default-value="`${userInfo.role}`"
            style="width: 60px"
            @change="handleAdminChange"
          >
            <a-select-option value="1"> 是 </a-select-option>
            <a-select-option value="2"> 否 </a-select-option>
          </a-select>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
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
      visible: false,
      addUserVisible: false,
      editUserVisible: false,
      defaultUserInfo: {
        id: 0,
        username: '',
        password: '',
        checkpwd: '',
        role: 2
      },
      userInfo: {
        id: 0,
        username: '',
        password: '',
        checkpwd: '',
        role: 2
      },
      userRules: {
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.username === '') {
                callback(new Error('请输入用户名'))
              }
              if ([...this.userInfo.username].length < 4 || [...this.userInfo.username].length > 22) {
                callback(new Error('用户名长度须满足4-22位'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.password === '') {
                callback(new Error('请输入密码'))
              }
              if ([...this.userInfo.password].length < 6 || [...this.userInfo.password].length > 20) {
                callback(new Error('密码长度须满足6-20位'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkpwd: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.checkpwd === '') {
                callback(new Error('请填写确认密码'))
              }
              if (this.userInfo.password !== this.userInfo.checkpwd) {
                callback(new Error('密码不一致，请重新输入'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      queryParams: {
        username: '',
        pageSize: 5,
        pageNum: 1
      },
      pageOptions: {
        pageSizeOptions: ['2', '5', '10'],
        defaultPageSize: 5,
        defaultCurrent: 1,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`
      }
    }
  },
  methods: {
    handleTablePageChange(pagination, filters, sorter) {
      var pager = { ...this.pageOptions }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParams.pageSize = pagination.pageSize
      this.queryParams.pageNum = pagination.current

      if (pagination.pageSize !== this.pageOptions.pageSize) {
        this.queryParams.pageNum = 1
        pager.current = 1
      }
      this.pageOptions = pager
      this.getUserList()
    },
    async deleteUser(userID) {
      console.log(userID)
      this.$confirm({
        title: '提示！',
        content: '确认删除此用户吗？删除操作无法恢复！',
        onOk: async () => {
          const ret = await this.$http.delete(`user/${userID}`)
          console.log(ret)
          if (ret.status !== 200) {
            this.$message.error(ret.error)
            return
          }
          this.$message.success('删除成功！')
          this.getUserList()
        },
        onCancel: () => { }
      })
    },
    addUser() {
      this.addUserVisible = true
    },
    addUserOk() {
      this.$refs.addUserModalRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('参数不符合要求，请检查后重新输入')
        }
        const ret = await this.$http.post('user/add', {
          username: this.userInfo.username,
          password: this.userInfo.password,
          role: Number(this.userInfo.role)
        })
        console.log(ret)
        if (ret.data.status !== 200) {
          this.$message.error(ret.data.err)
        }
        this.$message.success('新增用户成功！')
        this.addUserVisible = false
        this.$refs.addUserModalRef.resetFields()
        this.getUserList()
        this.userInfo = { ...this.defaultUserInfo }
      })
    },
    addUserCancle() {
      this.$refs.addUserModalRef.resetFields()
      this.addUserVisible = false
      this.userInfo = { ...this.defaultUserInfo }
    },
    editUser(actionData) {
      this.userInfo.id = actionData.ID
      this.userInfo.username = actionData.username
      this.userInfo.role = actionData.role
      this.editUserVisible = true
    },
    editUserOk() {
      this.$refs.editUserModalRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('参数不符合要求，请检查后重新输入')
        }
        const ret = await this.$http.put(`user/${this.userInfo.id}`, {
          username: this.userInfo.username,
          role: Number(this.userInfo.role)
        })
        console.log(ret)
        if (ret.data.status !== 200) {
          this.$message.error(ret.data.err)
        }
        this.$message.success('编辑用户成功！')
        this.editUserVisible = false
        this.$refs.editUserModalRef.resetFields()
        this.userInfo = { ...this.defaultUserInfo }
        this.getUserList()
      })
    },
    editUserCancle() {
      this.userInfo = { ...this.defaultUserInfo }
      this.$refs.editUserModalRef.resetFields()
      this.editUserVisible = false
    },
    handleAdminChange(value) {
      this.userInfo.role = value
    },
    async getUserList() {
      const { data: res } = await this.$http.get('users', {
        params: {
          username: this.queryParams.username,
          pageNum: this.queryParams.pageNum,
          pageSize: this.queryParams.pageSize
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.userlist = res.data.list
      this.pageOptions.total = res.data.total
    }
  }
}
</script>
