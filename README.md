[TOC]
##go-chat
使用Go基于WebSocket的通讯聊天软件。

### 功能列表：
* 群聊天
* 群好友列表
* 单人聊天
* 添加好友
* 文本消息
* 图片消息
* 文件发送
* 语音消息
* 视频消息
* 屏幕共享
* 视频聊天

##后端
[代码仓库](https://github.com/kone-net/go-chat)
go中协程是非常轻量级的。在每个client接入的时候，为每一个client开启一个协程，能够在单机实现更大的并发。同时go的channel，可以非常完美的解耦client接入和消息的转发等操作。

通过go-chat，可以掌握channel的和Select的配合使用，ORM框架的使用，web框架Gin的使用，配置管理，日志操作，还包括proto buffer协议的使用，等一些列项目中常用的技术。


### 后端技术和框架
* web框架Gin
* 长连接WebSocket
* 日志框架Uber的zap
* 配置管理viper
* ORM框架gorm
* 通讯协议Google的proto buffer
* makefile 的编写
* 数据库MySQL
* 图片文件二进制操作

##前端
基于react,UI和基本组件是使用ant design。可以很方便搭建前端界面。

界面选择单页框架可以更加方便写聊天界面，比如像消息提醒，可以在一个界面接受到消息进行提醒，不会因为换页面或者查看其他内容影响消息接受。
[前端代码仓库](https://github.com/kone-net/go-chat-web)：
https://github.com/kone-net/go-chat-web


### 前端技术和框架
* React
* AntDesign
* proto buffer的使用
* WebSocket
* 剪切板的文件读取和操作
* 聊天框发送文字显示底部
* FileReader对文件操作
* ArrayBuffer，Blob，Uint8Array之间的转换
* 获取摄像头视频（mediaDevices）
* 获取麦克风音频（Recorder）
* 获取屏幕共享（mediaDevices）


### 截图
![go-chat-panel](/static/screenshot/go-chat-panel.jpeg)

## 快速运行
### 运行go程序
go环境的基本配置
...
拉取后端代码
```shell
git clone https://github.com/kone-net/go-chat
```

进入目录
```shell
cd go-chat
```

拉取程序所需依赖
```shell
go mod download
```

MySQL创建数据库
```mysql
CREATE DATABASE chat;
```

修改数据库配置文件
```shell
vim config.toml

[mysql]
host = "127.0.0.1"
name = "chat"
password = "root1234"
port = 3306
table_prefix = ""
user = "root"

修改用户名user，密码password等信息。
```

创建表
```shell
将chat.sql里面的sql语句复制到控制台创建对应的表。
```

在user表里面添加初始化用户
```shell
手动添加用户。
```

运行程序
```shell
go run cmd/main.go
```

### 运行前端代码
配置React基本环境，比如nodejs
...

拉取代码
```shell
git clone https://github.com/kone-net/go-chat-web
```

安装前端基本依赖
```shell
npm install
```

如果后端地址或者端口号需要修改
```shell
修改src/common/param/Params.jsx里面的IP_PORT
```

运行前端代码默认启动端口是3000
```shell
npm start
```

访问前端入口
```
http://127.0.0.1:3000/login
```

## 代码结构
```
├── Makefile       代码编译，打包，结构化等操作
├── README.md
├── api
│   └── v1         controller类，对外的接口，如添加好友，查找好友等。所有http请求的入口
├── bin
│   └── chat       打包的二进制文件
├── chat.sql       整个项目的SQL
├── cmd            main函数入口，程序启动
├── common     
│   ├── constant   常量
│   └── util       工具类
├── config         配置初始化类
├── config.toml    配置文件
├── dao
│   └── pool       数据库连接池
├── errors         封装的异常类
├── global
│   └── log        封装的日志类，使用时不会出现第三方的包依赖
├── go.mod
├── go.sum
├── logs           日志文件
├── model          数据库模型，和表一一对应
│   ├── request    请求的实体类
│   ├── response   响应的实体类
├── protocol       消息协议
│   ├── message.pb.go  protoc buffer自动生成的文件
│   └── message.proto  定义的protoc buffer字段
├── response       全局响应，通过http请求的，都包含code，msg，data三个字段
├── router         gin和controller类进行绑定
├── server         WebSocket中消息的接受和转发的主要逻辑
├── service        controller调用的服务类
├── static         静态文件，图片等
│   ├── img
│   └── screenshot markdown用到的截图文件
└── test           测试文件
```

## Makefile
### 程序打包
在根目录下执行make命令
mac
```bash
make build-darwin

实际执行命令是Makefile下的
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/chat cmd/main.go
```

linux
```bash
make build

实际执行命令是Makefile下的
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/chat cmd/main.go
```

### 后端proto文件生成
如果修改了message.proto，就需要重新编译生成对应的go文件。
在根目录下执行
```bash
make proto

实际执行命令是Makefile下的
protoc --gogo_out=. protocol/*.proto
```

如果本地没有安装proto文件，需要先进行安装，不然找不到protoc命令。
使用gogoprotobuf

安装protobuf库文件
```bash
go get github.com/golang/protobuf/proto
```

安装protoc-gen-gogo
```bash
go get github.com/gogo/protobuf/protoc-gen-gogo
```

安装gogoprotobuf库文件
```bash
go get github.com/gogo/protobuf/proto
```

在根目录测试：
```bash
protoc --gogo_out=. protocol/*.proto
```

### 前端proto文件生成
前端需要安装protoc buffer库

```bash
npm install protobufjs
```

生成protoc的js文件到目录
```bash
npx pbjs -t json-module -w commonjs -o src/chat/proto/proto.js  src/chat/proto/*.proto

src/chat/proto/proto.js 是生成的文件的目录路径及其文件名称
src/chat/proto/*.proto  是自己写的字段等
```

## 代码说明
### WebSocket
该文件是gin的路由映射，将普通的get请求，Upgrader为socket连接
```go
// router/router.go
func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	server.Use(Cors())
	server.Use(Recovery)

	socket := RunSocekt

	group := server.Group("")
	{
        ...

		group.GET("/socket.io", socket)
	}
	return server
}
```

这部分对请求进行升级为WebSocket。
* c.Query("user")用户登录后，会获取用户的uuid，在连接到socket时会携带用户的uuid。
* 通过该uuid和connection进行关联。
* server.MyServer.Register <- client将每个client实例，通过channel进行传达，Server实例的Select会对该实例进行保存。
* client.Read()，client.Write()通过协程让每个client对自己独有的channel进行消息的读取和发送
```go
// router/socket.go
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunSocekt(c *gin.Context) {
	user := c.Query("user")
	if user == "" {
		return
	}
	log.Info("newUser", zap.String("newUser", user))
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil) //升级协议为WebSocket
	if err != nil {
		return
	}

	client := &server.Client{
		Name: user,
		Conn: ws,
		Send: make(chan []byte),
	}

	server.MyServer.Register <- client
	go client.Read()
	go client.Write()
}
```

这是Server的三个channel，
* 用户登录后，将用户和connection绑定存放在map中
* 用户离线后，将用户从map中剔除
* 所有消息，每个client将消息获取后放入该channel中，统一在这里进行消息的分发
* 分发消息：
    * 如果是单聊，直接根据前端发送的uuid找到对应的client进行发送。
    * 如果是群聊，需要在数据库查询该群所有的成员，在根据uuid找到对应的client进行发送。
    * 如果消息为普通文本消息，可以直接转发到对应的客户端。
    * 如果消息为视频文件，普通文件，照片之类的，需要先将文件进行保存，然后返回文件名称，前端根据名称调用接口获取文件。
```go
// server/server.go
func (s *Server) Start() {
	log.Info("start server", log.Any("start server", "start server..."))
	for {
		select {
		case conn := <-s.Register:
			log.Info("login", log.Any("login", "new user login in"+conn.Name))
			s.Clients[conn.Name] = conn
			msg := &protocol.Message{
				From:    "System",
				To:      conn.Name,
				Content: "welcome!",
			}
			protoMsg, _ := proto.Marshal(msg)
			conn.Send <- protoMsg

		case conn := <-s.Ungister:
			log.Info("loginout", log.Any("loginout", conn.Name))
			if _, ok := s.Clients[conn.Name]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Name)
			}

		case message := <-s.Broadcast:
			msg := &protocol.Message{}
			proto.Unmarshal(message, msg)
            ...
            ...
		}
	}
}

```

