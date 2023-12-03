/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 14:28:08
*/
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"

	pb "github.com/robberphex/grpc-in-memory/helloworld"
	"google.golang.org/grpc"
)

/*
在操作系统层面，pipe表示一个数据管道，而这个管道两端都在本程序中，可以很好的满足
我们的要求：基于内存的网络通信。

Golang也基于pipe提供了net.Pipe()函数创建了一个双向的、基于内存通信的管道，在能力上，
能够很好的满足gRPC对底层通信的要求。

但是net.Pipe仅仅产生了两个net.Conn，即只产生两个网络连接，没有之前提到的Listner，
也没有Dial方法。

于是结合Golang的channel，把net.Pipe包装成了Listner，也提供了Dial方法：

Listener.Accept()，只需要监听一个channel，客户端连接过来的时候，把连接通过channel传
递过来即可

Dial方法，调用Pipe，将一端通过channel给服务端（作为服务端连接），另一端作为客户端连接
*/
var ErrPipeListenerClosed = errors.New(`pipe listener already closed`)

type PipeListener struct {
	ch    chan net.Conn
	close chan struct{}
	done  uint32
	m     sync.Mutex
}

func ListenPipe() *PipeListener {
	return &PipeListener{
		ch:    make(chan net.Conn),
		close: make(chan struct{}),
	}
}

// Accept 等待客户端连接
func (l *PipeListener) Accept() (c net.Conn, e error) {
	select {
	case c = <-l.ch:
	case <-l.close:
		e = ErrPipeListenerClosed
	}
	return
}

// Close 关闭 listener.
func (l *PipeListener) Close() (e error) {
	if atomic.LoadUint32(&l.done) == 0 {
		l.m.Lock()
		defer l.m.Unlock()
		if l.done == 0 {
			defer atomic.StoreUint32(&l.done, 1)
			close(l.close)
			return
		}
	}
	e = ErrPipeListenerClosed
	return
}

// Addr 返回 listener 的地址
func (l *PipeListener) Addr() net.Addr {
	return pipeAddr(0)
}
func (l *PipeListener) Dial(network, addr string) (net.Conn, error) {
	return l.DialContext(context.Background(), network, addr)
}
func (l *PipeListener) DialContext(ctx context.Context, network, addr string) (conn net.Conn, e error) {
	// PipeListener是否已经关闭
	if atomic.LoadUint32(&l.done) != 0 {
		e = ErrPipeListenerClosed
		return
	}

	// 创建pipe
	c0, c1 := net.Pipe()
	// 等待连接传递到服务端接收
	select {
	case <-ctx.Done():
		e = ctx.Err()
	case l.ch <- c0:
		conn = c1
	case <-l.close:
		c0.Close()
		c1.Close()
		e = ErrPipeListenerClosed
	}
	return
}

type pipeAddr int

func (pipeAddr) Network() string {
	return `pipe`
}
func (pipeAddr) String() string {
	return `pipe`
}

// helloworld.GreeterServer 的实现
type server struct {
	// 为了后面代码兼容，必须聚合UnimplementedGreeterServer
	// 这样以后在proto文件中新增加一个方法的时候，这段代码至少不会报错
	pb.UnimplementedGreeterServer
}

// unary调用的服务端代码
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// 客户端流式调用的服务端代码
// 接收两个req，然后返回一个resp
func (s *server) SayHelloRequestStream(streamServer pb.Greeter_SayHelloRequestStreamServer) error {
	req, err := streamServer.Recv()
	if err != nil {
		log.Printf("error receiving: %v", err)
		return err
	}
	log.Printf("Received: %v", req.GetName())
	req, err = streamServer.Recv()
	if err != nil {
		log.Printf("error receiving: %v", err)
		return err
	}
	log.Printf("Received: %v", req.GetName())
	streamServer.SendAndClose(&pb.HelloReply{Message: "Hello " + req.GetName()})
	return nil
}

// 服务端流式调用的服务端代码
// 接收一个req，然后发送两个resp
func (s *server) SayHelloReplyStream(req *pb.HelloRequest, streamServer pb.Greeter_SayHelloReplyStreamServer) error {
	log.Printf("Received: %v", req.GetName())
	err := streamServer.Send(&pb.HelloReply{Message: "Hello " + req.GetName()})
	if err != nil {
		log.Printf("error Send: %+v", err)
		return err
	}
	err = streamServer.Send(&pb.HelloReply{Message: "Hello " + req.GetName() + "_dup"})
	if err != nil {
		log.Printf("error Send: %+v", err)
		return err
	}
	return nil
}

// 双向流式调用的服务端代码
func (s *server) SayHelloBiStream(streamServer pb.Greeter_SayHelloBiStreamServer) error {
	req, err := streamServer.Recv()
	if err != nil {
		log.Printf("error receiving: %+v", err)
		// 及时将错误返回给客户端，下同
		return err
	}
	log.Printf("Received: %v", req.GetName())
	err = streamServer.Send(&pb.HelloReply{Message: "Hello " + req.GetName()})
	if err != nil {
		log.Printf("error Send: %+v", err)
		return err
	}
	// 离开这个函数后，streamServer会关闭，所以不推荐在单独的goroute发送消息
	return nil
}

// 新建一个服务端实现
func NewServerImpl() *server {
	return &server{}
}

// 将一个服务实现转化为一个客户端
func serverToClient(svc *server) pb.GreeterClient {
	// 创建一个基于pipe的Listener
	pipe := ListenPipe()

	s := grpc.NewServer()
	// 注册Greeter服务到gRPC
	pb.RegisterGreeterServer(s, svc)
	if err := s.Serve(pipe); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// 客户端指定使用pipe作为网络连接
	clientConn, err := grpc.Dial(`pipe`,
		grpc.WithInsecure(),
		grpc.WithContextDialer(func(c context.Context, s string) (net.Conn, error) {
			return pipe.DialContext(c, `pipe`, s)
		}),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 基于pipe连接，创建gRPC客户端
	c := pb.NewGreeterClient(clientConn)
	return c
}

func main() {
	svc := NewServerImpl()
	c := serverToClient(svc)

	ctx := context.Background()

	// unary调用
	for i := 0; i < 5; i++ {
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: fmt.Sprintf("world_unary_%d", i)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}

	// 客户端流式调用
	for i := 0; i < 5; i++ {
		streamClient, err := c.SayHelloRequestStream(ctx)
		if err != nil {
			log.Fatalf("could not SayHelloRequestStream: %v", err)
		}
		err = streamClient.Send(&pb.HelloRequest{Name: fmt.Sprintf("SayHelloRequestStream_%d", i)})
		if err != nil {
			log.Fatalf("could not Send: %v", err)
		}
		err = streamClient.Send(&pb.HelloRequest{Name: fmt.Sprintf("SayHelloRequestStream_%d_dup", i)})
		if err != nil {
			log.Fatalf("could not Send: %v", err)
		}
		reply, err := streamClient.CloseAndRecv()
		if err != nil {
			log.Fatalf("could not Recv: %v", err)
		}
		log.Println(reply.GetMessage())
	}

	// 服务端流式调用
	for i := 0; i < 5; i++ {
		streamClient, err := c.SayHelloReplyStream(ctx, &pb.HelloRequest{Name: fmt.Sprintf("SayHelloReplyStream_%d", i)})
		if err != nil {
			log.Fatalf("could not SayHelloReplyStream: %v", err)
		}
		reply, err := streamClient.Recv()
		if err != nil {
			log.Fatalf("could not Recv: %v", err)
		}
		log.Println(reply.GetMessage())
		reply, err = streamClient.Recv()
		if err != nil {
			log.Fatalf("could not Recv: %v", err)
		}
		log.Println(reply.GetMessage())
	}

	// 双向流式调用
	for i := 0; i < 5; i++ {
		streamClient, err := c.SayHelloBiStream(ctx)
		if err != nil {
			log.Fatalf("could not SayHelloStream: %v", err)
		}
		err = streamClient.Send(&pb.HelloRequest{Name: fmt.Sprintf("world_stream_%d", i)})
		if err != nil {
			log.Fatalf("could not Send: %v", err)
		}
		reply, err := streamClient.Recv()
		if err != nil {
			log.Fatalf("could not Recv: %v", err)
		}
		log.Println(reply.GetMessage())
	}
}
