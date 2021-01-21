<template>
  <v-col>
    <v-card class="ma-3" v-for="item in articleList" :key="item.ID" link @click="$router.push(`article/${item.ID}`)">
      <v-row no-gutters>
        <!-- 题图 -->
        <v-col class="d-flex justify-center align-center mx-3" cols="2">
          <v-img max-height="100" max-width="100" :src="item.img"></v-img>
        </v-col>
        <!-- 文章信息 -->
        <v-col>
          <!-- 标题 -->
          <v-card-title>
            <v-chip color="pink" label class="mr-3 white--text">{{
              item.Category.name
            }}</v-chip
            >{{ item.title }}
          </v-card-title>
          <!-- 描述 -->
          <v-card-subtitle class="ma-2" v-text="item.desc"></v-card-subtitle>
          <v-divider></v-divider>
          <!-- 更新时间 -->
          <v-card-text>
            <v-icon>{{ "mdi-calendar-month" }}</v-icon>
            <span>{{ item.CreatedAt }}</span>
          </v-card-text>
        </v-col>
      </v-row>
    </v-card>

    <!-- pagenation -->
    <div class="text-center">
      <v-pagination
        v-model="quaryParam.pageNum"
        :length="Math.ceil(this.total / quaryParam.pageSize)"
        total-visible="7"
        :input="getArticleList"
      >
      </v-pagination>
    </div>
  </v-col>
</template>

<script>
export default {
  created() {
    this.getArticleList()
  },
  data() {
    return {
      articleList: {},
      quaryParam: {
        pageSize: 20,
        pageNum: 1
      },
      total: 0
    }
  },
  methods: {
    async getArticleList() {
      const { data: res } = await this.$http.get('article/list', {
        params: {
          pageSize: this.quaryParam.pageSize,
          pageNum: this.quaryParam.pageNum
        }
      })
      console.log(res)
      if (res.status != 200) return //this.$message.error(res.msg)
      this.articleList = res.data.list
      this.total = res.data.total
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
