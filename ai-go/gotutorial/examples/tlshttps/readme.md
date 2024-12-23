## 使用 golang 部署运行 tls 的 https 服务时，不用停机，高效证书下放，如何实现？

#### 第一部分

这篇文章主要介绍如何在应用 golang 语言开发 http/https 服务时，如何让 tls 自动获取证书，而不必在证书更新或重置以后，还要重启服务器来让业务重新起效，本文分成三部分，第一部分会介绍 tls 加密的常用加密算法进行分析总结，虽然与主干关系不特别大，但是该段络会帮你厘清一个日常使用中，非常容易被混淆的问题；第二部分会重点介绍如何部署一个不需要重启也能 tls 自动更新的高抽象度的 http 服务；第三部分会对整个文章进行总结，相信基于该文章的学习，你一定会对 tls 领域和流量监测、安全防护领域常见的算法有相对深刻的理解，也对如何高度抽象一个自签名的 golang 服务有全新的认识。那么文章开始！

我之前已经有文章分享过如果来做爬虫，用 python 的 scrapy 和基于 node 的 puppeteer 之间的优劣对比。也分享过我之前抓取某国外站点的时候，由于页面验签要正确的传输参数，在 post 请求中提交到后台才能返回正确的结果，这个在国内的也有很多站点有这种验签机制的存在，比如说我在写基于 puppeteer 的自动发布文章应用的时候，一些平台，像今日头条平台、像简书用浏览器缓存的 cookies 就可以避免重复登录，直接模拟浏览器登录发布文章就好了。但是像百度的百家号，只有 cookie 是无法成功验签的，还要获取下放的 token，才能完整成功完成发布百家号的流程，总之，就是“兵来将挡，水来土淹”。

但是你有考虑过，这些平台其实也是基于`tls`的验签机会来保护客户端和服务器端数据交互安全性的，但是遵循的底层原理却是各有不同的。比如说`JA3`指纹算法，它能基于`TLS`客户端与服务端之间握手消息内容生成一个指纹，具体来说，就是在进行`TLS`握手时，客户端会发送一些包含有关自身支持的加密套件、`TLS/SSL`版本等信息的消息给服务器，服务器会回应类似的消息，`JA3`也是根据这些传输消息生成的指纹。它是基于四层网络传输协议，在第四层，即传输层被使用的。而我上面举的例子，比如说向浏览器种`token`，把它添加到传输报文的报文头中，服务器对于浏览器提交带着的这个`token`进行校验，以确定其合法性，其实它作用的是应用层协议(第七层)中使用的[身份验证](https://cloud.tencent.com/product/mfas?from_column=20065&from=20065)机制，而并非传输层(四层网络传输协议的第 4 层)，这里要注意区分。

对于 tls 的生成，其实有很多算法，但是`JA3`算法被最广泛使用，那相比于其它算法，它有什么样的优势和劣势呢？我做了个图表进行总结，供大家参考：

| 算法                 | 优点                                                                                                                                                                                                                                                                                                                      | 缺点                                                                              |
| :------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | :-------------------------------------------------------------------------------- |
| JA3 指纹算法         | 可以识别 TLS 客户端版本；可以基于握手消息内容生成指纹，具有更高的精度；在不同设备和操作系统上的一致性较好； 它是一种开放标准，任何人都可以实现它并将其集成到自己的应用程序或工具中，这使它成为一个通用的、可扩展的方案； 可用来验证 TLS 是否被篡改，与 SSL 证书指纹不同，JA3 算法可以检测中间人攻击等网络层面的攻击行为； | 无法判断代理层的影响；无法识别使用自定义密码套件的客户端；只能用于 TLS 握手识别。 |
| SSL/TLS 证书指纹算法 | 不受代理层、客户端版本等因素的影响；可以识别采用自定义密码套件的客户端。                                                                                                                                                                                                                                                  | 无法识别中间人攻击；证书签发机构可能存在错误或欺诈。                              |
| HTTP 消息头指纹算法  | 可以识别代理层、CDN 等影响；适用范围广，可用于 HTTP 流量识别。                                                                                                                                                                                                                                                            | 可能存在误判；对于加密流量而言，只能识别应用层信息。                              |
| TCP/IP 指纹算法      | 可以识别代理层、NAT 等影响；可以在网络层识别流量。                                                                                                                                                                                                                                                                        | 具有较高的误判率；对于加密流量而言，只能识别网络层信息。                          |
| DNS 指纹算法         | 可以在域名解析阶段进行指纹识别；不受代理层等因素的影响。                                                                                                                                                                                                                                                                  | 无法识别加密流量；可能存在 DNS 缓存的干扰。                                       |

python 中有 JA3 算法也非常常用，特别是下列一些场景：

1.  识别恶意软件：通过 JA3 算法可以识别出具有特定 JA3 指纹的恶意软件，从而帮助网络安全人员及时发现和防范攻击。
2.  流量识别：JA3 算法可以用于流量识别和分类，帮助工程师进行流量监控、分析等操作。
3.  加密流量检测：由于 JA3 算法可以识别 TLS 客户端版本和加密套件，因此它可以被用来检测加密流量是否合法以及是否遵循最佳实践。
4.  网络侦查：JA3 算法可以用于识别用户访问网站的类型、客户端操作系统、浏览器版本等信息，帮助网络侦查人员了解对方的技术能力和行为习惯。
5.  安全策略制定：通过对 JA3 数据的统计和分析，可以了解不同客户端的使用情况，并据此制定相应的安全策略和措施，提高网络安全性。

#### 第二部分

那如何来部署 golang 服务，让其支持动态更新`TLS certificates`而无需停机？我们知道`Transport Layer Security(TLS)`是一种基于`SSLv3`的加密协议，用于在两个站点之间加密和解密流量。换言之，`TLS`确保你正在访问的站点和你之间数据的传输数据不被侦测到。这是通过相互交换数字证书来实现的：一个存在于`web`服务器上的私有证书，另一个通常随 web 浏览器分发的公共证书。

在生产环境，服务都是以安全方式运行，但服务验证经过一定周期就会过期。然后对于服务响应去验证、重新生成，同时不用停机，就可以重新使用生成的验签证书。这篇文章，演示一下`TLS`验证是在基于`golang`语言的`HTTPS`服务是如何使用的。

这篇教程有先要满足下面这些先决条件。

- 要对`客户端-服务端`模型要有基本理解
- `Golang`的基础知识

###### 配置 HTTP Server

开始这篇文章之前，先演示一个简单的`HTTP`服务，只需要使用`http.ListenAndServe`函数启动一个`HTTP`服务，再用`http.HandleFunc`函数对于特定`endpoint`注册一个`response handler`。

开始，配置`HTTP`服务：

```go
package main
import (
    "net/http"
    "fmt"
    "log"
)

func main(){
    mux := http.NewServeMux()
    mux.HandleFun("/",func(res http.ResponseWriter, req *http.Request){
        fmt.Fprint(res, "Running Http Servcer")
    })

    srv := &http.Server{
        Addr: fmt.Sprintf(":%d",8080),
        Handler: mux,
    }
    //在8080端口运行
    log.Fatal(srv.ListenAndServe())
}
```

上面例子，用`go run server.go`，会在`HTTP`服务的`8080`端口运行，浏览器输入`http://localhost:8080`，你就会看到`Hello World!`输出在屏幕上。

`srv.ListenAndServe()`调用了 go 语言`HTTP`服务的标准库配置，然而，你可以使用`Server`结构类型来定制`server`。

启动一个`HTTPS`服务，用配置调用函数`srv.ListenAndServeTLS(certFile,keyFile)`，与`srv.ListenAndServe()`函数相似。`ListenAndServe`和`ListenAndServeTLS`两个函数在`HTTP`包和`Server`结构中都是可用的。

`ListenAndServeTLS`与`ListenAndServe`函数类似，除了前者对`HTTPS`服务的支持。

```javascript
func ListenAndServeTLS(certFile string,keyFile string) error
```

从以上函数签名上可以看出，两者唯一的区别在于额外的`certFile`和`keyFile`参数，分别代表`SSL Certificate`文件的路径和`private key`文件。

###### 生成`private key`和`SSL certificate`

以下是生成根`key和certificate`的步骤:

- 1. 生成根`key`

```
openssl genrsa -des3 -out rootCA.key 4096
```

- 1. 创建和对根`certificate`做`self-signature(自签名)`

```
openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1024 -out rootCA.crt
```

接着，按下面的方式为每个服务生成`certificate`:

- 1. 创建`certificate key` `openssl genrsa -out localhost.key 2048`
- 1. 创建`certificate-signing request(CSR)`。`CSR`是在哪里指定你想生成的`certificate`的详情。根`key`的拥有者将执行`request`来生成`certificate`。当创建`CSR`时，重要的是指定提供 IP 地址的`Common Name`，或者服务的域名，否则`certificate`无法验证。

```javascript
openssl req -new -key localhost.key -out localhost.csr
```

- 1. 使用`TLS CSR`和密钥以及`CA`根密钥生成证书

```javascript
openssl x509 -req -in localhost.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out localhost.crt -days 500 -sha256
```

最后，遵循同样步骤为每个客户端生成`certificate`。

###### 配置一个 HTTPS Server

既然有`private key`和`certificate`文件，就可以对先前的 go 程序做修改了，这次用`ListenAndServeTLS`代替。

```go
package main

import(
    "net/http"
    "fmt"
    "log"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func( res http.ResponseWriter, req *http.Request ) {
        fmt.Fprint( res, "Running HTTPS Server!!" )
    })

    srv := &http.Server{
        Addr: fmt.Sprintf(":%d", 8443),
        Handler:           mux,
    }

    // run server on port "8443"
    log.Fatal(srv.ListenAndServeTLS("localhost.crt", "localhost.key"))
}
```

运行以上程序，将使用包含运行文件同级目录下的`localhost.crt`作为`certFile`，使用`localhost.key`作为`keyFile`启动一个`HTTPS`服务。再浏览器访问`https://localhost:8443`，或者[命令行工具](https://cloud.tencent.com/product/cli?from_column=20065&from=20065)`(CLI)`，可以看到如下输出：

```javascript
$ curl https://localhost:8443/ --cacert rootCA.crt --key client.key --cert client.crt

Running HTTPS Server!!
```

就是这样!这是大多数人启动 HTTPS 服务器必须做的事情。是 Go 管理 TLS 通信的默认行为和功能。

###### 配置 HTTPS 服务以自动更新证书

当运行以上的`HTTPS`服务，你把`certFile`和`keyFile`传给了`ListenAndServeTLS`函数，然而，如果因为`certificate`过期`certFile`和`keyFile`发生变化，服务需要重启来使变化生效，为了克服这种中断导致的的短暂服务中断，可以使用`net/http`包的`TLSConfig`。

`cryto/tls`包的`TLSConfig`结构会配置服务的`TLS`参数，包括服务证书等。所有`TLSConfig`的参数都是可选项，同时也要注意给`TLSconfig`的参数配置选项赋以空结构，就等同于赋个`nil`值给它。然而，配置`GetCertificate`字段却是相当有益的。

```go
type Config struct{
    GetCertificate func(*ClientHelloInfo) (*Certicate,error)
}
```

`TLSConfig`的`GetCertificate`字段会基于`ClientHelloInfo`返回证书。

我需要实现`GetCertificate`闭包函数，该函数使用`tls.LoadX509KeyPair(certFile string, keyFile string) 或者 tls.X509KeyPair(certFile []byte, keyFile []byte)`函数来获取证书，举两个例子：

```go
// 这是使用tls.LoadX509KeyPair(certFile string, keyFile string)函数的例子
func GetCertificate(certFile string, keyFile string) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
    cert, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
        log.Fatal(err)
    }
    return func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
        return &cert, nil
    }
}
// 这是使用tls.X509KeyPair的例子：
func GetCertificate(certData []byte, keyData []byte) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
    cert, err := tls.X509KeyPair(certData, keyData)
    if err != nil {
        log.Fatal(err)
    }
    return func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
        return &cert, nil
    }
}
// 以上两个函数使用后都会返回一个闭包，可以用作tls.config结构体中的GetCertificate字段。
```

现在我将使用`TLSConfig`字段值创建`Server`结构：

```go
package main

import (
    "crypto/tls"
    "fmt"
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func( res http.ResponseWriter, req *http.Request ) {
        fmt.Fprint( res, "Running HTTPS Server!!\n" )
    })
    srv := &http.Server{
        Addr: fmt.Sprintf(":%d", 8443),
        Handler:           mux,
        TLSConfig: &tls.Config{
            GetCertificate: func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
                // 总是获取最新的localhost.crt和localhost.key
                // 将证书文件保存在全局位置中，这样创建新证书时可以更新它们，并且该闭包函数可以引用它。
                cert, err := tls.LoadX509KeyPair("localhost.crt", "localhost.key")
                if err != nil {
                    return nil, err
                }
                return &cert, nil
            },
        },
    }

    // 在8443端口运行服务
    log.Fatal(srv.ListenAndServeTLS("", ""))
}
```

以上程序中，我实现了`GetCertificate`闭包函数，通过使用`LoadX509KeyPair`及证收和之前创建的私有文件，返回了一个类型为`Certificate`的`cert`对象。同时函数了一个`error`，方便调试和追踪。

由于我正在用`TLS`配置，为`srv`服务对象做预配置，我不需要给`srv.ListenAndServeTLS`函数调用提供`certFile`和`keyFile`。运行服务，它会像之前一样运行，但是区别点就在于，我从调用对象中抽象了所有的服务配置，因此这些配置即便更新，也会动态加载，而不必重启服务。

```javascript
$ curl https://localhost:8443/ --cacert rootCA.crt --key client.key --cert client.crt

Running HTTPS Server!!
```

#### 第三部分

好了，这篇有关如何抽象`TLS`服务配置，达到不需要重启服务就能加载变更证书的文章就分享至些，感谢阅读，我特别将可用于`tls`加密的指纹算法提到第一段来讲，并把`JA3`指纹算法在四层服务传输协议中的使用，和浏览器`token`验签属于应用层(七层网络传输服务)协议中使用的身份验证机制做了区分，方便你在今后使用的过程中有更深刻的理解。

四层服务传输协议（TCP/IP 模型）和七层服务传输协议（OSI 模型）的对比如下：

| OSI 模型        | TCP/IP 模型     |
| :-------------- | :-------------- |
| 应用层（7）     | 应用层（4）     |
| 表示层（6）     |                 |
| 会话层（5）     |                 |
| 传输层（4）     | 传输层（3）     |
| 网络层（3）     | 网际层（2）     |
| 数据链路层（2） | 网络接口层（1） |
| 物理层（1）     |                 |

其中，TCP/IP 模型将原本的“会话层、表示层、应用层”合并为一个应用层，而将原本的“网络层、数据链路层、物理层”分别放在了网际层、网络接口层两个层级中。

下面是四层服务传输协议（TCP/IP 模型）和七层服务传输协议（OSI 模型）每层常见的使用场景的对比图表：

| OSI 模型        | TCP/IP 模型     | 常见的使用场景             |
| :-------------- | :-------------- | :------------------------- |
| 应用层（7）     | 应用层（4）     | HTTP、FTP、SMTP 等应用程序 |
| 表示层（6）     |                 | 数据压缩和加密             |
| 会话层（5）     |                 | 远程访问和 RPC             |
| 传输层（4）     | 传输层（3）     | TCP 和 UDP 协议            |
| 网络层（3）     | 网际层（2）     | IP、ICMP、ARP 等           |
| 数据链路层（2） | 网络接口层（1） | 以太网、WiFi、ATM 等       |
| 物理层（1）     |                 | 传输介质及物理设备         |

在 OSI 模型中，每一层都有自己的功能和特点。应用层负责定义应用程序之间的交互规则；表示层用于对应用数据进行编码和解码；会话层为不同主机上的应用程序之间建立会话连接；传输层提供端到端的可靠[数据传输服务](https://cloud.tencent.com/product/dts?from_column=20065&from=20065)；网络层负责将数据包从源主机传输到目标主机；数据链路层管理网络节点之间的数据帧传输；物理层负责传输介质及物理设备。

在 TCP/IP 模型中，应用层包含了 OSI 模型的应用层、表示层和会话层的功能；传输层提供端到端的可靠数据传输服务；网际层负责将数据包从源主机传输到目标主机；网络接口层管理网络节点之间的数据帧传输。

总之，在网络通信中，无论是使用 OSI 模型还是 TCP/IP 模型，每一层都有各自的功能和特点，能够互相配合完成数据传输和网络通信的任务。

# go 语言实现双向 TLS 认证的 REST Service

# 服务器端代码如下

```go
package main

import (
    "fmt"
    "log"
    "flag"
    "net/http"
    "io/ioutil"
    "crypto/tls"
    "crypto/x509"
    "encoding/json"
    "github.com/gorilla/mux"
)

var (
    port       int
    hostname   string
    caroots    string
    keyfile    string
    signcert   string
)

func init() {
    flag.IntVar(&port,          "port",     8080,       "The host port on which the REST server will listen")
    flag.StringVar(&hostname,   "hostname", "0.0.0.0",  "The host name on which the REST server will listen")
    flag.StringVar(&caroots,    "caroot",   "",         "Path to file containing PEM-encoded trusted certificate(s) for clients")
    flag.StringVar(&keyfile,    "key",      "",         "Path to file containing PEM-encoded key file for service")
    flag.StringVar(&signcert,   "signcert", "",         "Path to file containing PEM-encoded sign certificate for service")
}

func startServer(address string, caroots string, keyfile string, signcert string, router *mux.Router) {
    pool := x509.NewCertPool()

    caCrt, err := ioutil.ReadFile(caroots)
    if err != nil {
        log.Fatalln("ReadFile err:", err)
    }
    pool.AppendCertsFromPEM(caCrt)

    s := &http.Server{
            Addr:    address,
            Handler: router,
            TLSConfig: &tls.Config{
                MinVersion: tls.VersionTLS12,
                ClientCAs:  pool,
                ClientAuth: tls.RequireAndVerifyClientCert,
            },
    }
    err = s.ListenAndServeTLS(signcert, keyfile)
    if err != nil {
        log.Fatalln("ListenAndServeTLS err:", err)
    }
}

func SayHello(w http.ResponseWriter, r *http.Request) {
    log.Println("Entry SayHello")
    res := map[string]string {"hello": "world"}

    b, err := json.Marshal(res)
    if err == nil {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")
        w.Write(b)
    }

    log.Println("Exit SayHello")
}

func main() {
    flag.Parse()

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/service/hello", SayHello).Methods("GET")

    var address = fmt.Sprintf("%s:%d", hostname, port)
    fmt.Println("Server listen on", address)
    startServer(address, caroots, keyfile, signcert, router)

    fmt.Println("Exit main")
}
```

其中 TLS 配置项 ClientAuth: tls.RequireAndVerifyClientCert 表明需要对客户端认证，也就是要完成服务器和客户端的双向认证。

# 生成服务端证书

- 生成服务端私钥
  `$ openssl genrsa -out server.key 2048`
  或者
  `$ openssl genrsa -des3 -out server.key 2048`
  此时需要用户输入密码，然后每次用到私钥的时候都需要再次输入密码。
  注意这个私钥非常重要，通常需要安全保存并且把读写权限改成 600
- 生成服务端证书请求文件
  `$ openssl req -new -key server.key -out server.csr -subj "/C=CN/ST=BJ/L=beijing/O=myorganization/OU=mygroup/CN=myname"`

注意这一步生成的是证书请求文件，不是证书文件，下面才会生成证书文件。

# 生成客户端端证书

这个过程和生成服务端证书一样

- 生成客户端私钥
  `$ openssl genrsa -out client.key 2048`
- 生成客户证书请求文件
  `$ openssl req -new -key client.key -out client.csr -subj "/C=CN/ST=BJ/L=beijing/O=myorganization/OU=mygroup/CN=myname"`

# 生成服务器和客户端经过签名的证书

证书请求文件 csr 生成以后，需要将其发送给 CA 认证机构进行签发以生成真正的证书文件，当然在我们例子里，我们使用 OpenSSL 对该证书进行自签发。

- 生成根证书私钥
  `$ openssl genrsa -out ca.key 2048`
- 生成根证书请求文件
  `$ openssl req -new -key ca.key -out ca.csr -subj "/C=CN/ST=BJ/L=beijing/O=myorganization/OU=mygroup/CN=myname"`
- 生成自签名的根证书文件
  `$ openssl x509 -req -days 365 -sha1 -extensions v3_ca -signkey ca.key -in ca.csr -out ca.cer`
- 利用已签名根证书生成服务端证书和客户端证书
  ** 生成服务端证书
  `$ openssl x509 -req -days 365 -sha1 -extensions v3_req -CA ca.cer -CAkey ca.key -CAcreateserial -in server.csr -out server.cer`
  ** 生成客户端证书
  `$ openssl x509 -req -days 365 -sha1 -extensions v3_req -CA ca.cer -CAkey ca.key -CAcreateserial -in client.csr -out client.cer`

注意，关于 extensions 参数值 v3_ca/v3_req 的含义请参考 openssl.cnf 配置文件

```csharp
$ locate openssl.cnf
$ vim /etc/pki/tls/openssl.cnf

[ v3_req ]
    basicConstraints = CA:FALSE
...
[ v3_ca ]
    basicConstraints = CA:true
```

其中最重要的区别是，标识这是不是一个 CA 证书。

# 编译运行服务端程序

```go
$ go build main.go
$ ./main -caroot ./ca.cer -key ./server.key -signcert ./server.cer
Server listen on 0.0.0.0:8080
```

# 运行客户端程序

```python
$ curl --cacert ./ca.cer --key ./client.key --cert ./client.cer https://localhost:8080/service/hello
curl: (60) Peer's certificate has an invalid signature.
More details here: http://curl.haxx.se/docs/sslcerts.html

curl performs SSL certificate verification by default, using a "bundle"
 of Certificate Authority (CA) public keys (CA certs). If the default
 bundle file isn't adequate, you can specify an alternate file
 using the --cacert option.
If this HTTPS server uses a certificate signed by a CA represented in
 the bundle, the certificate verification probably failed due to a
 problem with the certificate (it might be expired, or the name might
 not match the domain name in the URL).
If you'd like to turn off curl's verification of the certificate, use
 the -k (or --insecure) option.
```

很遗憾，你应该看到上面的错误信息，再查看服务端的日志：
`2017/09/27 22:42:49 http: TLS handshake error from [::1]:56168: remote error: tls: bad certificate`

提示证书无效，原因是我们的证书里 Commone Name 这个字段填的值是 myname，而当前服务器运行的域名是 localhost，他们不匹配，Common Name 是要授予证书的服务器域名或主机名。

我们修改修改服务器端证书，重新生成：

```objectivec
$ openssl req -new -key server.key -out server.csr -subj "/C=CN/ST=BJ/L=beijing/O=myorganization/OU=mygroup/CN=localhost"
$ openssl x509 -req -days 365 -sha1 -extensions v3_req -CA ca.cer -CAkey ca.key -CAcreateserial -in server.csr -out server.cer
```

再运行，看看是不是想要的结果：
(注意，根证书和客户端证书不需要重新生成)

```cpp
$ curl --cacert ./ca.cer --key ./client.key --cert ./client.cer https://localhost:8080/service/hello
{"hello":"world"}
```

这就是我们想要的结果。
同理，如果使用真实机器主机名或者域名，例如主机名 saturn，则

```objectivec
$ openssl req -new -key server.key -out server.csr -subj "/C=CN/ST=BJ/L=beijing/O=myorganization/OU=mygroup/CN=saturn"
$ openssl x509 -req -days 365 -sha1 -extensions v3_req -CA ca.cer -CAkey ca.key -CAcreateserial -in server.csr -out server.cer
$ curl --cacert ./ca.cer --key ./client.key --cert ./client.cer https://saturn:8080/service/hello
{"hello":"world"}
```

查看证书内容

```php
$ openssl x509 -in server.cer -text -noout 2>&1| head -n 15
Certificate:
    Data:
        Version: 1 (0x0)
        Serial Number: 15163366668719918823 (0xd26f19a5700c8ee7)
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: C=CN, ST=BJ, L=beijing, O=myorganization, OU=mygroup, CN=myname
        Validity
            Not Before: Sep 27 14:44:07 2017 GMT
            Not After : Sep 27 14:44:07 2018 GMT
        Subject: C=CN, ST=BJ, L=beijing, O=myorganization, OU=mygroup, CN=localhost
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:9e:f0:05:0f:1f:4d:43:36:65:86:36:5e:80:bb:
```

里面表示了当前证书信息，以及签发者的信息。

# 总结

每个节点(不管是客户端还是服务端)都有一个证书文件和 key 文件，他们用来互相加密解密；因为证书里面包含 public key，key 文件里面包含 private key；他们构成一对密钥对，是互为加解密的。

根证书是所有节点公用的，不管是客户端还是服务端，都要先注册根证书(通常这个过程是注册到操作系统信任的根证书数据库里面，在咱们这个例子里面没有这么做，因为这是一个临时的根证书，只在服务端和客户端命令行中指定了一下)，以示这个根证书是可信的， 然后当需要验证对方的证书时，因为待验证的证书是通过这个根证书签名的，我们信任根证书，所以推导出也可以信任对方的证书。

所以如果需要实现双向认证，那么每一端都需要三个文件

- {node}.cer: PEM certificate
  己方证书文件，将会被发给对方，让对方认证
- {node}..key: PEM RSA private key
  己方 private key 文件，用来解密经己方证书(因为包含己方 public key)加密的内容，这个加密过程一般是由对方实施的。
- ca.cer: PEM certificate
  根证书文件，用来验证对方发过来的证书文件，所有由同一个根证书签名的证书都应该能验证通过。

# https 示例

## 一.生成证书

通过 openssl 工具生成证书

### 1.1 安装 openssl

macos 通过 brew 安装

```bash
brew install openssl
```

### 1.2 生成跟证书私钥

```bash
openssl genrsa -out ca.key 4096
```

### 1.3 准备配置文件

```bash
vim ca.conf
```

内容如下

```bash
[ req ]
default_bits       = 4096
distinguished_name = req_distinguished_name

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = JiangSu
localityName                = Locality Name (eg, city)
localityName_default        = NanJing
organizationName            = Organization Name (eg, company)
organizationName_default    = Sheld
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = Ted CA Test
```

生成根证书签发申请文件(csr 文件)

```bash
openssl req -new -sha256 -out ca.csr -key ca.key -config ca.conf
```

该命令含义如下：

req——执行证书签发命令

-new——新证书签发请求

-key——指定私钥路径

-out——输出的 csr 文件的路径

### 1.4 自签发根证书(cer 文件)

```bash
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.crt
```

该命令的含义如下：

x509——生成 x509 格式证书

-req——输入 csr 文件

-days——证书的有效期（天）

-signkey——签发证书的私钥

-in——要输入的 csr 文件

-out——输出的 cer 证书文件

### 1.5 生成服务端私钥

```bash
openssl genrsa -out server.key 2048
```

### 1.6 准备配置文件，得到 server.conf

```bash
vim server.conf
```

内容如下

```bash
[ req ]
default_bits       = 2048
distinguished_name = req_distinguished_name
req_extensions     = req_ext

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = JiangSu
localityName                = Locality Name (eg, city)
localityName_default        = NanJing
organizationName            = Organization Name (eg, company)
organizationName_default    = Sheld
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = server.com

[ req_ext ]
subjectAltName = @alt_names

[alt_names]
DNS.1   = server.com
```

生成服务端证书申请文件

```bash
openssl req -new -sha256 -out server.csr -key server.key -config server.conf
```

参考 1.3 输入服务端证书信息

### 1.7 用 CA 证书签发服务端证书

```bash
openssl x509 -req -days 3650 -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.crt -extensions req_ext -extfile server.conf
```

这里有必要解释一下这几个参数：

-CA——指定 CA 证书的路径

-CAkey——指定 CA 证书的私钥路径

-CAserial——指定证书序列号文件的路径

-CAcreateserial——表示创建证书序列号文件(即上方提到的 serial 文件)，创建的序列号文件默认名称为-CA，指定的证书名称后加上.srl 后缀

### 1.8 生成客户端私钥

```bash
openssl genrsa -out client.key 2048
```

### 1.9 准备配置文件，得到 client.conf

```bash
vim client.conf
```

内容如下

```bash
[ req ]
default_bits       = 2048
distinguished_name = req_distinguished_name
req_extensions     = req_ext

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = HeNan
localityName                = Locality Name (eg, city)
localityName_default        = AnYang
organizationName            = Organization Name (eg, company)
organizationName_default    = Sheld_client
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = server.com

[ req_ext ]
subjectAltName = @alt_names

[alt_names]
DNS.1   = server.com
DNS.2    = localhost
```

生成客户端证书申请文件

```bash
openssl req -new -sha256 -out client.csr -key client.key -config client.conf
```

### 1.10 用跟证书签发客户端证书

```bash
openssl x509 -req -days 3650 -CA ca.crt -CAkey ca.key -CAserial ca.srl -in client.csr -out client.crt -extensions req_ext -extfile client.conf
```

需要注意的是，上方签发服务端证书时已经使用-CAcreateserial 生成过 ca.srl 文件，因此这里不需要带上这个参数了。

## 二. gin 框架实现 https，服务端不需要验证客户端证书的示例

### 2.1 项目引入 gin 框架

```bash
go get -u github.com/gin-gonic/gin
```

### 2.2 服务端代码

```bash
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.New()
    router.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "success",
        })
    })

    // 可以直接用
    //router.RunTLS("0.0.0.0:10679", "./certs/server.cer", "./certs/server.key")
    server := &http.Server{Addr: "0.0.0.0:10679", Handler: router}
    _ = server.ListenAndServeTLS("./certs/server.cer", "./certs/server.key")
}
```

### 2.3 安全提示

打开浏览器访问 https://localhost:10679/test

因为浏览器并不信任证书的颁发机构，浏览器会有安全提示，当然可以点高级直接强行访问页面，同样可以返回结果，很多文章也是到这里就结束了。强迫症的我还是要解决一下。

![img](../../../../images/c4efebcd2c69d0c5884e92dacc83da45.png)

### 2.4 给浏览器添加自己生成的 CA 证书

本文选择 firefox 浏览器，因为设置可以轻松添加信任 CA 证书。chrome 则是直接掉起系统的证书，没有找到友好的办法自己添加

我们把 CA 证书(1.4 节生成的 ca.csr)添加到 firefox 证书颁发机构中去，这样浏览器就信任我们自制的 CA 证书了。路径为：设置->隐私与安全->查看证书->导入

![img](../../../../images/be170bcbe51f474bb953ecc9e5f45bb0.png)

![img](../../../../images/208bf3d7e8337a556a1f62b76abcd04f.png)

由于我们证书配置的 common name 是[server.com](https://server.com/)，因此需要修改本地 hosts 文件，将下面文字添加到 hosts 文件中

```bash
127.0.0.1 server.com
```

再次访问https://server.com:10679/test不再阻止

![img](../../../../images/1df7a221c9342633d0261d51a3836038.png)

## 三.实现服务器和客户端双端验证

### 3.1 服务端代码

```go
package main
import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "log"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

var (
    caCert     string = "./certs/ca.crt"
    serverCert string = "./certs/server.crt"
    serverKey  string = "./certs/server.key"
)



func main() {
    router := gin.New()
    router.Use(gin.Logger())
    router.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "success",
        })
    })
    // 客户端CA证书
    certPool := x509.NewCertPool()
    ca, err := os.ReadFile(caCert)
    if err != nil {
        fmt.Printf("load ca err: %s", err)
        return
    }
    if ok := certPool.AppendCertsFromPEM(ca); !ok {
        fmt.Printf("certpool append ca fail.")
        return
    }
    // 可以直接用注释的代码代替最后两行
    //router.RunTLS("0.0.0.0:10679", "./cert/server.cer", "./cert/server.key")
    server := &http.Server{
        Addr:    "server.com:10679",
        Handler: router,
        TLSConfig: &tls.Config{
            ClientAuth: tls.RequireAndVerifyClientCert,
            //这里一定要注意，服务端设置ClientCAs，用于服务端验证客户端证书，客户端设置RootCAs，用户客户端验证服务端证书。设置错误或者设置反了都会造成认证不通过。
            //RootCAs:    certPool,
            ClientCAs: certPool,
        },
    }
    _ = server.ListenAndServeTLS(serverCert, serverKey)
}
```

### 3.2 客户端代码

```go
package main

import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func main() {
    pool := x509.NewCertPool()
    caCrt, err := os.ReadFile("./certs/ca.crt")
    if err != nil {
        log.Fatal("read ca.crt file error:", err.Error())
    }
    pool.AppendCertsFromPEM(caCrt)
    cliCrt, err := tls.LoadX509KeyPair("./certs/client.crt", "./certs/client.key")
    if err != nil {
        log.Fatalln("LoadX509KeyPair error:", err.Error())
    }
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{
            //这里一定要注意，服务端设置ClientCAs，用于服务端验证客户端证书，客户端设置RootCAs，用户客户端验证服务端证书。设置错误或者设置反了都会造成认证不通过。
            RootCAs: pool,
            //ClientCAs:    pool,
            Certificates: []tls.Certificate{cliCrt},
        },
    }
    client := &http.Client{Transport: tr}
    resp, err := client.Get("https://server.com:10679/test")
    if err != nil {
        fmt.Printf("get failed. | err: %s\n", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
```

## 四. 整理过程中遇到的 bug

### 4.1 tls: failed to verify certificate: x509: “[server.com](https://server.com/)” certificate is not standards compliant

remote error: tls: bad certificate

这两个报错可能就是客户端或服务端设置参数 ClientCAs、RootCAs 错误有关。

服务端设置 ClientCAs，里面保存客户端的 CA 证书 Pool，用于服务端验证客户端证书。

客户端设置 RootCAs，里面保存服务端的 CA 证书 Pool，用户客户端验证服务端证书。

设置错误或者设置反了都会造成认证不通过。

这就属于知道就很简单解决，但找不到错误就很崩溃，本人因为这个小 bug 竟然熬了一夜，说多了都是泪啊啊啊啊。。。

### 4.2 use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0

这个报错原因是生成证书没有开启 SAN 扩展，go 1.15 版本开始废弃 CommonName，因此推荐使用 SAN 证书。 这就是 1.6、1.9 两节配置文件最下面做的事情。重新把证书生成一下。当然，自己生成的证书可以随便改，如果线上证书出现这种情况，我看有的说设置下环境变量 GODEBUG 为 x509ignoreCN=0，不过我测试没有效果

这个问题的解决要感谢下面这个博主，我是看了这个帖子解决的问题。

最近在**golang 1.15+**版本上，用 gRPC 通过 TLS 实现数据传输加密时，遇到了一个问题：

原文:

**2021/02/02 16:17:04 Call Route err: rpc error: code = Unavailable desc = connection error: desc = "transport: authentication field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0"**

造成的原因是因为我用的证书，并没有开启 SAN 扩展（默认是没有开启 SAN 扩展）所生成的，所以在导致客户端和服务端无法建立连接， 所以我们要根据提示来解决这个问题：

按照提示所示，有 2 种方案可以解决:

### 方案一：设置 GODEBUG 为 `x509ignoreCN=0`

将 go 的 环境变量 GODEBUG 为 `x509ignoreCN=0，但是我设置了，还是未成功，你们可以自己试下，是否能成功。`

### 方案二：使用开启扩展 SAN 的证书

#### 1. 什么是 SAN

SAN(Subject Alternative Name) 是 SSL 标准 x509 中定义的一个扩展。使用了 SAN 字段的 SSL 证书，可以扩展此证书支持的域名，使得一个证书可以支持多个不同域名的解析。

#### 2.如何生成含有 SAN 的证书

2.1. 生成 CA 根证书

​ 2.1.1 准备 ca 配置文件，得到 ca.conf

```undefined
vim ca.conf
```

​ 内容如下:

```cobol
[ req ]
default_bits       = 4096
distinguished_name = req_distinguished_name

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = JiangSu
localityName                = Locality Name (eg, city)
localityName_default        = NanJing
organizationName            = Organization Name (eg, company)
organizationName_default    = Sheld
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = Ted CA Test
```

2.1.2 生成 ca 秘钥，得到 ca.key

```cobol
openssl genrsa -out ca.key 4096
```

2.1.3 生成 ca 证书签发请求，得到 ca.csr

```vbnet
openssl req \
  -new \
  -sha256 \
  -out ca.csr \
  -key ca.key \
  -config ca.conf
```

**配置文件中已经有默认值了，shell 交互时一路回车就行**

2.1.4 生成 ca 根证书，得到 ca.crt

```cobol
openssl x509 \
    -req \
    -days 3650 \
    -in ca.csr \
    -signkey ca.key \
    -out ca.crt
```

**2.2. 生成终端用户证书**

**2.**2.1 准备配置文件，得到 server.conf

```vbscript
vim server.conf
```

内容如下:

```cobol
[ req ]
default_bits       = 2048
distinguished_name = req_distinguished_name
req_extensions     = req_ext

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = JiangSu
localityName                = Locality Name (eg, city)
localityName_default        = NanJing
organizationName            = Organization Name (eg, company)
organizationName_default    = Sheld
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = www.eline.com


[ req_ext ]
subjectAltName = @alt_names

[alt_names]
DNS.1   = www.eline.com
IP      = 192.168.1.100
```

2.2.2 生成秘钥，得到 server.key

```cobol
openssl genrsa -out server.key 2048
```

2.2.3 生成证书签发请求，得到 server.csr

```vbscript
openssl req \
  -new \
  -sha256 \
  -out server.csr \
  -key server.key \
  -config server.conf
```

**配置文件中已经有默认值了，shell 交互时一路回车就行**

2.2.4 用 CA 证书生成终端用户证书，得到 server.crt

```cobol
openssl x509 \
  -req \
  -days 3650 \
  -CA ca.crt \
  -CAkey ca.key \
  -CAcreateserial \
  -in server.csr \
  -out server.pem\
  -extensions req_ext \
  -extfile server.conf
```

现在证书已经生成完毕， server.pem 和 server.key 正式我们需要的证书和密钥，我们用生成的证书验证一下:

启动服务端：

![img](../../../../images/20210202164701757.png)

启动成功了

客户端调用：

![img](../../../../images/20210202164734298.png)

看到调用成功了。

### 4.3 安全提示

参考 2.3 解决办法

### 4.4 证书 commonName

这个字段是比较重要的，不要随便配置。服务端、客户端证书认证都会验证 host name 是否与其一致，不一致会造成认证失败。这也是为什么要修改 hosts 并用[server.com](https://server.com/)来访问

# TLS 完全指南（二）：OpenSSL 操作指南

## 生成 RSA 秘钥对

以下 OpenSSL 的 genrsa[genrsa](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/openssl.md%23genrsa)命令生成一个 2048 bit 的公钥私 钥对，输出到文件 server.key 里[gist](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/openssl.md%23gist)：

```text
openssl genrsa -out server.key 2048
```

server.key 是 PEM 格式[pem](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/openssl.md%23pem)的：

```text
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-EDE3-CBC,DB98A9512DD7CBCF

yKTM+eoxBvptGrkEixhljqHSuE+ucTh3VqYQsgO6+8Wbh1docbFUKzLKHrferJBH
...
-----END RSA PRIVATE KEY-----
```

虽说文件头尾都标注着 RSA PRIVATE KEY，但实际上这个文件里既包括公钥也 包括私钥[genrsa](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/openssl.md%23genrsa)。

## 生成身份证申请

以下 OpenSSL 的 req 命令[req](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/openssl.md%23req)以上文中的 server.key 为输 入，生成一个身份证申请（CSR）文件 server.csr。

```text
openssl req -nodes -new -key server.key -subj "/CN=localhost" -out server.csr
```

这个 CSR 里的公钥是从 server.key 里提取出来的，域名是 localhost。 需要注意的是，如果将来我们启动一个 HTTPS 服务，使用这个 CSR 签署的身份 证，那么客户端必须可以通过域名 locahost 访问到这个 HTTPS 服务。

server.csr 文件也是 PEM 格式的，文件头尾标注为 CERTIFICATE REQUEST:

```text
-----BEGIN CERTIFICATE REQUEST-----
MIIC0TCCAbkCAQAwgYsxCzAJBgNVBAYTAlVTMQswCQYDVQQIEwJDQTERMA8GA1UE
...
-----END CERTIFICATE REQUEST-----
```

## 签署身份证

以下 OpenSSL 的 x509 命令[x509](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/openssl.md%23x509)用指定的私钥 server.key 签署 server.csr，输出身份证 server.crt：

```text
openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -out server.crt
```

server.crt 也是 PEM 格式的。文件头尾的标记为 CERTIFICATE:

```text
-----BEGIN CERTIFICATE-----
MIIDlDCCAnwCCQDQ1UvQyFD7jDANBgkqhkiG9w0BAQsFADCBizELMAkGA1UEBhMC
...
-----END CERTIFICATE-----
```

在这个例子里，用来签署 CSR 的私钥和 CSR 里的公钥是一对儿。也就是说这是一 个自签名（self-sign）的例子。

通常情况下，我们会用一个 CA 的私钥来签署一个 CSR。在这个为 Kubernetes apiserver 签署身份证的例子[sign](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/openssl.md%23sign)里，apiserver 的身份 证是用一个自签署的 CA 的私钥来签署的：

```text
$ openssl genrsa -out ca-key.pem 2048
$ openssl req -x509 -new -nodes -key ca-key.pem -days 10000 -out ca.pem -subj "/CN=kube-ca"

$ openssl genrsa -out apiserver-key.pem 2048
$ openssl req -new -key apiserver-key.pem -out apiserver.csr -subj "/CN=kube-apiserver" -config openssl.cnf
$ openssl x509 -req -in apiserver.csr -CA ca.pem -CAkey ca-key.pem -CAcreateserial -out apiserver.pem -days 365 -extensions v3_req -extfile openssl.cnf
```

这个例子里，ca-key.pem 是 Kubernetes 管理员自己创建的 CA 的私钥（其实是密 钥对）。第二个命令（openssl req -x509 ...）不经过生成 CSR 的过程，直接 输出 CA 的身份证。注意 CA 的域名是 kube-ca。

接下来三个命令分别创建 apiserver 的私钥，生成 apiserver 的 CSR，以及用 CA 的 private key（ca-key.pem）来签署 apiserver 的身份证。

注意，签署 apiserver.csr 得到 apiserver 的身份证（apiserver.pem） 的过程中，不仅需要 ca-key.pem，还需要 CA 的身份证 ca.pem，因为 apiserver.pem 里附上了 CA 的身份证。从而构成一个[信任链](./tls.md#信 任链)。

## HTTPS Server

现在我们有了 server.key 和 server.crt。我们可以写一个 HTTPS 服务程序， 它私藏 server.key，同时在与任何客户端程序首轮通信的时候通告自己的身 份证 server.crt。这里有几点需要注意：

1. 为了确保“私藏”一个文件，我们需要设置其文件系统访问权限为只有 owner 可 读[sign](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/openssl.md%23sign)：

   ```text
   chmod 400 server.key
   ```

2. 如果我们用 Go 语言来写这个 HTTPS server，只需要在调用 http.ListenAndServeTLS() 函数的时候，把 server.key 和 server.crt 的文件路径传递给它即可。整个程序[源码](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/server.go)如下：

   ```text
    package main

    import (
    	"io"
    	"log"
    	"net/http"
    )

    func main() {
    	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    		io.WriteString(w, "hello, world!\n")
    	})
    	if e := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil); e != nil {
    		log.Fatal("ListenAndServe: ", e)
    	}
    }
   ```

3. 因为我们在[前面这一步](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/openssl.md%23%E7%94%9F%E6%88%90%E8%BA%AB%E4%BB%BD%E8%AF%81%E7%94%B3%E8%AF%B7)生成 server.csr 的时 候，指定的域名是 localhost，所以必须确保 HTTPS 程序监听 localhost 虚拟网卡上的端口 443。上面程序里指定的监听地址是"443"，只有端口没 有标识网卡的 IP 地址或者域名，那么 ListenAndServerTLS 会让程序监听 本机所有网卡上的 443 端口。

   ```text
   sudo go run server.go &
   ```

4. 同样的原因，客户端必须通过 localhost 访问我们的 HTTPS 服务。在这 个例子里，localhost 域名意味着只有本机上执行的客户端才能访问。

## 访问 TLS 服务

### 用浏览器

我们可以通过浏览器访问我们的 HTTPS server。但是因为 server 的身份证是我 们自签署的，浏览器里没有[CA 的身份证](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/tls.md%23%E6%95%B0%E5%AD%97%E7%AD%BE%E5%90%8D%E5%92%8CCA)其中的公钥 可以验证 server.crt，所以浏览器会提示说它不信任我们的 HTTPS 服务。但是 如果用户表示信任，还是可以访问的。

要想消除浏览器的提示，最简单的办法就是把我们为 HTTPS 服务自签署的身份证 加入到浏览器里[chrome](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/openssl.md%23chrome)。

实际上，很多公司的运维团队都会生自命为本公司内部的 CA，成一个自签署的身 份证，加入到公司配发的电脑的操作系统或者浏览器里。而本公司内部的很多网 路服务（报销系统、人事管理系统、考评系统、各种计算资源上的 SSH 服务）都 用这个内部 CA 来签署。这样用公司的电脑，即可访问这些服务。

### 用 curl

类似的，我们可以通过加 -k 参数让 curl 信任我们的 HTTPS 服务器：

```text
$ /usr/local/Cellar/curl/7.49.1/bin/curl -k https://localhost
hello, world!
```

或者我们可以把我们自签署的身份证告诉 curl，让它用这个身份证验证我们的 HTTPS 服务：

```text
$ curl --cacert server.crt https://localhost
hello, world!
```

如果手边没有 server.crt，我们甚至可以用 openssl 工具找我们的 HTTPS server 要一份：

```text
$ /usr/local/Cellar/openssl/1.0.2h/bin/openssl s_client -showcerts -connect localhost:443 > cacert.pem
$ curl --cacert cacert.pem https://localhost
hello, world!
```

请注意，Mac OS X 自带的 openssl 版本太低，不支持 Go 语言 http package 里实现的高级加密算法，所以我们得用 Homebrew 安装新版本：

```text
brew update && brew install openssl
```

## 下一步

从上面操作可以看出来，TLS 的具体操作和 OpenSSL 息息相关。关于 OpenSSL 的 更多用法，可以参见[esse](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/openssl.md%23esse)。

当我们用 OpenSSL 创建了 CA 身份证并且用 CA 身份证给相关程序签署了各自的身 份证之后，我们就可以开工写作和调试这些程序了。在[下一篇](https://link.zhihu.com/?target=https%3A//github.com/k8sp/tls/blob/master/golang.md) 里，我们介绍如何用 Go 语言写 HTTPS 的 server 和 client。

# Golang（十）TLS 相关知识（一）基本概念原理

## 0. 前言

- 最近参与一个基于 BitTorrent 协议的 Docker 镜像分发加速插件的开发，主要参与补充 https 协议
- 学习了 TLS 相关知识，下面对之前的学习做一下简单总结
- 参考文献：[TLS 完全指南系列文章](https://cloud.tencent.com/developer/tools/blog-entry?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F26684050&source=article&objectId=2173771)

## 1. 基本原理

- TLS 依赖两种加密技术：
  - 对称加密（symmetric encryption）
  - 非对称加密（asymmetric encryption）

### 1.1 对称加密

- **加密方**和**解密方**共享同一个秘钥 K：

```javascript
加密：C = E(M, K)
解密：M = D(C, K)
```

- 攻击者且听到 K 后就可以作为一个中间人伪装成任意一个对象，实现**中间人攻击**

### 1.2 非对称加密

- 非对称加密利用成对的两个秘钥：K1 和 K2，加密者使用一个加密，解密者可以利用另一个解密：

```javascript
加密：C = E(M, K1)
解密：M = D(C, K2)
```

- **解密者**生成一对秘钥，私钥保存，公钥公开
- 但是**中间人**可以截获公钥，然后自己生成一对秘钥，把自己的公钥发送给**加密者**
- 用自己的私钥解密加密者的信息，然后用**解密者**的公钥加密发送给**解密者**
- 或者**中间人**收到**解密者**公钥加密的消息后，对消息破坏篡改，再发送给**解密者**
- 导致解密者无法正确解析密文

### 1.3 数字签名

- 光靠非对称加密很难确定信息发送方身份，因此发明了

  **数字签名**

  - 根据数字签名，接收方接收到信息之后，可以判断信息是否被破坏过，如果没有被破坏，就可以正确的解码。如果被破坏，就直接丢弃
  - 数字签名可以保证对信息的任何篡改都可以被发现，从而保证信息传输过程中的完整性
  - 数字签名是实现的关键技术就是哈希技术

- 加密者

  先对将要传输的信息进行哈希，得到一串独一无二的

  信息摘要

  - 哈希函数往往是不可逆的，无法根据哈希后的内容推断原文；同时，不同的原文，会造成不同的哈希结果，并且结果的差异是巨大甚至毫无规律的
  - 对原文进行再细微的修改，得到的哈希后的内容都会与未经修改的原文的哈希内容大相径庭，保证消息内容不被篡改

- 然后，

  加密者

  将信息摘要，用自己的私钥进行加密

  - 这样就保证只有加密者的公钥才能对信息摘要进行正确的解码，进而保证信息摘要一定是来自**加密者**
  - **加密后的信息摘要实际就是数字签名的内容**

- 最后，**加密者**再将数字签名附加到原文信息的后面，使用解密者公钥加密后发送给**解密者**

- 解密者

  接收到信息之后，首先使用自己的私钥解密报文，分别获得原文和数字签名

  - 利用**加密者**公钥对数字签名进行解密，得到信息摘要，如果成功解码，就说明数字签名是来自加密者

- 然后，**解密者**将信息原文进行哈希得到自己的信息摘要，与解码数字签名得到的信息摘要进行对比，如果相同，就说明原文信息完整，没有被篡改，反之，则确认信息被破坏了

- 目前为止，利用公钥和私钥以及数字签名，可以保证信息传输过程中的私密性和完整性

- 但还存在一个问题：就是公钥分发的问题，上述中间人劫持公钥的问题并没有解决

- 这个问题就需要数字证书和 CA 来解决了

### 1.4 数字证书和 CA

- 每个加密者或者接受者都有自己的私钥和公钥，如何判断对方的公钥是真实代表对方的是一个问题
- 实际我们会引入一个第三方机构，每个人都找这个真实可信的独立的第三方，请求真伪鉴别服务
- 第三方机构就是 CA（Certification Authorith，证书颁发机构）会给解密者创建一个数字证书
- 用户首先产生自己的密钥对，并将公钥及部分个人身份信息传送给认证中心
- 认证中心在核实身份后，将执行一些必要的步骤，以确信请求确实由用户发送而来
- 然后，认证中心将发给用户一个数字证书，该证书内包含用户的个人信息和他的公钥信息，同时还附有认证中心的签名信息”
- 公钥和身份信息（[域名](https://cloud.tencent.com/act/pro/domain-sales?from_column=20065&from=20065)或者 IP）合起来就是 CSR（certificate signing request，身份证申请）
- **实际过程可以看做 CA 利用自己的私钥对 CSR 加密，作为数字签名**
- **然后 CSR 连同 CA 的数字签名构成数字证书，也称为 CRT（CA signed certificate）**
- 在之后的发送中加密者将数字证书附在数字签名后
- 接收者收到后用 CA 的公钥解密获得，加密者的身份和公钥
- 攻击者不能通过 CA 的验证，无法生成可信任的证书，**解密者**接受后判断数字证书包含的身份信息不是**加密者**，因此会拒绝
- 当然，如果选择信任了错误的 CA，也会被攻击，通常浏览器中会内置靠谱 CA 的身份证（公钥）

### 1.4 信任链、根身份证和自签名

- CA 也分为不同级别，需要逐级验证
- 比如 CA1 不被大家信任，于是可以将身份信息和公钥发送给受信任的 CA2，获得自己的数字证书
- CA1 在给其他人签署数字证书时，会在后面附上自己的数字证书
- 这样接受者首先利用 CA2 的公钥验证 CA1，获得 CA1 的公钥后再验证发送者
- 这样逐级签署数字证书，形成了一条信任链
- 最终的根节点就是自签名证书，如 CA2 可以用自己的私钥把自己的公钥和域名加密，生成证书

### 1.5 应用场景：https 协议

- 首先，浏览器向[服务器](https://cloud.tencent.com/act/pro/promotion-cvm?from_column=20065&from=20065)发送加密请求

- 服务器将网页加密，连同自身的数字证书发送给浏览器

- 浏览器收到返回验证服务器身份，同时服务器也

  可以

  验证浏览器身份

  - 浏览器验证服务器是通过 TLS [身份验证](https://cloud.tencent.com/product/mfas?from_column=20065&from=20065)实现的，服务器验证浏览器是通过输入用户名密码实现的，通常服务器不会验证浏览器身份
  - 客户端（浏览器）的“证书管理器”，有“受信任的根证书颁发机构”列表。客户端会根据这张列表，查看解开数字证书的公钥是否在列表之内

![img](../../../../images/9926b106f36f57ddbaec7c1d963682da.png)

- - 如果数字证书记载的网址，与你正在浏览的网址不一致，就说明这张证书可能被冒用，浏览器会发出警告

![img](../../../../images/4b4ea15f8e314e6f055d17210769a6ee.png)

- - 如果这张数字证书不是由受信任的机构颁发的，浏览器会发出另一种警告

![img](../../../../images/039f1a0ae43163d29125565c007018ef.png)

- - 如果数字证书是可靠的，客户端就可以使用证书中的服务器公钥，对信息进行加密，然后与服务器交换加密信息

# Golang（十一）TLS 相关知识（二）OpenSSL 生成证书

## 0. 前言

- 接前一篇文章，上篇文章我们介绍了数字签名、数字证书等基本概念和原理
- 本篇我们尝试自己生成证书
- 参考文献：[TLS 完全指南（二）：OpenSSL 操作指南](https://cloud.tencent.com/developer/tools/blog-entry?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F26684071&source=article&objectId=2173773)

## 1. OpenSSL 简介

- OpenSSL 是一个开源项目，其组成主要包括三个组件：
  - openssl：多用途的[命令行工具](https://cloud.tencent.com/product/cli?from_column=20065&from=20065)
  - libcrypto：加密算法库
  - libssl：加密模块应用库，实现了 ssl 及 tls
- OpenSSL 主要用于秘钥证书管理、对称加密和非对称加密

### 1.1 指令

- 常用指令包括：**genrsa、req、x509**

#### 1.1.1 genrsa

- 主要用于生成私钥，选择算法、加密私钥使用的对称加密密码和秘钥长度
- 基本用法：**openssl genrsa [args] [numbits]**

```javascript
[args]:
    args1 对生成的私钥文件是否要使用加密算法进行对称加密:
        -des : CBC模式的DES加密
            -des3 : CBC模式的3DES加密
            -aes128 : CBC模式的AES128加密
        -aes192 : CBC模式的AES192加密
        -aes256 : CBC模式的AES256加密
    args2 对称加密密码
        -passout passwords
        其中passwords为对称加密(des、3des、aes)的密码(使用这个参数就省去了console交互提示输入密码的环节)
    args3 输出文件
        -out file : 输出证书私钥文件
[numbits]: 密钥长度，理解为私钥长度
```

- 生成一个 2048 位的 RSA 私钥，并用 des3 加密（密码为 123456），保存为 server.key 文件：**openssl genrsa -des3 -passout pass:123456 -out server.key 1024**

#### 1.1.2 req

- req 的基本功能主要有两个：生成证书请求和生成自签名证书（当然这并不是其全部功能，但是这两个最为常见）
- 基本用法：**openssl req [args] outfile**

```javascript
[args]
    args1 是输入输入文件格式：-inform arg
        -inform DER 使用输入文件格式为 DER
        -inform PEM 使用输入文件格式为 PEM
    args2 输出文件格式：-outform arg
        -outform DER 使用输出文件格式为 DER
        -outform PEM 使用输出文件格式为 PEM
    args3 是待处理文件
        -in inputfilepath
    args4 待输出文件
        -out outputfilepath
    args5 用于签名待生成的请求证书的私钥文件的解密密码
        -passin passwords
    args6 用于签名待生成的请求证书的私钥文件
        -key file
    args7 指定输入密钥的编码格式 -keyform arg
        -keyform DER
        -keyform NET
        -keyform PEM
    args8 生成新的证书请求
        -new
    args9 输出一个 X509 格式的证书,签名证书时使用
        -x509
    args10 使用 X509 签名证书的有效时间
        -days // -days 3650 有效期 10 年
    args11 生成一个 bits 长度的 RSA 私钥文件，用于签发，与-key互斥，生成证书请求或者自签名证书时自动生成密钥，然后生成的密钥名称由 -keyout 参数指定
        -newkey rsa:bits
    args12 设置 HASH 算法-[digest]，指定对创建请求时提供的申请者信息进行数字签名时指定的 hash 算法
        -md5
        -sha1 // 高版本浏览器开始不信任这种算法
        -md2
        -mdc2
        -md4
    args13 指定 openssl 配置文件,很多内容不容易通过参数配置，可以指定配置文件
        -config filepath
    args14 显示格式 txt（用于查看证书、私钥信息）
        -text
```

- 利用私钥生成证书请求 CSR：**openssl req -new -key server.key -out server.csr**
- 利用私钥生成自签名证书 CRT：**openssl req -new -x509 -days 3650 -key ca.key -out ca.crt**

#### 1.1.3 x509

- x509 可以实现显示证书的内容、转换其格式、给 CSR 签名等 X.509 证书的管理工作
- 基本用法：**openssl x509 [args]**

```javascript
[args]
    args1 是输入输入文件格式：-inform arg
        -inform DER 使用输入文件格式为 DER
        -inform PEM 使用输入文件格式为 PEM
    args2 输出文件格式:-outform arg
        -outform DER 使用输出文件格式为 DER
        -outform PEM 使用输出文件格式为 PEM
    args3 是待处理 X509 证书文件
        -in inputfilepath
    args4 待输出 X509 证书文件
        -out outputfilepath
    args5 表明输入文件是一个“请求签发证书文件（CSR）”，等待进行签发
        -req
    args6 签名证书的有效时间
        -days // -days 3650 有效期 10 年
    args7 指定用于签发请求证书的根 CA 证书
        -CA arg
    args8 根 CA 证书格式（默认是 PEM）
        -CAform arg
    args9 指定用于签发请求证书的 CA 私钥证书文件
        -CAkey arg
    args10 指定根 CA 私钥证书文件格式（默认为 PEM 格式）
        -CAkeyform arg
    args11 指定序列号文件（serial number file）
        -CAserial arg
    args12 如果序列号文件（serial number file）没有指定，则自动创建它
        -CAcreateserial
    args12 设置 HASH 算法-[digest]，指定对创建请求时提供的申请者信息进行数字签名时指定的 hash 算法
        -md5
        -sha1 // 高版本浏览器开始不信任这种算法
        -md2
        -mdc2
        -md4
```

- 使用根 CA 证书 ca.crt 和私钥 ca.key 对“请求签发证书” server.csr 进行签发，生成 x509 格式证书：**openssl x509 -req -days 3650 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out serverx509.crt**

## **2. 具体使用**

### 2.1 生成 RSA 秘钥对

- 使用 genrsa 生成 RSA 秘钥对：**openssl genrsa -out server.key 2048**

### 2.2 生成身份证申请

- 使用 req 命令，以之前的 server.key 为输入，生成一个身份证申请（CSR）文件：**openssl req -nodes -new -key server.key -subj "/CN=localhost" -out server.csr**
- CSR 的公钥从 server.key 中提取，[域名](https://cloud.tencent.com/act/pro/domain-sales?from_column=20065&from=20065)是 localhost
- 如果启动 https 服务，使用这个 CSR 签署的 CRT，客户端必须访问 localhost 才能访问到这个 HTTPS 服务
- 关于配置多个域名和 IP 的 CSR，后面会介绍

### 2.3 生成数字证书

- 使用 x509 使用指定私钥 server,key 签署 server.csr，输出数字证书（ CRT）：**openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -out server.crt**
- 此处使用自身的私钥签署 CSR

### 2.4 HTTPS 验证

- 生成证书后，我们可以编写一个 Golang 的 https 服务验证刚刚生成的证书
- 服务端代码如下：

```javascript
package main

import (
    "io"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        io.WriteString(w, "hello, world!\n")
    })
    if e := http.ListenAndServeTLS("0.0.0.0:5200", "/home/ao/Documents/certs/review/server.crt",
        "/home/ao/Documents/certs/review/server.key", nil); e != nil {
        log.Fatal("ListenAndServe: ", e)
    }
    //if e := http.ListenAndServe("0.0.0.0:5200", nil); e != nil {
    //    log.Fatal("ListenAndServe: ", e)
    //}
}
```

- 客户端代码如下：

```javascript
package main

import (
    "crypto/tls"
    "crypto/x509"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func loadCA(caFile string) *x509.CertPool {
    pool := x509.NewCertPool()

    ca, err := ioutil.ReadFile(caFile)
    if err != nil {
        log.Fatal("ReadFile: ", err)
    }
    pool.AppendCertsFromPEM(ca)
    return pool
}

func main() {
    //c := &http.Client{
    //    Transport: &http.Transport{
    //        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    //    }}
    c := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{RootCAs: loadCA("/home/ao/Documents/certs/review/server.crt")},
        }}

    resp, err := c.Get("https://localhost:5200")
    if err != nil {
        log.Fatal("http.Client.Get: ", err)
    }

    defer resp.Body.Close()
    io.Copy(os.Stdout, resp.Body)
}
```

- 启动服务端：**go run server.go**
- 启动客户端：**go run client.go**

```javascript
$ go run client.go
hello, world!
```

- 改变客户端请求域名为 127.0.0.1 时，客户端结果如下：

```javascript
$ go run client.go
2019/09/30 15:11:41 http.Client.Get: Get https://127.0.0.1:5200: x509: cannot validate certificate for 127.0.0.1 because it doesn't contain any IP SANs
exit status 1
```

- 服务端输出如下：

```javascript
2019/09/30 15:11:41 http: TLS handshake error from 127.0.0.1:33596: remote error: tls: bad certificate
```

- 若客户端保持 127.0.0.1 不变，改变 http.Transport 的 TLSClientConfig 为 insecure 配置时，可以正常返回：

```javascript
$ go run client.go
2019/09/30 15:11:41 http.Client.Get: Get https://127.0.0.1:5200: x509: cannot validate certificate for 127.0.0.1 because it doesn't contain any IP SANs
exit status 1
```

- 此时就像浏览器保持信任网站证书，继续选择浏览的动作是一样的

### 2.5 配置域名为 IP

- 上述我们使用 localhost，作为 CSR 里的域名，导致请求时必须使用域名
- 若我们使用本地 IP 作为域名呢：**openssl req -nodes -new -key server.key -subj "/CN=127.0.0.1" -out server.csr**
- **此时客户端请求 https://127.0.0.1:5200 返回：**

```javascript
$ go run client.go
2019/09/30 15:19:24 http.Client.Get: Get https://127.0.0.1:5200: x509: cannot validate certificate for 127.0.0.1 because it doesn't contain any IP SANs
exit status 1
```

- 启用 insecure 时正常返回
- 此处说明，我们不能简单通过 -subj "/CN=[IP]" 实现 CSR 中包含域名
- 重新生成 CSR 和 CRT

```javascript
$ openssl genrsa -out server.key 2048
$ echo subjectAltName = IP:127.0.0.1 > extfile.cnf
$ openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -extfile extfile.cnf -out server.crt
```

- 再次测试发现，请求 127.0.0.1 时可以了

### 2.6 不使用自签名证书

- 上述我们使用自签名证书，下面我们尝试模拟一个 CA 签署证书：
- 首先生成 CA 的秘钥和自签名证书，中间不生成 CSR：

```javascript
$ openssl genrsa -out ca.key 2048
$ openssl req -x509 -new -nodes -key ca.key -days 10000 -out ca.crt -subj "/CN=localhost.ca.com"
```

- 生成私钥、证书，并使用 CA 签名：

```javascript
$ openssl genrsa -out server.key 2048
$ openssl req -new -key server.key -subj "/CN=127.0.0.1" -out server.csr
$ openssl x509 -req -sha256 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -extfile extfile.cnf
```

- 经测试，正确

### 2.7 CSR 配置多个域名和 IP

- 上述提到 -subj 配置域名只能指定单个域名或者 IP
- 下面介绍一下多域名 CSR 配置
- 新增一个配置文件：

```javascript
[req]
distinguished_name = req_distinguished_name
req_extensions = v3_req

[req_distinguished_name]
countryName = CN
countryName_default = CN
stateOrProvinceName = Beijing
stateOrProvinceName_default = Beijing
localityName = Beijing
localityName_default = Beijing
organizationName = WangAo
organizationName_default = WangAo
organizationalUnitName = Dev
organizationalUnitName_default = Dev
commonName = test.openssl.com
commonName_default = test.openssl.com
commonName_max = 64

[v3_req]
basicConstraints = CA:TRUE
subjectAltName = @alt_names

[alt_names]
DNS.1 = test.openssl.com
IP.1 = 127.0.0.1
IP.2 = 10.0.2.15
```

- 我们为服务制定了一个域名两个 IP，地址
- 下面利用之前的 CA，重新签署数字证书：

```javascript
$  openssl genrsa -out server.key 2048
$  openssl req -nodes -new -key server.key -out server.csr -subj "/CN=test.openssl.com"
$ openssl x509 -req -sha256 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -extensions v3_req -extfile openssl.cnf
```

- CA 签署数字证书时制定了 -extensions 和 -extfile
- 分别向 127.0.0.1、10.0.2.15 和 test.openssl.com 请求，均可以成功

# Golang（十二）TLS 相关知识（三）理解并模拟简单代理

## 0. 前言

- 前面的介绍我们理解了数字签名等知识，同时学习了 OpenSSL 生成私钥和证书并验证
- 之前提过我们基于 BitTorrent 协议开发了一个 [docker](https://cloud.tencent.com/product/tke?from_column=20065&from=20065) 镜像分发加速插件
- 中间涉及到了配置 docker 的代理
- 下面我们简单介绍下 Golang 的 http.transport 配置了网络代理后的网络行为并编写一个简单的代理转发，加深理解代理转发行为

## 1. http.Transport 配置代理

- http 代理配置代码如下：

```javascript
func TLSTransport(caFile string) (*http.Transport, error) {
    tr := &http.Transport{TLSClientConfig: &tls.Config{}, Proxy: http.ProxyFromEnvironment}
    if len(caFile) == 0 {
        tr.TLSClientConfig.InsecureSkipVerify = true
        return tr, nil
    }

    ca, err := ioutil.ReadFile(caFile)
    if err != nil {
        return nil, fmt.Errorf("read CA file failed: %v", err)
    }
    pool := x509.NewCertPool()
    pool.AppendCertsFromPEM(ca)

    tr.TLSClientConfig.RootCAs = pool

    return tr, nil
}
```

- 上述代码制定了 Proxy 为 **http.ProxyFromEnvironment**
- 我们跟踪一下 http.ProxyFromEnvironment 的代码

```javascript
func ProxyFromEnvironment(req *Request) (*url.URL, error) {
    return envProxyFunc()(req.URL)
}
/////////////////////////////////////////////////////////////////////////
func envProxyFunc() func(*url.URL) (*url.URL, error) {
    envProxyOnce.Do(func() {
        envProxyFuncValue = httpproxy.FromEnvironment().ProxyFunc()
    })
    return envProxyFuncValue
}
/////////////////////////////////////////////////////////////////////////
func FromEnvironment() *Config {
    return &Config{
        HTTPProxy:  getEnvAny("HTTP_PROXY", "http_proxy"),
        HTTPSProxy: getEnvAny("HTTPS_PROXY", "https_proxy"),
        NoProxy:    getEnvAny("NO_PROXY", "no_proxy"),
        CGI:        os.Getenv("REQUEST_METHOD") != "",
    }
}
/////////////////////////////////////////////////////////////////////////
func (cfg *Config) ProxyFunc() func(reqURL *url.URL) (*url.URL, error) {
    // Preprocess the Config settings for more efficient evaluation.
    cfg1 := &config{
        Config: *cfg,
    }
    cfg1.init()
    return cfg1.proxyForURL
}
/////////////////////////////////////////////////////////////////////////
func (cfg *config) proxyForURL(reqURL *url.URL) (*url.URL, error) {
    var proxy *url.URL
    if reqURL.Scheme == "https" {
        proxy = cfg.httpsProxy
    }
    fmt.Printf("WangAo test: proxy: %+v", proxy)
    if proxy == nil {
        proxy = cfg.httpProxy
        if proxy != nil && cfg.CGI {
            return nil, errors.New("refusing to use HTTP_PROXY value in CGI environment; see golang.org/s/cgihttpproxy")
        }
    }
    if proxy == nil {
        return nil, nil
    }
    if !cfg.useProxy(canonicalAddr(reqURL)) {
        return nil, nil
    }

    return proxy, nil
}
```

- proxy 指定返回给定请求的代理的函数
- 如果函数返回一个非 nil 错误，请求将因提供的错误而中止
- 代理类型由 URL scheme 决定：支持 http、https 和 socks5
- 如果 scheme 为空，则假定为 http
- 如果 proxy 为 nil 或返回 nil 的 \*url.URL 类型，则不使用 proxy
- envProxyFunc 返回一个函数，函数读取环境变量确定代理地址
- FromEnvironment 可以看出代码主要读取 **HTTP_PROXY、HTTPS_PROXY、NO_PROXY 和 REQUEST_METHOD**
- ProxyFunc 中调用 config.init 方法解析环境变量，并返回实际解析 URL 并返回代理地址的函数
- 在 proxyForURL 中我们发现，对于 https 请求首选是采用 https 代理地址，若 https 代理地址为空或者请求为其他请求则采用 http 地址
- 若配置了 http 代理地址同时配置了 REQUEST_METHOD，返回空代理地址和错误信息
- 如果 http 代理也没有配置则返回空代理地址
- 解析请求信息若为 localhost 或者为回环地址不使用代理地址，否则返回配置的代理地址

## 2. 测试网络行为

- 上述我们简单读取了 http.ProxyFromEnvironment 读取环境变量确定代理地址的行为
- 下面我们简单介绍下测试代码
- 首先是 Server 端：

```javascript
package main

import (
    "bufio"
    "context"
    "fmt"
    "git.tencent.com/tke/p2p/pkg/util"
    "github.com/elazarl/goproxy"
    "github.com/gorilla/mux"
    "io"
    "k8s.io/klog"
    "log"
    "net"
    "net/http"
)

func main() {
    go func() {
        log.Println("Starting httpServer")
        router := mux.NewRouter().SkipClean(true)
        proxy := goproxy.NewProxyHttpServer()
        proxy.Verbose = true
        proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
            req.URL.Host = req.Host
            req.URL.Scheme = "http"
            proxy.ServeHTTP(w, req)
        })
        proxy.OnRequest(goproxy.ReqHostIs("test.openssl.com:1213")).HijackConnect(func(req *http.Request, client net.Conn, _ *goproxy.ProxyCtx) {
            var err error

            log.Printf("getHijhack: %+v", req.URL)
            defer func() {
                if err != nil {
                    klog.Errorf("Transfer HTTP CONNECT request failed: %+v, %v", req, err)
                    if _, writeErr := client.Write([]byte("HTTP/1.1 500 Cannot reach destination\r\n\r\n")); err != nil {
                        klog.Errorf("Write CONNECT failing header failed: %v", writeErr)
                    }
                }
                if closeErr := client.Close(); closeErr != nil {
                    klog.Errorf("Close client connection failed: %v", closeErr)
                }
            }()

            log.Println("before connectDial")
            remote, err := connectDial(proxy, "tcp", "127.0.0.1:1213")
            if remote != nil {
                log.Printf("==============> remote: %+v>%+v\n", remote.LocalAddr(), remote.RemoteAddr())
            }
            if err != nil {
                return
            }

            bufferedRemote := bufio.NewReadWriter(bufio.NewReader(remote), bufio.NewWriter(remote))
            bufferedClient := bufio.NewReadWriter(bufio.NewReader(client), bufio.NewWriter(client))

            errCh := make(chan error, 1)
            go func() {
                defer close(errCh)
                if _, reverseErr := io.Copy(bufferedRemote, bufferedClient); reverseErr != nil {
                    klog.Errorf("Transfer remote to client failed: %v", reverseErr)
                    errCh <- reverseErr
                }
            }()

            if _, transferErr := io.Copy(bufferedClient, bufferedRemote); transferErr != nil {
                klog.Errorf("Transfer client to remote failed: %v", transferErr)
                err = transferErr
            }

            if reverseErr := <-errCh; reverseErr != nil {
                err = reverseErr
            }
        })
        router.HandleFunc("/http", func(w http.ResponseWriter, r *http.Request) {
            log.Printf("1--------------------->http: /http >>>>>> req.URL: %+v", r.URL)
            cnt, err := w.Write([]byte(fmt.Sprintf("http: /http return response of req: %+v", r)))
            log.Printf("/http write: cnt: %v, err: %v", cnt, err)
        })
        router.HandleFunc("/https", func(w http.ResponseWriter, r *http.Request) {
            log.Printf("2--------------------->http: /https >>>>>>req.URL: %+v", r.URL)
            cnt, err := w.Write([]byte(fmt.Sprintf("http: /https return response of req: %+v", r)))
            log.Printf("/http write: cnt: %v, err: %v", cnt, err)
            //proxy.ServeHTTP(w, r)
        })
        router.NotFoundHandler = proxy
        if err := http.ListenAndServe(":1212", router); err != nil {
            log.Printf("httpServer err: %+v", err)
        }
    }()
    go func() {
        log.Println("Starting httpsServer")
        router := mux.NewRouter().SkipClean(true)
        proxy := goproxy.NewProxyHttpServer()
        proxy.Verbose = true
        proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
            req.URL.Host = req.Host
            req.URL.Scheme = "https"
            proxy.ServeHTTP(w, req)
        })
        if tr, err := util.TLSTransport("/home/ao/Documents/certs/review/server.crt"); err == nil {
            proxy.Tr = tr
        }
        router.HandleFunc("/https", func(w http.ResponseWriter, r *http.Request) {
            log.Printf("3--------------------->https: req: %+v", r)
            cnt, err := w.Write([]byte(fmt.Sprintf("https: /https return response of req: %+v", r)))
            log.Printf("/http write: cnt: %v, err: %v", cnt, err)
        })
        if err := http.ListenAndServeTLS(":1213", "/home/ao/Documents/certs/review/server.crt", "/home/ao/Documents/certs/review/server.key", router); err != nil {
            log.Printf("httsServer err: %+v", err)
        }
    }()
    select {}
}

func dial(proxy *goproxy.ProxyHttpServer, network, addr string) (c net.Conn, err error) {
    if proxy.Tr.DialContext != nil {
        return proxy.Tr.DialContext(context.Background(), network, addr)
    }
    return net.Dial(network, addr)
}

func connectDial(proxy *goproxy.ProxyHttpServer, network, addr string) (c net.Conn, err error) {
    if proxy.ConnectDial == nil {
        return dial(proxy, network, addr)
    }
    return proxy.ConnectDial(network, addr)
}
```

- 服务端启动了两个 goroutine，分别监听 http 和 https 请求
- http 监听地址为配置的代理地址
- https 为请求实际请求的地址，同时我们设置了拦截 CONNECT 方法的目标[域名](https://cloud.tencent.com/act/pro/domain-sales?from_column=20065&from=20065)
- 在拦截 CONNECT 方法之后的回调函数我们看到此时会和 https 监听地址交换数据转发给 https 地址
- 然后我们看一下 Client 端：

```javascript
package main

import (
    "fmt"
    "git.tencent.com/tke/p2p/pkg/util"
    "io/ioutil"
    "net/http"
)

func main() {
    tr, _ := util.TLSTransport("/home/ao/Documents/certs/review/server.crt")
    client := &http.Client{Transport: tr}

    req, _ := http.NewRequest("GET", "https://test.openssl.com:1213/https", nil)
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("err: %+v", err)
    } else {
        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Printf("resp: %+v=>%+v", resp.StatusCode, string(body))
    }
}
```

- Client 端很简单，我们只是制定了证书发送一个 https 请求
- 分别启动 Server 端和 Client 端我们看一下结果：

```javascript
$ go run server.go
2019/10/10 14:51:08 Starting httpsServer
2019/10/10 14:51:08 Starting httpServer
2019/10/10 14:51:33 [001] INFO: Running 1 CONNECT handlers
2019/10/10 14:51:33 [001] INFO: on 0th handler: &{3 0x69b280 <nil>} test.openssl.com:1213
2019/10/10 14:51:33 [001] INFO: Hijacking CONNECT to test.openssl.com:1213
2019/10/10 14:51:33 getHijhack: //test.openssl.com:1213
2019/10/10 14:51:33 before connectDial
2019/10/10 14:51:33 ==============> remote: 127.0.0.1:55062>127.0.0.1:1213
2019/10/10 14:51:33 3--------------------->https: req: &{Method:GET URL:/https Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Accept-Encoding:[gzip] User-Agent:[Go-http-client/1.1]] Body:{} GetBody:<nil> ContentLength:0 TransferEncoding:[] Close:false Host:test.openssl.com:1213 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:127.0.0.1:55062 RequestURI:/https TLS:0xc0000ae6e0 Cancel:<nil> Response:<nil> ctx:0xc000142510}
2019/10/10 14:51:33 /http write: cnt: 434, err: <nil>
2019/10/10 14:52:54 [002] INFO: Running 1 CONNECT handlers
2019/10/10 14:52:54 [002] INFO: on 0th handler: &{3 0x69b280 <nil>} test.openssl.com:1213
2019/10/10 14:52:54 [002] INFO: Hijacking CONNECT to test.openssl.com:1213
2019/10/10 14:52:54 getHijhack: //test.openssl.com:1213
2019/10/10 14:52:54 before connectDial
2019/10/10 14:52:54 ==============> remote: 127.0.0.1:55066>127.0.0.1:1213
2019/10/10 14:52:54 3--------------------->https: req: &{Method:GET URL:/https Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Accept-Encoding:[gzip] User-Agent:[Go-http-client/1.1]] Body:{} GetBody:<nil> ContentLength:0 TransferEncoding:[] Close:false Host:test.openssl.com:1213 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:127.0.0.1:55066 RequestURI:/https TLS:0xc000160160 Cancel:<nil> Response:<nil> ctx:0xc000154510}
2019/10/10 14:52:54 /http write: cnt: 434, err: <nil>

$ HTTP_PROXY=http://127.0.0.1:1212 go run connect.go
resp: 200=>return response of req: &{Method:GET URL:http://test.openssl.com:1213/https Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Accept-Encoding:[gzip] User-Agent:[Go-http-client/1.1]] Body:{} GetBody:<nil> ContentLength:0 TransferEncoding:[] Close:false Host:test.openssl.com:1213 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:127.0.0.1:34024 RequestURI:http://test.openssl.com:1213/https TLS:<nil> Cancel:<nil> Response:<nil> ctx:0xc00011cba0}%

$ HTTPS_PROXY=http://127.0.0.1:1212 go run connect.go
resp: 200=>return response of req: &{Method:GET URL:http://test.openssl.com:1213/https Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Accept-Encoding:[gzip] User-Agent:[Go-http-client/1.1]] Body:{} GetBody:<nil> ContentLength:0 TransferEncoding:[] Close:false Host:test.openssl.com:1213 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:127.0.0.1:34024 RequestURI:http://test.openssl.com:1213/https TLS:<nil> Cancel:<nil> Response:<nil> ctx:0xc00011cba0}%
```

- 设定 HTTPS_PROXY 和 HTTP_PROXY 启动得到的结果都是一样的
- 当我们把 Client 端请求的结果设为 http 时，我们再看一下 Server 和 Client 的输出：

![img](../../../../images/fc7460811abbe392647b481d79e114b3.png)

```javascript
$ go run server.go
2019/10/10 14:55:21 Starting httpsServer
2019/10/10 14:55:21 Starting httpServer
2019/10/10 14:55:23 2--------------------->http: /https >>>>>>req.URL: http://test.openssl.com:1213/https
2019/10/10 14:55:23 /http write: cnt: 482, err: <nil>

$ HTTP_PROXY=http://127.0.0.1:1212 go run client.go
resp: 200=>http: /https return response of req: &{Method:GET URL:http://test.openssl.com:1213/https Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Accept-Encoding:[gzip] User-Agent:[Go-http-client/1.1]] Body:{} GetBody:<nil> ContentLength:0 TransferEncoding:[] Close:false Host:test.openssl.com:1213 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:127.0.0.1:34630 RequestURI:http://test.openssl.com:1213/https TLS:<nil> Cancel:<nil> Response:<nil> ctx:0xc00007bdd0}%

$ HTTPS_PROXY=http://127.0.0.1:1212 go run client.go
resp: 400=>Client sent an HTTP request to an HTTPS server.
```

- 同样分别设定 HTTP_PROXY 和 HTTPS_PROXY 发送 http 请求得到结果是不同的
- 配置了 HTTP_PROXY 的请求被代理收到但是没有发出 CONNECT 方法，而是以 http 方式直接请求的
- 配置了 HTTPS_PROXY 的请求没有被代理收到，但是由于 https 服务端同样的 IP 地址，被 https 服务端直接收到。但由于是 http 请求直接返回 400 了
