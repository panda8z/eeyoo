<template>
  <div>
    <h3>{{ id ? "编辑文章" : "新增文章" }}</h3>
    <a-card>
      <a-form-model :model="articleInfo">
        <!-- 文章标题 -->
        <a-form-model-item label="文章标题">
          <a-input style="width: 300px" v-model="articleInfo.title" />
        </a-form-model-item>
        <!-- 文章分类 -->
        <a-form-model-item label="文章分类">
          <a-select
            style="width: 200px"
            v-model="articleInfo.cid"
            placeholder="请选择分类"
            @change="handleCategoryChange"
          >
            <a-select-option
              v-for="item in categoryList"
              :key="item.ID"
              :value="item.ID"
              >{{ item.name }}</a-select-option
            >
          </a-select>
          <!-- 文章描述 -->
        </a-form-model-item>
        <a-form-model-item label="文章描述">
          <a-input v-model="articleInfo.desc" />
        </a-form-model-item>
        <!-- 文章缩略图 -->
        <a-form-model-item label="文章缩略图">
          <a-upload
            name="file"
            :action="uploadURL"
            :headers="headers"
            listType="picture"
            :default-file-list="uploadFileList"
          >
            <a-button> <a-icon type="upload" /> Upload </a-button>
          </a-upload>
        </a-form-model-item>
        <!-- 文章内容 -->
        <a-form-model-item label="文章内容">
          <a-input v-model="articleInfo.content" />
        </a-form-model-item>
        <!-- 提交和取消按钮 -->
        <a-form-model-item :wrapper-col="{ span: 14, offset: 4 }">
          <a-button type="primary" @click="onSubmit(articaleInfo.ID)">
            <span>{{ id ? "更新" : "提交" }}</span>
          </a-button>
          <a-button style="margin-left: 10px" @click="cancleAdd">
            取消
          </a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>
import { URL } from '../../plugin/http'
export default {
  props: ['id'],
  created() {
    console.log('写文章', this.id)
    this.getArticleInfo()
    this.getCategoryList()
  },
  data() {
    return {
      articleInfo: {
        ID: 0,
        title: '',
        cid: undefined,
        desc: '',
        content: '',
        img: ''
      },
      categoryList: [],
      uploadURL: URL + '/upload',
      headers: {},
      uploadFileList: []
    }
  },
  methods: {
    handleCategoryChange(value) {
      console.log(value)
      this.articleInfo.cid = value
    },
    getArticleInfo(id) {
      if (id) {
        // 根据ID获取文章信息
      } else {
        // 没有ID就不获取文章信息
      }
    },
    async getCategoryList() {
      const { data: res } = await this.$http.get('cates')
      console.log(res)
      if (res.status !== 200) return this.$message.error(res.msg)
      this.categoryList = res.data.list
    },
    onSubmit(id) {
      if (id) {
        // 有ID就 更新article
      } else {
        // 没ID就 新建article
      }
    },
    cancleAdd() { }
  }
}
</script>
