import { createApp } from "vue";
import App from "./App.vue";
import "./css/master.css";
import { ScanPluginDB, GetPluginDBPath } from "../wailsjs/go/main/App";
import { LogPrint } from "../wailsjs/runtime/runtime";

createApp(App).mount("#app");
