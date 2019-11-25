package gateway

import (
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"strconv"
	"time"
)

// 	WebSocket服务端
type WSServer struct {
	server *http.Server
	curConnId uint64
}

type CheckPacket struct {
	mid string
	sKey string

}

var (
	G_wsServer *WSServer

	wsUpgrader = websocket.Upgrader{
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func handleConnect(resp http.ResponseWriter, req *http.Request) {
	var (
		err error
		wsSocket *websocket.Conn
		connId uint64
		wsConn *WSConnection
	)

	// WebSocket握手
	if wsSocket, err = wsUpgrader.Upgrade(resp, req, nil); err != nil {
		return
	}

	// 首次校验
	cp, err := firstCheckPacket(wsSocket)
	if err != nil {
		_ = wsSocket.Close()
		return
	}

	// 连接唯一标识
	//connId = atomic.AddUint64(&G_wsServer.curConnId, 1)

	// 初始化WebSocket的读写协程
	wsConn = InitWSConnection(cp, wsSocket)

	// 开始处理websocket消息
	wsConn.WSHandle()
}

func firstCheckPacket(wsSocket *websocket.Conn) (cp CheckPacket, err error) {
	// 1. 从socket中读取消息
	// 2. 解密
	// 3. 解析出mid
	//
	// 二阶段登陆认证：
	// 1. # 使用xor解密数据
	// 上行数据
	// v=版本号
	// mid=9621c58b61358fe082579b14d58d2486
	// m2=zsf
	// crc32

	// 下行数据
	// skey=随机密钥64位编码 // 每个
	// errno=错误码
	// crc32

	// 2
	// 上行
	// rsa 非对称加密 [公钥加密，私钥解密]
	// aes + skey 对称加密上一步密文

	// 下行
	// 下行数据直接使用rsa非对称加密 [私钥加密，公钥解密


	return
}

func InitWSServer() (err error) {
	var (
		mux *http.ServeMux
		server *http.Server
		listener net.Listener
	)

	// 路由
	mux = http.NewServeMux()
	mux.HandleFunc("/connect", handleConnect)

	// HTTP服务
	server = &http.Server{
		ReadTimeout: time.Duration(G_config.WsReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(G_config.WsWriteTimeout) * time.Millisecond,
		Handler: mux,
	}

	// 监听端口
	if listener, err = net.Listen("tcp", ":" + strconv.Itoa(G_config.WsPort)); err != nil {
		return
	}

	// 赋值全局变量
	G_wsServer = &WSServer{
		server: server,
		curConnId: uint64(time.Now().Unix()),
	}

	// 拉起服务
	go server.Serve(listener)

	return
}