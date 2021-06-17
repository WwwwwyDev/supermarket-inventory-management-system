import { createRouter, createWebHashHistory } from 'vue-router'
import Login from '../components/Login.vue'
import Home from '../components/home.vue'

const routes = [
	
	{
		path: '/login',
		component: Login
	},
	{
		path: '/home',
		component: Home
	},
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

// 挂载路由导航守卫
router.beforeEach((to, from, next) => {
  if (to.path == '/login') return next();// 访问路径为登录
  // 获取flag
  const flagStr = window.sessionStorage.getItem("username");// session取值
  if (!flagStr) return next('/login');// 没登录去登录
  if (flagStr && to.path == '/') return next('/home');// 登录去home
  next();
})

export default router
