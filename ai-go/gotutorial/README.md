# @Go 教程(go tutorial)
<details>
  <summary><h2>目录</h2></summary>

[介绍](#介绍)
- [Go语言](#Go语言)
- [main函数和init函数](#main函数和init函数)
- [内置函数](#内置函数)
- [内置类型](#内置类型)
- [包](#包)
- [运算符](#运算符)
- [命名规则](#命名规则)
- [命令](#命令)

[基础](#基础)
- [类型](#类型)
    - [基本类型](#基本类型)
        - [整型](#整型)
        - [浮点型](#浮点型)
        - [数组](#数组)
        - [复数](#复数)
        - [布尔值](#布尔值)
        - [byte和rune类型](#byte和rune类型)
        - [字符串](#字符串)
        - [类型转换](#类型转换)
    - [值的类型](#值的类型)
    - [变量](#变量)
    - [常量](#常量)
    - [指针](#指针)
    - [引用类型](#引用类型)
        - [切片](#切片)
        - [容器](#容器)
        - [管道](#管道)
    - [自定义类型](#自定义类型)
- [流程控制](#流程控制)
    - [if/else条件语句](#if/else条件语句)
    - [for循环语句](#for循环语句)
    - [switch条件语句](#switch条件语句)
    - [select条件语句](#select条件语句)
    - [循环控制](#循环控制)

[函数](#函数)
- [函数定义](#函数定义)
- [参数](#参数)
- [返回值](#返回值)
- [匿名函数](#名函数)
- [闭包和递归](#闭包和递归)
- [延迟调用](#延迟调用)
- [异常处理](#异常处理)

[方法](#方法)
- [方法定义](#方法定义)
- [匿名字段](#匿名字段)
- [表达式](#表达式)
- [方法集](#方法集)
- [自定义error](#自定义error)

[接口](#接口)
- [接口定义](#接口定义)
- [执行机制](#执行机制)
- [接口转换](#接口转换)
- [接口技巧](#接口技巧)

[并发](#并发)
- [并发介绍](#并发介绍)
- [goroutine](#goroutine)
- [chan](#chan)
- [WaitGroup](#WaitGroup)
- [Context](#Context)

[正则](#正则)

[网络编程](#网络编程)
- [http编程](#http编程)
- [TCP编程](#TCP编程)
- [UDP编程](#UDP编程)

</details>

### 一、介绍

### 1 Go语言

#### 1.1 来历
    很久以前，有一个IT公司，这公司有个传统，允许员工拥有20%自由时间来开发实验性项目。在2007的某一天，公司的几个大牛，正在用c++开发一些比较繁琐但是核心的工作，主要包括庞大的分布式集群，大牛觉得很闹心，后来c++委员会来他们公司演讲，说c++将要添加大概35种新特性。这几个大牛的其中一个人，名为：Rob Pike，听后心中一万个xxx飘过，“c++特性还不够多吗？简化c++应该更有成就感吧”。于是乎，Rob Pike和其他几个大牛讨论了一下，怎么解决这个问题，过了一会，Rob Pike说要不我们自己搞个语言吧，名字叫“go”，非常简短，容易拼写。其他几位大牛就说好啊，然后他们找了块白板，在上面写下希望能有哪些功能（详见文尾）。接下来的时间里，大牛们开心的讨论设计这门语言的特性，经过漫长的岁月，他们决定，以c语言为原型，以及借鉴其他语言的一些特性，来解放程序员，解放自己，然后在2009年，go语言诞生。

#### 1.2 思想
    Less can be more 大道至简,小而蕴真 让事情变得复杂很容易，让事情变得简单才难 深刻的工程文化

#### 1.3 指导原则
    Go语言通过减少关键字的数量（25 个）来简化编码过程中的混乱和复杂度。干净、整齐和简洁的语法也能够提高程序的编译速度，因为这些关键字在编译过程中少到甚至不需要符号表来协助解析。
    
    这些方面的工作都是为了减少编码的工作量，甚至可以与 Java 的简化程度相比较。
    
    Go 语言有一种极简抽象艺术家的感觉，因为它只提供了一到两种方法来解决某个问题，这使得开发者们的代码都非常容易阅读和理解。众所周知，代码的可读性是软件工程里最重要的一部分（ 译者注：代码是写给人看的，不是写给机器看的 ）。
    
    这些设计理念没有建立其它概念之上，所以并不会因为牵扯到一些概念而将某个概念复杂化，他们之间是相互独立的。
    
    Go 语言有一套完整的编码规范，你可以在 Go 语言编码规范 页面进行查看。
    
    它不像 Ruby 那样通过实现过程来定义编码规范。作为一门具有明确编码规范的语言，它要求可以采用不同的编译器如 gc 和 gccgo（第 2.1 节）进行编译工作，这对语言本身拥有更好的编码规范起到很大帮助。
    
    LALR 是 Go 语言的语法标准，你也可以在 src/cmd/internal/gc/go.y 中查看到，这种语法标准在编译时不需要符号表来协助解析

#### 1.4 特性
    Go 语言从本质上（程序和结构方面）来实现并发编程。
    
    因为 Go 语言没有类和继承的概念，所以它和 Java 或 C++ 看起来并不相同。但是它通过接口（interface）的概念来实现多态性。Go 语言有一个清晰易懂的轻量级类型系统，在类型之间也没有层级之说。因此可以说这是一门混合型的语言。
    
    在传统的面向对象语言中，使用面向对象编程技术显得非常臃肿，它们总是通过复杂的模式来构建庞大的类型层级，这违背了编程语言应该提升生产力的宗旨。
    
    函数是 Go 语言中的基本构件，它们的使用方法非常灵活。在第六章，我们会看到 Go 语言在函数式编程方面的基本概念。
    
    Go 语言使用静态类型，所以它是类型安全的一门语言，加上通过构建到本地代码，程序的执行速度也非常快。
    
    作为强类型语言，隐式的类型转换是不被允许的，记住一条原则：让所有的东西都是显式的。
    
    Go 语言其实也有一些动态语言的特性（通过关键字 var），所以它对那些逃离 Java 和 .Net 世界而使用 Python、Ruby、PHP 和 JavaScript 的开发者们也具有很大的吸引力。
    
    Go 语言支持交叉编译，比如说你可以在运行 Linux 系统的计算机上开发运行 Windows 下运行的应用程序。这是第一门完全支持 UTF-8 的编程语言，这不仅体现在它可以处理使用 UTF-8 编码的字符串，就连它的源码文件格式都是使用的 UTF-8 编码。Go 语言做到了真正的国际化！

#### 1.5 用途
    Go 语言被设计成一门应用于搭载 Web 服务器，存储集群或类似用途的巨型中央服务器的系统编程语言。对于高性能分布式系统领域而言，Go 语言无疑比大多数其它语言有着更高的开发效率。它提供了海量并行的支持，这对于游戏服务端的开发而言是再好不过了。
    
    Go 语言一个非常好的目标就是实现所谓的复杂事件处理（CEP），这项技术要求海量并行支持，高度的抽象化和高性能。当我们进入到物联网时代，CEP 必然会成为人们关注的焦点。
    
    但是 Go 语言同时也是一门可以用于实现一般目标的语言，例如对于文本的处理，前端展现，甚至像使用脚本一样使用它。
    
    值得注意的是，因为垃圾回收和自动内存分配的原因，Go 语言不适合用来开发对实时性要求很高的软件。
    
    越来越多的谷歌内部的大型分布式应用程序都开始使用 Go 语言来开发，例如谷歌地球的一部分代码就是由 Go 语言完成的。
    
    如果你想知道一些其它组织使用Go语言开发的实际应用项目，你可以到 使用 Go 的组织 页面进行查看。出于隐私保护的考虑，许多公司的项目都没有展示在这个页面。我们将会在第 21 章讨论到一个使用 Go 语言开发的大型存储区域网络（SAN）案例。
    
    在 Chrome 浏览器中内置了一款 Go 语言的编译器用于本地客户端（NaCl），这很可能会被用于在 Chrome OS 中执行 Go 语言开发的应用程序。

#### 1.6 发展目标
    Go 语言的主要目标是将静态语言的安全性和高效性与动态语言的易开发性进行有机结合，达到完美平衡，从而使编程变得更加有乐趣，而不是在艰难抉择中痛苦前行。
    
    因此，Go 语言是一门类型安全和内存安全的编程语言。虽然 Go 语言中仍有指针的存在，但并不允许进行指针运算。
    
    Go 语言的另一个目标是对于网络通信、并发和并行编程的极佳支持，从而更好地利用大量的分布式和多核的计算机，这一点对于谷歌内部的使用来说就非常重要了。设计者通过 goroutine 这种轻量级线程的概念来实现这个目标，然后通过 channel 来实现各个 goroutine 之间的通信。他们实现了分段栈增长和 goroutine 在线程基础上多路复用技术的自动化。
    
    这个特性显然是 Go 语言最强有力的部分，不仅支持了日益重要的多核与多处理器计算机，也弥补了现存编程语言在这方面所存在的不足。
    
    Go 语言中另一个非常重要的特性就是它的构建速度（编译和链接到机器代码的速度），一般情况下构建一个程序的时间只需要数百毫秒到几秒。作为大量使用 C++ 来构建基础设施的谷歌来说，无疑从根本上摆脱了 C++ 在构建速度上非常不理想的噩梦。这不仅极大地提升了开发者的生产力，同时也使得软件开发过程中的代码测试环节更加紧凑，而不必浪费大量的时间在等待程序的构建上。
    
    依赖管理是现今软件开发的一个重要组成部分，但是 C 语言中“头文件”的概念却导致越来越多因为依赖关系而使得构建一个大型的项目需要长达几个小时的时间。人们越来越需要一门具有严格的、简洁的依赖关系分析系统从而能够快速编译的编程语言。这正是 Go 语言采用包模型的根本原因，这个模型通过严格的依赖关系检查机制来加快程序构建的速度，提供了非常好的可量测性。

### 2 main函数和init函数
#### 2.1 init 函数
go语言中init函数用于包(package)的初始化，该函数是go语言的一个重要特性。
有下面的特征：
```
1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
2 每个包可以拥有多个init函数
3 包的每个源文件也可以拥有多个init函数
4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)
5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
```


#### 2.2 main 函数

Go语言程序的默认入口函数(主函数)：func main()
函数体用｛｝一对括号包裹。
```go
package main

func main(){
	//函数体
}
```



#### 2.3 init 函数和 main 函数的异同

相同点：两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用。
不同点：
    1、init可以应用于任意包中，且可以重复定义多个。
    2、main函数只能用于main包中，且只能定义一个。

两个函数的执行顺序：
```
对同一个go文件的init()调用顺序是从上到下的。
对同一个package中不同文件是按文件名字符串比较“从小到大”顺序调用各文件中的init()函数。
对于不同的package，如果不相互依赖的话，按照main包中"先import的后调用"的顺序调用其包中的init()，如果package存在依赖，则先调用最早被依赖的package中的init()，最后调用main函数。

如果init函数中使用了println()或者print()你会发现在执行过程中这两个不会按照你想象中的顺序执
行。这两个函数官方只推荐在测试环境中使用，对于正式环境不要使用。
```



### 3 内置函数

Go 语言拥有一些不需要进行导入操作就可以使用的内置函数。它们有时可以针对不同的类型进行操作，例如：len、cap 和 append，或必须用于系统级的操作，例如：panic。因此，它们需要直接获得编译器的支持。
```go
append  		-- 用来追加元素到数组、slice中,返回修改后的数组、slice
close   		-- 主要用来关闭channel
delete    		-- 从map中删除key对应的value
panic    		-- 停止常规的goroutine  （panic和recover：用来做错误处理）
recover 		-- 允许程序定义goroutine的panic动作
imag    		-- 返回complex的实部   （complex、real imag：用于创建和操作复数）
real    		-- 返回complex的虚部
make    		-- 用来分配内存，返回Type本身(只能应用于slice, map, channel)
new        		-- 用来分配内存，主要用来分配值类型，比如int、struct。返回指向Type的指针
cap        		-- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
copy    		-- 用于复制和连接slice，返回复制的数目
len        		-- 来求长度，比如string、array、slice、map、channel ，返回长度
print、println 	-- 底层打印函数，在部署环境中建议使用 fmt 包
```
内置接口error
```go
type error interface { //只要实现了Error()函数，返回值为String的都实现了err接口
    Error()    String
}
```
### 4 内置类型
值类型：
```go
bool
int(32 or 64), int8, int16, int32, int64
uint(32 or 64), uint8(byte), uint16, uint32, uint64
float32, float64
string
complex64, complex128
array    -- 固定长度的数组
```


引用类型(指针类型)：

```go
slice   -- 序列数组(最常用)
map     -- 映射
chan    -- 管道
```
### 5 包
#### 5.1 源文件
```
编码:源码文件必须是 UTF-8 格式，否则会导致编译器出错。
结束:语句以 ";" 结束，多数时候可以省略。
注释: 持 "//"、"/**/" 两种注释方式，不能嵌套。
命名:采用 camelCasing 风格（驼峰命名法），不建议使用下划线。中文命名,编译环境module为on时编译会报错。
```


#### 5.2 工作空间

Golang 工作空间 ：编译工具对源码目录有严格要求，每个工作空间 (workspace) 必须由 bin、pkg、src 三个目录组成。
```
可在 GOPATH 环境变量列表中添加多个工作空间，但不能和 GOROOT 相同。
export GOPATH=$HOME/projects/golib:$HOME/projects/go
通常 go get使用第一个工作空间保存下载的第三方库。
Golang目前有150个标准的包，覆盖了几乎所有的基础库。
```


#### 5.3 包结构

Golang 包结构 ：所有代码都必须组织在 package 中。
```
• 源文件头部以 "package <name>" 声明包名称。
• 包由同一目录下的多个源码文件组成。
• 包名类似 namespace，与包所在目录名、编译文件名无关。 
• 目录名最好不用 main、all、std 这三个保留名称。
• 可执行文件必须包含 package main，入口函数 main。
```
说明:os.Args 返回命令行参数，os.Exit 终止进程。要获取正确的可执行文件路径，可用 filepath.Abs(exec.LookPath(os.Args[0]))。



package 基本的管理单元：
```go
同一个package下面，可以有非常多的不同文件，只要每个文件的头部 都有如:"package xxx"的相同name
就可以在主方法中使用 xxx.Method()调用不同文件中的方法了。
文件夹名字可以和这个package 名称不一致，
比如我有个文件夹名字是mypackage,其中包含了a.go,b.go, c.go三个文件:
mypackage	
  | --a.go
  | --b.go
  | --c.go

比如a.go中有 Saya(),b.go中有Sayb() 而几个文件共同的package name 确是testpackage
所以在 主函数中调用a.go 和b.go文件中的各自方法只要用，testpackage.Saya() ,testpackage.Sayb()即可
还有默认的init方法，在import进来的时候就去执行了，而且允许每个文件中都有init()这个方法，当然是每个都会执行。
```


导出包：

在Go中，包中成员以名称首字母大小写决定访问权限。首字母大写的名称是被导出的。
在导入包之后，你只能访问包所导出的名字，任何未导出的名字是不能被包外的代码访问的。
Foo 和 FOO 都是被导出的名称。名称 foo 是不会被导出的。
```go
• public: 首字母大写，可被包外访问。
• internal: 首字母小写，仅包内成员可以访问。
该规则适用于全局变量、全局常量、类型、结构字段、函数、方法等。
```



导入包 ：

使用包成员前，必须先用 import 关键字导入，但不能形成导入循环。
import "相对目录/包主文件名"
相对目录是指从<workspace>/pkg/<os_arch>开始的子目录，以标准库为例:
```go
import "fmt"      ->  /usr/local/go/pkg/darwin_amd64/fmt.a
import "os/exec"  ->  /usr/local/go/pkg/darwin_amd64/os/exec.a
```
在导入时，可指定包成员访问方式。比如对包重命名，以避免同名冲突。



import的用法：

```go
import "fmt"   // 最常用的一种形式（系统包）
import "./test" // 导入同一目录下test包中的内容（相对路径）
import "shorturl/model" // 加载gopath/src/shorturl/model模块（绝对路径）
import f "fmt"	// 导入fmt，并给他启别名ｆ
import . "fmt" 	// 将fmt启用别名"."，这样就可以直接使用其内容，而不用再添加fmt。如fmt.Println可以直接写成Println
import  _ "fmt" // 表示不使用该包，而是只是使用该包的init函数，并不显示的使用该包的其他内容。注意：这种形式的import，当import时就执行了fmt包中的init函数，而不能够使用该包的其他函数。

或
import (
    "fmt"
    "./texst"
    _ "fmt"
)
```


自定义路径 ：可通过 meta 设置为代码库设置自定义路径。
server.go

```go
package main

import (
    "fmt"
    "net/http" 
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, `<meta name="go-import"
                   content="test.com/qyuhen/test git https://github.com/qyuhen/test">`)
}

func main() {
    http.HandleFunc("/qyuhen/test", handler)
    http.ListenAndServe(":80", nil)
}
```
该示例使用自定义域名 test.com 重定向到 github。
```go
#go get -v test.com/qyuhen/test

Fetching https://test.com/qyuhen/test?go-get=1
https fetch failed.
Fetching http://test.com/qyuhen/test?go-get=1
Parsing meta tags from http://test.com/qyuhen/test?go-get=1 (status code 200)
get "test.com/qyuhen/test": found meta tag http://test.com/qyuhen/test?go-get=1
test.com/qyuhen/test (download)
test.com/qyuhen/test
```
如此，该库就有两个有效导入路径，可能会导致存储两个本地副本。为此，可以给库添加专门的 "import comment"。当 go get 下载完成后，会检查本地存储路径和该注释是否一致。
```go
github.com/qyuhen/test/abc.go

package test // import "test.com/qyuhen/test"
func Hello() {
    println("Hello, Custom import path!")
}
```
如继续用 github 路径，会导致 go build 失败。
```go
#go get -v github.com/qyuhen/test

github.com/qyuhen/test (download)
package github.com/qyuhen/test
    imports github.com/qyuhen/test
    imports github.com/qyuhen/test: expects import "test.com/qyuhen/test"
```
这就强制包用户使用唯一路径，也便于日后将包迁移到其他位置。



Golang初始化,初始化函数:

```
• 每个源文件都可以定义一个或多个初始化函数。 
• 编译器不保证多个初始化函数执行次序。
• 初始化函数在单一线程被调 ，仅执行一次。 
• 初始化函数在包所有全局变量初始化后执行。 
• 在所有初始化函数结束后才执行 main.main()。 
• 无法调用初始化函数。
```

因为无法保证初始化函数执行顺序，因此全局变量应该直接用 var 初始化。
```go
var now = time.Now()

func init() {
    fmt.Printf("now: %v\n", now)
}

func init() {
    fmt.Printf("since: %v\n", time.Now().Sub(now))
}
```
可在初始化函数中使用 goroutine，可等待其结束。
```go
package main

import (
    "fmt"
    "time"
)

var now = time.Now()

func main() {
    fmt.Println("main:", int(time.Now().Sub(now).Seconds()))
}

func init() {
    fmt.Println("init:", int(time.Now().Sub(now).Seconds()))
    w := make(chan bool)
    go func() {
        time.Sleep(time.Second * 3)
        w <- true
    }()
    <-w 
}
```
输出:
```
init: 0
main: 3
```
#### 5.3 文档
Golang 文档 ：扩展工具 godoc 能自动提取注释生成帮助文档。
```
• 仅和成员相邻 (中间没有空行) 的注释被当做帮助信息。 
• 相邻行会合并成同一段落，用空行分隔段落。
• 缩进表示格式化文本，比如示例代码。
• 自动转换 URL 为链接。
• 自动合并多个源码文件中的 package 文档。 
• 无法显式 package main 中的成员文档。
```
package
```
• 建议用专门的 doc.go 保存 package 帮助信息。
• 包文档第一整句 (中英文句号结束) 被当做 packages 列表说明。
```
只要 Example 测试函数名称符合以下规范即可：
| | 格式 | 示例 |
|--|----|----|
|package| Example, Example_suffix|Example_test|
|func	|ExampleF, ExampleF_suffix|	ExampleHello|
|type	|ExampleT, ExampleT_suffix|	ExampleUser, ExampleUser_copy|
|method	|ExampleT_M, ExampleT_M_suffix|	ExampleUser_ToString|

说明:使用 suffix 作为示例名称，其首字母必须小写。如果文件中仅有一个 Example 函数，且调用了该文件中的其他成员，那么示例会显示整个文件内容，而不仅仅是测试函数自己。

非测试源码文件中以 BUG(author) 开始的注释，会在帮助文档 Bugs 节点中显示。
```
// BUG(yuhen): memory leak.
```



### 6 运算符

全部运算符、分隔符，以及其他符号。
|符号（说明）|||||||||
|--|--|--|--|--|--|--|--|--|
|+（加号）	|&	|+=	|&=	|&&（逻辑与）	|==（等于）| !=（不等于） |	(|	)|
|-（减号）	|\|	|-=	|\|=	|\|\|（逻辑或）	|<（小于）| <=（小于等于） | [ | ] |
|*（乘号）	|^	|*=|	^=	|<-	|>（大于）| >=（大于等于） |	{|	}|
|/（除号）	|<<	|/=	|<<=|	++|	=	|:=	|,|	;|
|%（取模）	|>>	|%=	|>>=|	--	|!（逻辑非）	|...	|.|	:|
|&^|		&^=	||||||||

算术运算符

| 符号 |   说明    | 示例 |
| :--: | :-------: | :--: |
|  +   |   加号    |      |
|  -   |   减号    |      |
|  *   |   乘号    |      |
|  /   |   除号    |      |
|  %   | 取模/求余 |      |
|  ++  |   自增    |      |
|  --  |   自减    |      |

位运算符
位运算符对整数在内存中的二进制位进行操作。下表列出了位运算符 &, |, 和 ^ 的计算：

| p    | q    | p & q | p \| q | p ^ q |
| :--- | :--- | :---- | :----- | :---- |
| 0    | 0    | 0     | 0      | 0     |
| 0    | 1    | 0     | 1      | 1     |
| 1    | 1    | 1     | 1      | 0     |
| 1    | 0    | 0     | 1      | 1     |

假定 A = 60; B = 13; 其二进制数转换为：

```
A = 0011 1100
B = 0000 1101

-----------------
A&B = 0000 1100
A|B = 0011 1101
A^B = 0011 0001
```



Go 语言支持的位运算符如下表所示。假定 A 为60，B 为13：

| 运算符 | 说明                                                         | 实例                                   |
| :----- | :----------------------------------------------------------- | :------------------------------------- |
| &      | 按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 | (A & B) 结果为 12, 二进制为 0000 1100  |
| \|     | 按位或运算符"\|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或 | (A \| B) 结果为 61, 二进制为 0011 1101 |
| ^      | 按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。 | (A ^ B) 结果为 49, 二进制为 0011 0001  |
| <<     | 左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 | A << 2 结果为 240 ，二进制为 1111 0000 |
| >>     | 右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。 | A >> 2                                 |
```go
a := 0
a |= 1 << 2 		// 0000100: 在 bit2 设置标志位。
a |= 1 << 6 		// 1000100: 在 bit6 设置标志位 
a = a &^ (1 << 6) 	// 0000100: 清除 bit6 标志位。
```

比较运算符

| 符号 | 说明                                           | 示例                                   |
| :--: | :--------------------------------------------- | :------------------------------------- |
|  :=  | 简单的赋值运算符，将一个表达式的值赋给一个左值 | C := A + B 将 A + B 表达式结果赋值给 C |
|  =   | 简单的赋值运算符，将一个表达式的值赋给一个左值 | C = A + B 将 A + B 表达式结果赋值给 C  |
|  +=  | 相加后再赋值                                   | C += A 等于 C = C + A                  |
|  -=  | 相减后再赋值                                   | C -= A 等于 C = C - A                  |
|  *=  | 相乘后再赋值                                   | C *= A 等于 C = C * A                  |
|  /=  | 相除后再赋值                                   | C /= A 等于 C = C / A                  |
|  %=  | 求余后再赋值                                   | C %= A 等于 C = C % A                  |
|  &=  | 按位与后赋值                                   | C &= 2 等于 C = C & 2                  |
| \|=  | 按位或后赋值                                   | C \|= 2 等于 C = C \| 2                |
|  ^=  | 按位异或后赋值                                 | C ^= 2 等于 C = C ^ 2                  |
| >>=  | 右移后赋值                                     | C >>= 2 等于 C = C >> 2                |
| <<=  | 左移后赋值                                     | C <<= 2 等于 C = C << 2                |

括号

|         符号          |             说明             | 示例 |
| :-------------------: | :--------------------------: | :--: |
| (                   ) | 小括号："("前括号，")"后括号 |      |
| [                  ]  | 中括号："["前括号，"]"后括号 |      |
| {                   } | 大括号："{"前括号，"}"后括号 |      |

逻辑运算符

| 符号 | 说明                                                         | 示例              |
| :--: | :----------------------------------------------------------- | :---------------- |
|  &&  | 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 | (A && B) 为 False |
| \|\| | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。 | (A\|\|B)为True    |
|  ！  | 逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。 | !(A && B) 为 True |

关系运算符

下表列出了所有Go语言的关系运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述                                                         | 实例              |
| :----: | :----------------------------------------------------------- | :---------------- |
|   ==   | 检查两个值是否相等，如果相等返回 True 否则返回 False。       | (A == B) 为 False |
|   !=   | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。   | (A != B) 为 True  |
|   >    | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。   | (A > B) 为 False  |
|   <    | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。   | (A < B) 为 True   |
|   >=   | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 | (A >= B) 为 False |
|   <=   | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 | (A <= B) 为 True  |

### 7 命名规则
正如命名在其它语言中的地位，它在 Go 中同样重要。有时它们甚至会影响语义： 例如，某个名称在包外是否可见，就取决于其首个字符是否为大写字母。 因此有必要花点时间来讨论 Go 程序中的命名约定。
#### 7.1 包名
 当一个包被导入后，包名就会成了内容的访问器。在以下代码
```go
import "bytes"
```
​		之后，被导入的包就能通过 bytes.Buffer 来引用了。 若所有人都以相同的名称来引用其内容将大有裨益， 这也就意味着包应当有个恰当的名称：其名称应该简洁明了而易于理解。按照惯例， 包应当以小写的单个单词来命名，且不应使用下划线或驼峰记法。err 的命名就是出于简短考虑的，因为任何使用该包的人都会键入该名称。 不必担心引用次序的冲突。包名就是导入时所需的唯一默认名称， 它并不需要在所有源码中保持唯一，即便在少数发生冲突的情况下， 也可为导入的包选择一个别名来局部使用。 无论如何，通过文件名来判定使用的包，都是不会产生混淆的。
​		另一个约定就是包名应为其源码目录的基本名称。在 src/encoding/base64 中的包应作为 "encoding/base64" 导入，其包名应为 base64， 而非 encoding_base64 或 encodingBase64。
包的导入者可通过包名来引用其内容，因此包中的可导出名称可以此来避免冲突。 （请勿使用 import . 记法，它可以简化必须在被测试包外运行的测试， 除此之外应尽量避免使用。）例如，bufio 包中的缓存读取器类型叫做 Reader 而非 BufReader，因为用户将它看做 bufio.Reader，这是个清楚而简洁的名称。 此外，由于被导入的项总是通过它们的包名来确定，因此 bufio.Reader 不会与 io.Reader 发生冲突。同样，用于创建 ring.Ring 的新实例的函数（这就是 Go 中的构造函数）一般会称之为 NewRing，但由于 Ring 是该包所导出的唯一类型，且该包也叫 ring，因此它可以只叫做 New，它跟在包的后面，就像 ring.New。使用包结构可以帮助你选择好的名称。

#### 7.2 获取器
Go 并不对获取器（getter）和设置器（setter）提供自动支持。 你应当自己提供获取器和设置器，通常很值得这样做，但若要将 Get 放到获取器的名字中，既不符合习惯，也没有必要。若你有个名为 owner （小写，未导出）的字段，其获取器应当名为 Owner（大写，可导出）而非 GetOwner。大写字母即为可导出的这种规定为区分方法和字段提供了便利。 若要提供设置器方法，SetOwner 是个不错的选择。两个命名看起来都很合理：
```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```
#### 7.3 接口命名
按照约定，只包含一个方法的接口应当以该方法的名称加上 - er 后缀来命名，如 Reader、Writer、 Formatter、CloseNotifier 等。

诸如此类的命名有很多，遵循它们及其代表的函数名会让事情变得简单。 Read、Write、Close、Flush、 String 等都具有典型的签名和意义。为避免冲突，请不要用这些名称为你的方法命名， 除非你明确知道它们的签名和意义相同。反之，若你的类型实现了的方法， 与一个众所周知的类型的方法拥有相同的含义，那就使用相同的命名。 请将字符串转换方法命名为 String 而非 ToString。

#### 7.4 驼峰 命名
最后，Go 中的约定是使用 MixedCaps 或 mixedCaps 而不是下划线来编写多个单词组成的命名

### 8 命令
已安装了golang环境，你可以在命令行执行go命令查看相关的Go语言命令：
```go
# go

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        work        workspace maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildconstraint build constraints
        buildmode       build modes
        c               calling between Go and C
        cache           build and test caching
        environment     environment variables
        filetype        file types
        go.mod          the go.mod file
        gopath          GOPATH environment variable
        gopath-get      legacy GOPATH go get
        goproxy         module proxy protocol
        importpath      import path syntax
        modules         modules, module versions, and more
        module-get      module-aware go get
        module-auth     module authentication using go.sum
        packages        package lists and patterns
        private         configuration for downloading non-public code
        testflag        testing flags
        testfunc        testing functions
        vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.
```
go env用于打印Go语言的环境信息。
go run命令可以编译并运行命令源码文件。
go get可以根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装。
go build命令用于编译我们指定的源码文件或代码包以及它们的依赖包。
go install用于编译并安装指定的代码包及它们的依赖包。
go clean命令会删除掉执行其它命令时产生的一些文件和目录。
go doc命令可以打印附于Go语言程序实体上的文档。我们可以通过把程序实体的标识符作为该命令的参数来达到查看其文档的目的。
go test命令用于对Go语言编写的程序进行测试。
go list命令的作用是列出指定的代码包的信息。
go fix会把指定代码包的所有Go语言源码文件中的旧版本代码修正为新版本的代码。
go vet是一个用于检查Go语言源码中静态错误的简单工具。
go tool pprof命令来交互式的访问概要文件的内容。
用户go get无法下载翻墙的时候的包，可以用gopm下载

什么是gopm
在nodejs中我们有npm，可以通过npm来下载安装一些依赖包。在go中也开发了类似的东西，那就是gopm。这玩意儿是七牛开发的。在这里说下，七牛公司大部分程序都是用go语言编写的，所以开发出这么一个方便的东西肯定也是合情合理的。
gopm安装：
```go
go get github.com/gpmgo/gopm
go install github.com/gpmgo/gopm
```
通过这个命令来安装插件，默认的会存放到GOBIN，如果没有配置%GOBIN%环境变量，那么会默认安装到%GOPATH%下的bin目录，为了我们操作方便，我们把GOBIN加到%PATH%下。
使用方法：
```go
NAME:
   Gopm - Go Package Manager

USAGE:
   Gopm [global options] command [command options] [arguments...]

VERSION:
   0.8.8.0307 Beta

COMMANDS:
   list         list all dependencies of current project
   gen          generate a gopmfile for current Go project
   get          fetch remote package(s) and dependencies
   bin          download and link dependencies and build binary
   config       configure gopm settings
   run          link dependencies and go run
   test         link dependencies and go test
   build        link dependencies and go build
   install      link dependencies and go install
   clean        clean all temporary files
   update       check and update gopm resources including itself
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --noterm, -n         disable color output
   --strict, -s         strict mode
   --debug, -d          debug mode
   --help, -h           show help
   --version, -v        print the version
```
下载示例：
```
gopm get -g -v -u golang.org/x/tools/cmd/goimports
```

### 基础
### 1.1 类型
#### 1.1.1 基本类型
Golang 更明确的数字类型命名，支持 Unicode，支持常用数据结构。

| 类型          | 长度(字节) | 默认值 | 说明                                      |
| :------------ | :--------- | :----- | :---------------------------------------- |
| bool          | 1          | false  |                                           |
| byte          | 1          | 0      | uint8                                     |
| rune          | 4          | 0      | Unicode Code Point, int32                 |
| int, uint     | 4或8       | 0      | 32 或 64 位                               |
| int8, uint8   | 1          | 0      | -128 ~ 127, 0 ~ 255，byte是uint8 的别名   |
| int16, uint16 | 2          | 0      | -32768 ~ 32767, 0 ~ 65535                 |
| int32, uint32 | 4          | 0      | -21亿~ 21亿, 0 ~ 42亿，rune是int32 的别名 |
| int64, uint64 | 8          | 0      |                                           |
| float32       | 4          | 0.0    |                                           |
| float64       | 8          | 0.0    |                                           |
| complex64     | 8          |        |                                           |
| complex128    | 16         |        |                                           |
| uintptr       | 4或8       |        | 以存储指针的 uint32 或 uint64 整数        |
| array         |            |        | 值类型                                    |
| struct        |            |        | 值类型                                    |
| string        |            | ""     | UTF-8 字符串                              |
| slice         |            | nil    | 引用类型                                  |
| map           |            | nil    | 引用类型                                  |
| channel       |            | nil    | 引用类型                                  |
| interface     |            | nil    | 接口                                      |
| function      |            | nil    | 函数                                      |

支持八进制、 六进制，以及科学记数法。标准库 math 定义了各数字类型取值范围。
```go
a, b, c, d := 071, 0x1F, 1e9, math.MinInt16
```
空指针值 nil，而非C/C++ NULL。

##### 1.1.1.1 整型
```go
整型分为以下两个大类： 按长度分为：int8、int16、int32、int64对应的无符号整型：uint8、uint16、uint32、uint64

其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。
```
##### 1.1.1.2 浮点型
```go
Go语言支持两种浮点型数：float32和float64。这两种浮点型数据格式遵循IEEE 754标准： float32 的浮点数的最大范围约为3.4e38，可以使用常量定义：math.MaxFloat32。 float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。
```
##### 1.1.1.3 数组
```go
1. 数组：是同一种数据类型的固定长度的序列。
2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
for i := 0; i < len(a); i++ {
}
for index, v := range a {
}
5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
8.指针数组 [n]*T，数组指针 *[n]T。
```
数组初始化：
1、一维数组
```go
全局：
    var arr0 [5]int = [5]int{1, 2, 3}
    var arr1 = [5]int{1, 2, 3, 4, 5}
    var arr2 = [...]int{1, 2, 3, 4, 5, 6}
    var str = [5]string{3: "hello world", 4: "tom"}

局部：
    a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
```
代码：
```go
package main

import (
	"fmt"
)

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

func main() {
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)
}
```
结果：
```
[1 2 3 0 0] [1 2 3 4 5] [1 2 3 4 5 6] [   hello world tom]
[1 2 0] [1 2 3 4] [0 0 100 0 200] [{user1 10} {user2 20}]
```

2、多维数组
```go
全局
    var arr0 [5][3]int
    var arr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

局部：
    a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
```
代码：
```go
package main

import (
	"fmt"
)

var arr0 [5][3]int
var arr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func main() {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
	fmt.Println(arr0, arr1)
	fmt.Println(a, b)
}
```
结果：
```
[[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]] [[1 2 3] [7 8 9]]
[[1 2 3] [4 5 6]] [[1 1] [2 2] [3 3]]
```

##### 1.1.1.4 复数
```
complex64和complex128

复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。
```

##### 1.1.1.5 布尔值
```
Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值

注意：
    布尔类型变量的默认值为false。
    Go 语言中不允许将整型强制转换为布尔型.
    布尔型无法参与数值运算，也无法与其他类型进行转换。
```

##### 1.1.1.6 byte和rune类型
```go
组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：
    var a := '中'
    var b := 'x'

Go 语言的字符有以下两种：
    uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
    rune类型，代表一个 UTF-8字符。

当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。 Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾
// 遍历字符串
func traversalString() {
    s := "pprof.cn博客"
    for i := 0; i < len(s); i++ { //byte
        fmt.Printf("%v(%c) ", s[i], s[i])
    }
    fmt.Println()
    for _, r := range s { //rune
        fmt.Printf("%v(%c) ", r, r)
    }
    fmt.Println()
}

输出：
112(p) 112(p) 114(r) 111(o) 102(f) 46(.) 99(c) 110(n) 229(å) 141() 154() 229(å) 174(®) 162(¢)
112(p) 112(p) 114(r) 111(o) 102(f) 46(.) 99(c) 110(n) 21338(博) 23458(客)

因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。

字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
```

##### 1.1.1.7 字符串
Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用UTF-8编码。 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符，例如：
```go
s1 := "hello"
s2 := "你好"
```



字符串转义：

Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

| 转义 | 含义                               |
| ---- | ---------------------------------- |
| \r   | 回车符（返回行首）                 |
| \n   | 换行符（直接跳到下一行的同列位置） |
| \t   | 制表符                             |
| \'   | 单引号                             |
| \"   | 双引号                             |
| \    | 反斜杠                             |

举个例子，我们要打印一个Windows平台下的一个文件路径：

```go
package main
import (
    "fmt"
)
func main() {
    fmt.Println("str := \"c:\\pprof\\main.exe\"")
}
```



多行字符串：

Go语言中要定义一个多行字符串时，就必须使用`反引号`字符：

```go
s1 := `第一行
第二行
第三行
`
fmt.Println(s1)
```

反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。



字符串的常用操作：

| 方法                                | 介绍           |
| ----------------------------------- | -------------- |
| len(str)                            | 求长度         |
| +或fmt.Sprintf                      | 拼接字符串     |
| strings.Split                       | 分割           |
| strings.Contains                    | 判断是否包含   |
| strings.HasPrefix,strings.HasSuffix | 前缀/后缀判断  |
| strings.Index(),strings.LastIndex() | 子串出现的位置 |
| strings.Join(a[]string, sep string) | join操作       |



修改字符串：
要修改字符串，需要先将其转换成`[]rune或[]byte`，完成后再转换为`string`。无论哪种转换，都会重新分配内存，并复制字节数组。

```go
    func changeString() {
        s1 := "hello"
        // 强制类型转换
        byteS1 := []byte(s1)
        byteS1[0] = 'H'
        fmt.Println(string(byteS1))

        s2 := "博客"
        runeS2 := []rune(s2)
        runeS2[0] = '狗'
        fmt.Println(string(runeS2))
    }
```

##### 1.1.1.8 类型转换
变量在定义时没有明确的初始化时会赋值为 零值 。
```go
数值类型为 0 ，
布尔类型为 false ，
字符串为 "" （空字符串）。
```


Golang 不支持隐式类型转换，即便是从窄向宽转换也不行。

```go
package main

var b byte = 100
// var n int = b
// ./main.go:5:5: cannot use b (type byte) as type int in assignment
var n int = int(b) // 显式转换
func main() {

}
```

同样不能将其他类型当 bool 值使用。

```go
package main

func main() {
	a := 100
	if a { // Error: non-bool a (type int) used as if condition
		println("true")
	}
}
```



类型转换：

```go
类型转换用于将一种数据类型的变量转换为另外一种类型的变量。Go 语言类型转换基本格式如下：
表达式 T(v) 将值 v 转换为类型 T 。
type_name(expression)
type_name 为类型，expression 为表达式。
```

实例：

```go
将整型转化为浮点型，并计算结果，将结果赋值给浮点型变量：
package main
import "fmt"

func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}

输出结果：
mean 的值为: 3.400000


package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 42
	fmt.Printf("i value is : %v , type is : %v \n", i, reflect.TypeOf(i))
	var f float64 = float64(i)
	fmt.Printf("f value is : %v , type is : %v \n", f, reflect.TypeOf(f))
	var u uint = uint(f)
	fmt.Printf("u value is : %v , type is : %v \n", u, reflect.TypeOf(u))
}

输出结果：
i value is : 42 , type is : int 
f value is : 42 , type is : float64 
u value is : 42 , type is : uint 


更加简单的形式：
package main
import (
	"fmt"
	"reflect"
)

func main() {
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Printf("i value is : %v , type is : %v \n", i, reflect.TypeOf(i))
	fmt.Printf("f value is : %v , type is : %v \n", f, reflect.TypeOf(f))
	fmt.Printf("u value is : %v , type is : %v \n", u, reflect.TypeOf(u))
}

输出结果：
i value is : 42 , type is : int 
f value is : 42 , type is : float64 
u value is : 42 , type is : uint 
```



类型推导

在定义一个变量却并不显式指定其类型时（使用 := 语法或者 var = 表达式语法）【全局变量不适用】， 变量的类型由（等号）右侧的值推导得出。

当右值定义了类型时，新变量的类型与其相同：

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int
	j := i // j 也是一个 int
	fmt.Printf("i type is : %v\n", reflect.TypeOf(i))
	fmt.Printf("j type is : %v\n", reflect.TypeOf(j))
}

输出结果：
i type is : int
j type is : int


变量就可能是 int 、 float64 或 complex128 :
package main
import (
	"fmt"
	"reflect"
)

func main() {
	i := 42           
	f := 3.142        
	g := 0.867 + 0.5i 
	fmt.Printf("i type is : %v\n", reflect.TypeOf(i))
	fmt.Printf("f type is : %v\n", reflect.TypeOf(f))
	fmt.Printf("g type is : %v\n", reflect.TypeOf(g))
}

输出结果：
i type is : int
f type is : float64
g type is : complex128
```



Go各种类型转换及函数的高级用法

字符串转整形

```
将字符串转换为 int 类型 
strconv.ParseInt(str,base,bitSize)
str：要转换的字符串 
base：进位制（2 进制到 36 进制） 
bitSize：指定整数类型（0:int、8:int8、16:int16、32:int32、64:int64） 
返回转换后的结果和转换时遇到的错误 
如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
ParseUint 功能同 ParseInt 一样，只不过返回 uint 类型整数
```

Atoi 相当于 ParseInt(s, 10, 0)
通常使用这个函数，而不使用 ParseInt
该方法的源码是：

```
// Itoa is shorthand for FormatInt(i, 10).
func Itoa(i int) string {
    return FormatInt(int64(i), 10)
}

可以看出是FormatInt方法的简单实现。
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	i, ok := strconv.ParseInt("1000", 10, 0)
	if ok == nil {
		fmt.Printf("ParseInt , i is %v , type is %v\n", i, reflect.TypeOf(i))
	}

	ui, ok := strconv.ParseUint("100", 10, 0)
	if ok == nil {
		fmt.Printf("ParseUint , ui is %v , type is %v\n", ui, reflect.TypeOf(i))
	}

	oi, ok := strconv.Atoi("100")
	if ok == nil {
		fmt.Printf("Atoi , oi is %v , type is %v\n", oi, reflect.TypeOf(i))
	}

}
```

输出结果：

```
ParseInt , i is 1000 , type is int64
ParseUint , ui is 100 , type is int64
Atoi , oi is 100 , type is int64
```

整形转字符串

```
FormatInt int 型整数 i 转换为字符串形式 
strconv.FormatInt.(i,base)
FormatUint 将 uint 型整数 i 转换为字符串形式 
strconv.FormatUint.(i,base)
base：进位制（2 进制到 36 进制） 
大于 10 进制的数，返回值使用小写字母 ‘a’ 到 ‘z’

Itoa 相当于 FormatInt(i, 10)
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var i int64
	i = 0x100
	str := strconv.FormatInt(i, 10) // FormatInt第二个参数表示进制，10表示十进制。
	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))
}
```

输出结果：

```
256
string
```

AppendInt 将 int 型整数 i 转换为字符串形式，并追加到 []byte 的尾部
strconv.AppendInt([]byte, i, base)
AppendUint 将 uint 型整数 i 转换为字符串形式，并追加到 dst 的尾部
strconv.AppendUint([]byte, i, base)
i：要转换的字符串
base：进位制
返回追加后的 []byte

```
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := make([]byte, 0)
	b = strconv.AppendInt(b, -2048, 16)
	fmt.Printf("%s\n", b)
}
```

输出结果：

```
-800
```

字节转32位整形

```
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	b := []byte{0x00, 0x00, 0x03, 0xe8}
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	fmt.Println(x)
}
// 其中binary.BigEndian表示字节序，相应的还有little endian。通俗的说法叫大端、小端。
```

输出结果：

```
1000
```

32位整形转字节

```
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

func main() {
	var x int32
	x = 106
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	b := bytesBuffer.Bytes()
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
}
```

输出结果：

```
[0 0 0 106]
[]uint8
```

字节转字符串

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	b := []byte{97, 98, 99, 100}
	str := string(b)
	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))
}
```

输出结果：

```go
abcd
string
```

字符串转字节

```go
package main

import (
	"fmt"
)

func main() {
	str := "abcd"
	b := []byte(str)
	fmt.Println(b)
}
```

输出结果：

```go
[97 98 99 100]
```

字符串转布尔值 ParseBool

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, err := strconv.ParseBool("1")
	fmt.Printf("string 1 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("t")
	fmt.Printf("string t 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("T")
	fmt.Printf("string T 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("true")
	fmt.Printf("string true 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("True")
	fmt.Printf("string True 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("TRUE")
	fmt.Printf("string TRUE 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("TRue")
	fmt.Printf("string TRue 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("")
	fmt.Printf("string '' 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("0")
	fmt.Printf("string 0 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("f")
	fmt.Printf("string f 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("F")
	fmt.Printf("string F 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("false")
	fmt.Printf("string false 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("False")
	fmt.Printf("string False 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("FALSE")
	fmt.Printf("string FALSE 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("FALse")
	fmt.Printf("string FALse 转 bool ：%v , err is : %v\n", b, err)
	b, err = strconv.ParseBool("abc")
	fmt.Printf("string abc 转 bool ：%v , err is : %v\n", b, err)
}
```

输出结果：

```go
string 1 转 bool ：true , err is : <nil>
string t 转 bool ：true , err is : <nil>
string T 转 bool ：true , err is : <nil>
string true 转 bool ：true , err is : <nil>
string True 转 bool ：true , err is : <nil>
string TRUE 转 bool ：true , err is : <nil>
string TRue 转 bool ：false , err is : strconv.ParseBool: parsing "TRue": invalid syntax
string '' 转 bool ：false , err is : strconv.ParseBool: parsing "": invalid syntax
string 0 转 bool ：false , err is : <nil>
string f 转 bool ：false , err is : <nil>
string F 转 bool ：false , err is : <nil>
string false 转 bool ：false , err is : <nil>
string False 转 bool ：false , err is : <nil>
string FALSE 转 bool ：false , err is : <nil>
string FALse 转 bool ：false , err is : strconv.ParseBool: parsing "FALse": invalid syntax
string abc 转 bool ：false , err is : strconv.ParseBool: parsing "abc": invalid syntax

ParseBool 将字符串转换为布尔值 
它接受真值：1, t, T, TRUE, true, True 
它接受假值：0, f, F, FALSE, false, False. 
其它任何值都返回一个错误
```

布尔值转换为字符串 FormatBool

```go
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	t := strconv.FormatBool(true)
	f := strconv.FormatBool(false)
	fmt.Printf("t is %v , t type is %v\n", t, reflect.TypeOf(t))
	fmt.Printf("f is %v , f type is %v\n", f, reflect.TypeOf(f))
}
```

输出结果：

```go
t is true , t type is string
f is false , f type is string
```

AppendBool 将布尔类型转换为字符串
然后将结果追加到 []byte 的尾部，返回追加后的 []byte

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	rst := []byte{}
	fmt.Printf("[]byte{} is %s\n", rst)
	rst = strconv.AppendBool(rst, true)
	fmt.Printf("appended true []byte{} is %s\n", rst)
	rst = strconv.AppendBool(rst, false)
	fmt.Printf("appended false []byte{} is %s\n", rst)
}
```

输出结果：

```go
[]byte{} is 
appended true []byte{} is true
appended false []byte{} is truefalse
```

将字符串转换为浮点数
strconv.ParseFloat(str,bitSize)
str：要转换的字符串
bitSize：指定浮点类型（32:float32、64:float64）
如果 str 是合法的格式，而且接近一个浮点值，
则返回浮点数的四舍五入值（依据 IEEE754 的四舍五入标准）
如果 str 不是合法的格式，则返回“语法错误”
如果转换结果超出 bitSize 范围，则返回“超出范围”

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "0.12345678901234567890"
	f, err := strconv.ParseFloat(s, 32)
	fmt.Println(f, err)
	fmt.Println(float32(f))
	fmt.Println("-----")
	f, err = strconv.ParseFloat(s, 64)
	fmt.Println(f, err)
	fmt.Println(float64(f))
	fmt.Println("-----")
	str := "abcd"
	f, err = strconv.ParseFloat(str, 32)
	fmt.Println(f, err)
}
```

输出结果：

```go
0.12345679104328156 <nil>
0.12345679
-----
0.12345678901234568 <nil>
0.12345678901234568
-----
0 strconv.ParseFloat: parsing "abcd": invalid syntax
```

将浮点数转换为字符串值

```go
strconv.FormatFloat(f,fmt,prec,bitSize)
f：要转换的浮点数 
fmt：格式标记（b、e、E、,f、g、G） 
prec：精度（数字部分的长度，不包括指数部分） 
bitSize：指定浮点类型（32:float32、64:float64）

格式标记： 
‘b’ (-ddddp±ddd，二进制指数) 
‘e’ (-d.dddde±dd，十进制指数) 
‘E’ (-d.ddddE±dd，十进制指数) 
‘f’ (-ddd.dddd，没有指数) 
‘g’ (‘e’:大指数，’f’:其它情况) 
‘G’ (‘E’:大指数，’f’:其它情况)

如果格式标记为 ‘e’，’E’和’f’，则 prec 表示小数点后的数字位数 
如果格式标记为 ‘g’，’G’，则 prec 表示总的数字位数（整数部分+小数部分）
package main

import (
	"fmt"
	"strconv"
)

func main() {
	f := 100.12345678901234567890123456789
	fmt.Println(strconv.FormatFloat(f, 'b', 5, 32))
	fmt.Println(strconv.FormatFloat(f, 'e', 5, 32))
	fmt.Println(strconv.FormatFloat(f, 'E', 5, 32))
	fmt.Println(strconv.FormatFloat(f, 'f', 5, 32))
	fmt.Println(strconv.FormatFloat(f, 'g', 5, 32))
	fmt.Println(strconv.FormatFloat(f, 'G', 5, 32))
	fmt.Println(strconv.FormatFloat(f, 'b', 30, 32))
	fmt.Println(strconv.FormatFloat(f, 'e', 30, 32))
	fmt.Println(strconv.FormatFloat(f, 'E', 30, 32))
	fmt.Println(strconv.FormatFloat(f, 'f', 30, 32))
	fmt.Println(strconv.FormatFloat(f, 'g', 30, 32))
	fmt.Println(strconv.FormatFloat(f, 'G', 30, 32))
}
```

输出结果：

```go
13123382p-17
1.00123e+02
1.00123E+02
100.12346
100.12
100.12
13123382p-17
1.001234588623046875000000000000e+02
1.001234588623046875000000000000E+02
100.123458862304687500000000000000
100.1234588623046875
100.1234588623046875
```

AppendFloat 将浮点数 f 转换为字符串值，并将转换结果追加到 []byte 的尾部
返回追加后的 []byte

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	f := 100.12345678901234567890123456789
	b := make([]byte, 0)
	b = strconv.AppendFloat(b, f, 'f', 5, 32)
	b = append(b, " "...)
	b = strconv.AppendFloat(b, f, 'e', 5, 32)
	fmt.Printf("%s\n", b)
}
```

输出结果：

```
100.12346 1.00123e+02
```

Quote 将字符串 s 转换为“双引号”引起来的字符串
其中的特殊字符将被转换为“转义字符”
不可显示的字符”将被转换为“转义字符”

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.Quote(`C:\Windows`))
}
```

输出结果：

```go
"C:\\Windows"
```

AppendQuote 将字符串 s 转换为“双引号”引起来的字符串，
并将结果追加到 []byte 的尾部，返回追加后的 []byte
其中的特殊字符将被转换为“转义字符”

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := `C:\Windows`
	b := make([]byte, 0)
	b = strconv.AppendQuote(b, s)
	fmt.Printf("%s\n", b)
}
```

输出结果：

```go
"C:\\Windows"
```

QuoteToASCII 将字符串 s 转换为“双引号”引起来的 ASCII 字符串
“非 ASCII 字符”和“特殊字符”将被转换为“转义字符”

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	asc := strconv.QuoteToASCII("Hello 世界！")
	fmt.Println(asc)
}
```

输出结果：

```go
"Hello \u4e16\u754c\uff01"
```

AppendQuoteToASCII 将字符串 s 转换为“双引号”引起来的 ASCII 字符串，
并将结果追加到 []byte 的尾部，返回追加后的 []byte
非 ASCII 字符”和“特殊字符”将被转换为“转义字符”

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "Hello 世界！"
	b := make([]byte, 0)
	b = strconv.AppendQuoteToASCII(b, s)
	fmt.Printf("%s\n", b)
}
```

输出结果：

```go
"Hello \u4e16\u754c\uff01"
```

QuoteRune 将 Unicode 字符转换为“单引号”引起来的字符串
特殊字符”将被转换为“转义字符”

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := strconv.QuoteRune('哈')
	fmt.Println(str)
}
```

输出结果：

```go
'哈'
```

AppendQuoteRune 将 Unicode 字符转换为“单引号”引起来的字符串，
并将结果追加到 []byte 的尾部，返回追加后的 []byte
特殊字符”将被转换为“转义字符”

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := make([]byte, 0)
	b = strconv.AppendQuoteRune(b, '哈')
	fmt.Printf("%s\n", b)
}
```

输出结果：

```go
'哈'
```

QuoteRuneToASCII 将 Unicode 字符转换为“单引号”引起来的 ASCII 字符串
“非 ASCII 字符”和“特殊字符”将被转换为“转义字符”

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	asc := strconv.QuoteRuneToASCII('哈')
	fmt.Println(asc)
}
```

输出结果：

```go
'\u54c8'
```

AppendQuoteRune 将 Unicode 字符转换为“单引号”引起来的 ASCII 字符串，
并将结果追加到 []byte 的尾部，返回追加后的 []byte
“非 ASCII 字符”和“特殊字符”将被转换为“转义字符”

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := make([]byte, 0)
	b = strconv.AppendQuoteRuneToASCII(b, '哈')
	fmt.Printf("%s\n", b)
}
```

输出结果：

```go
'\u54c8'
```

CanBackquote 判断字符串 s 是否可以表示为一个单行的“反引号”字符串
字符串中不能含有控制字符（除了 \t）和“反引号”字符，否则返回 false

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := strconv.CanBackquote("C:\\Windows\n")
	fmt.Printf("\\n is %v\n", b)
	b = strconv.CanBackquote("C:\\Windows\r")
	fmt.Printf("\\r is %v\n", b)
	b = strconv.CanBackquote("C:\\Windows\f")
	fmt.Printf("\\f is %v\n", b)
	b = strconv.CanBackquote("C:\\Windows\t")
	fmt.Printf("\\t is %v\n", b)
	b = strconv.CanBackquote("C:\\Windows`")
	fmt.Printf("` is %v\n", b)
}
```

输出结果：

```go
\n is false
\r is false
\f is false
\t is true
` is false
```

UnquoteChar 将 s 中的第一个字符“取消转义”并解码

s：转义后的字符串
quote：字符串使用的“引号符”（用于对引号符“取消转义”）
value： 解码后的字符
multibyte：value 是否为多字节字符
tail： 字符串 s 除去 value 后的剩余部分
error： 返回 s 中是否存在语法错误

参数 quote 为“引号符”
如果设置为单引号，则 s 中允许出现 ' 字符，不允许出现单独的 ' 字符
如果设置为双引号，则 s 中允许出现 " 字符，不允许出现单独的 " 字符
如果设置为 0，则不允许出现 ' 或 " 字符，可以出现单独的 ' 或 " 字符

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := `\"大\\家\\好！\"`
	c, mb, sr, _ := strconv.UnquoteChar(s, '"')
	fmt.Printf("%-3c %v\n", c, mb)
	for ; len(sr) > 0; c, mb, sr, _ = strconv.UnquoteChar(sr, '"') {
		fmt.Printf("%-3c %v\n", c, mb)
	}

}
```

输出结果：

```go
"   false
"   false
大   true
\   false
家   true
\   false
好   true
！   true
```

Unquote 将“带引号的字符串” s 转换为常规的字符串（不带引号和转义字符）
s 可以是“单引号”、“双引号”或“反引号”引起来的字符串（包括引号本身）
如果 s 是单引号引起来的字符串，则返回该该字符串代表的字符

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	sr, err := strconv.Unquote("\"大\t家\t好！\"")
	fmt.Println(sr, err)
	sr, err = strconv.Unquote(`'大家好！'`)
	fmt.Println(sr, err)
	sr, err = strconv.Unquote("'好'")
	fmt.Println(sr, err)
	sr, err = strconv.Unquote("大\\t家\\t好！")
	fmt.Println(sr, err)
}
```

输出结果：

```go
大	家	好！ <nil>
 invalid syntax
好 <nil>
 invalid syntax
```

IsPrint 判断 Unicode 字符 r 是否是一个可显示的字符
可否显示并不是你想象的那样，比如空格可以显示，而\t则不能显示

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.IsPrint('a'))
	fmt.Println(strconv.IsPrint('好'))
	fmt.Println(strconv.IsPrint(' '))
	fmt.Println(strconv.IsPrint('\t'))
	fmt.Println(strconv.IsPrint('\n'))
	fmt.Println(strconv.IsPrint(0))
}
```

输出结果：

```go
true
true
true
false
false
false
```

### 函数

```go

```

### 方法
```go

```

### 接口
```go

```

### 并发
```go

```

### 正则
```go

```

### 网络编程
```go

```