<template>
  <a-layout class="container">
    <!-- 侧边导航 -->
    <Nav></Nav>

    <a-layout>
      <!-- 头部 -->
      <Header></Header>

      <!-- 内容 -->
      <a-layout-content>
        <router-view :key="$router.path"></router-view>
      </a-layout-content>

      <!-- 页脚 -->
      <Footer></Footer>
    </a-layout>
  </a-layout>
</template>

<script>
import Nav from '../components/admin/Nav'
import Header from '../components/admin/Header'
import Footer from '../components/admin/Footer'
export default {
  components: {
    Nav,
    Footer,
    Header
  },
  methods: {
  }
}
</script>

<style scoped>
.container {
  height: 100%;
}
</style>
