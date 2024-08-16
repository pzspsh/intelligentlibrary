# ssh教程
#### ssh安装
```shell
# Ubuntu 和 Debian
$ sudo apt install openssh-client

# CentOS 和 Fedora
$ sudo dnf install openssh-clients
```
#### ssh登录
```shell
$ ssh hostname # hostname是主机名，它可以是域名，也可能是 IP 地址或局域网内部的主机名。不指定用户名的情况下，将使用客户端的当前用户名
$ ssh user@hostname # 指定用户名

$ ssh -l username host # 用户名也可以使用ssh的-l参数指定，这样的话，用户名和主机名就不用写在一起

$ ssh -p 8821 foo.com # ssh 默认连接服务器的22端口，-p参数可以指定其他端口
```
#### 执行远程命令
```shell
$ ssh username@hostname command # 命令会使得 SSH 在登录成功后，立刻在远程主机上执行命令command
例子：
    $ ssh foo@server.example.com cat /etc/hosts

1、ssh 客户端不会提供互动式的 Shell 环境，而是直接远程命令的执行结果输出在命令行。但是，有些命令需要互动式的 Shell 环境，这时就要使用-t参数
# 报错
$ ssh remote.server.com emacs
emacs: standard input is not a tty

# 不报错
$ ssh -t server.example.com emacs
```

#### 修改密码
```shell
1、修改登录密码
passwd  # 打开终端执行该命令，修改你要设置的密码
2、开启root远程密码,使用pubkey登录到普通用户,然后执行以下命令：
$ sudo passwd # 输入你要修改的密码
Changing password for user root.
New password: # 输入设置的新密码
Retype new password: # 确认新密码
passwd: all authentication tokens updated successfully.

3、登录root用户
$ su root
Password:  # 输入密码

4、修改sshd配置
vi /etc/ssh/sshd_config # 修改一些内容
PermitRootLogin yes #允许root登录
PasswordAuthentication yes #允许密码登录

5、重启sshd服务
service sshd restart
```



# ssh 命令行配置项

ssh 命令有很多配置项，修改它的默认行为。

**-c**

`-c`参数指定加密算法。

```
$ ssh -c blowfish,3des server.example.com
# 或者
$ ssh -c blowfish -c 3des server.example.com
```

上面命令指定使用加密算法`blowfish`或`3des`。

**-C**

`-C`参数表示压缩数据传输。

```
$ ssh -C server.example.com
```

**-D**

`-D`参数指定本机的 Socks 监听端口，该端口收到的请求，都将转发到远程的 SSH 主机，又称动态端口转发，详见《端口转发》一章。

```
$ ssh -D 1080 server
```

上面命令将本机 1080 端口收到的请求，都转发到服务器`server`。

**-f**

`-f`参数表示 SSH 连接在后台运行。

**-F**

`-F`参数指定配置文件。

```
$ ssh -F /usr/local/ssh/other_config
```

上面命令指定使用配置文件`other_config`。

**-i**

`-i`参数用于指定私钥，意为“identity_file”，默认值为`~/.ssh/id_dsa`。注意，对应的公钥必须存放到服务器，详见《密钥登录》一章。

```
$ ssh -i my-key server.example.com
```

**-l**

`-l`参数指定远程登录的账户名。

```
$ ssh -l sally server.example.com
# 等同于
$ ssh sally@server.example.com
```

**-L**

`-L`参数设置本地端口转发，详见《端口转发》一章。

```
$ ssh  -L 9999:targetServer:80 user@remoteserver
```

上面命令中，所有发向本地`9999`端口的请求，都会经过`remoteserver`发往 targetServer 的 80 端口，这就相当于直接连上了 targetServer 的 80 端口。

**-m**

`-m`参数指定校验数据完整性的算法（message authentication code，简称 MAC）。

```
$ ssh -m hmac-sha1,hmac-md5 server.example.com
```

上面命令指定数据校验算法为`hmac-sha1`或`hmac-md5`。

**-N**

`-N`参数用于端口转发，表示建立的 SSH 只用于端口转发，不能执行远程命令，这样可以提供安全性，详见《端口转发》一章。

**-o**

`-o`参数用来指定一个配置命令。

```
$ ssh -o "Keyword Value"
```

举例来说，配置文件里面有如下内容。

```
User sally
Port 220
```

通过`-o`参数，可以把上面两个配置命令从命令行传入。

```
$ ssh -o "User sally" -o "Port 220" server.example.com
```

使用等号时，配置命令可以不用写在引号里面，但是等号前后不能有空格。

```
$ ssh -o User=sally -o Port=220 server.example.com
```

**-p**

`-p`参数指定 SSH 客户端连接的服务器端口。

```
$ ssh -p 2035 server.example.com
```

上面命令连接服务器的2035端口。

**-q**

`-q`参数表示安静模式（quiet），不向用户输出任何警告信息。

```
$ ssh –q foo.com
root’s password:
```

上面命令使用`-q`参数，只输出要求用户输入密码的提示。

**-R**

`-R`参数指定远程端口转发，详见《端口转发》一章。

```
$ ssh -R 9999:targetServer:902 local
```

上面命令需在跳板服务器执行，指定本地计算机`local`监听自己的 9999 端口，所有发向这个端口的请求，都会转向 targetServer 的 902 端口。

**-t**

`-t`参数在 ssh 直接运行远端命令时，提供一个互动式 Shell。

```
$ ssh -t server.example.com emacs
```

**-v**

`-v`参数显示详细信息。

```
$ ssh -v server.example.com
```

`-v`可以重复多次，表示信息的详细程度，比如`-vv`和`-vvv`。

```
$ ssh -vvv server.example.com
# 或者
$ ssh -v -v -v server.example.com
```

上面命令会输出最详细的连接信息。

**-V**

`-V`参数输出 ssh 客户端的版本。

```
$ ssh –V
ssh: SSH Secure Shell 3.2.3 (non-commercial version) on i686-pc-linux-gnu
```

上面命令输出本机 ssh 客户端版本是`SSH Secure Shell 3.2.3`。

**-X**

`-X`参数表示打开 X 窗口转发。

```
$ ssh -X server.example.com
```

**-1，-2**

`-1`参数指定使用 SSH 1 协议。

`-2`参数指定使用 SSH 2 协议。

```
$ ssh -2 server.example.com
```

**-4，-6**

`-4`指定使用 IPv4 协议，这是默认值。

```
$ ssh -4 server.example.com
```

`-6`指定使用 IPv6 协议。

```
$ ssh -6 server.example.com
```

# 配置文件位置

SSH 客户端的全局配置文件是`/etc/ssh/ssh_config`，用户个人的配置文件在`~/.ssh/config`，优先级高于全局配置文件。

除了配置文件，`~/.ssh`目录还有一些用户个人的密钥文件和其他文件。下面是其中一些常见的文件。

- `~/.ssh/id_ecdsa`：用户的 ECDSA 私钥。
- `~/.ssh/id_ecdsa.pub`：用户的 ECDSA 公钥。
- `~/.ssh/id_rsa`：用于 SSH 协议版本2 的 RSA 私钥。
- `~/.ssh/id_rsa.pub`：用于SSH 协议版本2 的 RSA 公钥。
- `~/.ssh/identity`：用于 SSH 协议版本1 的 RSA 私钥。
- `~/.ssh/identity.pub`：用于 SSH 协议版本1 的 RSA 公钥。
- `~/.ssh/known_hosts`：包含 SSH 服务器的公钥指纹。
- 

# 主机设置

用户个人的配置文件`~/.ssh/config`，可以按照不同服务器，列出各自的连接参数，从而不必每一次登录都输入重复的参数。下面是一个例子。

```
Host *
     Port 2222

Host remoteserver
     HostName remote.example.com
     User neo
     Port 2112
```

上面代码中，`Host *`表示对所有主机生效，后面的`Port 2222`表示所有主机的默认连接端口都是2222，这样就不用在登录时特别指定端口了。这里的缩进并不是必需的，只是为了视觉上，易于识别针对不同主机的设置。

后面的`Host remoteserver`表示，下面的设置只对主机`remoteserver`生效。`remoteserver`只是一个别名，具体的主机由`HostName`命令指定，`User`和`Port`这两项分别表示用户名和端口。这里的`Port`会覆盖上面`Host *`部分的`Port`设置。

以后，登录`remote.example.com`时，只要执行`ssh remoteserver`命令，就会自动套用 config 文件里面指定的参数。 单个主机的配置格式如下。

```
$ ssh remoteserver
# 等同于
$ ssh -p 2112 neo@remote.example.com
```

`Host`命令的值可以使用通配符，比如`Host *`表示对所有主机都有效的设置，`Host *.edu`表示只对一级域名为`.edu`的主机有效的设置。它们的设置都可以被单个主机的设置覆盖。

## 配置命令的语法

ssh 客户端配置文件的每一行，就是一个配置命令。配置命令与对应的值之间，可以使用空格，也可以使用等号。

```
Compression yes
# 等同于
Compression = yes
```

`#`开头的行表示注释，会被忽略。空行等同于注释。

## 主机配置命令

下面是 ssh 客户端的一些主要配置命令，以及它们的范例值。

- `AddressFamily inet`：表示只使用 IPv4 协议。如果设为`inet6`，表示只使用 IPv6 协议。
- `BindAddress 192.168.10.235`：指定本机的 IP 地址（如果本机有多个 IP 地址）。
- `CheckHostIP yes`：检查 SSH 服务器的 IP 地址是否跟公钥数据库吻合。
- `Ciphers blowfish,3des`：指定加密算法。
- `Compression yes`：是否压缩传输信号。
- `ConnectionAttempts 10`：客户端进行连接时，最大的尝试次数。
- `ConnectTimeout 60`：客户端进行连接时，服务器在指定秒数内没有回复，则中断连接尝试。
- `DynamicForward 1080`：指定动态转发端口。
- `GlobalKnownHostsFile /users/smith/.ssh/my_global_hosts_file`：指定全局的公钥数据库文件的位置。
- `Host server.example.com`：指定连接的域名或 IP 地址，也可以是别名，支持通配符。`Host`命令后面的所有配置，都是针对该主机的，直到下一个`Host`命令为止。
- `HostKeyAlgorithms ssh-dss,ssh-rsa`：指定密钥算法，优先级从高到低排列。
- `HostName myserver.example.com`：在`Host`命令使用别名的情况下，`HostName`指定域名或 IP 地址。
- `IdentityFile keyfile`：指定私钥文件。
- `LocalForward 2001 localhost:143`：指定本地端口转发。
- `LogLevel QUIET`：指定日志详细程度。如果设为`QUIET`，将不输出大部分的警告和提示。
- `MACs hmac-sha1,hmac-md5`：指定数据校验算法。
- `NumberOfPasswordPrompts 2`：密码登录时，用户输错密码的最大尝试次数。
- `PasswordAuthentication no`：指定是否支持密码登录。不过，这里只是客户端禁止，真正的禁止需要在 SSH 服务器设置。
- `Port 2035`：指定客户端连接的 SSH 服务器端口。
- `PreferredAuthentications publickey,hostbased,password`：指定各种登录方法的优先级。
- `Protocol 2`：支持的 SSH 协议版本，多个版本之间使用逗号分隔。
- `PubKeyAuthentication yes`：是否支持密钥登录。这里只是客户端设置，还需要在 SSH 服务器进行相应设置。
- `RemoteForward 2001 server:143`：指定远程端口转发。
- `SendEnv COLOR`：SSH 客户端向服务器发送的环境变量名，多个环境变量之间使用空格分隔。环境变量的值从客户端当前环境中拷贝。
- `ServerAliveCountMax 3`：如果没有收到服务器的回应，客户端连续发送多少次`keepalive`信号，才断开连接。该项默认值为3。
- `ServerAliveInterval 300`：客户端建立连接后，如果在给定秒数内，没有收到服务器发来的消息，客户端向服务器发送`keepalive`消息。如果不希望客户端发送，这一项设为`0`。
- `StrictHostKeyChecking yes`：`yes`表示严格检查，服务器公钥为未知或发生变化，则拒绝连接。`no`表示如果服务器公钥未知，则加入客户端公钥数据库，如果公钥发生变化，不改变客户端公钥数据库，输出一条警告，依然允许连接继续进行。`ask`（默认值）表示询问用户是否继续进行。
- `TCPKeepAlive yes`：客户端是否定期向服务器发送`keepalive`信息。
- `User userName`：指定远程登录的账户名。
- `UserKnownHostsFile /users/smith/.ssh/my_local_hosts_file`：指定当前用户的`known_hosts`文件（服务器公钥指纹列表）的位置。
- `VerifyHostKeyDNS yes`：是否通过检查 SSH 服务器的 DNS 记录，确认公钥指纹是否与`known_hosts`文件保存的一致。

#### /etc/ssh/sshd_config配置文件
```shell
参数选项                                                        说明
Port 22                                                         SSH 预设使用 22 这个 port，您也可以使用多的 port ！
Protocol 2,1                                                    选择的 SSH 协议版本，可以是 1 也可以是 2 ，如果要同时支持两者，就必须要使用 2,1 这个分隔了！
ListenAddress 0.0.0.0                                           监听的主机适配卡！举个例子来说，如果您有两个 IP，分别是 192.168.0.100 及 192.168.2.20 ，那么只想要开放 192.168.0.100 时，就可以写如同下面的样式：
ListenAddress 192.168.0.100                                     只监听来自 192.168.0.100 这个 IP 的SSH联机。如果不使用设定的话，则预设所有接口均接受 SSH
PidFile /var/run/sshd.pid                                       可以放置 SSHD 这个 PID 的档案！左列为默认值
LoginGraceTime 600                                              当使用者连上 SSH server 之后，会出现输入密码的画面，在该画面中，在多久时间内没有成功连上 SSH server ，就断线！时间为秒！
Compression yes                                                 是否可以使用压缩指令？
HostKey /etc/ssh/ssh_host_key                                   SSH version 1 使用的私钥
HostKey /etc/ssh/ssh_host_rsa_key                               SSH version 2 使用的 RSA 私钥
HostKey /etc/ssh/ssh_host_dsa_key                               SSH version 2 使用的 DSA 私钥
KeyRegenerationInterval 3600                                    由前面联机的说明可以知道， version 1 会使用 server 的 Public Key ，每隔一段时间来重新建立一次！时间为秒！
ServerKeyBits 768                                               Server key 的长度！
SyslogFacility AUTH                                             当有人使用 SSH 登入系统的时候，SSH会记录信息
LogLevel INFO                                                   登录记录的等级---》全部
PermitRootLogin no                                              是否允许 root 登入！预设是允许的，但是建议设定成 no！
UserLogin no                                                    在 SSH 底下本来就不接受 login 这个程序的登入！
StrictModes yes                                                 当使用者的 host key 改变之后，Server 就不接受联机
RSAAuthentication yes                                           是否使用纯的 RSA 认证！？仅针对 version 1 ！
PubkeyAuthentication yes                                        是否允许 Public Key ？只有 version 2
AuthorizedKeysFile   .ssh/authorized_keys                       设定若要使用不需要密码登入的账号时，那么那个账号的存放档案所在档名！
RhostsAuthentication no                                         本机系统不使用 .rhosts ， .rhosts 不安全！
IgnoreRhosts yes                                                是否取消使用 ~/.ssh/.rhosts 来做为认证！
RhostsRSAAuthentication no                                      针对 version 1 ，使用 rhosts 档案在/etc/hosts.equiv配合 RSA 演算方式来进行认证！
HostbasedAuthentication no                                      这个项目与上面的项目类似，不过是给 version 2 使用的！
IgnoreUserKnownHosts no                                         是否忽略家目录内的 ~/.ssh/known_hosts 这个档案所记录的主机内容
PasswordAuthentication yes                                      密码验证当然是需要的！
PermitEmptyPasswords no                                         上面那一项如果设定为 yes 的话，这一项就最好设定为 no ，这个项目在是否允许以空的密码登入！
ChallengeResponseAuthentication yes                             挑战任何的密码认证！所以，任何 login.conf 规定的认证方式，均可适用！
PAMAuthenticationViaKbdInt yes                                  是否启用其它的 PAM 模块！启用这个模块将会导致 PasswordAuthentication 设定失效！

与Kerberos 有关的参数设定！底下不用设定
KerberosAuthentication no
KerberosOrLocalPasswd yes
KerberosTicketCleanup yes
KerberosTgtPassing no

有关在 X-Window 底下使用的相关设定
X11Forwarding yes
X11DisplayOffset 10
X11UseLocalhost yes

PrintMotd no                                                    登入后是否显示出一些信息呢？例如上次登入的时间、地点等，预设是 yes ，但是，如果为了安全，可以考虑改为 no ！
PrintLastLog yes                                                显示上次登入的信息！预设也是 yes 
KeepAlive yes                                                   一般而言，如果设定这项目的话，那么 SSH Server 会传送KeepAlive 的讯息给 Client 端，以确保两者的联机正常！在这个情况下，任何一端死掉后， SSH 可以立刻知道！而不会有僵尸程序的发生！
UsePrivilegeSeparation yes                                      使用者的权限设定项目！
MaxStartups 10                                                  同时允许几个尚未登入的联机画面
DenyUsers *                                                     设定受抵挡的使用者名称
AllowUsers *                                                    设定允许的使用者名称
```