# [python实现调用摄像头或打开视频文件](https://www.cnblogs.com/april0315/p/13702425.html)



目录：

（一）调用摄像头或打开视频文件代码实现

（二）说明和补充

（一）调用摄像头或打开视频文件代码实现

```
 1 # -*- coding=GBK -*-
 2 import cv2 as cv
 3  
 4  
 5 #打开摄像头获取图片
 6 def video_demo():
 7     capture = cv.VideoCapture(0)#打开摄像头，0代表的是设备id，如果有多个摄像头，可以设置其他数值
 8     while True:
 9         ret, frame = capture.read() #读取摄像头,它能返回两个参数，第一个参数是bool型的ret，其值为True或False，代表有没有读到图片；第二个参数是frame，是当前截取一帧的图片
10         frame = cv.flip(frame, 1)#翻转 等于0:逆时针180度旋转， 大于0:正常 ，小于0上下颠倒
11         cv.imshow("video", frame)
12         if cv.waitKey(10) & 0xFF == ord('q'): #键盘输入q退出窗口，不按q点击关闭会一直关不掉 也可以设置成其他键。
13             break
14  
15  
16 video_demo()
17 cv.destroyAllWindows()
```

（二）代码实现说明和补充

### 1. c = cv.waitKey(40)  if c == 27 和cv.waitKey(10) & 0xFF == ord('q')  两者之一是必须要否则会报错，c == 27 时是用esc关闭的 ，点窗口的×是不能关闭视频窗口的。

2.函数：VideoCapture(0)

​          打开摄像头，0代表的是设备id，如果有多个摄像头，可以设置其他数值

​          也可以是视频文件地址，调用视频文件，如果要播放要设置帧的循环

3.函数：read() 

  读取摄像头,它能返回两个参数，第一个参数是bool型的ret，其值为True或False，代表有没有读到图片；第二个参数是frame，是当前截取一帧的图片

4.函数：frame = cv.flip(frame, 1)

​          表示翻转    

​           等于0:逆时针180度旋转， 

​          *大于0:正常 ，*

```
   小于0上下颠倒
```





```python
import winreg
import ctypes

# 如果从来没有开过代理 有可能健不存在 会报错
INTERNET_SETTINGS = winreg.OpenKey(winreg.HKEY_CURRENT_USER,
                                   r'Software\Microsoft\Windows\CurrentVersion\Internet Settings',
                                   0, winreg.KEY_ALL_ACCESS)
# 设置刷新
INTERNET_OPTION_REFRESH = 37
INTERNET_OPTION_SETTINGS_CHANGED = 39
internet_set_option = ctypes.windll.Wininet.InternetSetOptionW


def set_key(name, value):
    # 修改键值
    _, reg_type = winreg.QueryValueEx(INTERNET_SETTINGS, name)
    winreg.SetValueEx(INTERNET_SETTINGS, name, 0, reg_type, value)


# 启用代理
set_key('ProxyEnable', 1)  # 启用
set_key('ProxyOverride', u'*.local;<local>')  # 绕过本地
set_key('ProxyServer', u'127.0.0.1:8888')  # 代理IP及端口，将此代理修改为自己的代理IP
internet_set_option(0, INTERNET_OPTION_REFRESH, 0, 0)
internet_set_option(0, INTERNET_OPTION_SETTINGS_CHANGED, 0, 0)
# 停用代理
set_key('ProxyEnable', 0)  # 停用
internet_set_option(0, INTERNET_OPTION_REFRESH, 0, 0)
internet_set_option(0, INTERNET_OPTION_SETTINGS_CHANGED, 0, 0)
```





### 常见代理网站

[https://www.kuaidaili.com/free/](https://www.oschina.net/action/GoToLink?url=https%3A%2F%2Fwww.kuaidaili.com%2Ffree%2F)

[http://www.xicidaili.com/](https://www.oschina.net/action/GoToLink?url=http%3A%2F%2Fwww.xicidaili.com%2F)

常见agent

**一、为何要设置User Agent**

​    有一些网站不喜欢被爬虫程序访问，所以会检测连接对象，如果是爬虫程序，也就是非人点击访问，它就会不让你继续访问，所以为了要让程序可以正常运行，需要隐藏自己的爬虫程序的身份。此时，我们就可以通过设置User Agent的来达到隐藏身份的目的，User Agent的中文名为用户代理，简称UA。

​    User Agent存放于Headers中，服务器就是通过查看Headers中的User Agent来判断是谁在访问。在Python中，如果不设置User Agent，程序将使用默认的参数，那么这个User Agent就会有Python的字样，如果服务器检查User Agent，那么没有设置User Agent的Python程序将无法正常访问网站。

​    Python允许我们修改这个User Agent来模拟浏览器访问，它的强大毋庸置疑。

**二、常见的User Agent**

**1.Android**

- Mozilla/5.0 (Linux; Android 4.1.1; Nexus 7 Build/JRO03D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19
- Mozilla/5.0 (Linux; U; Android 4.0.4; en-gb; GT-I9300 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30
- Mozilla/5.0 (Linux; U; Android 2.2; en-gb; GT-P1000 Build/FROYO) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1

**2.Firefox**

- Mozilla/5.0 (Windows NT 6.2; WOW64; rv:21.0) Gecko/20100101 Firefox/21.0
- Mozilla/5.0 (Android; Mobile; rv:14.0) Gecko/14.0 Firefox/14.0

**3.Google Chrome**

- Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.94 Safari/537.36
- Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Mobile Safari/535.19

**4.iOS**

- Mozilla/5.0 (iPad; CPU OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3
- Mozilla/5.0 (iPod; U; CPU like Mac OS X; en) AppleWebKit/420.1 (KHTML, like Gecko) Version/3.0 Mobile/3A101a Safari/419.3

​    上面列举了Andriod、Firefox、Google Chrome、iOS的一些User Agent，直接copy就能用。

**三、设置User Agent的方法**

![img](https://static.oschina.net/uploads/space/2018/0114/231414_aZPu_2856757.png)

通过代理隐藏自己的ip信息

设置代理字典

```python
proxyDict = {
                      "http"  : self.http_proxy,
                      "https" : self.https_proxy,
                      "ftp"   : self.ftp_proxy
                    }
```

有密码的代理

```python
proxies = {
    "http": "http://user:pass@10.10.1.10:3128/"
}
```

第一种requests模块

```python
# 作者：十四君
# 链接：https://www.zhihu.com/question/23825711/answer/129293723
# 来源：知乎
# 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

# encoding=utf8
import requests
import sys

type = sys.getfilesystemencoding()
s = requests.session()
proxie = {
    # 'http': 'http://122.193.14.102:80'
    # 'http': 'http://61.135.217.7:80'
    # 'http': "http://61.155.164.109:3128"
    # 'https': "https:175.8.85.61:8118",
    # "http": "http://61.178.238.122:63000"
    # 'https': "https://125.88.177.128:3128",
    "http": "http://118.193.107.174:80"
}

url = 'http://www.ahaoboy.cn:888'

print(url)
response = s.get(url, verify=False, proxies=proxie, timeout=20)
print(response.text)
```

 

第二种urlopen模块

```python
from urllib import request

if __name__ == "__main__":
    # 访问网址
    url = 'http://www.ahaoboy.cn:888/'
    # 这是代理IP
    proxy = {
        # 'http': '106.46.136.112:808'
        # 'https': "https://112.112.236.145:9999",
        "http": "http://118.193.107.174:80"
    }
    # 创建ProxyHandler
    proxy_support = request.ProxyHandler(proxy)
    # 创建Opener
    opener = request.build_opener(proxy_support)
    # 添加User Angent
    opener.addheaders = [('User-Agent',
                          'Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36')]
    # 安装OPener
    request.install_opener(opener)
    # 使用自己安装好的Opener
    response = request.urlopen(url)
    # 读取相应信息并解码
    html = response.read().decode("utf-8")
    # 打印信息
    print(html)
```



# python设置win系统代理

1 人赞同了该文章

采集数据的时候有时候会需要频繁修改系统代理, 调用python标准库来实现是个不错的方法

## Windows修改代理注册表位置:

> HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Internet Settings

## 该目录下有三个键值：

- ProxyEnable *是否开启代理 1:开启 0:关闭*
- ProxyOverride *不代理的ip,一般是本地*
- ProxyServer *代理服务器ip和端口*

## *用*winreg修改完以后一定要刷新才会生效,就好比要点击应用或确定

```python
import winreg
import ctypes
import get_proxies

# 如果从来没有开过代理 有可能健不存在 会报错
INTERNET_SETTINGS = winreg.OpenKey(winreg.HKEY_CURRENT_USER,
                                   r'Software\Microsoft\Windows\CurrentVersion\Internet Settings',
                                   0, winreg.KEY_ALL_ACCESS)
# 设置刷新
INTERNET_OPTION_REFRESH = 37
INTERNET_OPTION_SETTINGS_CHANGED = 39
internet_set_option = ctypes.windll.Wininet.InternetSetOptionW


def set_key(name, value):
    # 修改键值
    _, reg_type = winreg.QueryValueEx(INTERNET_SETTINGS, name)
    winreg.SetValueEx(INTERNET_SETTINGS, name, 0, reg_type, value)


# 启用代理
def start():
    stop()  # 先关闭代理,请求的代理一般来自api,如果前一个代理ip失效或者没加入白名单,会请求失败
    proxy = get_proxies()
    ip_port = proxy['http'].split("//", 1)[1]  # 形式: 12.145.32.68:8888
    set_key('ProxyEnable', 1)  # 启用
    # 本地链接不代理
    set_key('ProxyOverride',
            u'localhost;127.*;10.*;172.16.*;172.17.*;172.18.*;172.19.*;172.20.*;172.21.*;172.22.*;172.23.*;172.24.*;172.25.*;172.26.*;172.27.*;172.28.*;172.29.*;172.30.*;172.31.*;192.168.*;127.0.0.1"""')  
    set_key('ProxyServer', u'{}'.format(ip_port))  # 代理IP及端口，将此代理修改为自己的代理IP
    internet_set_option(0, INTERNET_OPTION_REFRESH, 0, 0)
    internet_set_option(0, INTERNET_OPTION_SETTINGS_CHANGED, 0, 0)
    print(f'当前代理: {ip_port}')


def stop():
    # 停用代理
    set_key('ProxyEnable', 0)  # 停用
    internet_set_option(0, INTERNET_OPTION_REFRESH, 0, 0)
    internet_set_option(0, INTERNET_OPTION_SETTINGS_CHANGED, 0, 0)
```



##### proxy.py

```python
# !/usr/bin/python
# _*_ coding: utf-8 _*_
# @File     : proxy.py

import sys
import socket
import threading
#回复消息，原样返回
def replyMessage(conn):
  while True:
    data = conn.recv(1024)
    conn.send(data)
    if data.decode().lower() == 'bye':
      break
  conn.close()
def main():
  sockScr = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
  sockScr.bind(('', port))
  sockScr.listen(200)
  while True:
    try:
      conn, addr = sockScr.accept()
      #只允许特定主机访问本服务器
      if addr[0] != onlyYou:
        conn.close()
        continue
      #创建并启动线程
      t = threading.Thread(target=replyMessage, args=(conn,))
      t.start()
    except:
      print('error')
if __name__ == '__main__':
  try:
    #获取命令行参数
    port = int(sys.argv[1])
    onlyYou = sys.argv[2]
    main()
  except:
    print('Must give me a number as port')
```

##### proxyServer.py

```python
# !/usr/bin/python
# _*_ coding: utf-8 _*_
# @File     : proxyServer.py
import sys
import socket
import threading
def middle(conn, addr):
  #面向服务器的Socket
  sockDst = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
  sockDst.connect((ipServer,portServer))
  while True:
    data = conn.recv(1024).decode()
    print('收到客户端消息：'+data)
    if data == '不要发给服务器':
      conn.send('该消息已被代理服务器过滤'.encode())
      print('该消息已过滤')
    elif data.lower() == 'bye':
      print(str(addr)+'客户端关闭连接')
      break
    else:
      sockDst.send(data.encode())
      print('已转发服务器')
      data_fromServer = sockDst.recv(1024).decode()
      print('收到服务器回复的消息：'+data_fromServer)
      if data_fromServer == '不要发给客户端':
        conn.send('该消息已被代理服务器修改'.encode())
        print('消息已被篡改')
      else:
        conn.send(b'Server reply:'+data_fromServer.encode())
        print('已转发服务器消息给客户端')
  conn.close()
  sockDst.close()
def main():
  sockScr = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
  sockScr.bind(('', portScr))
  sockScr.listen(200)
  print('代理已启动')
  while True:
    try:
      conn, addr = sockScr.accept()
      t = threading.Thread(target=middle, args=(conn, addr))
      t.start()
      print('新客户：'+str(addr))
    except:
      pass
if __name__ == '__main__':
  try:
    #(本机IP地址,portScr)<==>(ipServer,portServer)
    #代理服务器监听端口
    portScr = int(sys.argv[1])
    #服务器IP地址与端口号
    ipServer = sys.argv[2]
    portServer = int(sys.argv[3])
    main()
  except:
    print('Sth error')
```

##### proxyClient.py

```python
# !/usr/bin/python
# _*_ coding: utf-8 _*_
# @File     : proxyClient.py

import sys
import socket
def main():
  sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
  sock.connect((ip, port))
  while True:
    data = input('What do you want to ask:')
    sock.send(data.encode())
    print(sock.recv(1024).decode())
    if data.lower() == 'bye':
      break
  sock.close()
if __name__ == '__main__':
  try:
    #代理服务器的IP地址和端口号
    ip = sys.argv[1]
    port = int(sys.argv[2])
    main()
  except:
    print('Sth error')
```


https://github.com/pzspsh/Code-management-library/blob/main/images/1648447422850.png
https://github.com/pzspsh/Code-management-library/blob/main/images/1648447443658.png
https://github.com/pzspsh/Code-management-library/blob/main/images/1648447405011.png


#### 使用socks5代理访问

##### client.py

```python
# -*- coding: utf-8 -*-
# time: 2022-3-28 17:27:49
# desc: 测试使用socks5代理访问

import socket
import socks
import requests

# 设置代理
# socks.set_default_proxy(socks.SOCKS5, "10.0.36.74", 2019)
# 如果使用账号密码验证，那么使用下面这行连接方式
socks.set_default_proxy(socks.SOCKS5, "10.0.36.74", 2019, username='panzhongsheng', password='123456')
socket.socket = socks.socksocket

# 测试访问 重庆大学
test_url = 'https://www.hao123.com/'
html = requests.get(test_url, timeout=8)
html.encoding = 'utf-8'
print(html.text)

```

##### server.py

```python
# -*- coding: utf-8 -*-

import select
import socket
import struct
import platform
from socketserver import StreamRequestHandler as Tcp, ThreadingTCPServer

SOCKS_VERSION = 5  # socks版本

"""
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++

    一、客户端认证请求
        +----+----------+----------+
        |VER | NMETHODS | METHODS  |
        +----+----------+----------+
        | 1  |    1     |  1~255   |
        +----+----------+----------+
    二、服务端回应认证
        +----+--------+
        |VER | METHOD |
        +----+--------+
        | 1  |   1    |
        +----+--------+
    三、客户端连接请求(连接目的网络)
        +----+-----+-------+------+----------+----------+
        |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
        +----+-----+-------+------+----------+----------+
        | 1  |  1  |   1   |  1   | Variable |    2     |
        +----+-----+-------+------+----------+----------+
    四、服务端回应连接
        +----+-----+-------+------+----------+----------+
        |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
        +----+-----+-------+------+----------+----------+
        | 1  |  1  |   1   |  1   | Variable |    2     |
        +----+-----+-------+------+----------+----------+

++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
"""


class DYProxy(Tcp):
    # 用户认证 用户名/密码
    username = 'pan'
    password = '123456'

    def handle(self):
        print("客户端：", self.client_address, " 请求连接！")
        """
        一、客户端认证请求
            +----+----------+----------+
            |VER | NMETHODS | METHODS  |
            +----+----------+----------+
            | 1  |    1     |  1~255   |
            +----+----------+----------+
        """
        # 从客户端读取并解包两个字节的数据
        header = self.connection.recv(2)
        VER, NMETHODS = struct.unpack("!BB", header)
        # 设置socks5协议，METHODS字段的数目大于0
        assert VER == SOCKS_VERSION, 'SOCKS版本错误'

        # 接受支持的方法
        # 无需认证：0x00    用户名密码认证：0x02
        # assert NMETHODS > 0
        methods = self.IsAvailable(NMETHODS)
        # 检查是否支持该方式，不支持则断开连接
        if 0 not in set(methods):
            self.server.close_request(self.request)
            return

        """
        二、服务端回应认证
            +----+--------+
            |VER | METHOD |
            +----+--------+
            | 1  |   1    |
            +----+--------+
        """
        # 发送协商响应数据包 
        self.connection.sendall(struct.pack("!BB", SOCKS_VERSION, 0))

        # 校验用户名和密码
        # if not self.VerifyAuth():
        #    return

        """
        三、客户端连接请求(连接目的网络)
            +----+-----+-------+------+----------+----------+
            |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
            +----+-----+-------+------+----------+----------+
            | 1  |  1  |   1   |  1   | Variable |    2     |
            +----+-----+-------+------+----------+----------+
        """
        version, cmd, _, address_type = struct.unpack("!BBBB", self.connection.recv(4))
        assert version == SOCKS_VERSION, 'socks版本错误'
        if address_type == 1:  # IPv4
            # 转换IPV4地址字符串（xxx.xxx.xxx.xxx）成为32位打包的二进制格式（长度为4个字节的二进制字符串）
            address = socket.inet_ntoa(self.connection.recv(4))
        elif address_type == 3:  # Domain
            domain_length = ord(self.connection.recv(1)[0])
            address = self.connection.recv(domain_length)
        port = struct.unpack('!H', self.connection.recv(2))[0]

        """
        四、服务端回应连接
            +----+-----+-------+------+----------+----------+
            |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
            +----+-----+-------+------+----------+----------+
            | 1  |  1  |   1   |  1   | Variable |    2     |
            +----+-----+-------+------+----------+----------+
        """
        # 响应，只支持CONNECT请求
        try:
            if cmd == 1:  # CONNECT
                remote = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
                remote.connect((address, port))
                bind_address = remote.getsockname()
                print('已建立连接：', address, port)
            else:
                self.server.close_request(self.request)
            addr = struct.unpack("!I", socket.inet_aton(bind_address[0]))[0]
            port = bind_address[1]
            reply = struct.pack("!BBBBIH", SOCKS_VERSION, 0, 0, address_type, addr, port)
        except Exception as err:
            print(err)
            # 响应拒绝连接的错误
            reply = self.ReplyFaild(address_type, 5)
        self.connection.sendall(reply)  # 发送回复包

        # 建立连接成功，开始交换数据
        if reply[1] == 0 and cmd == 1:
            self.ExchangeData(self.connection, remote)
        self.server.close_request(self.request)

    def IsAvailable(self, n):
        """ 
        检查是否支持该验证方式 
        """
        methods = []
        for i in range(n):
            methods.append(ord(self.connection.recv(1)))
        return methods

    def VerifyAuth(self):
        """
        校验用户名和密码
        """
        version = ord(self.connection.recv(1))
        assert version == 1
        username_len = ord(self.connection.recv(1))
        username = self.connection.recv(username_len).decode('utf-8')
        password_len = ord(self.connection.recv(1))
        password = self.connection.recv(password_len).decode('utf-8')
        if username == self.username and password == self.password:
            # 验证成功, status = 0
            response = struct.pack("!BB", version, 0)
            self.connection.sendall(response)
            return True
        # 验证失败, status != 0
        response = struct.pack("!BB", version, 0xFF)
        self.connection.sendall(response)
        self.server.close_request(self.request)
        return False

    def ReplyFaild(self, address_type, error_number):
        """ 
        生成连接失败的回复包 
        """
        return struct.pack("!BBBBIH", SOCKS_VERSION, error_number, 0, address_type, 0, 0)

    def ExchangeData(self, client, remote):
        """ 
        交换数据 
        """
        while True:
            # 等待数据
            rs, ws, es = select.select([client, remote], [], [])
            if client in rs:
                data = client.recv(4096)
                if remote.send(data) <= 0:
                    break
            if remote in rs:
                data = remote.recv(4096)
                if client.send(data) <= 0:
                    break


def getip():
    try:
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        s.connect(('8.8.8.8', 80))
        ip = s.getsockname()[0]
    except:
        ip = "127.0.0.1"
    finally:
        s.close()
        return ip


if __name__ == '__main__':
    # 服务器上创建一个TCP多线程服务，监听2019端口
    ip = "0.0.0.0"
    sysType = platform.system()
    if sysType == "Windows":
        hostname = socket.gethostname()
        ip = socket.gethostbyname(hostname)
    elif sysType == "Linux":
        ip = getip()
    elif sysType == "Darwin":
        ip = socket.gethostbyname(socket.gethostname())
    Server = ThreadingTCPServer(('0.0.0.0', 2019), DYProxy)
    print("**********************************************************")
    print("************************* DYPROXY ************************")
    print("*************************   1.0   ************************")
    print(f"********************  IP:{ip}  ******************")
    print("***********************  PORT:2019  **********************")
    print("**********************************************************")
    Server.serve_forever()

```

server_login.py

```python
# -*- coding: utf-8 -*-

import select
import socket
import struct
from socketserver import StreamRequestHandler as Tcp, ThreadingTCPServer

SOCKS_VERSION = 5                           # socks版本

"""
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++

    一、客户端认证请求
        +----+----------+----------+
        |VER | NMETHODS | METHODS  |
        +----+----------+----------+
        | 1  |    1     |  1~255   |
        +----+----------+----------+
    二、服务端回应认证
        +----+--------+
        |VER | METHOD |
        +----+--------+
        | 1  |   1    |
        +----+--------+
    三、客户端连接请求(连接目的网络)
        +----+-----+-------+------+----------+----------+
        |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
        +----+-----+-------+------+----------+----------+
        | 1  |  1  |   1   |  1   | Variable |    2     |
        +----+-----+-------+------+----------+----------+
    四、服务端回应连接
        +----+-----+-------+------+----------+----------+
        |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
        +----+-----+-------+------+----------+----------+
        | 1  |  1  |   1   |  1   | Variable |    2     |
        +----+-----+-------+------+----------+----------+

++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
"""

class DYProxy(Tcp):
    # 用户认证 用户名/密码
    username = 'pan'
    password = '123456'

    def handle(self):
        print("客户端：", self.client_address, " 请求连接！")
        """
        一、客户端认证请求
            +----+----------+----------+
            |VER | NMETHODS | METHODS  |
            +----+----------+----------+
            | 1  |    1     |  1~255   |
            +----+----------+----------+
        """
        # 从客户端读取并解包两个字节的数据
        header = self.connection.recv(2)
        VER, NMETHODS = struct.unpack("!BB", header)
        # 设置socks5协议，METHODS字段的数目大于0
        assert VER == SOCKS_VERSION, 'SOCKS版本错误'
        
        # 接受支持的方法
        # 无需认证：0x00    用户名密码认证：0x02
        assert NMETHODS > 0
        methods = self.IsAvailable(NMETHODS)
        # 检查是否支持该方式，不支持则断开连接
        if 2 not in set(methods):
            self.server.close_request(self.request)
            return
        
        """
        二、服务端回应认证
            +----+--------+
            |VER | METHOD |
            +----+--------+
            | 1  |   1    |
            +----+--------+
        """
        # 发送协商响应数据包 
        self.connection.sendall(struct.pack("!BB", SOCKS_VERSION, 2))
        
        # 校验用户名和密码
        if not self.VerifyAuth():
            return
        

        """
        三、客户端连接请求(连接目的网络)
            +----+-----+-------+------+----------+----------+
            |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
            +----+-----+-------+------+----------+----------+
            | 1  |  1  |   1   |  1   | Variable |    2     |
            +----+-----+-------+------+----------+----------+
        """
        version, cmd, _, address_type = struct.unpack("!BBBB", self.connection.recv(4))
        assert version == SOCKS_VERSION, 'socks版本错误'
        if address_type == 1:       # IPv4
            # 转换IPV4地址字符串（xxx.xxx.xxx.xxx）成为32位打包的二进制格式（长度为4个字节的二进制字符串）
            address = socket.inet_ntoa(self.connection.recv(4))
        elif address_type == 3:     # Domain
            domain_length = ord(self.connection.recv(1)[0])
            address = self.connection.recv(domain_length)
        port = struct.unpack('!H', self.connection.recv(2))[0]

        """
        四、服务端回应连接
            +----+-----+-------+------+----------+----------+
            |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
            +----+-----+-------+------+----------+----------+
            | 1  |  1  |   1   |  1   | Variable |    2     |
            +----+-----+-------+------+----------+----------+
        """
        # 响应，只支持CONNECT请求
        try:
            if cmd == 1:  # CONNECT
                remote = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
                remote.connect((address, port))
                bind_address = remote.getsockname()
                print('已建立连接：', address, port)
            else:
                self.server.close_request(self.request)
            addr = struct.unpack("!I", socket.inet_aton(bind_address[0]))[0]
            port = bind_address[1]
            reply = struct.pack("!BBBBIH", SOCKS_VERSION, 0, 0, address_type, addr, port)
        except Exception as err:
            print(err)
            # 响应拒绝连接的错误
            reply = self.ReplyFaild(address_type, 5)
        self.connection.sendall(reply)      # 发送回复包

        # 建立连接成功，开始交换数据
        if reply[1] == 0 and cmd == 1:
            self.ExchangeData(self.connection, remote)
        self.server.close_request(self.request)


    def IsAvailable(self, n):
        """ 
        检查是否支持该验证方式 
        """
        methods = []
        for i in range(n):
            methods.append(ord(self.connection.recv(1)))
        return methods


    def VerifyAuth(self):
        """
        校验用户名和密码
        """
        version = ord(self.connection.recv(1))
        assert version == 1
        username_len = ord(self.connection.recv(1))
        username = self.connection.recv(username_len).decode('utf-8')
        password_len = ord(self.connection.recv(1))
        password = self.connection.recv(password_len).decode('utf-8')
        if username == self.username and password == self.password:
            # 验证成功, status = 0
            response = struct.pack("!BB", version, 0)
            self.connection.sendall(response)
            return True
        # 验证失败, status != 0
        response = struct.pack("!BB", version, 0xFF)
        self.connection.sendall(response)
        self.server.close_request(self.request)
        return False


    def ReplyFaild(self, address_type, error_number):
        """ 
        生成连接失败的回复包 
        """
        return struct.pack("!BBBBIH", SOCKS_VERSION, error_number, 0, address_type, 0, 0)


    def ExchangeData(self, client, remote):
        """ 
        交换数据 
        """
        while True:
            # 等待数据
            rs, ws, es = select.select([client, remote], [], [])
            if client in rs:
                data = client.recv(4096)
                if remote.send(data) <= 0:
                    break
            if remote in rs:
                data = remote.recv(4096)
                if client.send(data) <= 0:
                    break


if __name__ == '__main__':
    # 服务器上创建一个TCP多线程服务，监听2019端口
    Server = ThreadingTCPServer(('0.0.0.0', 2019), DYProxy)
    print("**********************************************************")
    print("************************* DYPROXY ************************")
    print("*************************   1.0   ************************")
    print("********************  IP:xxx.xxx.xx.xx  ******************")
    print("***********************  PORT:2019  **********************")
    print("**********************************************************")
    Server.serve_forever();

```
![Image text](https://github.com/pzspsh/Code-management-library/blob/main/images/1648451697430.png)
