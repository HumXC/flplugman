本项目使用 [Vue.js](https://cn.vuejs.org/) 构建。

![Vue.js](https://img.shields.io/badge/vue.js-35495e.svg?style=for-the-badge&logo=vue.js&logoColor=4FC08D)

### 运行

`npm run dev`

> [!WARNING]
> 在运行 `npm run dev` 之前，请先使用 `npm install` 安装依赖。

> [!NOTE]
> 在 `package.json`，`scripts` 下的 `dev` 选项中已默认加入了 `--host 0.0.0.0` 参数，以允许通过 IP 地址访问。
> ```json
>"scripts": {
>   "dev": "vite --host 0.0.0.0",
>   "build": "run-p type-check \"build-only {@}\" --",
>   "preview": "vite preview",
>   "build-only": "vite build",
>   "type-check": "vue-tsc --build --force"
>}
> ```

### 构建

`npm run build`