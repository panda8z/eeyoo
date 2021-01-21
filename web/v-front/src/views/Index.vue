<template>
  <v-app app>
    <Header />
    <v-main class="grey lighten-3">
      <v-container>
        <v-row>
          <!-- 导航栏 -->
          <v-col cols="4">
            <Nav></Nav>
          </v-col>
          <!-- 文章列表 -->
          <v-col>
            <v-sheet min-height="80vh" rounded="lg">
              <router-view></router-view>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
    <Footer />
  </v-app>
</template>

<script>
import Header from '../components/Header';
import Footer from '../components/Footer';
import Nav from '../components/Nav';
export default {
  components: { Header, Footer, Nav }
}
</script>

<style scope>
</style>
