# docker 创建的elasticsearch容器添加密码
### elasticsearch安装
```
docker pull docker.io/elasticsearch:版本号
docker pull docker.io/elasticsearch:7.1.1
docker pull docker.io/elasticsearch:latest # 表示最新版本

docker images # 查看镜像
```
### 运行并启动容器elasticsearch
```
docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:latest

参数说明：
elasticsearch:latest # 表示你安装的容器对应版本号
-d # 后台启动
--name # 容器名称
-p # 映射端口
-e # 设置环境变量
discovery.type=single-node # 单机运行

如果启动不了，可以加大内存设置：-e ES_JAVA_OPTS="-Xms512m -Xmx512m"

docker ps 查看容器运行状态
```

## 设置密码
```
1、首先进入容器：
    docker exec -it 容器id号 bash

2、启用认证添加编辑如下：
    vim config/elasticsearch.yml
    添加内容：
    http.cors.enabled: true
    http.cors.allow-origin: "*"
    http.cors.allow-headers: Authorization
    xpack.security.enabled: true
    xpack.security.transport.ssl.enabled: true
    保存后退出重启elasticsearch容器

3、设置用户密码：
    1.再次进入容器：
    docker exec -it elasticsearch /bin/bash
    2.输入命令：
    ./bin/elasticsearch-setup-passwords interactive
    出现内容大概如下：
    Initiating the setup of passwords for reserved users elastic,apm_system,kibana,logstash_system,beats_system,remote_monitoring_user.
You will be prompted to enter passwords as the process progresses.
Please confirm that you would like to continue [y/N]
    3.选择y就可以设置你想设置的密码
    4.退出容器再次重启容器就可以使用你设置的密码
```