package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 创建一个 upgrader 对象，用于升级 HTTP 连接为 WebSocket 连接
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 将 HTTP 连接升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}

	defer conn.Close()

	// 处理 WebSocket 会话
	for {
		// 读取客户端发送的消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read message failed:", err)
			break
		}

		// 处理客户端发送的消息
		fmt.Println("Received message: %s", message)

		// 回复客户端消息
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write message failed:", err)
			break
		}
	}
}
