import { createApp } from 'vue'
import App from './App.vue'
import { setupRouter } from '@/router'
import naive from 'naive-ui'
import 'vfonts/Lato.css'
import { setupStore } from '@/store'
import '@/assets/style/main.scss'

const init = async () => {
  const app = createApp(App)
  app.use(naive)
  setupRouter(app)
  setupStore(app)
  app.mount('#app')
}

init()
