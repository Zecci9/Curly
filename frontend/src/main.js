import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './assets/base.css'
import './assets/app.css'
import { installRegisteredPlugins } from './plugins/registry'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
installRegisteredPlugins(app, { router, pinia })

app.mount('#app')
