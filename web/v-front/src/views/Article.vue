<template >
  <div>
    <div class="d-flex justify-center pa-3 text-h4 font-weight-bold">{{articleInfo.title}}</div>
    <v-divider class="pa-3 ma-3"></v-divider>
    <v-alert class="ma-4" elevation="1" color="indigo" dark border="left" outlined>{{articleInfo.desc}}</v-alert>
    <div class="content ma-5 pa-3 text-justify" v-html="articleInfo.content"></div>
  </div>
</template>
<script>
export default {
  props: ['id'],
  created() {
    console.log(this.id)
    if (this.id) {
      this.getArticleInfo()
    }
  },
  data() {
    return {
      articleInfo: {}

    }
  },
  methods: {
    async getArticleInfo() {
      const { data: res } = await this.$http.get(`article/info/${this.id}`)
      console.log(res)
      if (res.status !== 200) return //
      this.articleInfo = res.data
    },
  }
}
</script>
<style scope>
</style>