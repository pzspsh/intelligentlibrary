一、RPC工作流程：
![img](https://github.com/pzspsh/intelligentlibrary/blob/main/images/RPC工作流程.png)
二、go支持三个级别的RPC( HTTP，TCP，JSONRPC)

三、实现http的RPC实例：

3.1 GO RPC的函数只有符合以下条件才能被远程访问

函数必须是首字母是大写
必须有两个首字母大写的参数
第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
函数还要有一个返回值error
func (t *T)MethodName(arg1 T1, returnArg *T2)error