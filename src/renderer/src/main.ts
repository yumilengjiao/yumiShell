//清除默认样式
import './styles/reset.scss'
//引入dark主题
import '@opentiny/vue-theme/dark-theme-index.css'
//引入base主题
import './styles/base.scss'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'

const pinia = createPinia()
const app = createApp(App)
app.use(pinia)
app.mount('#app')
