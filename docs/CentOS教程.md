# CentOS教程
#### CentOS 使用firewalld打开关闭防火墙与端口
#### 1、firewalld的基本使用
```shell
systemctl start firewalld  # 启动 
systemctl stop firewalld  # 关闭
systemctl status firewalld  # 查看状态
systemctl disable firewalld  # 开机禁用 
systemctl enable firewalld  # 开机启用 
```


#### 2.systemctl是CentOS7的服务管理工具中主要的工具，它融合之前service和chkconfig的功能于一体。
```shell
systemctl start firewalld.service   # 启动一个服务
systemctl stop firewalld.service   # 关闭一个服务
systemctl restart firewalld.service   # 重启一个服务
systemctl status firewalld.service   # 显示一个服务的状态
systemctl enable firewalld.service   # 在开机时启用一个服务
systemctl disable firewalld.service   # 在开机时禁用一个服务
systemctl is-enabled firewalld.service   # 查看服务是否开机启动
systemctl list-unit-files|grep enabled   # 查看已启动的服务列表
systemctl --failed   # 查看启动失败的服务列表
```

#### 3.配置firewalld-cmd
```shell
firewall-cmd --version  # 查看版本
firewall-cmd --help   # 查看帮助
firewall-cmd --state   # 显示状态
firewall-cmd --zone=public --list-ports   # 查看所有打开的端口
firewall-cmd --reload   # 更新防火墙规则
firewall-cmd --get-active-zones    # 查看区域信息
firewall-cmd --get-zone-of-interface=eth0   # 查看指定接口所属区域
firewall-cmd --panic-on    # 拒绝所有包
firewall-cmd --panic-off   # 取消拒绝状态
firewall-cmd --query-panic   # 查看是否拒绝

如何开启一个端口:
    firewall-cmd --zone=public --add-port=80/tcp --permanent （–permanent永久生效，没有此参数重启后失效） # 添加
    firewall-cmd --reload   # 重新载入 
    firewall-cmd --zone= public --query-port=80/tcp  # 查看
    firewall-cmd --zone= public --remove-port=80/tcp --permanent  # 删除
```

### 1、基本命令操作
```shell

```

### 2、镜像安装
