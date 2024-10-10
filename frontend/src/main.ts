import { createApp } from "vue";
import App from "./App.vue";
import "./css/master.css";
import { ScanPluginDB, GetPluginDBPath } from "../wailsjs/go/main/App";
import { LogPrint } from "../wailsjs/runtime/runtime";
import router from './router'
const app = createApp(App)

app.use(router)

app.mount('#app')