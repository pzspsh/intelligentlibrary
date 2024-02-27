```go
package main

import (
	"fmt"
	"math"
	"net"
	"strings"
	"time"
)

const (
	StatusStart = 1 // 启动动中
	StatusRun   = 2 // 运行中
	StatusStop  = 3 // 已关闭
)

type Service struct {
	Name    string
	Addr    string
	Port    uint32
	Timeout time.Duration
	Status  byte //1:启动，2：运行，3：关闭
	UdpConn *net.UDPConn
	//Router  udp.IRouter
	msgHandle iudp.IMsgHandle
	ExitChan  chan bool

	MaxWorkerTaskLen uint32               // 队列最大任务数
	WorkerPoolSize   uint32               //业务工作Worker池的数量
	TaskQueue        []chan iudp.IRequest //Worker负责取任务的消息队列
}

/**
udp server 初始化
*/
func NewServer(name, addr string, port uint32, timeout time.Duration) iudp.IServer {
	if strings.TrimSpace(name) == "" {
		name = "SPA Serve"
	}
	poolSize := config.Cfg.Spa.WorkPoolSize
	if poolSize <= 0 {
		poolSize = 10
	}
	taskLen := config.Cfg.Spa.WorkTaskLen
	if taskLen <= 0 {
		taskLen = 1024
	}
	return &Service{
		Name:             name,
		Addr:             addr,
		Port:             port,
		Timeout:          timeout,
		Status:           StatusStart,
		UdpConn:          nil,
		ExitChan:         make(chan bool),
		msgHandle:        NewMsgHandle(),
		MaxWorkerTaskLen: taskLen,
		WorkerPoolSize:   poolSize,
		TaskQueue:        make([]chan iudp.IRequest, taskLen),
	}
}

//将消息交给TaskQueue,由worker进行处理
func (s *Service) sendMsgToTaskQueue(request iudp.IRequest) {
	//根据requestId来分配当前的连接应该由哪个worker负责处理
	//轮询的平均分配法则
	//得到需要处理此条连接的workerID
	workerID := request.GetId() % s.WorkerPoolSize
	log.Printf("Add Request id:%d to worker id:%d", request.GetId(), workerID)
	//将请求消息发送给任务队列
	s.TaskQueue[workerID] <- request
}

//启动服务
func (s *Service) Start() {
	address := fmt.Sprintf("%s:%d", s.Addr, s.Port)
	log.Printf("[start] %s :%s.", s.Name, address)
	var addr *net.UDPAddr
	var err error
	if addr, err = net.ResolveUDPAddr("udp", address); err != nil {
		log.Fatalf("udp addr resolve err:%s", err)
	}
	s.UdpConn, err = net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("udp spa server listen err:%s", err)
	}

	go s.startReader()

	s.Status = StatusRun
}

//启动业务服务
func (s *Service) startReader() {
	var reqId uint32
	data := pool.BufferSmallPool.Get().([]byte) //make([]byte, 1<<16)
	defer pool.BufferSmallPool.Put(data)
	for {
		n, cAddr, err := s.UdpConn.ReadFromUDP(data)
		if err != nil {
			if s.Status == StatusStop {
				return
			}
			log.Errorf("reader from udp err:%s", err)
			continue
		}
		reqId++
		if reqId >= math.MaxInt32 {
			reqId = 1
		}
		log.Printf("udp server reader:%s,%d", string(data[:n]), len(data[:n]))
		dp := NewDataPack()
		msg, err := dp.Unpack(data[:n])
		if err != nil {
			log.Errorf("client:%s udp data of unpack err:%s", cAddr.String(), err)
			continue
		}
		req := NewRequest(reqId, s.UdpConn, cAddr, msg)
		//go s.msgHandle.DoMsgHandler(req)
		s.sendMsgToTaskQueue(req)
	}
}

//启动一个Worker工作流程
func (s *Service) startOneWorker(workerID int, taskQueue chan iudp.IRequest) {
	//log.Printf("%s Worker ID:%d  is started.", s.Name, workerID)
	//不断的等待队列中的消息
	for {
		select {
		//有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			s.msgHandle.DoMsgHandler(request)
		case <-s.ExitChan:
			//log.Printf("%s service stopped", s.Name)
			return
		}
	}
}

//启动worker工作池
func (s *Service) StartWorkerPool() {
	//遍历需要启动worker的数量，依此启动
	go func() {
		for i := 0; i < int(s.WorkerPoolSize); i++ {
			//一个worker被启动
			//给当前worker对应的任务队列开辟空间
			s.TaskQueue[i] = make(chan iudp.IRequest, s.MaxWorkerTaskLen)
			//启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
			go s.startOneWorker(i, s.TaskQueue[i])
		}
	}()
}

//启动业务服务
func (s *Service) Serve() {
	s.StartWorkerPool()
	s.Start()
}

//停止服务
func (s *Service) Stop() {
	if s.Status == StatusStop {
		return
	}
	s.Status = StatusStop
	//通知读业务，此连接已关闭
	s.ExitChan <- true
	// 关闭通道
	close(s.ExitChan)
	//关闭连接
	s.UdpConn.Close()
	log.Printf("%s stopped.", s.Name)
}

//添加路由处理
func (s *Service) AddRouter(msgId uint32, router iudp.IRouter) {
	//s.Router = router
	s.msgHandle.AddRouter(msgId, router)
}

func (s *Service) AddPEP(pep ipolicyengine.IService) {
	s.msgHandle.AddPEP(pep)
}

```


client.go
```go

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"
	
)


type TLServe struct {
	Name    string
	IP      string
	Port    uint32
	Timeout time.Duration
}

type Connection struct {
	//1. tls 服务器参数
	TLServe
	
	Conn *tls.Conn
	//4.读channel
	readChan map[uint32]chan itls.IResponse
	//5. conn 是否已关闭
	isClosed    bool
	isConnected chan bool
	
	//7.服务是否退出
	isExit   bool
	exitChan chan bool
	
	l        sync.Mutex
}

/**

 */
func NewConnection(name, ip string, port uint32, timeout time.Duration) itls.IConnection {
	return &Connection{
		TLServe: TLServe{Name: name, IP: ip, Port: port, Timeout: timeout},
		Conn:        nil,
		readChan:    make(map[uint32]chan itls.IResponse),
		spa:         nil,
		isClosed:    true,
		isConnected: make(chan bool, 1),
		isExit:      false,
		exitChan:    make(chan bool, 1),
	}
}


/**
tls 连接
需要加互斥锁 ，防止其他应用请求进来 重复处理 用户数据同步
*/
func (c *Connection) connect(message itls.IMessage) error {
	c.l.Lock()
	defer c.l.Unlock()
	if c.isClosed == false {
		return nil
	}
	var tryCnt int = 3
SPAStart:
	if err := c.knock(); err != nil {
		time.Sleep(time.Second * time.Duration(3-tryCnt))
		if tryCnt > 0 {
			tryCnt--
			goto SPAStart
		}
		log.Errorf("spa has knock 3 times but failed,please check network situation:%s", err)
		return err
	}

	address := fmt.Sprintf("%s:%d", c.IP, c.Port)
	var err error
	tryCnt = 3
TLSStart:
	c.Conn, err = tls.Dial("tcp", address, base.ClientTLSConfig())
	if err != nil {
		time.Sleep(time.Millisecond * 500 * time.Duration(3-tryCnt))
		if tryCnt > 0 {
			tryCnt--
			goto TLSStart
		}
		log.Errorf("client tls  dial  failed:%s", err)
		return err
	}
	c.isConnected <- true
	if c.GetHasLogin() && message.GetMsgId() != TypeControllerLogin {
		data, err := c.userInfoPack(c)
		if err != nil {
			log.Errorf("client tls  data pack  failed:%s", err)
			return err
		}
		typ := uint32(TypeControllerRegion)
		msg := NewMessage(typ, data)
		c.AddRouter(typ)
		wRes := c.write(msg)
		if wRes.GetCode() != 200 {
			return fmt.Errorf(wRes.GetMsg())
		}
		res := <-c.readChan[msg.GetMsgId()]
		if res.GetCode() != 200 {
			log.Errorf("client tls  user info sync  failed:%s", res.GetMsg())
			return fmt.Errorf(res.GetMsg())
		}
		c.SetLogin(true)
	}
	c.isClosed = false
	log.Debug("tls tunnel:%s established.", c.Conn.RemoteAddr())
	return nil
}



/**

 */
func (c *Connection) GetName() string {
	return c.Name
}

/**
路由添加
*/
func (c *Connection) AddRouter(typ uint32) {
	if _, ok := c.readChan[typ]; ok {
		return
	}
	c.readChan[typ] = make(chan itls.IResponse)
}


/**
客户端请求流量
*/
func (c *Connection) write(message itls.IMessage) itls.IResponse {
	dp := NewDataPack()
	reqMsg, err := dp.Pack(message)
	if err != nil {
		return c.responseBad(fmt.Sprintf("pack message id:%d  pack err:%s", message.GetMsgId(), err), message)
	}

	if err := c.Conn.SetWriteDeadline(time.Now().Add(time.Second * 5)); err != nil {
		log.Errorf("client tls  set write dead line err:%s", err)
		return c.responseBad(fmt.Sprintf("message id:%d  set write dead line  err:%s", message.GetMsgId(), err), message)
	}
	if _, err := c.Conn.Write(reqMsg); err != nil {
		log.Errorf("socket writing data err:%s", err)
		c.disconnect()
		return c.responseBad(fmt.Sprintf("message id:%d  write to client err:%s", message.GetMsgId(), err), message)
	}
	return c.responseOk(message)
}

/**
服务端响应流量
*/
func (c *Connection) read() itls.IResponse {

	//bufConn := bufio.NewReader(c.Conn)
	dp := NewDataPack()
	headBuf := pool.BufferDelimiterPool.Get().([]byte)
	defer pool.BufferDelimiterPool.Put(headBuf)
	//headBuf := make([]byte, dp.GetHeadLen())
	//_, err := bufConn.Read(headBuf)
	_, err := c.Conn.Read(headBuf)
	if err != nil {
		log.Errorf("socket read head err:%s", err)
		c.disconnect()
		return c.responseBad(err.Error(), nil)
	}

	msg, err := dp.Unpack(headBuf)
	if err != nil {
		return c.responseBad(fmt.Sprintf("conn socket header unpack  err:%s ", err), nil)
	}

	var body []byte
	if msg.GetDataLen() > 0 { //拿到数据部分 读取字节长度
		bufConn := bufio.NewReaderSize(c.Conn, int(msg.GetDataLen()))

		body = make([]byte, msg.GetDataLen())
		n, err := bufConn.Read(body)
		if err != nil {
			log.Errorf("socket read body err:%s", err)
			c.disconnect()
			return c.responseBad(err.Error(), nil)
		}
		if uint32(n) < msg.GetDataLen() {
			a := n
			y := 0
			for i := 0; i < 10000; i++ {
				y, err = bufConn.Read(body[a:])
				if err != nil {
					log.Errorf("socket read body err:%s", err)
					c.disconnect()
					return c.responseBad(err.Error(), nil)
				}
				a = a + y
				if msg.GetDataLen() <= uint32(a) {
					break
				}
			}
		}
	}
	//拿到数据部分
	msg.SetData(body)
	return c.responseOk(msg)
}

/**
客户端请求
考虑互斥锁
*/
func (c *Connection) Request(message itls.IMessage) (response itls.IResponse) {
	//1.判断隧道是否可用
	if !c.Available() {
		if err := c.connect(message); err != nil {
			return c.response(408, "网络请求超时！", message)
		}
	}

	//2.
	wRes := c.write(message)
	if wRes.GetCode() != 200 {
		return c.responseBad(fmt.Sprintf("connect controller err:%s", wRes.GetMsg()), message)
	}
	timeout, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	select {
	//2.读取消息头
	case r := <-c.readChan[message.GetMsgId()]:
		return r
	case <-timeout.Done():
		return c.response(408, "网络请求超时", message)
	}

}

/**
客户端响应
*/
func (c *Connection) responseOk(data itls.IMessage) (response itls.IResponse) {
	res := &Response{}
	res.SetCode(200)
	res.SetMsg("ok")
	if data != nil {
		res.SetData(data.GetMsgId(), data.GetData())
	}
	return res
}

/**
客户端响应
*/
func (c *Connection) responseBad(message string, data itls.IMessage) (response itls.IResponse) {
	res := &Response{}
	res.SetCode(500)
	res.SetMsg(message)
	if data != nil {
		res.SetData(data.GetMsgId(), data.GetData())
	}
	return res
}

/**
综合性
*/
func (c *Connection) response(code int, message string, data itls.IMessage) (response itls.IResponse) {
	res := &Response{}
	res.SetCode(code)
	res.SetMsg(message)
	if data != nil {
		res.SetData(data.GetMsgId(), data.GetData())
	}
	return res
}

/**
监控读线程
*/
func (c *Connection) StartReader() {
	//log.Printf("tls connection started.")
	for {
		select {
		case connected := <-c.isConnected:
			for {
				//1.连接掉线或未启动 退出
				if connected == true {
					response := c.read()
					if response.GetCode() != 200 {
						break
					}
					c.readChan[response.GetMsgId()] <- response
					continue
				}
				if connected == false {
					//等待 新连接 到来
					log.Printf("waiting new conn ")
					break
				}
			}
		case <-c.exitChan:
			log.Printf("socket reader service stop.")
			return
		}
	}
}

//1. 启动连接，让当前连接开始工作
func (c *Connection) Start() {
	go c.StartReader()
}

func (c *Connection) Logout() bool {

	return false
}

/**
主动断开连接
*/
func (c *Connection) disconnect() {
	//1.当前连接已关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	c.isConnected <- false
	//关闭连接
	if err := c.Conn.Close(); err != nil {
		log.Errorf("conn shutdown err:%s", err)
		return
	}
	c.Conn = nil
	log.Printf("conn socket closed.")
}

//2.服务终止
func (c *Connection) Stop() {
	//1.当前连接已关闭
	if c.isExit == true {
		return
	}

	c.isExit = true
	//关闭连接
	c.disconnect()
	//通知读业务，此连接已关闭
	c.exitChan <- true
	// 关闭通道
	close(c.isConnected)
	close(c.exitChan)
}

//3.获取当前连接
func (c *Connection) GetConn() net.Conn {
	return c.Conn
}

//5.获取客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}




//服务端代码

package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"
)

type Connection struct {
	//1.当前连接的socket套接字
	Conn net.Conn
	//2.当前连接的sessionID,全局唯一
	ConnID uint32
	//3.当前连接状态
	isClosed bool
	//4.连接关联的处理方法 :废弃
	MsgHandler itls.IMsgHandle
	//5.策略执行单元
	PEP ipolicyengine.IService
	//6.连接停止channel
	ExitChan chan bool
	//7.无缓冲通道，用于读写俩个goroutine 信息同步
	msgChan chan []byte
	stopCall imonitor.ICallback
}

/**

 */
func NewConnection(conn net.Conn, connID uint32, handle itls.IMsgHandle, pep ipolicyengine.IService, callback imonitor.ICallback) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		//Router: router,
		MsgHandler: handle,
		PEP:        pep,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte, 10),
		stopCall:   callback,
	}
}

// 策略执行单元
func (c *Connection) GetPEP() ipolicyengine.IService {
	return c.PEP
}

//启动read引擎
func (c *Connection) StartReader() {
	//log.Printf("Reader Goroutine:%d is  running", c.ConnID)
	//defer log.Printf("client:%s reader connId:%d goroutine exit.", c.RemoteAddr().String(), c.ConnID)
	bufConn := bufio.NewReader(c.Conn)
	for {
		select {
		case <-c.ExitChan:
			return
		default:
			//1.创建拆包解包对象
			dp := &DataPack{}
			//2.读取消息头
			headBuf := make([]byte, dp.GetHeadLen())
			_, err := bufConn.Read(headBuf)
			//_, err := io.ReadFull(c.Conn, headBuf)
			if err != nil {
				log.Errorf("conn socket read head err:%s ", err)
				c.Stop()
				return
			}
			//对message 结构
			msg, err := dp.Unpack(headBuf)
			if err != nil {
				log.Errorf("conn socket header unpack  err:%s ", err)
				c.Stop()
				return
			}
			var data []byte
			if msg.GetDataLen() > 0 { //拿到数据部分 读取字节长度
				data = make([]byte, msg.GetDataLen())
				_, err := bufConn.Read(data)
				//_, err := io.ReadFull(c.Conn, data)
				if err != nil {
					log.Errorf("conn socket read body data  err:%s ", err)
					c.Stop()
					return
				}
			}
			//拿到数据部分
			msg.SetData(data)
			req := Request{
				connection: c,
				msg:        msg,
			}
			//把客户端流量数据转化成请求
			c.MsgHandler.SendMsgToTaskQueue(&req)
		}
	}
}

//启动写引擎
func (c *Connection) StartWriter() {
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				log.Errorf("send to client err:%s", err)
				c.Stop()
				time.Sleep(time.Second)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

//1. 启动服务
func (c *Connection) Start() {
	defer log.Printf("client:%s connection socket exit.", c.RemoteAddr().String())

	// 读 goroutine
	go c.StartReader()
	//写 goroutine
	go c.StartWriter()

	select {
	case <-c.ExitChan: //进程退出
		return
	}
}

//2.终止服务
func (c *Connection) Stop() {
	//1.当前连接已关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	//执行回调函数
	c.stopCall.Call(c)
	//通知读业务，此连接已关闭
	c.ExitChan <- true
	// 关闭通道
	close(c.ExitChan)
	//关闭连接
	c.Conn.Close()
}

//3.获取当前连接
func (c *Connection) GetConn() net.Conn {
	return c.Conn
}

//4.获取当前连接id
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//5.获取客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//服务端消息写回客户端
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("conn closed when send msg ")
	}

	//1.data []byte 封装
	dp := NewDataPack()
	msg, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		c.Stop()
		return errors.New(fmt.Sprintf("pack message id:%d  pack err:%s", msgId, err))
	}
	c.msgChan <- msg
	return nil
}

func (c *Connection) GetLogin() bool {
	return c.isLogin
}

func (c *Connection) SetLogin(f bool) {
	c.isLogin = f
}

func (c *Connection) GetMsgAuth() bool {
	return c.isMsgAuth
}

func (c *Connection) SetMsgAuth(f bool) {
	c.isMsgAuth = f
}

/**
账号登录&短信认证同时具备 才可同步应用到客户端
*/
func (c *Connection) Available() bool {
	if c.isLogin && c.isMsgAuth {
		return true
	}
	return false
}

```