# Go context 包讲解

##### 目录

- Context 包到底是干嘛用的？

- context 原理

- 什么时候应该使用 Context？

- 如何创建 Context？

- 主协程通知有子协程，子协程又有多个子协程

- context 核心接口

- - emptyCtx 结构体
  - Backgroud
  - TODO
  - valueCtx 结构体
  - WithValue 向 context 添加值
  - Value 向 context 取值
  - 示例

- WithCancel 可取消的 context

- - cancelCtx 结构体

- WithDeadline-超时取消 context

-
- WithTimeout-超时取消 context

- - timerCtx 结构体

- 总结核心原理

## Context 包到底是干嘛用的？

我们会在用到很多东西的时候都看到 context 的影子，比如 gin 框架，比如 grpc，这东西到底是做啥的？
大家都在用，没几个知道这是干嘛的，知其然而不知其所以然

原理说白了就是：

- 当前协程取消了，可以通知所有由它创建的子协程退出
- 当前协程取消了，不会影响到创建它的父级协程的状态
- 扩展了额外的功能：超时取消、定时取消、可以和子协程共享数据

## context 原理

这就是 context 包的核心原理，链式传递 context，基于 context 构造新的 context

## 什么时候应该使用 Context？

- 每一个 RPC 调用都应该有超时退出的能力，这是比较合理的 API 设计
- 不仅仅是超时，你还需要有能力去结束那些不再需要操作的行为
- context.Context 是 Go 标准的解决方案
- 任何函数可能被阻塞，或者需要很长时间来完成的，都应该有个 context.Context

## 如何创建 Context？

在 RPC 开始的时候，使用 context.Background()

有些人把在 main() 里记录一个 context.Background()，然后把这个放到服务器的某个变量里，然后请求来了后从这个变量里继承 context。这么做是不对的。直接每个请求，源自自己的 context.Background() 即可。

如果你没有 context，却需要调用一个 context 的函数的话，用 context.TODO()

如果某步操作需要自己的超时设置的话，给它一个独立的 sub-context（如前面的例子）

## 主协程通知有子协程，子协程又有多个子协程

```go
package main
import (
    "context"
    "fmt"
    "time"
)
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    //缓冲通道预先放置10个消息
    messages := make(chan int, 10)
    defer close(messages)
    for i := 0; i < 10; i++ {
        messages <- i
    }
    //启动3个子协程消费messages消息
    for i := 1; i <= 3; i++ {
        go child(i, ctx, messages)
    }
    time.Sleep(3 * time.Second) //等待子协程接收一半的消息
    cancel() //结束前通知子协程
    time.Sleep(2 * time.Second) //等待所有的子协程输出
    fmt.Println("主协程结束")
}
//从messages通道获取信息，当收到结束信号的时候不再接收
func child(i int, ctx context.Context, messages <-chan int) {
    //基于父级的context建立context
    newCtx, _ := context.WithCancel(ctx)
    go childJob(i, "a", newCtx)
    go childJob(i, "b", newCtx)
Consume:
    for {
        time.Sleep(1 * time.Second)
        select {
        case <-ctx.Done():
            fmt.Printf("[%d]被主协程通知结束...\n", i)
            break Consume
        default:
            fmt.Printf("[%d]接收消息: %d\n", i, <-messages)
        }
    }
}
//任务
func childJob(parent int, name string, ctx context.Context) {
    for {
        time.Sleep(1 * time.Second)
        select {
        case <-ctx.Done():
            fmt.Printf("[%d-%v]被结束...\n", parent, name)
            return
        default:
            fmt.Printf("[%d-%v]执行\n", parent, name)
        }
    }
}
```

运行结果如下

![img](../../../../../images/5bcfa0f95efbf00f48261a09bb6f0021.jpg)

可以看到，改成 context 包还是顺利的通过子协程退出了
主要修改了几个地方，再 ctx 向下传递

![img](../../../../../images/571d86fc073eb921f1e5973d60e6e5e5.jpg)

基于上层 context 再构建当前层级的 context

![img](../../../../../images/5d4ca754cdce9aaf73407d7f979d7eaf.jpg)

监听 context 的退出信号，

![img](../../../../../images/6a2856bff55a3f06da90ea0802664770.jpg)

这就是 context 包的核心原理，链式传递 context，基于 context 构造新的 context

## context 核心接口

```
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

Deadline 返回绑定当前 context 的任务被取消的截止时间；如果没有设定期限，将返回 ok == false。

Done 当绑定当前 context 的任务被取消时，将返回一个关闭的 channel；如果当前 context 不会被取消，将返回 nil。

Err 如果 Done 返回的 channel 没有关闭，将返回 nil;如果 Done 返回的 channel 已经关闭，将返回非空的值表示任务结束的原因。如果是 context 被取消，Err 将返回 Canceled；如果是 context 超时，Err 将返回 DeadlineExceeded。

Value 返回 context 存储的键值对中当前 key 对应的值，如果没有对应的 key,则返回 nil。

### emptyCtx 结构体

实现了 context 接口，emptyCtx 没有超时时间，不能取消，也不能存储额外信息，所以 emptyCtx 用来做根节点，一般用 Background 和 TODO 来初始化 emptyCtx

### Backgroud

通常用于主函数，初始化以及测试，作为顶层的 context

> context.Background()

### TODO

不确定使用什么用 context 的时候才会使用

### valueCtx 结构体

```go
type valueCtx struct{ Context key, val interface{} }
```

valueCtx 利用 Context 的变量来表示父节点 context，所以当前 context 继承了父 context 的所有信息
valueCtx 还可以存储键值。

### WithValue 向 context 添加值

可以向 context 添加键值

```go
func WithValue(parent Context, key, val interface{}) Context {
    if key == nil {
        panic("nil key")
    }
    if !reflect.TypeOf(key).Comparable() {
        panic("key is not comparable")
    }
    return &amp;valueCtx{parent, key, val}
}
```

添加键值会返回创建一个新的 valueCtx 子节点

### Value 向 context 取值

```go
func (c *valueCtx) Value(key interface{}) interface{} {
    if c.key == key {
        return c.val
    }
    return c.Context.Value(key)
}
```

可以用来获取当前 context 和所有的父节点存储的 key

如果当前的 context 不存在需要的 key，会沿着 context 链向上寻找 key 对应的值，直到根节点

### 示例

```go
package main
import (
	"context"
	"fmt"
	"time"
)
func main() {
	ctx := context.WithValue(context.Background(), "name1", "root1")

	//第一层
	go func(parent context.Context) {
		ctx = context.WithValue(parent, "name2", "root2")
		//第二层
		go func(parent context.Context) {
			ctx = context.WithValue(parent, "name3", "root3")
			//第三层
			go func(parent context.Context) {
				//可以获取所有的父类的值
				fmt.Println(ctx.Value("name1"))
				fmt.Println(ctx.Value("name2"))
				fmt.Println(ctx.Value("name3"))
				//不存在
				fmt.Println(ctx.Value("name4"))
			}(ctx)
		}(ctx)
	}(ctx)
	time.Sleep(1 * time.Second)
	fmt.Println("end")
}
```

运行结果

![img](../../../../../images/6359eb23adf0d4a16ff12b7729b3a74a.jpg)

可以看到，子 context 是可以获取所有父级设置过的 key

## WithCancel 可取消的 context

用来创建一个可取消的 context，返回一个 context 和一个 CancelFunc，调用 CancelFunc 可以触发 cancel 操作。

```go
package main
import (
	"context"
	"fmt"
	"time"
)
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//第一层
	go func(parent context.Context) {
		ctx, _ := context.WithCancel(parent)
		//第二层
		go func(parent context.Context) {
			ctx, _ := context.WithCancel(parent)
			//第三层
			go func(parent context.Context) {
				waitCancel(ctx, 3)
			}(ctx)
			waitCancel(ctx, 2)
		}(ctx)
		waitCancel(ctx, 1)
	}(ctx)
	// 主线程给的结束时间
	time.Sleep(2 * time.Second)
	cancel() // 调用取消context
	time.Sleep(1 * time.Second)
}
func waitCancel(ctx context.Context, i int) {
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Printf("%d end\n", i)
			return
		default:
			fmt.Printf("%d do\n", i)
		}
	}
}
```

结果：

![img](../../../../../images/081888bc4b10077b530e3d50a143075d.jpg)

### cancelCtx 结构体

```go
type cancelCtx struct {
    Context
    mu sync.Mutex
    done chan struct{}
    children map[canceler]struct{}
    err error
}
type canceler interface {
    cancel(removeFromParent bool, err error)
    Done() &lt;-chan struct{}
}
```

## WithDeadline-超时取消 context

返回一个基于 parent 的可取消的 context，并且过期时间 deadline 不晚于所设置时间 d

## WithTimeout-超时取消 context

创建一个定时取消 context，和 WithDeadline 差不多，WithTimeout 是相对时间

### timerCtx 结构体

timerCtx 是基于 cancelCtx 的 context 精英，是一种可以定时取消的 context，过期时间的 deadline 不晚于所设置的时间 d

示例：

```go
package main
import (
    "context"
    "fmt"
    "time"
)
func main() {
    // 设置超时时间
    ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
    //第一层
    go func(parent context.Context) {
        ctx, _ := context.WithCancel(parent)
        //第二层
        go func(parent context.Context) {
            ctx, _ := context.WithCancel(parent)
            //第三层
            go func(parent context.Context) {
                waitCancel(ctx, 3)
            }(ctx)
            waitCancel(ctx, 2)
        }(ctx)
        waitCancel(ctx, 1)
    }(ctx)

    <-ctx.Done()
    // 给时间调用end
    time.Sleep(time.Second)
}
func waitCancel(ctx context.Context, i int) {
    for {
        time.Sleep(time.Second)
        select {
        case <-ctx.Done():
            fmt.Printf("%d end\n", i)
            return
        default:
            fmt.Printf("%d do\n", i)
        }
    }
}
```

运行结果：

```vbnet
1 do
3 do
2 do
1 end
3 end
2 end
```

可以看到，虽然我们没有调用 cancel 方法，5 秒后自动调用了，所有的子 goroutine 都已经收到停止信号

## 总结核心原理

- Done 方法返回一个 channel
- 外部通过调用<-channel 监听 cancel 方法
- cancel 方法会调用 close(channel)
  当调用 close 方法的时候，所有的 channel 再次从通道获取内容，会返回零值和 false

```
res,ok := &lt;-done:
```

- 过期自动取消，使用了 time.AfterFunc 方法，到时调用 cancel 方法

```
  c.timer = time.AfterFunc(dur, func() {
   c.cancel(true, DeadlineExceeded)
  })
```
