// main.js 网站的启动开关。它做三件事：创建Vue应用、装上路由、挂载到那个#appdiv上。你可以理解为"按下电源键"。
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './style.css'

const app = createApp(App)
app.use(router)
app.mount('#app')
