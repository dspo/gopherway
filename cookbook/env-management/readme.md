# 环境管理
## 使用代理下载国内不方便下载的资源
可利用代理下载国内不方便下载的包，如[goproxy.io](https://goproxy.io)是一个稳定可信赖的代理站点。  
* Linux/macOS
    ```bash
    export GO111MODULE=ON
    export GOPROXY=https://goproxy.io
    go get -v -u github.com/micro/micro
    ```
* Windows
    ```bash
    set GO111MODULE=ON
    set GOPROXY=https://goproxy.io
    go get -v -u github.com/micro/micro
    ```
    或
    ```bash
    # Enable the go modules feature
    $env:GO111MODULE=on
    # Set the GOPROXY environment variable
    $env:GOPROXY=https://goproxy.io
    go get -v -u github.com/micro/micro
    ```
    或 Go version >= 1.13 时
    ```bash
    go env -w GOPROXY=https://goproxy.io,direct
    # Set environment variable allow bypassing the proxy for selected modules
    go env -w GOPRIVATE=*.corp.example.com
    ```
注意 `go get` 相当于 `git clone` + `go install` ，mod模式下的 `go get` 仍会编译安装二进制文件到 `%GOPATH%/bin` 。  
    
## 使用mod进行环境管理
mod管理环境依赖是在go 1.11中引入的，首先在项目中用`go mod init`新建一个`go.mod`文件，在文件中添加依赖，如
```mod
module gopherway

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/gorilla/websocket v1.4.0
	github.com/hashicorp/consul v1.5.3
	github.com/micro/cli v0.2.0
	github.com/micro/go-api v0.7.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.8.0
	github.com/micro/go-web v1.0.0
	github.com/nats-io/nats-server/v2 v2.0.2 // indirect
	github.com/prometheus/common v0.6.0
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80 // indirect
	google.golang.org/grpc v1.22.1 // indirect
)
```
设置`GO111MODULE=on`，切换到mod模式，`GO111MODULE=off`则为gopath模式。  
`go mod download`会下载所有依赖到`mod/pkg`目录下。  
```bash
> set GO111MODULE=on
> go mod download
```
对于国内不方便下载的内容，可设置 `GOPROXY` 。