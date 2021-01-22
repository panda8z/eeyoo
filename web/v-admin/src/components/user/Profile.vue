<template>
  <div>
    <a-card>
      <a-form-model
        ref="profileForm"
        :model="userProfile"
        :rules="userProfileRules"
      >
        <a-form-model-item label="昵称" prop="name">
          <a-input class="in" v-model="userProfile.name" />
        </a-form-model-item>
        <a-form-model-item label="个人描述" prop="desc">
          <a-input type="textarea" v-model="userProfile.desc" />
        </a-form-model-item>
        <a-form-model-item label="昵称" prop="name">
          <a-input class="in" v-model="userProfile.name" />
        </a-form-model-item>
        <a-form-model-item label="昵称" prop="name">
          <a-input class="in" v-model="userProfile.name" />
        </a-form-model-item>
        <a-form-model-item label="QQ" prop="qqchat">
          <a-input class="in" v-model="userProfile.qqchat" />
        </a-form-model-item>
        <a-form-model-item label="WeChat" prop="wechat">
          <a-input class="in" v-model="userProfile.wechat" />
        </a-form-model-item>
        <a-form-model-item label="Email" prop="email">
          <a-input class="in" v-model="userProfile.email" />
        </a-form-model-item>
        <a-form-model-item label="Facebook" prop="facebook">
          <a-input class="in" v-model="userProfile.facebook" />
        </a-form-model-item>
        <a-form-model-item label="Twitter" prop="twitter">
          <a-input class="in" v-model="userProfile.twitter" />
        </a-form-model-item>
        <!-- 头像 -->
        <a-form-model-item label="头像" prop="img">
          <a-upload
            name="file"
            :action="uploadURL"
            :headers="headers"
            listType="picture"
            @change="uploadAvatar"
          >
            <a-button> <a-icon type="upload" /> {{ "上传图片" }} </a-button>
            <br />
            <template v-if="userProfile.avatar">
              <img
                :src="userProfile.avatar"
                alt=""
                style="margin: 10px; width: 100px; height: 100px"
              />
            </template>
          </a-upload>
        </a-form-model-item>
        <!-- 头像背景 -->
        <a-form-model-item label="头像背景" prop="img">
          <a-upload
            name="file"
            :action="uploadURL"
            :headers="headers"
            listType="picture"
            @change="uploadAvatarBg"
          >
            <a-button> <a-icon type="upload" /> {{ "上传图片" }} </a-button>
            <br />
            <template v-if="userProfile.img">
              <img
                :src="userProfile.img"
                alt=""
                style="margin: 10px; width: 100px; height: 100px"
              />
            </template>
          </a-upload>
        </a-form-model-item>
        <a-form-model-item>
          <a-button type="danger" style="margin-right: 15px">更新</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>
import { URL } from '../../plugin/http'
export default {
  created() {
    this.getUserProfile()
  },
  data() {
    return {
      uploadURL: URL + '/upload',
      headers: {
        Authorization: 'Beare ' + window.sessionStorage.getItem('token')
      },
      uploadFileList: [],
      userProfile: {

      },
      userProfileRules: {

      }

    }
  },
  methods: {
    async getUserProfile() {
      const { data: res } = await this.$http.get('profile/1')
      console.log(res)
      if (res.status !== 200) {
        this.$message.error(res.msg)
        return
      }
      this.userProfile = res.data
    },
    uploadAvatar(info) {
      console.log('uploadAvatar info: ', info)
      if (info.file.status !== 'uploading') {

      }
      if (info.file.status === 'done' && info.file.response.status === 200) {
        this.$message.success('图片上传成功')
        const imgUrl = info.file.response.data.url
        this.userProfile.avatar = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error('图片上传失败')
      }
    },
    uploadAvatarBg(info) {
      console.log('uploadAvatarBg info: ', info)
      if (info.file.status !== 'uploading') {

      }
      if (info.file.status === 'done' && info.file.response.status === 200) {
        this.$message.success('图片上传成功')
        const imgUrl = info.file.response.data.url
        this.userProfile.img = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error('图片上传失败')
      }
    }
  }
}
</script>

<style scoped>
.in {
  width: 300px;
}
</style>
