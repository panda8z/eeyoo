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
          <a-button type="primary" @click="$router.push('/admin/writearticle')"
            >新增</a-button
          >
        </a-col>
      </a-row>
      <!-- 文章列表 -->
      <a-table
        rowKey="name"
        :columns="columns"
        :pagination="pageOptions"
        :dataSource="articlelist"
        @change="handleTablePageChange"
      >
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
  { title: 'ID', dataIndex: 'ID', width: '10%', key: 'id', align: 'center' },
  { title: '标题', dataIndex: 'title', width: '20%', key: 'title', align: 'center' },
  { title: '描述', dataIndex: 'desc', width: '20%', key: 'desc', align: 'center' },
  { title: '分类', dataIndex: 'cid', width: '10%', key: 'cid', align: 'center' },
  { title: '缩略图', dataIndex: 'img', width: '25%', key: 'img', align: 'center' },
  { title: '操作', width: '15%', key: 'action', align: 'center', scopedSlots: { customRender: 'action' } }
]
export default {
  created() {
    this.getArticleList()
  },
  data() {
    return {
      articlelist: [],
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
      const { data: res } = await this.$http.get('article/list', this.queryParams)
      console.log(res)
      if (res.status !== 200) return this.$message.error(res.msg)
    },
    deleteArticle(articleID) {

    },
    editArticle(info) {
      this.$router.push(`admin/writearticle/${info.ID}`)
    }
  }
}
</script>
