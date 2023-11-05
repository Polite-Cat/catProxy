## catProxy

一个简单、轻量、高效的基于配置的反向代理服务，使用Golang语言编写。

###  使用场景

- 同一IP绑定多个域名的同一端口。
- 负载均衡。

### 特性

1. 支持同一端口多个代理。
2. 支持多个端口。

### 开始使用

- 改写`main.go`修改配置文件位置

  ```go
  line38:
  err := confbox.Load("I:\\ReverseProxy\\test.yaml", &list)
  ```

- 编译

  ```go
  go build main.go
  ```

- 修改配置文件

  ```yaml
  point_lists:
      - port: 80
        points:
          local.polite.cat: http://localhost:3001
          local.simon-app.cn: http://localhost:3001
      - port: 8080
        points:
          local.polite.cat: http://localhost:3002
          local.simon-app.cn: http://localhost:3002
  ```

  `介绍：`将local.polite.cat与local.simon-app.cn的80端口转发到localhost的3001端口；将local.polite.cat与local.simon-app.cn的8080端口转发到localhost:3002端口；

