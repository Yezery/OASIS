package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// var APIKey = "Bearer sk-sXW9AAjZrbPg2zZbRw2pyMVQjCWVtGD6g5NsIkWByAfjW3si"

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
type User struct {
	Username string `json:"username"`
}

//	type Message struct {
//		Role    string `json:"role"`
//		Content string `json:"content"`
//	}
//
//	type MakePayload struct {
//		Model      string     `json:"model"`
//		Messages   [1]Message `json:"messages"`
//		Max_tokens int64      `json:"max_tokens"`
//		Stream     bool       `json:"stream"`
//	}
type PrivateChat struct {
	ClientID string
	Conn     *websocket.Conn
}

func NewChatController() *ChatController {
	return &ChatController{
		wsPool: &WebSocketPool{
			clients: make(map[string]*websocket.Conn),
		},
	}
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

// func (cc *ChatController) WebSocketHandler(c *gin.Context) {
// 	user := new(User)
// 	user.Username = c.Param("username")

// 	// 升级为 WebSocket 连接
// 	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer func() {
// 		// 关闭 WebSocket 连接
// 		err := ws.Close()
// 		if err != nil {
// 			log.Println("Failed to close WebSocket connection:", err)
// 		}
// 		// 删除 WebSocket 连接池中的对应连接
// 		cc.wsPool.clientsMutex.Lock()
// 		delete(cc.wsPool.clients, user.Username)
// 		cc.wsPool.clientsMutex.Unlock()
// 	}()

// 	// 将连接添加到连接池
// 	cc.wsPool.clientsMutex.Lock()
// 	cc.wsPool.clients[user.Username] = ws
// 	cc.wsPool.clientsMutex.Unlock()
// 	fmt.Println(cc.wsPool.clients)
// 	cc.getNowArrayNumber()
// 	log.Printf("有新用户加入,username={%s}, 当前在线人数为：{%d}\n", user.Username, len(cc.wsPool.clients))

// 	// 处理 WebSocket 消息
// 	for {
// 		_, p, err := ws.ReadMessage()
// 		if err != nil {
// 			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
// 				// 连接已正常关闭或正在关闭

// 			} else {
// 				cc.getNowArrayNumber()
// 				fmt.Println(websocket.CloseNormalClosure)
// 				fmt.Println(websocket.CloseGoingAway)
// 				fmt.Println("other")
// 			}
// 			break
// 		}

//			messageData := new(messageData)
//			err = json.Unmarshal(p, messageData)
//			if err != nil {
//				log.Printf("Failed to unmarshal message data: %v", err)
//				continue
//			}
//			if toConn, ok := cc.wsPool.clients[messageData.To]; ok {
//				createPrivateChat(user.Username, toConn).SendMessage(*messageData)
//			} else {
//				log.Printf("User %s is not online", messageData.To)
//			}
//		}
//	}
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

// func (cc *ChatController) GPTSocketHandler(c *gin.Context) {
// 	user := new(User)
// 	user.Username = c.Param("username")

// 	// 升级为 WebSocket 连接
// 	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer func() {
// 		// 关闭 WebSocket 连接
// 		err := ws.Close()
// 		if err != nil {
// 			log.Println("Failed to close WebSocket connection:", err)
// 		}
// 	}()
// 	// 将连接添加到连接池
// 	cc.wsPool.clientsMutex.Lock()
// 	cc.wsPool.clients[user.Username] = ws
// 	cc.wsPool.clientsMutex.Unlock()
// 	fmt.Println(cc.wsPool.clients)
// 	cc.getNowArrayNumber()
// 	log.Printf("有新用户加入,username={%s}, 当前在线人数为：{%d}\n", user.Username, len(cc.wsPool.clients))
// 	// 处理 WebSocket 消息
// 	for {
// 		_, p, err := ws.ReadMessage()
// 		if err != nil {
// 			fmt.Println(err)
// 			// if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
// 			// 	// 连接已正常关闭或正在关闭

// 			// } else {
// 			// 	cc.getNowArrayNumber()
// 			// 	fmt.Println(websocket.CloseNormalClosure)
// 			// 	fmt.Println(websocket.CloseGoingAway)
// 			// 	fmt.Println("other")
// 			// }
// 			// break
// 		}

// 		messageData := new(messageData)

// 		err = json.Unmarshal(p, messageData)
// 		if err != nil {
// 			log.Printf("Failed to unmarshal message data: %v", err)
// 			continue
// 		}

// 		var messager Message
// 		messager.Content = messageData.Content
// 		messager.Role = messageData.Role
// 		var messageArray [1]Message
// 		messageArray[0] = messager
// 		data := MakePayload{
// 			"gpt-3.5-turbo",
// 			messageArray,
// 			500,
// 			true,
// 		}
// 		fmt.Println(data)
// 		payload, _ := json.MarshalIndent(data, "", " ")
// 		messageData.Content = processMessage(string(payload))
// 		fmt.Println(messageData.Content)
// 		createPrivateChat("GPT", cc.wsPool.clients[user.Username]).SendMessage(*messageData)
// 		// cc.BroadcastMessage()
// 		// var messager Message
// 		// if err := c.BindJSON(&messager); err != nil {
// 		// 	utils.SendResponse(c.Writer, http.StatusBadRequest, err)
// 		// 	panic(err)
// 		// }
// 		// var messageArray [1]Message
// 		// messageArray[0] = messager
// 		// data := MakePayload{
// 		// 	"gpt-3.5-turbo",
// 		// 	messageArray,
// 		// 	500,
// 		// 	true,
// 		// }
// 		// payload, _ := json.MarshalIndent(data, "", " ")

// 		// processMessage(strings.NewReader(string(payload)))

// 		// if err != nil {
// 		// 	log.Printf("Failed to unmarshal message data: %v", err)
// 		// 	continue
// 		// }
// 		// if toConn, ok := cc.wsPool.clients[messageData.To]; ok {
// 		// 	createPrivateChat(user.Username, toConn).SendMessage(*messageData)
// 		// } else {
// 		// 	log.Printf("User %s is not online", messageData.To)
// 		// }
// 	}

// }
// func processMessage(message string) string {

// 	url := "https://api.chatanywhere.com.cn/v1/chat/completions"
// 	method := "POST"
// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, strings.NewReader(message))

// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
// 	req.Header.Add("Authorization", APIKey)
// 	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
// 	req.Header.Add("Content-Type", "application/json")
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
// 	fmt.Println(string(body))
// 	return string(body)
// }

func (pc *PrivateChat) SendMessage(data interface{}) {
	messageData, err := json.Marshal(data)
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
