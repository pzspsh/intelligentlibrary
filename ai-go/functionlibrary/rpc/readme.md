一、什么是RPC

远程过程调用（Remote Procedure Call，缩写为 RPC）是一个计算机通信协议。该协议允许运行于一台计算机的程序调用另一台计算机的子程序，而程序员无需额外地为这个交互作用编程。如果涉及的软件采用面向对象编程，那么远程过程调用亦可称作远程调用或远程方法调用。

用通俗易懂的语言描述就是：RPC允许跨机器、跨语言调用计算机程序方法。打个比方，我用go语言写了个获取用户信息的方法getUserInfo，并把go程序部署在阿里云服务器上面，现在我有一个部署在腾讯云上面的php项目，需要调用golang的getUserInfo方法获取用户信息，php跨机器调用go方法的过程就是RPC调用。

二、RPC工作流程：

![img](https://github.com/pzspsh/intelligentlibrary/blob/main/images/RPC工作流程.png)

三、go支持三个级别的RPC( HTTP，TCP，JSONRPC)

四、实现RPC实例：

3.1 GO RPC的函数只有符合以下条件才能被远程访问:

    函数必须是首字母是大写
    必须有两个首字母大写的参数
    第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
    函数还要有一个返回值error
    func (t *T)MethodName(arg1 T1, returnArg *T2)error

3.2 golang中如何实现RPC

在golang中实现RPC非常简单，有封装好的官方库和一些第三方库提供支持。Go RPC可以利用tcp或http来传递数据，可以对要传递的数据使用多种类型的编解码方式。golang官方的net/rpc库使用encoding/gob进行编解码，支持tcp或http数据传输方式，由于其他语言不支持gob编解码方式，所以使用net/rpc库实现的RPC方法没办法进行跨语言调用。

golang官方还提供了net/rpc/jsonrpc库实现RPC方法，JSON RPC采用JSON进行数据编解码，因而支持跨语言调用。但目前的jsonrpc库是基于tcp协议实现的，暂时不支持使用http进行数据传输。

除了golang官方提供的rpc库，还有许多第三方库为在golang中实现RPC提供支持，大部分第三方rpc库的实现都是使用protobuf进行数据编解码，根据protobuf声明文件自动生成rpc方法定义与服务注册代码，在golang中可以很方便的进行rpc服务调用。