import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'node:path'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(process.cwd(), 'src'),
    },
  },
  server: {
    host: '0.0.0.0',
    port: 31098,
  },
})


//开发工具的配置文件，不是网站本身的代码。
//Vite是一个构建工具，类似于一个"翻译机"——因为浏览器看不懂.vue文件，Vite负责把它翻译成浏览器能懂的普通HTML/CSS/JS。
//这个文件告诉Vite：端口用31098、@符号代表src目录等配置

