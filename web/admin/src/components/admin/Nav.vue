<template>
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="logo">
      <span>{{ collapsed ? "EEYOO" : "EEYOO Admin" }}</span>
    </div>
    <a-menu mode="inline" theme="dark" @click="jumpPage">
      <a-menu-item key="index">
        <a-icon type="dashboard" />
        <span>仪表盘</span>
      </a-menu-item>
      <a-sub-menu>
        <span slot="title">
          <a-icon type="diff" />
          <span>文章管理</span>
        </span>
        <a-menu-item key="writearticle">
          <span>
            <a-icon type="edit" />
            <span>写文章</span>
          </span>
        </a-menu-item>
        <a-menu-item key="articlelist">
          <span>
            <a-icon type="unordered-list" />
            <span>文章列表</span>
          </span>
        </a-menu-item>
      </a-sub-menu>
      <a-menu-item key="catelist">
        <a-icon type="funnel-plot" />
        <span>分类列表</span>
      </a-menu-item>
      <a-menu-item key="userlist">
        <a-icon type="profile" />
        <span>用户列表</span>
      </a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data() {
    return {
      collapsed: false
    }
  },
  methods: {
    jumpPage(item) {
      console.log(item)
      this.$router.push('/admin/' + item.key).catch(err => err)
    }
  }
}
</script>

<style  scoped>
.logo {
  height: 32px;
  margin: 16px;
  background-color: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
