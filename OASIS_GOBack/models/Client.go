package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketPool struct {
	clients      map[string]*websocket.Conn // 客户端连接池
	clientsMutex sync.Mutex                 // 连接池互斥锁
}
type ChatController struct {
	wsPool *WebSocketPool
}
type messageData struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

func NewChatController() *ChatController {
	return &ChatController{
		wsPool: &WebSocketPool{
			clients: make(map[string]*websocket.Conn),
		},
	}
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type User struct {
	Username string `json:"username"`
}

func (cc *ChatController) getNowArrayNumber() {
	result := make(map[string]interface{})
	array := make([]User, 0)
	for key := range cc.wsPool.clients {
		username := key
		user := User{Username: username}
		array = append(array, user)
	}
	result["users"] = array

	resultJSON, err := json.Marshal(result)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return
	}
	fmt.Println("==============")
	fmt.Println(cc.wsPool.clients)
	fmt.Println("==============")
	cc.BroadcastMessage(resultJSON)
}

func (cc *ChatController) WebSocketHandler(c *gin.Context) {
	user := new(User)
	user.Username = c.Param("username")

	// 升级为 WebSocket 连接
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		// 关闭 WebSocket 连接
		err := ws.Close()
		if err != nil {
			log.Println("Failed to close WebSocket connection:", err)
		}
		// 删除 WebSocket 连接池中的对应连接
		cc.wsPool.clientsMutex.Lock()
		delete(cc.wsPool.clients, user.Username)
		cc.wsPool.clientsMutex.Unlock()
	}()

	// 将连接添加到连接池
	cc.wsPool.clientsMutex.Lock()
	cc.wsPool.clients[user.Username] = ws
	cc.wsPool.clientsMutex.Unlock()
	fmt.Println(cc.wsPool.clients)
	cc.getNowArrayNumber()
	log.Printf("有新用户加入,username={%s}, 当前在线人数为：{%d}\n", user.Username, len(cc.wsPool.clients))

	// 处理 WebSocket 消息
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				// 连接已正常关闭或正在关闭

			} else {
				cc.getNowArrayNumber()
				fmt.Println(websocket.CloseNormalClosure)
				fmt.Println(websocket.CloseGoingAway)
				fmt.Println("other")
			}
			break
		}

		messageData := new(messageData)
		err = json.Unmarshal(p, messageData)
		if err != nil {
			log.Printf("Failed to unmarshal message data: %v", err)
			continue
		}
		if toConn, ok := cc.wsPool.clients[messageData.To]; ok {
			createPrivateChat(user.Username, toConn).SendMessage(*messageData)
		} else {
			log.Printf("User %s is not online", messageData.To)
		}
	}
}

type PrivateChat struct {
	ClientID string
	Conn     *websocket.Conn
}

func (pc *PrivateChat) SendMessage(data interface{}) {
	fmt.Println(data)
	messageData, err := json.Marshal(data)
	fmt.Println(messageData)
	if err != nil {
		log.Printf("Failed to marshal message data: %v", err)
		return
	}
	err = pc.Conn.WriteMessage(websocket.TextMessage, messageData)
	if err != nil {
		log.Printf("Failed to send message to client %s: %v", pc.ClientID, err)
	}
}

// 根据实际情况初始化 PrivateChat 对象
func createPrivateChat(clientID string, conn *websocket.Conn) *PrivateChat {
	return &PrivateChat{
		ClientID: clientID,
		Conn:     conn,
	}
}

func (cc *ChatController) BroadcastMessage(message []byte) {
	// 遍历连接池，并向每个连接发送消息
	cc.wsPool.clientsMutex.Lock()
	defer cc.wsPool.clientsMutex.Unlock()

	for _, client := range cc.wsPool.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Failed to send message to client:", err)
		}
	}
}
