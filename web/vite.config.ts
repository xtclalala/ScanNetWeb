import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import VueSetupExtend from 'vite-plugin-vue-setup-extend'
import { resolve } from 'path'

const NODE_ENV = process.env.VITE_USER_NODE_ENV || 'development'
const config = loadEnv(NODE_ENV, './')
const wsServiceUrl = config.VITE_GLOB_SERVICE_URL.replace('http', 'ws').replace('https', 'wss')
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), VueSetupExtend()],
  resolve: {
    alias: {
      // 配置导包使用的快捷方式
      '@': resolve('./src'),
      '@r': resolve('./src/router'),
      '#': resolve('./types'),
    },
  },
  base: './',
  server: {
    // 前端端口
    port: 3000,
    host: '0.0.0.0',
    proxy: {
      [config.VITE_GLOB_API_URL_PREFIX]: {
        target: config.VITE_GLOB_SERVICE_URL,
        changeOrigin: true,
        rewrite: (path) => path.replace(new RegExp('^' + config.VITE_GLOB_API_URL_PREFIX), ''),
      },
      [config.VITE_GLOB_WS_URL_PREFIX]: {
        target: wsServiceUrl,
        changeOrigin: true,
      },
    },
  },
  css: {
    preprocessorOptions: {
      less: {
        // 支持内联 JavaScript
        javascriptEnabled: true,
      },
    },
  },

  build: {
    assetsDir: './',
    minify: 'terser',
    terserOptions: {
      compress: {
        // eslint-disable-next-line camelcase
        drop_console: true,
        // eslint-disable-next-line camelcase
        drop_debugger: true,
      },
    },
  },
})
