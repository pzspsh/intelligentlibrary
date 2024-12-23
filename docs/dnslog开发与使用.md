# dnslog开发与使用
#### 在渗透测试过程中， 经常会使用到dnslog服务器，用于dns嗅探以及OOB，当前网络上有许多公开的dnslog平台工具可供使用，主要有
```
http://www.dnslog.cn
http://ceye.io
```

### interactsh的使用方法
#### nuclei是一个开源的基于模板的漏洞扫描工具， 对于RCE类漏洞， nuclei会使用dnslog进行检测，其中nuclei默认支持的dnslog服务器为:
```
oast.pro,
oast.live,
oast.site,
oast.online,
oast.fun,
oast.me
```
在安全测试过程中， 这些公开的dnslog地址可能会被封禁掉，无法正常执行dns嗅探。 因此， 可以自己搭建一个dnslog服务器。 这里输出一个操作指导。