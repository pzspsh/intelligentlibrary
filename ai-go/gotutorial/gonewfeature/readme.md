# Go的新特性
### Go 1.22

#### for循环

Go在1.22版本之前，for 循环迭代器的变量是一个单一变量，使用不当，会导致意想不到的行为，可能会造成共享循环变量的问题。 如依旧要使用旧版本，可以主动配置 GOEXPERIMENT=loopvar

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for _, v := range nums {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(2 * time.Second)
}
```

同样的一段代码，在1.22之前打印出来的数据可能为 `5 5 5 5 5` 1.22版本打印出来可能为 `1 2 3 4 5` (无序)



**1、Go1.22之后整数类型的范围进行循环迭代**

```go
package main

import "fmt"

func main() {
	for i := range 5 {
		fmt.Println(i)
	}
}
```

**2、math/rand 版本 2**

```golang
package main

import (
	"math/rand/v2"
)

func main() {
	rand.Intn(100)
}
```

**3、切片拼接**

```golang
package main

import (
	"fmt"
	"slices"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}
	slice3 := []int{11, 12, 13, 14, 15}
	
	// 1.22之前
	merged := append(slice1, slice2...)
	merged = append(merged, slice3...)
	fmt.Println(merged)

	// 1.22版本
	merged2 := slices.Concat(slice1, slice2, slice3)
	fmt.Println(merged2)

}
```



**4、增强了http路由功能**

```golang
mux := http.NewServeMux()
mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) { //v1.22可以直接在路由中指定允许的请求方法
	// if r.Method != "GET" {   //v1.22之前需要自己限制请求方法
	// 	fmt.Fprint(w, "warn: 只支持GET方法")
	// } else {
	fmt.Fprint(w, "你好 "+r.PathValue("name")) //从restful风格的url中获取参数
	// }
})

if err := http.ListenAndServe("127.0.0.1:5678", mux); err != nil {
	panic(err)
}
```



# **Go 1.22新特性前瞻**

## 1. 语言变化

Go 1.22的语言特性变化主要是**围绕for loop的**。

### 1.1 loopvar试验特性转正

在[Go 1.21版本](https://tonybai.com/2023/08/20/some-changes-in-go-1-21)中，作为试验特性loopvar在Go 1.22中正式转正。如果你还不知道这个特性是啥，我们来看一下下面这个最能说明问题的示例：

```go
package main
  
import (
    "fmt"
    "sync"
)

func main() {
    sl := []int{11, 12, 13, 14, 15}
    var wg sync.WaitGroup
    for i, v := range sl {
        wg.Add(1)
        go func() {
            fmt.Printf("%d : %d\n", i, v)
            wg.Done()
        }()
    }
    wg.Wait()
}
```

我们分别用Go 1.22rc1和Go 1.21.0来运行上面这段代码：

```go
// 使用go 1.22rc1的运行结果：
$go run for_range.go
4 : 15
1 : 12
0 : 11
3 : 14
2 : 13

// 使用go 1.21.0的运行结果：
$go run for_range.go 
4 : 15
4 : 15
4 : 15
4 : 15
4 : 15
```

之所以存在差异，是因为Go 1.22版本开始，for range语句中声明的循环变量（比如这里的i和v）不再是整个loop一份(loop var per loop)，而是每次iteration都会有自己的变量(loop var per-iteration)，这样在Go 1.22中，for range中的goroutine启动的[闭包函数](https://tonybai.com/2021/08/09/when-variables-captured-by-closures-are-recycled-in-go/)中捕获的变量是loop var per-iteration，这样才会输出5个不同的索引值和对应的切片值。

那传统的3-clause的for loop呢？其中的循环变量的语义是否也发生变化了呢？我们看下面示例：

```go
package main
  
import (
    "fmt"
    "sync"
)

func main() {
    sl := []int{11, 12, 13, 14, 15}
    var wg sync.WaitGroup
    for i := 0; i < len(sl); i++ {
        wg.Add(1)
        go func() {
            v := sl[i]
            fmt.Printf("%d : %d\n", i, v)
            wg.Done()
        }()
    }
    wg.Wait()
}
```

我们依然分别用Go 1.22rc1和Go 1.21.0版本运行这段代码，得到的结果如下：

```go
// 使用go 1.22rc1的运行结果：
$go run classic_for_loop.go 
0 : 11
4 : 15
2 : 13
3 : 14
1 : 12

// 使用go 1.21.0的运行结果：
$go run classic_for_loop.go 
panic: runtime error: index out of range [5] with length 5

goroutine 20 [running]:
main.main.func1()
	/Users/tonybai/test/go/go1.22-foresight/lang/for-range/classic_for_loop.go:14 +0xc9
created by main.main in goroutine 1
	/Users/tonybai/test/go/go1.22-foresight/lang/for-range/classic_for_loop.go:13 +0x7f
panic: runtime error: index out of range [5] with length 5

goroutine 19 [running]:
main.main.func1()
	/Users/tonybai/test/go/go1.22-foresight/lang/for-range/classic_for_loop.go:14 +0xc9
created by main.main in goroutine 1
	/Users/tonybai/test/go/go1.22-foresight/lang/for-range/classic_for_loop.go:13 +0x7f
exit status 2
```

从输出结果来看，3-clause的for语句中声明的循环变量也变成了loop var per-iteration了。

在Go 1.22之前，go vet工具在遇到像上面代码那样在闭包中引用循环变量的情况时会给出警告，但由于Go 1.22的这个语义修正，go vet对于Go 1.22及以后版本(根据go.mod中的指示)的类似Go代码将不再报错。

不过就像Russ Cox在[spec: less error-prone loop variable scoping](https://github.com/golang/go/issues/60078)这一issue中提及那样，该特性落地可能会带来不兼容问题，即对存量代码行为的破坏性改变。为此Go团队提供了一个[名为bisect的工具](https://pkg.go.dev/golang.org/x/tools/cmd/bisect)，该工具可以检测出存量代码在for loop语义发生变更后是否会导致问题。不过该工具似乎只能与go test一起使用，也就是说你只能对那些被[Go测试](https://tonybai.com/2023/07/16/the-guide-of-go-testing-with-testify-package/)覆盖到的for loop进行检测。

目前[spec: less error-prone loop variable scoping](https://github.com/golang/go/issues/60078)这一issue还处于open状态，也没有放入Go 1.22 milestone中，不知道后续是否还会存在变数！

### 1.2 range支持整型表达式

在Go 1.22版本中，for range后面的range表达式除了支持传统的像数组、切片、map、channel等表达式外，**还支持放置整型表达式**，比如下面这个例子：

```go
package main

import (
    "fmt"
)

func main() {  
    n := 5   
    for i := range n {
        fmt.Println(i) 
    }  
}  
```

我们知道：for range会在执行伊始对range表达式做一次求值，这里对n求值结果为5。按照新增的for range后接整型表达式的语义，对于整数值n，for range每次迭代值会从0到n-1按递增顺序进行。上面代码中的for range会从0迭代到4(5-1)，我们执行一下上述代码就可以印证这一点：

```go
$go run main.go
0
1
2
3
4
```

如果n <= 0，则循环不运行任何迭代。

这个新语法特性，可以理解为是一种“语法糖”，是下面等价代码的“语法糖”：

```go
for i := 0; i < 5; i++ {
	fmt.Println(i) 
}
```

不过，迭代总是从0开始，似乎限制了该语法糖的使用范围。

### 1.3 试验特性：range-over-function iterators

在for range支持整型表达式的时候，[Go团队也考虑了增加函数迭代器(iterator)](https://github.com/golang/go/issues/61405)，不过前者语义清晰，实现简单。后者展现形式、语义和实现都非常复杂，于是在Go 1.22中，函数迭代器以试验特性提供，通过GOEXPERIMENT=rangefunc可以体验该功能特性。

在没有函数迭代器之前，我们实现一个通用的反向迭代切片的函数可能是像这样：

```go
func Backward(s []E) func(func(int, E) bool) {
    return func(yield func(int, E) bool) {
        for i := len(s)-1; i >= 0; i-- {
            if !yield(i, s[i]) {
                return
            }
        }
        return
    }
}
```

下面是在Go 1.21.0版本中使用上面Backward函数的方式：

```go
func main() {
    sl := []string{"hello", "world", "golang"}
    Backward(sl)(func(i int, s string) bool {
        fmt.Printf("%d : %s\n", i, s)
        return true
    })
}
```

我们用Go 1.21.0运行一下上述示例：

```
$go run backward_iterate_slice_old.go
2 : golang
1 : world
0 : hello
```

在以前版本中，这种对切片、数组或map中进行元素迭代的情况在实际开发中非常常见，也比较模式化，但基于目前语法，使用起来非常不便。于是Go团队提出将它们[与for range结合在一起的提案](https://github.com/golang/go/issues/61405)。有了range-over-function iterator机制后，我们就可以像下面这样使用Backward泛型函数了：

```
func main() {
    sl := []string{"hello", "world", "golang"}
    for i, s := range Backward(sl) {
        fmt.Printf("%d : %s\n", i, s)
    }
}
```

相比于上面的老版本代码，这也的代码更简洁清晰了，使用Go 1.22rc1的运行结果也与老版本别无二致：

```
$GOEXPERIMENT=rangefunc  go run backward_iterate_slice_new.go
2 : golang
1 : world
0 : hello
```

但代价就是要理解什么样原型的函数才能与for range一起使用实现函数迭代，这的确有些复杂，本文就不展开说了，有兴趣的童鞋可以先看看有关[range-over-function iterator的wiki](https://go.dev/wiki/RangefuncExperiment)先行了解一下。

## 2. 编译器、运行时与工具链

### 2.1 继续增强[PGO优化](https://github.com/golang/go/issues/61577)

自[Go 1.20版本引入PGO](https://tonybai.com/2023/02/08/some-changes-in-go-1-20)(profile-guided optimization)后，PGO这种优化技术带来的优化效果就得到了持续的提升：Go 1.20实测性能提升仅为1.05%；[Go 1.21版本发布](https://tonybai.com/2023/08/20/some-changes-in-go-1-21)时，官方的数据是2%~7%，而Go 1.21编译器自身在PGO优化过后编译速度提升约6%。

在Go 1.22中，官方给出的数字则是2%~14%，这14%的提升想必是来自Google内部的某个实际案例。

### 2.2 inline和devirtualize

在Go 1.22中，Go编译器可以更灵活的运用devirtualize和inline对代码进行优化了。

在面向对象的编程中，虚拟函数是一种在运行时动态确定调用的函数。当调用虚拟函数时，编译器通常会为其生成一段额外的代码，用于在运行时确定要调用的具体函数。这种动态调度的机制使得程序可以根据实际对象类型来执行相应的函数，但也带来了一定的性能开销。通过devirtualize优化技术，编译器会尝试在编译时确定调用的具体函数，而不是在运行时进行动态调度。这样可以避免运行时的开销，并**允许编译器进行更多的优化**。

对应到Go来说，就是在编译阶段**将使用接口进行的方法调用转换为通过接口的实际类型的实例直接调用该方法**。

关于内联优化，今年Austin Clements发起了[inline大修项目](https://github.com/golang/go/issues/61502)，对Go编译器中的内联优化过程进行全面调整，目标是在Go 1.22中拥有更有效的、具有启发能力的内联，为后续内联的进一步增强奠定基础。该大修的成果目前以GOEXPERIMENT=newinliner试验特性的形式在Go 1.22中提供。

### 2.3 运行时

运行时的变化主要还是来自[GC](https://tonybai.com/2023/06/13/understand-go-gc-overhead-behind-the-convenience/)。

Go 1.22中，运行时会将基于类型的垃圾回收的元数据放在每个堆对象附近，从而可以将Go程序的CPU性能提高1-3%。同时，通过减少重复的元数据的优化，内存开销也将降低约1%。不确定减少重复元数据(metadata)这一优化是否[来自对unique包的讨论](https://github.com/golang/go/issues/62483#issuecomment-1800913220)。

### 2.4 工具链

在Go工具链改善方面，首当其冲的要数go module相关工具了。

在Go 1.22中，go work增加了一个与go mod一致的特性：支持vendor。通过go work vendor，可以将workspace中的依赖放到vendor目录下，同时在构建时，如果module root下有vendor目录，那么默认的构建是go build -mod=vendor，即基于vendor的构建。

go mod init在Go 1.22中将不再考虑GOPATH时代的包依赖工具的配置文件了，比如Gopkg.lock。在Go 1.22版本之前，如果go module之前使用的是类似[dep这样的工具来管理包依赖](https://tonybai.com/2017/06/08/first-glimpse-of-dep/)，go mod init会尝试读取dep配置文件来生成go.mod。

go vet工具取消了对loop变量引用的警告，增加了对空append的行为的警告(比如：slice = append(slice))、增加了deferring time.Since的警告以及在log/slog包的方法调用时key-value pair不匹配的警告。

## 3. 标准库

最后，我们来看看标准库的变化。每次Go发布新版本，标准库都是占更新的大头儿，这里无法将所有变更点一一讲解，仅说说几个重要的变更点。

### 3.1 增强http.ServerMux表达能力

Go内置电池，从诞生伊始就内置了强大的http库，不过长期以来http原生的ServeMux表达能力比较单一，不支持通配符等，这也是Go社区长期以来一直使用像[gorilla/mux](https://github.com/gorilla/mux/)、[httprouter](https://github.com/julienschmidt/httprouter)等第三方路由库的原因。

今年log/slog的作者Jonathan Amsterdam又创建了新的提案：[net/http: enhanced ServeMux routing](https://github.com/golang/go/issues/61410)，提高http.ServeMux的表达能力。在新提案中，[新的ServeMux将支持如下路由策略](https://pkg.go.dev/net/http@go1.22rc1#ServeMux)(来自http.ServeMux的官方文档)：

- “/index.html"路由将匹配任何主机和方法的路径”/index.html"；
- "GET /static/“将匹配路径以”/static/"开头的GET请求；
- "example.com/"可以与任何指向主机为"example.com"的请求匹配；
- “example.com/{$}“会匹配主机为"example.com”、路径为”/“的请求，即"example.com/”；
- "/b/{bucket}/o/{objectname…}“匹配第一段为"b”、第三段为"o"的路径。名称"bucket"表示第二段，"objectname"表示路径的其余部分。

下面就是基于上面的规则编写的示例代码：

```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/index.html", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintln(w, `match /index.html`)
    })
    mux.HandleFunc("GET /static/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintln(w, `match "GET /static/"`)
    })
    mux.HandleFunc("example.com/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintln(w, `match "example.com/"`)
    })
    mux.HandleFunc("example.com/{$}", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintln(w, `match "example.com/{$}"`)
    })
    mux.HandleFunc("/b/{bucket}/o/{objectname...}", func(w http.ResponseWriter, req *http.Request) {
        bucket := req.PathValue("bucket")
        objectname := req.PathValue("objectname")
        fmt.Fprintln(w, `match /b/{bucket}/o/{objectname...}`+":"+"bucket="+bucket+",objectname="+objectname)
    })

    http.ListenAndServe(":8080", mux)
}
```

我们使用curl对上述示例进行一个测试(前提是在/etc/hosts中设置example.com为127.0.0.1)：

```go
$curl localhost:8080/index.html
match /index.html

$curl example.com:8080/static/abc
match "example.com/"

$curl localhost:8080/static/abc
match "GET /static/"

$curl example.com:8080/          
match "example.com/{$}"

$curl example.com:8080/b/mybucket/o/myobject/tonybai
match "example.com/"

$curl localhost:8080/b/mybucket/o/myobject/tonybai
match /b/{bucket}/o/{objectname...}:bucket=mybucket,objectname=myobject/tonybai
```

从测试情况来看，不同路由设置之间存在交集，这就需要路由匹配优先级规则。新版Go ServeMux规定：如果一个请求有两个或两个以上的模式匹配，则更具体(specific)的模式优先。如果P1符合P2请求的严格子集，也就是说，如果P2符合P1及更多的所有请求，那么P1就比P2更具体。

举个例子："/images/thumbnails/“比”/images/“更具体，因此两者都可以注册。前者匹配以”/images/thumbnails/“开头的路径，后者则匹配”/images/"子树中的任何其他路径。

如果两者都不更具体，那么模式就会发生冲突。为了向后兼容，这一规则有一个例外：如果两个模式发生冲突，而其中一个模式有主机(host)，另一个没有，那么有主机的模式优先(比如上面测试中的第二次curl执行)。如果通过ServeMux.Handle或ServeMux.HandleFunc设置的模式与另一个已注册的模式发生冲突，这些函数就会panic。

增强后的ServeMux可能会影响向后兼容性，使用GODEBUG=httpmuxgo121=1可以保留原先的ServeMux行为。

### 3.2 增加math/rand/v2包

在日常开发中，我们多会在生成随机数的场景下使用math/rand包，其他时候使用的较少。但Go 1.22中新增了math/rand/v2包，我之所以将这个列为Go 1.22版本标准库的一次重要变化，是因为这是标准库第一次为某个包建立v2版本包，[按照Russ Cox的说法](https://github.com/golang/go/issues/61716)，这次v2包的创建，为标准库中的其他可能的v2包树立了榜样。创建math/rand/v2可以使Go团队能够在一个相对不常用且风险较低的包中解决工具问题（如gopls、goimports等对v2包的支持），然后再转向更常用、风险更高的包，如sync/v2或encoding/json/v2等。

[新增rand/v2包的直接原因](https://github.com/golang/go/issues/61716)是清理math/rand并修复其中许多悬而未决的问题，特别是使用过时的生成器、慢速算法以及与crypto/rand冲突的问题，这里就不针对v2包举具体的示例了，对该包感兴趣的同学可以自行阅读该包的在线文档，并探索如何使用v2包。

同时，该提案也为标准库中的v2包的创建建立了一种模式，即v2包是原始包的子目录，并且以原始包的API为起点，每个偏离点都要有明确的理由。

想当初，go module刚落地到Go中时，Go module支持两种识别major的两种方式，一种是通过branch或tag号来识别，另外一种就是利用vN目录来定义新包。当时还不是很理解为什么要有vN目录这种方式，现在从math/rand/v2包的增加来看，足以体现出当初module设计时的前瞻性考量了。

### 3.3 [大修Go execution tracer](https://github.com/golang/go/issues/60773)

Go Execution Tracer是解决Go应用性能方面“疑难杂症”的杀手锏级工具，它可以提供Go程序在一段时间内发生的情况的即时视图。这些信息对于了解程序随时间推移的行为非常宝贵，可辅助开发人员对应用进行性能改进。我曾在《[通过实例理解Go Execution Tracer](https://tonybai.com/2021/06/28/understand-go-execution-tracer-by-example)》中对其做过系统的说明。

不过当前版本的Go Execution Tracer在原理和使用方面还存在诸多问题，Google的Michael Knyszek在年初发起了[Execution tracer overhaul的提案](https://github.com/golang/go/issues/60773)，旨在对Go Execution Tracer进行改进，使Go Execution Tracer可扩展到大型Go部署的Go执行跟踪。具体目标如下：

- 使跟踪解析所需的内存占用量仅为当前的一小部分。
- 支持可流式传输的跟踪，以便在无需存储的情况下进行分析。
- 实现部分自描述的跟踪，以减少跟踪消费者的升级负担。
- 修复长期存在的错误，并提供一条清理实现的路径。

在近一年的时间里，Knyszek与Felix Geisendorfer、Nick Ripley、Michael Pratt等一起实现了该提案的目标。

鉴于篇幅，这里就不对新版Tracer的使用做展开说明，有兴趣的童鞋可结合《[通过实例理解Go Execution Tracer](https://tonybai.com/2021/06/28/understand-go-execution-tracer-by-example)》中的使用方法自行体验新版Tracer。

> 注：[新版Tracer的设计文档](https://go.googlesource.com/proposal/+/ac09a140c3d26f8bb62cbad8969c8b154f93ead6/design/60773-execution-tracer-overhaul.md) - https://go.googlesource.com/proposal/+/ac09a140c3d26f8bb62cbad8969c8b154f93ead6/design/60773-execution-tracer-overhaul.md

### 3.4 其他

- “出尔反尔” - [syscall包：取消弃用(undeprecate)](https://github.com/golang/go/issues/60797)

自[Go 1.4版本](https://tonybai.com/2014/11/04/some-changes-in-go-1-4/)以来，syscall包新特性就已经被冻结，并在[Go 1.11版本](https://tonybai.com/2018/11/19/some-changes-in-go-1-11/)中被标记为不推荐使用(deprecate)。Go团队推荐gopher使用golang.org/x/sys/unix或golang.org/x/sys/windows。syscall包的大多数功能都能被golang.org/x/sys包替代，除了下面这几个：

```
syscall.SysProcAttr（类型os/exec.Cmd.SysProcAttr)
syscall.Signal（参考文献os.Signal)
syscall.WaitStatus（参考文献os.(*ProcessState).Sys)
syscall.Stat_t
... ...
```

由于syscall包已经弃用，IDE等工具在开发人员使用上述内容时总是得到警告！这引发了众多开发人员的抱怨。为此，在Go 1.22版本中，syscall取消了弃用状态，但其功能特性依旧保持冻结，不再添加新特性。

- TCPConn to UnixConn：支持zerocopy

[gnet](https://tonybai.com/2021/07/31/io-multiplexing-model-tcp-stream-protocol-parsing-practice-in-go/)作者Andy Pan的提案：[TCPConn to UnixConn：支持zerocopy](https://github.com/golang/go/issues/58808)在Go 1.22落地，具体内容可以看一下[原始提案issue](https://github.com/golang/go/issues/58808)。

- 新增go/version包

在Go 1.21版本发布后，Go团队对Go语言的版本规则做了调整，并明确了[Go语言的向前兼容性和toolchain规则](https://tonybai.com/2023/09/10/understand-go-forward-compatibility-and-toolchain-rule/)，Go 1.22中增加go/version包实现了按照上述版本规则的Go version判断，这个包既用于go工具链，也可以用于Gopher自行开发的工具中。

## 4. 小结

Go 1.22版本具有至少两点重要的里程碑意义：

- 通过对loopvar语义的修正，开启了Go已有“语法坑”的fix之路
- 通过math/rand/v2包树立了Go标准库建立vN版本的模式

“语法坑”fix是否能得到社区正向反馈还是一个未知数，其导致的兼容性问题势必会成为Go社区在升级到Go 1.22版本的重要考虑因素，即便决定升级到Go 1.22，严格的代码审查和测试也是必不可少的。