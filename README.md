[TOC]


## go-chat
使用Go基于WebSocket的通讯聊天软件。

### 功能列表：
* 登录注册
* 修改头像
* 群聊天
* 群好友列表
* 单人聊天
* 添加好友
* 添加群组
* 文本消息
* 剪切板图片
* 图片消息
* 文件发送
* 语音消息
* 视频消息
* 屏幕共享（基于图片）
* 视频通话（基于WebRTC的p2p视频通话）
* 分布式部署（通过kafka全局消息队列，统一消息传递，可以水平扩展系统）

## 后端
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

## 前端
基于react,UI和基本组件是使用ant design。可以很方便搭建前端界面。

界面选择单页框架可以更加方便写聊天界面，比如像消息提醒，可以在一个界面接受到消息进行提醒，不会因为换页面或者查看其他内容影响消息接受。
[前端代码仓库](https://github.com/kone-net/go-chat-web)：
https://github.com/kone-net/go-chat-web


### 前端技术和框架
* React
* Redux状态管理
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
* WebRTC的p2p视频通话


### 截图
* 语音，文字，图片，视频消息
![go-chat-panel](/assets/screenshot/go-chat-panel.jpeg)

* 视频通话
![video-chat](/assets/screenshot/video-chat.png)

* 屏幕共享
![screen-share](/assets/screenshot/screen-share.png)

## 消息协议
### protocol buffer协议
```go
syntax = "proto3";
package protocol;

message Message {
    string avatar = 1;       //头像
    string fromUsername = 2; // 发送消息用户的用户名
    string from = 3;         // 发送消息用户uuid
    string to = 4;           // 发送给对端用户的uuid
    string content = 5;      // 文本消息内容
    int32 contentType = 6;   // 消息内容类型：1.文字 2.普通文件 3.图片 4.音频 5.视频 6.语音聊天 7.视频聊天
    string type = 7;         // 如果是心跳消息，该内容为heatbeat
    int32 messageType = 8;   // 消息类型，1.单聊 2.群聊
    string url = 9;          // 图片，视频，语音的路径
    string fileSuffix = 10;  // 文件后缀，如果通过二进制头不能解析文件后缀，使用该后缀
    bytes file = 11;         // 如果是图片，文件，视频等的二进制
}
```
### 选择协议原因
通过消息体能看出，消息大部分都是字符串或者整型类型。通过json就可以进行传输。那为什么要选择google的protocol buffer进行传输呢？
* 一方面传输快
是因为protobuf序列化后的大小是json的10分之一，是xml格式的20分之一，但是性能却是它们的5~100倍.
* 另一方面支持二进制
当我们看到消息体最后一个字段，是定义的bytes，二进制类型。
我们在传输图片，文件，视频等内容的时候，可以将文件直接通过socket消息进行传输。
当然我们也可以将文件先通过http接口上传后，然后返回路径，再通过socket消息进行传输。但是这样只能实现固定大小文件的传输，如果我们是语音电话，或者视频电话的时候，就不能传输流。

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

进入目录
```shell
cd go-chat-web
```

安装前端基本依赖
```shell
npm install
```

如果后端地址或者端口号需要修改
放在服务器运行时一定需要修改后端地址
```shell
修改src/chat/common/param/Params.jsx里面的IP_PORT
```

运行前端代码默认启动端口是3000
```shell
npm start
```

访问前端入口
```
http://127.0.0.1:3000/login
```

### 分布式部署
* 拉取代码
将代码拉取到服务器，运行make build构建后端代码。
* 构建后端服务镜像
进入目录deployments/docker
通过目录下的Dockerfile构建镜像
```
docker build -t konenet/gochat:1.0 .
```
* 部署服务
需要部署nginx进行反向代理，mysql保存数据，1个或者多个后端服务。
* 在config.toml中配置分布式消息队列
将msgChannelType中的channelType修改为kafka，就为分布式消息队列。需要填写消息队列对应的地址和topic
```toml
appName = "chat_room"

[mysql]
host = "mysql8"
name = "go-chat-message"
password = "thepswdforroot"
port = 3306
tablePrefix = ""
user = "root"

[log]
level = "debug"
path = "logs/chat.log"

[staticPath]
filePath = "web/static/file/"

[msgChannelType]
channelType = "kafka"

kafkaHosts = "kafka:9092"
kafkaTopic = "go-chat-message"
```
* 启动服务
通过deployments/docker下的docker-compose.yml进行启动。
```
docker-compose up -d
```
* 注意：分布式部署后，上传的文件视频等，可能会因为负载到不同的机器上，导致文件找不到的情况，所以需要一个在线或者分布式文件服务器。

## 代码结构
```
├── Makefile             代码编译，打包，结构化等操作
├── README.md
├── api                  controller类，对外的接口，如添加好友，查找好友等。所有http请求的入口
│   └── v1
├── assets
│   └── screenshot       系统使用到的资源，markdown用到的截图文件
├── bin                  打包的二进制文件
├── chat.sql             整个项目的SQL
├── cmd
│   └── main.go          main函数入口，程序启动
├── config
│   └── toml_config.go   系统全局的配置文件配置类
├── config.toml          配置文件
├── deployments
│   └── docker           docker构建镜像，docker-compose.yml等文件
├── go.mod
├── go.sum
├── internal
│   ├── dao              数据库
│   ├── kafka            kafka消费者和生产者
│   ├── model            数据库模型，和表一一对应
│   ├── router           gin和controller类进行绑定
│   ├── server           WebSocket中消息的接受和转发的主要逻辑
│   └── service          调用的服务类
├── logs
├── pkg
│   ├── common           常量,工具类
│   ├── errors           封装的异常类
│   ├── global           封装的日志类，使用时不会出现第三方的包依赖
│   └── protocol         protoc buffer自动生成的文件,定义的protoc buffer字段
├── test
│   └── kafka_test.go
└── web
    └── static           上传的文件等
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

### 剪切板图片上传
上传剪切板的文件，首先我们需要获取剪切板文件。
如以下代码：
* 通过在聊天输入框，绑定粘贴命令，获取粘贴板的内容。
* 我们只获取文件信息，其他文字信息过滤掉。
* 先获取文件的blob格式。
* 通过FileReader，将blob转换为ArrayBuffer格式。
* 将ArrayBuffer内容转换为Uint8Array二进制，放在消息体。
* 通过protobuf将消息转换成对应协议。
* 通过socket进行传输。
* 最后，将本地的图片追加到聊天框里面。
```javascript
bindParse = () => {
        document.getElementById("messageArea").addEventListener("paste", (e) => {
            var data = e.clipboardData
            if (!data.items) {
                return;
            }
            var items = data.items

            if (null == items || items.length <= 0) {
                return;
            }

            let item = items[0]
            if (item.kind !== 'file') {
                return;
            }
            let blob = item.getAsFile()

            let reader = new FileReader()
            reader.readAsArrayBuffer(blob)

            reader.onload = ((e) => {
                let imgData = e.target.result

                // 上传文件必须将ArrayBuffer转换为Uint8Array
                let data = {
                    fromUsername: localStorage.username,
                    from: this.state.fromUser,
                    to: this.state.toUser,
                    messageType: this.state.messageType,
                    content: this.state.value,
                    contentType: 3,
                    file: new Uint8Array(imgData)
                }
                let message = protobuf.lookup("protocol.Message")
                const messagePB = message.create(data)
                socket.send(message.encode(messagePB).finish())

                this.appendImgToPanel(imgData)
            })

        }, false)
    }
```

### 上传录制的视频
上传语音同原理
* 获取视频调用权限。
* 通过mediaDevices获取视频流，或者音频流，或者屏幕分享的视频流。
* this.recorder.start(1000)设定每秒返回一段流。
* 通过MediaRecorder将流转换为二进制，存入dataChunks数组中。
* 松开按钮后，将dataChunks中的数据合成一段二进制。
* 通过FileReader，将blob转换为ArrayBuffer格式。
* 将ArrayBuffer内容转换为Uint8Array二进制，放在消息体。
* 通过protobuf将消息转换成对应协议。
* 通过socket进行传输。
* 最后，将本地的视频，音频追加到聊天框里面。

**特别注意: 获取视频，音频，屏幕分享调用权限，必须是https协议或者是localhost，127.0.0.1 本地IP地址，所有本地测试可以开启几个浏览器，或者分别用这两个本地IP进行2tab测试**
```javascript
/**
     * 当按下按钮时录制视频
     */
    dataChunks = [];
    recorder = null;
    startVideoRecord = (e) => {
        navigator.getUserMedia = navigator.getUserMedia ||
            navigator.webkitGetUserMedia ||
            navigator.mozGetUserMedia ||
            navigator.msGetUserMedia; //获取媒体对象（这里指摄像头）

        let preview = document.getElementById("preview");
        this.setState({
            isRecord: true
        })

        navigator.mediaDevices
            .getUserMedia({
                audio: true,
                video: true,
            }).then((stream) => {
                preview.srcObject = stream;
                this.recorder = new MediaRecorder(stream);

                this.recorder.ondataavailable = (event) => {
                    let data = event.data;
                    this.dataChunks.push(data);
                };
                this.recorder.start(1000);
            });
    }

    /**
     * 松开按钮发送视频到服务器
     * @param {事件} e 
     */
    stopVideoRecord = (e) => {
        this.setState({
            isRecord: false
        })

        let recordedBlob = new Blob(this.dataChunks, { type: "video/webm" });

        let reader = new FileReader()
        reader.readAsArrayBuffer(recordedBlob)

        reader.onload = ((e) => {
            let fileData = e.target.result

            // 上传文件必须将ArrayBuffer转换为Uint8Array
            let data = {
                fromUsername: localStorage.username,
                from: this.state.fromUser,
                to: this.state.toUser,
                messageType: this.state.messageType,
                content: this.state.value,
                contentType: 3,
                file: new Uint8Array(fileData)
            }
            let message = protobuf.lookup("protocol.Message")
            const messagePB = message.create(data)
            socket.send(message.encode(messagePB).finish())
        })

        this.setState({
            comments: [
                ...this.state.comments,
                {
                    author: localStorage.username,
                    avatar: this.state.user.avatar,
                    content: <p><video src={URL.createObjectURL(recordedBlob)} controls autoPlay={false} preload="auto" width='200px' /></p>,
                    datetime: moment().fromNow(),
                },
            ],
        }, () => {
            this.scrollToBottom()
        })
        if (this.recorder) {
            this.recorder.stop()
            this.recorder = null
        }
        let preview = document.getElementById("preview");
        preview.srcObject.getTracks().forEach((track) => track.stop());
        this.dataChunks = []
    }
```
