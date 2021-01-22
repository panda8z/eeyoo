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
import Profile from '../components/user/Profile'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/',
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
      { path: 'userlist', component: UserList },
      { path: 'profile', component: Profile }
    ]
  }
]

const router = new VueRouter({
  routes
})

// 配置路由导航守卫
router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!token && to.path === '/') {
    next('/login')
  } else {
    next()
  }
})

export default router
