<template>
  <div>
    <v-app-bar app color="primary" flat>
      <v-avatar size="40px" color="#333"></v-avatar>
      <v-container class="py-0 fill-height">
        <v-btn text color="white">首页</v-btn>
        <v-btn v-for="item in categoryList" :key="item.id" text color="white">{{
          item.name
        }}</v-btn>
      </v-container>
      <v-spacer></v-spacer>
      <v-responsive max-width="260px" color="white">
        <v-text-field
          dense
          flat
          hide-details
          solo-inverted
          rounded
          dark
        ></v-text-field>
      </v-responsive>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  created() {
    this.getCateList()
  },
  data() {
    return {
      categoryList: {}
    }
  },
  methods: {
    // 获取 分类列表
    async getCateList() {
      const { data: res } = await this.$http.get('cates', { pageSize: 20, pagezNum: 1 })
      console.log(res)
      if (res.status !== 200) return this.$message.error(res.message)
      this.categoryList = res.data.list
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
