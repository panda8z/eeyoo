<template>
  <div>
    <a-card>
      <h3>{{ id ? "编辑文章" : "新增文章" }}</h3>
      <a-form-model
        :model="articleInfo"
        ref="articleForm"
        :rules="articleRules"
        :hideRequiredMark="true"
      >
        <!-- 文章标题 -->
        <a-form-model-item label="文章标题" prop="title">
          <a-input style="width: 300px" v-model="articleInfo.title" />
        </a-form-model-item>
        <!-- 文章分类 -->
        <a-form-model-item label="文章分类" prop="cid">
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
        <a-form-model-item label="文章描述" prop="desc">
          <a-input v-model="articleInfo.desc" />
        </a-form-model-item>
        <!-- 文章缩略图 -->
        <a-form-model-item label="文章缩略图" prop="img">
          <a-upload
            name="file"
            :action="uploadURL"
            :headers="headers"
            listType="picture"
            :default-file-list="uploadFileList"
            @change="uploadImgs"
          >
            <a-button> <a-icon type="upload" /> {{ "上传图片" }} </a-button>
            <br />
            <template v-if="id">
              <img
                :src="articleInfo.img"
                alt=""
                style="margin: 10px; width: 100px; height: 100px"
              />
            </template>
          </a-upload>
        </a-form-model-item>
        <!-- 文章内容 -->
        <a-form-model-item label="文章内容" prop="content">
          <!-- <a-input v-model="articleInfo.content" /> -->
          <mavon-editor
            ref="md"
            v-model="articleInfo.content"
            @imgAdd="$imgAdd"
            @imgDel="$imgDel"
          />
        </a-form-model-item>
        <!-- 提交和取消按钮 -->
        <a-form-model-item :wrapper-col="{ span: 14, offset: 0 }">
          <a-button type="primary" @click="onSubmit()">
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
    if (this.id) {
      this.getArticleInfo(this.id)
    }
    this.getCategoryList()
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
  },
  data() {
    return {
      defaultArticleInfo: {
        ID: undefined,
        title: '',
        cid: undefined,
        desc: '',
        content: '',
        img: ''
      },
      articleInfo: {
        ID: undefined,
        title: '',
        cid: undefined,
        desc: '',
        content: '',
        img: ''
      },
      categoryList: [],
      uploadURL: URL + '/upload',
      headers: {
        Authorization: 'Beare ' + window.sessionStorage.getItem('token')
      },
      uploadFileList: [],
      articleRules: {
        title: [{ required: true, message: '请输入文章标题', trigger: 'blur' }],
        cid: [{ required: true, message: '请输入分类', trigger: 'change' }],
        desc: [
          { required: true, message: '请输入文章描述', trigger: 'blur' },
          { max: 120, message: '文章描述最多120个字符', trigger: 'change' }],
        content: [{ required: true, message: '请输入文章内容', trigger: 'blur' }],
        img: [{ required: true, message: '请上传文章图拍呢', trigger: 'change' }]
      }
    }
  },
  methods: {
    handleCategoryChange(value) {
      console.log(value)
      this.articleInfo.cid = value
    },
    async getArticleInfo(id) {
      if (id) {
        // 根据ID获取文章信息
        const { data: res } = await this.$http.get(`article/info/${id}`)
        console.log('getArticleInfo:', res)
        if (res.status !== 200) return this.$message.error(res.msg)
        this.articleInfo = res.data
      } else {
        // 没有ID就不获取文章信息
      }
    },
    async getCategoryList() {
      const { data: res } = await this.$http.get('cates')
      console.log('getCategoryList:', res)
      if (res.status !== 200) return this.$message.error(res.msg)
      this.categoryList = res.data.list
    },
    uploadImgs(info) {
      console.log('uploadImgs info: ', info)
      if (info.file.status !== 'uploading') {

      }
      if (info.file.status === 'done') {
        this.$message.success('图片上传成功')
        const imgUrl = info.file.response.data.url
        this.articleInfo.img = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error('图片上传失败')
      }
    },
    async submitArticle() {
      if (this.id) {
        // 有ID就 更新article
        this.articleInfo.ID = this.id
        const { data: res } = await this.$http.put(`article/${this.id}`, this.articleInfo)
        console.log('add article:', res)
      } else {
        // 没ID就 新建article
        console.log(this.articleInfo)
        const { data: res } = await this.$http.post('article/add', this.articleInfo)
        console.log('add article:', res)
        if (res.status !== 200) {
          return this.$message.error('提交失败：' + res.msg)
        }
        this.$message.success('提交成功')
        this.articleInfo = this.defaultArticleInfo
        this.$refs.articleForm.resetFields()
      }
    },
    onSubmit() {
      this.$refs.articleForm.validate(valid => {
        console.log(valid)
        if (!valid) {
          return this.$message.error('参数不符合要求，请检查后重新输入')
        }
        this.submitArticle()
      })
    },
    cancleAdd() {
      this.$refs.articleForm.resetFields()
    },
    async uploadEditorImg(file) {
      const data = new FormData()
      const { data: res } = await this.$http.post('/upload', data, { headers: { 'Content-Type': 'multipart/form-data' } })
      if (res.status !== 200) return this.$message.error(res.msg)
      return res.data.url
    },
    $imgDel() { },
    $imgAdd(pos, file) {
      console.log(pos, file)
      this.uploadEditorImg(file).then((url) => {
        this.$refs.md.$img2Url(pos, url)
      })
    }
  }// methods
}// export
</script>
