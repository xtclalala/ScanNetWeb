package ws

import (
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"sync"
	"time"
)

type WsMessage struct {
	Title   string `json:"title"`
	Context string `json:"context"`
	State   string `json:"state"`
	// 任务对应的前端路由
	Task string `json:"task"`
}

var (
	// wsMessageCh 消息通道
	wsMessageCh = make(chan WsMessage)
	// websocket客户端链接池
	client *websocket.Conn = nil
	// 互斥锁，防止程序对统一资源同时进行读写
	mux sync.Mutex
)

// websocket Upgrader
var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func DeleteClient() error {
	// 关闭websocket链接
	conn, exist := getClient()
	if exist {
		conn.Close()
		deleteClient()
		return nil
	} else {
		return errors.New("not found client")
	}
}

// WsHandler 处理ws请求
func WsHandler(w http.ResponseWriter, r *http.Request) {
	// 创建一个定时器用于服务端心跳
	pingTicker := time.NewTicker(time.Second * 60)
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 把与客户端的链接添加到客户端链接池中
	_ = addClient(conn)
	// 设置客户端关闭ws链接回调函数
	conn.SetCloseHandler(func(code int, text string) error {
		deleteClient()
		return nil
	})

	go ReadMessage(conn)
	for {
		select {
		case message := <-wsMessageCh:
			// 从消息通道接收消息，然后推送给前端
			err = conn.WriteJSON(message)
			if err != nil {
				log.Println(err)
				conn.Close()
				deleteClient()
				return
			}
		case <-pingTicker.C:
			// 服务端心跳:每60秒ping一次客户端，查看其是否在线
			conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
			err = conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				log.Println("send ping err:", err)
				conn.Close()
				deleteClient()
				return
			}
		}
	}
}

func ReadMessage(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		print(string(msg))
		if err != nil {
			conn.Close()
			deleteClient()
			return
		}
	}

}

// 将客户端添加到客户端链接池
func addClient(conn *websocket.Conn) error {
	mux.Lock()
	if client != nil {
		mux.Unlock()
		return errors.New("client was exist")
	}
	client = conn
	mux.Unlock()
	return nil
}

// 获取指定客户端链接
func getClient() (conn *websocket.Conn, exist bool) {
	if client == nil {
		return nil, false
	}
	return client, true
}

// 删除客户端链接
func deleteClient() {
	client = nil
}

func PushMessage(data WsMessage) {
	if _, exist := getClient(); !exist {
		return
	}
	wsMessageCh <- data
}

func NewMessage(title, context, state, task string) WsMessage {
	return WsMessage{
		Title:   title,
		Context: context,
		State:   state,
		Task:    task,
	}
}
