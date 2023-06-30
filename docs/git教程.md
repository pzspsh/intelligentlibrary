# Git教程
```shell
1、克隆URL：
    git clone 目标url 别名
2、
```

```shell
https://www.ipaddress.com/ 用该url进行查询
140.82.112.4   github.com
151.101.193.194 github.global.ssl.fastly.net

linux版
vim /etc/hosts # 添加以上内容
centos: nmcli c reload  # 

windows版
windows/system32/drivers/etc/hosts

ipconfig /flushdns
```
问题：fatal: 无法访问 'https://github.com/kubernetes/kubernetes/'：OpenSSL SSL_read: error:0A000126:SSL routines::unexpected eof while reading, errno 0
git config --global http.sslVerify "false"
取消设置
git config --global --unset http.sslVerify
git config --global --unset http.proxy
git config --global --unset https.proxy
