# 域名https证书服务部署

#### 1、下载OpenSSL源码，[官网](http://www.openssl.org/)
#### 2、解压OpenSSL 找到\openssl-1.0.1j\apps目录，拷贝demoCA目录和openssl.cnf文件到Openssl的安装目录下的bin目录下（如：D:\OpenSSL-Win64\bin\）
#### 3、在D:\OpenSSL-Win64\bin目录下，创建ca、jks、server、client四个目录。 
#### 4、生成 server.key和server.crt （注：因为懒得加环境变量，我是在D:\OpenSSL-Win64\bin目录下执行的openssl）
##### 生成 server.key
```bash
openssl genrsa -out server.key 2048
```

##### 生成 server.crt
```bash
openssl req -new -x509 -key server.key -out server.crt -days 3650
```
按照给出的提示依次输入：CN  BJ等，如下：
```bash
Country Name (2 letter code) [AU]:CN
State or Province Name (full name) [Some-State]:BJ
Locality Name (eg, city) []:BJ
Organization Name (eg, company) [Internet Widgits PtyLtd]:ple
Organizational Unit Name (eg, section) []:live
Common Name (eg, YOUR name) []:root
Email Address []:
```

#### 5、将生成的server.crt和server.key剪切复制到第一步的go程序所在目录，运行go程序

#### 6、打开浏览器访问 localhost:8081
```bash
什么？竟然提示该网页无法正常运作！
哦，低级错误，应该访问：https://localhost:8081
又提示“您的连接不是私密连接”，忽略，终于显示Hi, This is an example of https service in golang!
也算是基本打通了https。
然而，浏览器地址栏上大大的红色“不安全”字样看着不舒服。怎么办？
办法1：花钱申请证书或免费证书
办法2：自签证书，在客户端，将证书导入到“受信任的根证书颁发机构”存储区中
另外，还可以双向验证。
```

#### 7、符合Chrome58+的证书制作（费了不少劲）
```bash
首先设置windows环境变量，我的是 D:\OpenSSL-Win64\bin
主要参考了https://blog.csdn.net/wdydxf/article/details/54576063，可能是openssl版本不同，稍做修改。如不明白，请看原文
下面的操作, 我将建立一个 MyRootCA 的根证书颁发机构, 然后为一个域名是 myserver.com 签发证书
```

##### 将 C:\OpenSSL-Win64\bin\openssl.cfg 复制到 F:\SSLTest\, 并在F:\SSLTest\执行以下命令
```bash
mkdir demoCA\private demoCA\newcerts
type nul > demoCA\index.txt
echo 01 > demoCA\serial
```

##### 生成CA自签名证书
```bash
openssl req -new -x509 -newkey rsa:2048 -days 3650 -keyout demoCA\private\MyRootCA.key -out demoCA\MyRootCA.crt -passout pass:123456 -config openssl.cnf
```
依次输入 CN bj0 bj1 bj2 MyRootCA 空（回车）

##### 生成用户的 RSA 密钥对
```bash
openssl genrsa -des3 -out myserver.com._has_passwd.key -passout pass:123456
```

##### 删除私钥中的密码(否则golang程序无法启动)
```bash
openssl rsa -in myserver.com._has_passwd.key -out myserver.com.key
```

##### 生成用户证书请求
```bash
openssl req -new -days 1825 -key myserver.com.key -out myserver.com.csr -config openssl.cnf
```
依次输入 123456 CN bj0 bj1 bj2 bjbj myserver.com 空（回车） 空（回车） 空（回车）（注意与6.2的相关输入内容保持一样）

##### 使用 CA 签发用户证书
```bash
openssl ca -in myserver.com.csr -out myserver.com.crt -cert demoCA\MyRootCA.crt -keyfile demoCA\private\MyRootCA.key -extensions v3_req -config openssl.cnf
```
依次输入 123456 y y

##### 验证(在C:\Windows\System32\drivers\etc\hosts中新增一行)
```bash
127.0.0.1 myserver.com
```
然后将 MyRootCA.crt 安装到 受信任的根证书颁发机构
将myserver.com.crt和myserver.com.key放到golang源代码目录下
访问：https://myserver.com:8081/
终于显示连接是安全的了。
 补充：有的源码不是用.crt 和.key格式的证书和私钥，而都是用.pem后缀的。有的说法是.crt后缀可直接改为.pem，而.key可以转换格式，
如：openssl rsa -in id_rsa -text > private.pem


重点来了:
    自签名不推荐，自己是自己拥有域名的话，直接去腾讯云签一个时长一年的专业CA机构签发的dv类型证书就完事了[腾讯云](https://cloud.tencent.com/product/ssl)

还有这里https://wzfou.com/letsencrypt/和https://www.cnblogs.com/tv151579/p/8268356.html介绍的：letsencrypt免费SSL证书
最后，听说go https 在速度上不如Nginx + go http 方式

##### 参考：
```bash
https://segmentfault.com/a/1190000013287122
https://segmentfault.com/a/1190000009666888
https://segmentfault.com/a/1190000016249967
https://blog.csdn.net/zhaotengfei36520/article/details/41962077
https://studygolang.com/articles/9267
https://studygolang.com/articles/9959
https://blog.51cto.com/colinzhouyj/1566438
https://blog.csdn.net/luyangbin01/article/details/50972693
https://blog.csdn.net/sunhuansheng/article/details/82902185
https://blog.csdn.net/mixika99/article/details/79009521
https://blog.csdn.net/huplion/article/details/52892869
https://douhan.li/?p=79
https://www.cnblogs.com/274914765qq/p/4672108.html
https://blog.csdn.net/c_base_jin/article/details/81229643
http://www.mamicode.com/info-detail-1938899.html
https://blog.csdn.net/wdydxf/article/details/54576063
https://blog.csdn.net/qq_37049781/article/details/84837342
```