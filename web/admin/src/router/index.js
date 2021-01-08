import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login'
import Admin from '../views/Admin'

// pages
import Index from '../components/admin/Index'
import ArticleList from '../components/article/ArticleList'
import WriteArticle from '../components/article/WriteArticle'
import CategoryList from '../components/category/CategoryList'
import UserList from '../components/user/UserList'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/admin',
    name: 'admin',
    component: Admin,
    children: [
      { path: 'index', component: Index },
      { path: 'articlelist', component: ArticleList },
      {
        path: 'writearticle',
        component: WriteArticle
      },
      {
        path: 'writearticle/:id',
        component: WriteArticle,
        props: true
      },
      { path: 'catelist', component: CategoryList },
      { path: 'userlist', component: UserList }
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router
