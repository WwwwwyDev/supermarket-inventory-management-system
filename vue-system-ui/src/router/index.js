import {
	createRouter,
	createWebHashHistory
} from 'vue-router'
import Login from '../components/login.vue'
import Home from '../components/home.vue'
import Welcome from '../components/welcome.vue'
import Goodsa from '../components/admin/goodsa.vue'
import Pda from '../components/admin/pda.vue'
import Pla from '../components/admin/pla.vue'
import Suppliera from '../components/admin/suppliera.vue'
import Staffa from '../components/admin/staffa.vue'
import Goods from '../components/common/goods.vue'
import Plc from '../components/common/pl.vue'
import Pdc from '../components/common/pd.vue'
import Nf from '../components/noFound.vue'
const routes = [


	{
		path: '/login',
		component: Login
	},
	{
		path: '/home',
		component: Home,
		redirect: '/welcome',
		children: [{
				path: '/welcome',
				component: Welcome,
			},
			{
				path: '/goodsa',
				component: Goodsa
			},
			{
				path: '/pda',
				component: Pda
			},
			{
				path: '/pla',
				component: Pla
			},
			{
				path: '/suppliera',
				component: Suppliera
			},
			{
				path: '/staffa',
				component: Staffa
			},
			{
				path: '/pd',
				component: Pdc
			},
			{
				path: '/pl',
				component: Plc
			},
			{
				path: '/goods',
				component: Goods
			},
		],
	},
	{
		path: "/404",
		component: Nf
	},
	{
		path: "/:catchAll(.*)",
		redirect: "/404"
	},
]

const router = createRouter({
	history: createWebHashHistory(),
	routes
})

// 挂载路由导航守卫
router.beforeEach((to, from, next) => {
	if (to.path == '/login') return next(); // 访问路径为登录
	// 获取flag
	const flagStr = window.sessionStorage.getItem("username"); // session取值
	if (!flagStr) return next('/login'); // 没登录去登录
	next();
})

export default router
