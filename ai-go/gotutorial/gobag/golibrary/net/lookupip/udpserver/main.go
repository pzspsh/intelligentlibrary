/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 00:08:18
*/
package main

// import (
// 	"fmt"
// 	"math"
// 	"net"
// 	"strings"
// 	"time"

// 	"github.com/labstack/gommon/log"
// )

// const (
// 	StatusStart = 1 // 启动动中
// 	StatusRun   = 2 // 运行中
// 	StatusStop  = 3 // 已关闭
// )

// type Service struct {
// 	Name    string
// 	Addr    string
// 	Port    uint32
// 	Timeout time.Duration
// 	Status  byte //1:启动，2：运行，3：关闭
// 	UdpConn *net.UDPConn
// 	//Router  udp.IRouter
// 	msgHandle udp.IMsgHandle
// 	ExitChan  chan bool

// 	MaxWorkerTaskLen uint32               // 队列最大任务数
// 	WorkerPoolSize   uint32               //业务工作Worker池的数量
// 	TaskQueue        []chan iudp.IRequest //Worker负责取任务的消息队列
// }

// /*
// *
// udp server 初始化
// */
// func NewServer(name, addr string, port uint32, timeout time.Duration) iudp.IServer {
// 	if strings.TrimSpace(name) == "" {
// 		name = "SPA Serve"
// 	}
// 	poolSize := config.Cfg.Spa.WorkPoolSize
// 	if poolSize <= 0 {
// 		poolSize = 10
// 	}
// 	taskLen := config.Cfg.Spa.WorkTaskLen
// 	if taskLen <= 0 {
// 		taskLen = 1024
// 	}
// 	return &Service{
// 		Name:             name,
// 		Addr:             addr,
// 		Port:             port,
// 		Timeout:          timeout,
// 		Status:           StatusStart,
// 		UdpConn:          nil,
// 		ExitChan:         make(chan bool),
// 		msgHandle:        NewMsgHandle(),
// 		MaxWorkerTaskLen: taskLen,
// 		WorkerPoolSize:   poolSize,
// 		TaskQueue:        make([]chan iudp.IRequest, taskLen),
// 	}
// }

// // 将消息交给TaskQueue,由worker进行处理
// func (s *Service) sendMsgToTaskQueue(request iudp.IRequest) {
// 	//根据requestId来分配当前的连接应该由哪个worker负责处理
// 	//轮询的平均分配法则
// 	//得到需要处理此条连接的workerID
// 	workerID := request.GetId() % s.WorkerPoolSize
// 	log.Printf("Add Request id:%d to worker id:%d", request.GetId(), workerID)
// 	//将请求消息发送给任务队列
// 	s.TaskQueue[workerID] <- request
// }

// // 启动服务
// func (s *Service) Start() {
// 	address := fmt.Sprintf("%s:%d", s.Addr, s.Port)
// 	log.Printf("[start] %s :%s.", s.Name, address)
// 	var addr *net.UDPAddr
// 	var err error
// 	if addr, err = net.ResolveUDPAddr("udp", address); err != nil {
// 		log.Fatalf("udp addr resolve err:%s", err)
// 	}
// 	s.UdpConn, err = net.ListenUDP("udp", addr)
// 	if err != nil {
// 		log.Fatalf("udp spa server listen err:%s", err)
// 	}

// 	go s.startReader()

// 	s.Status = StatusRun
// }

// // 启动业务服务
// func (s *Service) startReader() {
// 	var reqId uint32
// 	data := pool.BufferSmallPool.Get().([]byte) //make([]byte, 1<<16)
// 	defer pool.BufferSmallPool.Put(data)
// 	for {
// 		n, cAddr, err := s.UdpConn.ReadFromUDP(data)
// 		if err != nil {
// 			if s.Status == StatusStop {
// 				return
// 			}
// 			log.Errorf("reader from udp err:%s", err)
// 			continue
// 		}
// 		reqId++
// 		if reqId >= math.MaxInt32 {
// 			reqId = 1
// 		}
// 		log.Printf("udp server reader:%s,%d", string(data[:n]), len(data[:n]))
// 		dp := NewDataPack()
// 		msg, err := dp.Unpack(data[:n])
// 		if err != nil {
// 			log.Errorf("client:%s udp data of unpack err:%s", cAddr.String(), err)
// 			continue
// 		}
// 		req := NewRequest(reqId, s.UdpConn, cAddr, msg)
// 		//go s.msgHandle.DoMsgHandler(req)
// 		s.sendMsgToTaskQueue(req)
// 	}
// }

// // 启动一个Worker工作流程
// func (s *Service) startOneWorker(workerID int, taskQueue chan iudp.IRequest) {
// 	//log.Printf("%s Worker ID:%d  is started.", s.Name, workerID)
// 	//不断的等待队列中的消息
// 	for {
// 		select {
// 		//有消息则取出队列的Request，并执行绑定的业务方法
// 		case request := <-taskQueue:
// 			s.msgHandle.DoMsgHandler(request)
// 		case <-s.ExitChan:
// 			//log.Printf("%s service stopped", s.Name)
// 			return
// 		}
// 	}
// }

// // 启动worker工作池
// func (s *Service) StartWorkerPool() {
// 	//遍历需要启动worker的数量，依此启动
// 	go func() {
// 		for i := 0; i < int(s.WorkerPoolSize); i++ {
// 			//一个worker被启动
// 			//给当前worker对应的任务队列开辟空间
// 			s.TaskQueue[i] = make(chan iudp.IRequest, s.MaxWorkerTaskLen)
// 			//启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
// 			go s.startOneWorker(i, s.TaskQueue[i])
// 		}
// 	}()
// }

// // 启动业务服务
// func (s *Service) Serve() {
// 	s.StartWorkerPool()
// 	s.Start()
// }

// // 停止服务
// func (s *Service) Stop() {
// 	if s.Status == StatusStop {
// 		return
// 	}
// 	s.Status = StatusStop
// 	//通知读业务，此连接已关闭
// 	s.ExitChan <- true
// 	// 关闭通道
// 	close(s.ExitChan)
// 	//关闭连接
// 	s.UdpConn.Close()
// 	log.Printf("%s stopped.", s.Name)
// }

// // 添加路由处理
// func (s *Service) AddRouter(msgId uint32, router iudp.IRouter) {
// 	//s.Router = router
// 	s.msgHandle.AddRouter(msgId, router)
// }

// func (s *Service) AddPEP(pep ipolicyengine.IService) {
// 	s.msgHandle.AddPEP(pep)
// }
