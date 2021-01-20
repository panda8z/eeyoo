<template>
  <v-app app>
    <Header />
    <v-main> <v-container></v-container> </v-main>
    <Footer />
  </v-app>
</template>

<script>
import Header from '../components/Header';
import Footer from '../components/Footer';
export default {
  components: { Header, Footer }
}
</script>

<style scope>
</style>
