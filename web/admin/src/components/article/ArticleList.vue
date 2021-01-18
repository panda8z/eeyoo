<template>
  <div>
    <h3>文章列表</h3>
    <a-card>
      <a-row :gutter="20">
        <!-- 查找文章输入框和按钮 -->
        <a-col :span="6">
          <a-input-search
            allowClear
            placeholder="请输入文章名查找"
            v-model="queryParams.title"
            enter-button
            @search="getArticleList"
          />
        </a-col>
        <!-- 添加文章按钮 -->
        <a-col :span="4">
          <a-button type="primary" @click="$router.push('/writearticle')"
            >新增</a-button
          >
        </a-col>

        <a-col :span="6" :offset="4">
          <a-select
            defaultValue="请选择分类"
            style="width: 200px"
            @change="onCateSelectChange"
          >
            <a-select-option
              v-for="item in catelist"
              :key="item.ID"
              :value="item.ID"
            >
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-col>
      </a-row>
      <!-- 文章列表 -->
      <a-table
        rowKey="ID"
        :columns="columns"
        :pagination="pageOptions"
        :dataSource="articlelist"
        @change="handleTablePageChange"
      >
        <span slot="img" slot-scope="img">
          <img style="width: 100px; height: 120px" :src="img" alt="" />
        </span>
        <template slot="action" slot-scope="actionData">
          <div class="actionSlot">
            <a-button
              type="primary"
              style="margin-right: 15px"
              @click="editArticle(actionData)"
            >
              编辑
            </a-button>
            <a-button type="danger" @click="deleteArticle(actionData.ID)">
              删除
            </a-button>
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
const columns = [
  { title: 'ID', dataIndex: 'ID', width: '5%', key: 'id', align: 'center' },
  { title: '标题', dataIndex: 'title', width: '15%', key: 'title', align: 'center' },
  { title: '描述', dataIndex: 'desc', width: '25%', key: 'desc', align: 'center' },
  { title: '分类', dataIndex: 'Category.name', width: '10%', key: 'cid', align: 'center' },
  { title: '缩略图', dataIndex: 'img', width: '25%', key: 'img', align: 'center', scopedSlots: { customRender: 'img' } },
  { title: '操作', width: '20%', key: 'action', align: 'center', scopedSlots: { customRender: 'action' } }
]
export default {
  created() {
    this.getArticleList()
    this.getCateList()
  },
  data() {
    return {
      articlelist: [],
      catelist: [],
      columns,
      addArticleVisible: false,
      editArticleVisible: false,
      articleInfo: {
        ID: 0,
        title: '',
        cid: 0,
        desc: '',
        content: '',
        img: ''
      },
      queryParams: {
        title: '',
        cid: 0,
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
      this.getArticleList()
    },
    async getArticleList() {
      const { data: res } = await this.$http.get('article/list', {
        params: {
          title: this.queryParams.title,
          pageNum: this.queryParams.pageNum,
          pageSize: this.queryParams.pageSize
        }
      })
      console.log(res)
      if (res.status !== 200) return this.$message.error(res.msg)
      this.articlelist = res.data.list
      this.pageOptions.total = res.data.total
    },
    deleteArticle(articleID) {
      console.log('文章ID: ', articleID)
      this.$confirm({
        title: '提示！',
        content: '确认删除此文章吗？删除操作无法恢复！',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`article/${articleID}`)
          console.log('deleteArticle:', res)
          if (res.status !== 200) {
            this.$message.error(res.msg)
            return
          }
          this.$message.success('删除成功！')
          this.getArticleList()
        },
        onCancel: () => { }
      })
    },
    editArticle(info) {
      this.$router.push(`/writearticle/${info.ID}`)
    },
    async getCateList() {
      const { data: res } = await this.$http.get('cates')
      console.log('getCateList: ', res)
      if (res.status !== 200) return this.$message.error(res.message)
      this.catelist = res.data.list
    },
    async onCateSelectChange(cateID) {
      console.log('onCateSelectChange value:', cateID)
      const { data: res } = await this.$http.get(`article/cate/${cateID}`, {
        params: {
          pageNum: this.queryParams.pageNum,
          pageSize: this.queryParams.pageSize
        }
      })
      console.log(res)
      if (res.status !== 200) {
        this.$message.error(res.msg)
      }
      this.articlelist = res.data.list
      this.pageOptions.total = res.data.total
    }
  }
}
</script>
