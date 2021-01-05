<template>
  <div>
    <h3>分类列表</h3>
    <a-card>
      <a-row :gutter="20">
        <!-- 添加分类按钮 -->
        <a-col :span="4">
          <a-button type="primary" @click="addCate">新增</a-button>
        </a-col>
      </a-row>
      <!-- 分类列表 -->
      <a-table
        rowKey="name"
        :columns="columns"
        :pagination="pageOptions"
        :dataSource="catelist"
        @change="handleTablePageChange"
      >
        <template slot="action" slot-scope="actionData">
          <div class="actionSlot">
            <a-button
              type="primary"
              style="margin-right: 15px"
              @click="editCate(actionData)"
            >
              编辑
            </a-button>
            <a-button type="danger" @click="deleteCate(actionData.ID)">
              删除
            </a-button>
          </div>
        </template>
      </a-table>
    </a-card>
    <!-- 添加分类弹窗 -->
    <a-modal
      title="新增分类"
      width="25%"
      closable
      destroyOnClose
      :visible="addCateVisible"
      @ok="addCateOk"
      @cancel="addCateCancle"
    >
      <a-form-model ref="addCateModalRef" :model="cateInfo" :rules="cateRules">
        <a-form-model-item label="分类名" prop="catename">
          <a-input v-model="cateInfo.name" />
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑分类弹窗 -->
    <a-modal
      title="编辑分类"
      width="25%"
      closable
      destroyOnClose
      :visible="editCateVisible"
      @ok="editCateOk"
      @cancel="editCateCancle"
    >
      <a-form-model ref="editCateModalRef" :model="cateInfo" :rules="cateRules">
        <a-form-model-item label="ID">
          <a-input :disabled="true" v-model="cateInfo.id" />
        </a-form-model-item>
        <a-form-model-item label="分类名" prop="catename">
          <a-input v-model="cateInfo.name" />
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  { title: 'ID', dataIndex: 'ID', width: '10%', key: 'id', align: 'center' },
  { title: '分类名', dataIndex: 'name', width: '20%', key: 'name', align: 'center' },
  { title: '操作', width: '30%', key: 'action', align: 'center', scopedSlots: { customRender: 'action' } }
]
export default {
  created() {
    this.getCateList()
  },
  data() {
    return {
      catelist: [],
      columns,
      addCateVisible: false,
      editCateVisible: false,
      defaultCateInfo: {
        id: 0,
        name: ''
      },
      cateInfo: {
        id: 0,
        name: ''
      },
      cateRules: {
        catename: [
          {
            validator: (rule, value, callback) => {
              if (this.cateInfo.name === '') {
                callback(new Error('请输入分类名'))
              }
              if ([...this.cateInfo.name].length < 2 || [...this.cateInfo.name].length > 8) {
                callback(new Error('分类名长度须满足2-8位'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      queryParams: {
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
      this.getCateList()
    },
    async deleteCate(cateID) {
      console.log(cateID)
      this.$confirm({
        title: '提示！',
        content: '确认删除此分类吗？删除操作无法恢复！',
        onOk: async () => {
          const ret = await this.$http.delete(`cate/${cateID}`)
          console.log(ret)
          if (ret.status !== 200) {
            this.$message.error(ret.error)
            return
          }
          this.$message.success('删除成功！')
          this.getCateList()
        },
        onCancel: () => { }
      })
    },
    addCate() {
      this.addCateVisible = true
    },
    addCateOk() {
      this.$refs.addCateModalRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('参数不符合要求，请检查后重新输入')
        }
        const ret = await this.$http.post('cate/add', {
          name: this.cateInfo.name,
          password: this.cateInfo.password,
          role: Number(this.cateInfo.role)
        })
        console.log(ret)
        if (ret.data.status !== 200) {
          this.$message.error(ret.data.err)
        }
        this.$message.success('新增分类成功！')
        this.addCateVisible = false
        this.$refs.addCateModalRef.resetFields()
        this.getCateList()
        this.cateInfo = { ...this.defaultCateInfo }
      })
    },
    addCateCancle() {
      this.$refs.addCateModalRef.resetFields()
      this.addCateVisible = false
      this.cateInfo = { ...this.defaultCateInfo }
    },
    editCate(actionData) {
      this.cateInfo.id = actionData.ID
      this.cateInfo.name = actionData.name
      this.cateInfo.role = actionData.role
      this.editCateVisible = true
    },
    editCateOk() {
      this.$refs.editCateModalRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('参数不符合要求，请检查后重新输入')
        }
        const ret = await this.$http.put(`cate/${this.cateInfo.id}`, {
          id: Number(this.cateInfo.id),
          name: this.cateInfo.name
        })
        console.log(ret)
        if (ret.data.status !== 200) {
          this.$message.error(ret.data.err)
        }
        this.$message.success('编辑分类成功！')
        this.editCateVisible = false
        this.$refs.editCateModalRef.resetFields()
        this.cateInfo = { ...this.defaultCateInfo }
        this.getCateList()
      })
    },
    editCateCancle() {
      this.cateInfo = { ...this.defaultCateInfo }
      this.$refs.editCateModalRef.resetFields()
      this.editCateVisible = false
    },
    handleAdminChange(value) {
      this.cateInfo.role = value
    },
    async getCateList() {
      const { data: res } = await this.$http.get('cates', {
        params: {
          pageNum: this.queryParams.pageNum,
          pageSize: this.queryParams.pageSize
        }
      })
      console.log(res)
      if (res.status !== 200) return this.$message.error(res.message)
      this.catelist = res.data.list
      this.pageOptions.total = res.data.total
    }
  }
}
</script>
