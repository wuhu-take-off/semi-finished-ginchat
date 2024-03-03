package models

//消息

import (
	"fmt"
	"ginchat/utils"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"sync"
)

type Message struct {
	gorm.Model
	FormId   uint   //发送者
	TargetId uint   //接受者
	Type     string //发送类型 群聊,私聊,广播..
	Media    string //消息类型 文字,图片,音频...
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

func InitMessage() {
	err := utils.GetDB().AutoMigrate(Message{})
	if err != nil {
		panic(err)
	}
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

func Chat(write http.ResponseWriter, request http.Request) {
	query := request.URL.Query()
	//校验token
	token := query.Get("token")

	userId, _ := strconv.ParseInt(query.Get("userId"), 10, 64)
	targetId := query.Get("targetId")
	context := query.Get("context")
	msgType := query.Get("type")
	isvalida := true
	conn, err := (&websocket.Upgrader{
		//token校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(write, &request, nil)
	if err != nil {
		panic(err)
	}

	//获取conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	//绑定userid和node
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	//发送消息
	go sendProc(node)
	//接受消息
	go recvProc(node)
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		boradMsg(data)
		fmt.Println("[ws]<<<<", data)
	}
}

func boradMsg(data []byte) {

}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
}
