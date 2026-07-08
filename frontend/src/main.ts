import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import './assets/main.css' // ← WAJIB ADA INI

const app = createApp(App)
app.use(createPinia())
app.mount('#app')