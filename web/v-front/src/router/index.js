import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from '../views/Index'
import ArticleList from '../components/ArticleList'
import Article from '../views/Article'

Vue.use(VueRouter)
const routes = [
  {
    path: '/',
    name: 'index',
    component: Index,
    children: [
      { path: '/', name: 'index.article-list', component: ArticleList, meta: { title: '欢迎来到 eeyoo 博客' } },
      { path: 'article/:id', name: 'article', component: Article }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  routes,
  base: process.env.BASE_URL

})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})

export default router