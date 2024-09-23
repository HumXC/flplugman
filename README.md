## 开发

### 准备工作

1.  安装 Go (1.18+), NPM (Node 18.3+)
   
> [!NOTE]
> 配置 Go 镜像(可选): 
> 
> `go env -w  GOPROXY=https://goproxy.cn,direct`

1.  安装 Wails Cli

    ```shell
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    ```
> [!NOTE]
> `.gitignore` 忽略了 wails 生成的绑定代码 (frontend/wailsjs)，所以初次克隆时，前端代码会出现找不到部分导入的情况。可以手动执行 `wails generate module` 生成 `wailsjs` 目录，或者运行以下任意一条命令会触发 `wailsjs` 的生成。

调试: `wails dev`

构建: `wails build`
