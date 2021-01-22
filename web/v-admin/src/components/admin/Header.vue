<template>
  <a-layout-header class="header">
    <a-button type="danger" @click="logout">退出</a-button>
  </a-layout-header>
</template>

<script>
export default {
  methods: {
    logout() {
      window.sessionStorage.removeItem('token')
      this.$router.push('/login')
    }
  }
}
</script>

<style  scoped>
.header {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
