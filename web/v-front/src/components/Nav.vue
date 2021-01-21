<template>
  <!-- 导航栏 卡片 -->
  <v-card class="mx-auto" max-width="300">
    <!-- 头部:头像区背景图 -->
    <v-img src="../assets/nav-bg.jpeg">
      <v-card-title>
        <v-col align="center">
          <!-- 圆形头像 -->
          <v-avatar size="130" color="grey">
            <img src="../assets/arvatar.png" alt />
          </v-avatar>
          <!-- 博主名字 -->
          <div class="ma-4 white--text">eeyoo</div>
        </v-col>
      </v-card-title>
      <!-- 分割线 -->
      <v-divider></v-divider>
    </v-img>

    <!-- 关于信息 -->
    <v-col>
      <div class="ma-3">About Me:</div>
      <div class="ma-3">Developer Programmer Designer</div>
    </v-col>

    <v-divider color="indigo"></v-divider>

    <!-- 社交列表 -->
    <v-list nav dense>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{ `mdi-qqchat` }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey-text">9XXXXXXXXX0</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{ `mdi-wechat` }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey-text">9XXXXXXXXX0</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{ `mdi-sina-weibo` }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey-text">9XXXXXXXXX0</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{ `mdi-facebook` }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey-text">9XXXXXXXXX0</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{ `mdi-email` }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey-text">9XXXXXXXXX0</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon>{{ `mdi-twitter` }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey-text">9XXXXXXXXX0</v-list-item-content>
      </v-list-item>
    </v-list>
  </v-card>
</template>

<script>
export default {

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
