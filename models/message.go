package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type Message struct {
	gorm.Model //gorm包里的属性
	Name       string
	FormId     int64  //发送者
	TargetId   int64  //接收者
	Type       int    //发送类型
	Media      int    //消息类型
	Content    string //消息内容
	Pic        string
	Url        string
	Desc       string
	Amount     int //其他数字统计
}

func (table *Message) TableName() string {
	return "message" //数据库表名
}

// 客户端连接信息
type Node struct {
	Conn      *websocket.Conn // WebSocket 连接
	DataQueue chan []byte     // 用于发送数据的通道
	GroupSets set.Interface   // 用于存储所属的群组信息的集合
}

// 定义了一个名为 clientMap 的变量，映射，用来存储客户端的信息。
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	/*msgType := query.Get("type")
	targetId := query.Get("targetId")
	context := query.Get("context")*/
	isvalida := true

	//升级连接
	conn, err := (&websocket.Upgrader{
		//权限校验         接收请求  升级连接
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
	}

	//获取conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	//用户关系
	//userId跟node 绑定关系
	//阻塞其他协程的读和写操作
	rwLocker.Lock()
	//便于 用户通过ID 查找对应的连接信息
	//clientMap映射，用来存储客户端的信息。
	clientMap[userId] = node
	rwLocker.Unlock()

	go sendProc(node)

	go recvProc(node)

	sendMsg(userId, []byte("欢迎进入聊天室"))
}

// 别人给我发送消息
func sendProc(node *Node) {
	for {
		select { //用来监听多个通道的消息  监听node.DataQueue
		case data := <-node.DataQueue: //如果有消息可读，就赋值给data
			err := node.Conn.WriteMessage(websocket.TextMessage, data) //将数据发送给websocket  D
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws] <<<<<", data)
	}
}

var upsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	upsendChan <- data
}

func init() {
	go upSendProc()
	go upRecvProc()
}

// 完成udp数据发送协程
func upSendProc() {
	addr := &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 100),
		Port: 3000,
	}
	//协议类型  本机默认nil 由本机自己选择
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		select { //用来监听多个通道的消息  监听node.DataQueue
		case data := <-upsendChan: //如果有消息可读，就赋值给data
			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}

// 完成udp数据接收协程
func upRecvProc() {
	//创建udp监听器  监听0，0，0，0（监听所有可用的网络接口）和端口号为3000的
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n])
	}
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	//将JSON格式的数据解析成对应的Go数据结构
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1:
		sendMsg(msg.TargetId, data)
	}
}
func sendMsg(userID int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userID]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
