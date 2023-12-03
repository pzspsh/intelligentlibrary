/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:40:31
*/
package main
import ( 
	"encoding/json" 
	"fmt" 
	"github.com/gorilla/websocket" 
	"log" 
	"net" 
	"net/http"
	 "strconv" 
	 "sync" 
	 "time" 
	 ) 
type Node struct { 
	Conn *websocket.Conn //连接 
	Addr string //客户端地址 
	DataQueue chan []byte //消息 
	UserId string 
	TagId string 
	}

var clients = make(map[*websocket.Conn]bool) 
var broadcast = make(chan Message) 
var upgrader = websocket.Upgrader{} 
var userId int 
var tagId int 

type Message struct { 
	Content string `json:"content"` //消息内容 
	UserId int `json:"userId"` 
	TagUserId int `json:"tagUserId"` 
	Type int `json:"type"` //发送类型 1私聊 2群聊 3心跳 
	Media int `json:"media"` //消息类型 1文字 2表情包 3语音 4图片 /表情包 
	CreateTime uint64 `json:"createTime"` //创建时间 
	ReadTime uint64 `json:"readTime"` //读取时间 
	} 

func main() { 
	http.HandleFunc("/ws", handleWebSocket) 
	log.Println("WebSocket server started on localhost:8000") 
	err := http.ListenAndServe(":8000", nil) 
	if err != nil { log.Fatal("ListenAndServe: ", err) 
	} 
} 

// 读写锁 
var rwLocker sync.RWMutex

// 映射关系 
var clientMapsss map[string]*Node = make(map[string]*Node, 0) 

func handleWebSocket(w http.ResponseWriter, r *http.Request) { 
	query := r.URL.Query() 
	userId := query.Get("userId") 
	
	//userId, _ = strconv.Atoi(get) 
	fmt.Println("userId >>>>>>>", userId) 

	//tagId := query.Get("tagId") 
	// tagId, _ = strconv.Atoi(getTagId) 
	fmt.Println("tagId >>>>>>>", tagId)
	conn, err := upgrader.Upgrade(w, r, nil) 
	if err != nil { 
		log.Println(err) 
		return 
		} 
	node := &Node{ 
		Conn: conn, 
		Addr: conn.RemoteAddr().String(),//客户端地址 
		DataQueue: make(chan []byte, 50), 
		} 
		//go handleMessages(node) 
		//clients[conn] = true 
		//4. userid 跟 node绑定 并加锁 
	rwLocker.Lock() 
	clientMapsss[userId] = node 
	rwLocker.Unlock() 
	go sendProc(node) 
	go recvProc(node) 
	//for { 
		// var msg Message 
		// err := conn.ReadJSON(&msg) 
		// if err != nil { 
			// log.Println(err) 
			// delete(clients, conn) 
			// break 
			// } 
			//
		// broadcast <- msg 
		//}
		//conn.Close() 
} 

func sendProc(node *Node) { 
	for { 
		select { 
		case msgsss := <-node.DataQueue: 
		fmt.Println(msgsss) 
		//for i := range clientMapsss { 
			fmt.Println("sendsendsendsden") 
			//fmt.Println(clientMapsss[i].UserId) 
			//fmt.Println(clientMapsss[i].TagId) 
			fmt.Println("node.TagId", node.TagId) 
			fmt.Println("node.UserId", node.UserId) 
			node.Conn.WriteJSON(msgsss) 
			//err := node.Conn.WriteMessage(websocket.TextMessage, data) 
			//if err != nil { 
				// fmt.Println(err) 
				// return 
				//} 
				} 
				} 
				} 

var udpsendChan chan []byte = make(chan []byte, 1024) 
func broadMsg(data []byte) { 
	udpsendChan <- data 
	}

func init() { 
	go udpSendProc() 
	go udpRecvProc() 
	fmt.Println("init goroutine ") 

} 

// 完成udp数据发送协程 
func udpSendProc() { 
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{ 
		IP: net.IPv4(192, 168, 203, 1),
		 Port: 3000, 
		 }) 
defer con.Close() 
if err != nil { 
	fmt.Println(err) 
	} 
	for {
		select { 
		case data := <-udpsendChan: 
		fmt.Println("udpSendProc data :", string(data)) 
		_, err := con.Write(data) 
		if err != nil { 
			fmt.Println(err) return 
			} 
			} 
			} 
			}

// 完成udp数据接收协程 
func udpRecvProc() { 
	con, err := net.ListenUDP("udp", &net.UDPAddr{ 
		IP: net.IPv4zero, 
		Port: 3000, 
		}) 
	if err != nil { 
		fmt.Println(err) 
		} 
	defer con.Close() 
	for { 
		var buf [512]byte 
		n, err := con.Read(buf[0:]) 
		if err != nil { 
			fmt.Println(err) 
			return 
			} 
			fmt.Println("udpRecvProc data :", string(buf[0:n])) 
			dispatch(buf[0:n]) 
			} 
			} 

func recvProc(node *Node) { 
	for { 
		_, data, err := node.Conn.ReadMessage() 
		if err != nil { 
			fmt.Println(err) 
			return 
			}
		msg := Message{} 
		err = json.Unmarshal(data, &msg) 
		if err != nil { 
			fmt.Println(err) 
			} 
	fmt.Println("recvProcrecvproc") 
	fmt.Println(node.UserId) 
	fmt.Println(node.TagId) 
	fmt.Println(">>>>>>>", msg) 
	//node.Conn.WriteJSON(msg) 
	//node.Conn.WriteJSON() 
	//心跳检测 msg.Media == -1 || msg.Type == 3 
	if msg.Type == 3 { 
		currentTime := uint64(time.Now().Unix()) 
		fmt.Println(currentTime) 
		//node.Heartbeat(currentTime) 
		} else { 
			//dispatch(data) 
			//分发消息存储 
			broadMsg(data) //todo 将消息广播到局域网 
			fmt.Println("[ws] recvProc <<<<< ", string(data)) 
			} 
			//node.Conn.Close() 
			} 
			} // 后端调度逻辑处理 
			

func dispatch(data []byte) { 
	msg := Message{} 
	msg.CreateTime = uint64(time.Now().Unix()) 
	err := json.Unmarshal(data, &msg) 
	if err != nil { 
		fmt.Println(err) 
		return 
		} 
	switch msg.Type { 
		case 1: //私信 
		fmt.Println("dispatch data :", string(data)) 
		sendMsg(msg.TagUserId, data) 
		case 2: 
		//群发 
		//sendGroupMsg(msg.TargetId, data) 
		//发送的群ID ，消息内容 
		// case 4: 
		// 心跳 // node.Heartbeat() 
		//case 4: // 
		} 
		} 
func sendMsg(tagUserId int, msg []byte) { 
	rwLocker.RLock() 
	node := clientMapsss[strconv.FormatInt(int64(tagUserId), 10)]
	rwLocker.RUnlock() 
	jsonMsg := Message{} 
	err2 := json.Unmarshal(msg, &jsonMsg) 
	if err2 != nil { 
		return 
		} 
	err := node.Conn.WriteJSON(jsonMsg) 
	if err != nil { 
		return 
		} 
		//node.DataQueue <- msg 
		//r, err := utils.Red.Get(ctx, "online_"+userIdStr).Result() 
		//if err != nil { 
			// fmt.Println(err) 
			//} 
			//if r != "" { 
				// if ok { 
					// fmt.Println("sendMsg >>> userID: ", userId, " msg:", string(msg)) 
					// node.DataQueue <- msg 
					// } 
					//} 
					//var key string 
					//if userId > jsonMsg.UserId { 
						// key = "msg_" + userIdStr + "_" + targetIdStr 
						//} else { 
							// key = "msg_" + targetIdStr + "_" + userIdStr 
							//} 
							//res, err := utils.Red.ZRevRange(ctx, key, 0, -1).Result() 
							//if err != nil { 
								// fmt.Println(err) 
								//} 
								//score := float64(cap(res)) + 1 
								//ress, e := utils.Red.ZAdd(ctx, key, &redis.Z{score, msg}).Result() //jsonMsg res, e := utils.Red.Do(ctx, "zadd", key, 1, jsonMsg).Result() 
					//备用 后续拓展 记录完整msg 
					//if e != nil { // fmt.Println(e) 
						//} //fmt.Println(ress) 
					}
