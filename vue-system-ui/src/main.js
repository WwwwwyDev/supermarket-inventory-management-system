import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import installElementPlus from './plugins/element'
import './assets/css/global.css'
import axios from'axios'
// 设置请求的根路径
axios.defaults.baseURL = "http://localhost:9000/"


const app = createApp(App)
//挂载http
app.config.globalProperties.$http=axios
installElementPlus(app)
app.use(router).mount('#app')
