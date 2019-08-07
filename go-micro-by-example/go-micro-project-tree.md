# 一个典型的全micro生态的项目结构
>本处引用的项目是[micro-in-cn/tutorials/microservice-in-micro](https://github.com/micro-in-cn/tutorials/tree/master/microservice-in-micro)

1. ## 项目架构图
    ![项目架构-服务抽象结构](/img/design.png)

0. ## 结构树
```
gopath/company/project
│  README.md:"
│
===============================================================================================
├─docs                              - 持久化文本
│      schema.sql
│  
===============================================================================================              
├─basic >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> - 基础模块 <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
│  │  basic.go                      - 提供初始化接口 
│  │  options.go
│  │  
│  ├─common                         - 公共设置 其他服务可以从此处获取公共设置
│  │      app.go                    - 提供获取配置的接口
│  │      const.go                  - 定义常量
│  │      consul.go                 - 提供获取consul配置的接口
│  │      
│  └─config
│          config.go                - 提供获取配置器的接口
│          options.go
│
===============================================================================================          
├─config-grpc-srv >>>>>>>>>>>>>>>>>>>>>>>>>>>>> - micro配置 <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
│  │  main.go                       - 基于tcp的grpc服务 加载、监听配置文件 
│  │                                    其他服务请求配置信息
│  │  
│  └─conf                           - micro持久化配置
│          micro.yml                - 配置文件
│
===============================================================================================          
├─plugins >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> - 插件 <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
│  ├─breaker                        - 容错：中断
│  │  │  breaker.go
│  │  │  
│  │  └─http
│  │          http.go
│  │          
│  ├─db                             - 实现数据库连接池
│  │      db.go
│  │      mysql.go
│  │      
│  ├─jwt                            - json web tokens
│  │      jwt.go                    - jwt 配置接口
│  │      
│  ├─redis                          - redis
│  │      redis.go                  - redis 配置接口
│  │      
│  ├─session                        - session
│  │      gorilla.go                - session 接口
│  │      
│  └─tracer                         - 链路跟踪
│      ├─jaeger
│      │      jaeger.go
│      │      
│      └─opentracing
│          ├─std2micro
│          │      std2micro.go
│          │      
│          └─stdhttp
│                  stdhttp.go
│
===============================================================================================
├─auth >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> - 服务层：鉴权 <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
│  │  Dockerfile
│  │  main.go                       - 从配置中心获取配置 初始化配置 新建服务 
│  │                                    向服务治理中心暴露命名空间 注册服务 启动服务
│  │  Makefile
│  │  plugin.go
│  │  README.md
│  │  
│  ├─handler                        - 处理器
│  │      auth.go                   - 提供服务接口：鉴权
│  │      
│  ├─model                          - 模型
│  │  │  model.go                   - 初始化各模型
│  │  │  
│  │  └─access                      - 一个具体的模型 模型要实现proto/* 下的service interface
│  │          access.go             - 定义服务interface、struct，但不在此处实现insterface 
│  │                                    提供获取服务的方法
│  │          access_internal.go    - 为服务实现一些私有方法
│  │          access_token.go       - 实现proto/auth中的services interfaces
│  │          
│  └─proto                          - protobuf协议的go实现
│      └─auth                       - 一个基于protobuf协议的模型
│              auth.micro.go        - 提供基于protobuf的接口
│              auth.pb.go           - 提供消息体及其方法
│              auth.proto           - 定义protobuf服务模型（service）、
│                                       消息体（message）、接口服务（rpc）
│
===============================================================================================      
├─inventory-srv >>>>>>>>>>>>>>>>>>>>>>>>>>>>>> - 服务层：inventory <<<<<<<<<<<<<<<<<<<<<<<<<<<<<
│  │  Dockerfile
│  │  main.go
│  │  Makefile
│  │  plugin.go
│  │  README.md
│  │  
│  ├─handler
│  │      inventory.go
│  │      
│  ├─model
│  │  │  model.go
│  │  │  
│  │  └─inventory
│  │          inventory.go
│  │          inventory_post.go
│  │          
│  └─proto
│      └─inventory
│              inventory.micro.go
│              inventory.pb.go
│              inventory.proto
│
===============================================================================================      
├─orders-srv >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> - 服务层：orders <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
│  │  Dockerfile
│  │  main.go
│  │  Makefile
│  │  orders-srv
│  │  plugin.go
│  │  README.md
│  │  
│  ├─handler
│  │      orders.go
│  │      
│  ├─model
│  │  │  model.go
│  │  │  
│  │  └─orders
│  │          orders.go
│  │          orders_get.go
│  │          orders_post.go
│  │          
│  ├─proto
│  │  └─orders
│  │          orders.micro.go
│  │          orders.pb.go
│  │          orders.proto
│  │          
│  └─subscriber
│          pay.go                           - 广播订单状态
│
===============================================================================================          
├─orders-web >>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Web层：orders <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< 
│  │  Dockerfile
│  │  main.go
│  │  Makefile
│  │  plugin.go
│  │  README.md
│  │  
│  └─handler
│          handler.go
│          wrapper.go
│
===============================================================================================          
├─payment-srv >>>>>>>>>>>>>>>>>>>>>>>>>> - 服务层：payment <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
│  │  Dockerfile                         
│  │  main.go
│  │  Makefile
│  │  plugin.go
│  │  README.md
│  │  
│  ├─handler
│  │      payment.go
│  │      
│  ├─model
│  │  │  model.go
│  │  │  
│  │  └─payment
│  │          payment.go
│  │          payment_evt.go
│  │          payment_post.go
│  │          
│  └─proto
│      └─payment
│              payment.micro.go
│              payment.pb.go
│              payment.proto
│
===============================================================================================              
├─payment-web >>>>>>>>>>>>>>>>>>>>>>>>>> - Web层：payment <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
│  │  Dockerfile
│  │  main.go
│  │  Makefile
│  │  plugin.go
│  │  README.md
│  │  
│  └─handler
│          handler.go
│          wrapper.go
│
===============================================================================================
├─user-srv ======================== - 服务层-用户 负责内部逻辑交互 向Web提供RPC服务 ================
│  │  Dockerfile
│  │  main.go                   - 新建微服务 注册处理器 向服务治理中心暴露微服务
│  │  Makefile
│  │  plugin.go
│  │  README.md
│  │  
│  ├─handler                    - 具体的处理器 与model配合使用 model定义模型，handler实现方法
│  │      service.go            - 具体的处理器 实现相应模型的interface，会调用相应的业务接口
│  │                                此处handler是微服务handler，不是go原生HTTPhandler
│  │      
│  ├─model                      - 模型
│  │  │  model.go               - 初始化目录下所有模型
│  │  │  
│  │  └─user                    - 一个具体的模型：user
│  │          user.go           - 定义模型interface、struct，但struct没有在此处实现interface
│  │          user_get.go       - 具体的业务接口 handler会调用业务接口 
│  │                                数据库操作、算法操作、本地IO、外部资源
│  │          
│  └─proto                      - protobuf协议的go实现
│      └─user                   - 基于protobuf协议的微服务模型
│              user.micro.go    - 提供基于protobuf的micro生态的服务接口（含服务端、客户端） 
│                                   由protoc-gen-micro生成
│              user.pb.go       - 微服务模型（struct）、消息体（struct）、服务接口（method） 
│                                   由protoc-gen-go生成
│              user.proto       - 定义protobuf协议模型（service）、消息体（message）、接口服务（rpc）
│
===============================================================================================              
├─user-web ======================== - Web层-用户 暴露HTTP接口给外部（如前端） =====================
│  │  Dockerfile
│  │  main.go                   - 新建micro web service（兼容go原生server mux） 
│  │                                注册HTTP handler 向外部暴露暴露http接口
│  │  Makefile
│  │  plugin.go
│  │  README.md
│  │  
│  └─handler                    - 处理器
│          handler.go           - 具体的处理器（原生HTTP handler）
│                                   建立微服务客户端，利用目标微服务命名空间向服务治理中心调用微服务
│          
└─utils
│   └─slices
│           slice.go
│           slice_test.go
===============================================================================================              
├─micro
       main.go
```