## k8s安装
### 开发环境安装
```bash
mkdir -p $GOPATH/src/k8s.io
cd $GOPATH/src/k8s.io
git clone https://github.com/kubernetes/kubernetes
cd kubernetes
make


cd $GOPATH/src/k8s.io/kubernetes
export KUBERNETES_PROVIDER=local
hack/install-etcd.sh # 使用源码自带脚本安装etcd
export PATH=$GOPATH/src/k8s.io/kubernetes/third_party/etcd:$PATH

hack/local-up-cluster.sh  # 则可以用hack/local-up-cluster.sh来启动一个本地集群：

# 打开另外一个终端，配置kubectl：
cd $GOPATH/src/k8s.io/kubernetes
export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig
运行命令：
cluster/kubectl.sh
cluster/kubectl.sh cluster-info
cluster/kubectl.sh get pods -n kube-system

查看日志：
cat /tmp/kube-apiserver.log

可参考：
http://docs.kubernetes.org.cn/109.html
https://zhuanlan.zhihu.com/p/332751754
```

docker环境安装
```bash
git clone https://github.com/kubernetes/kubernetes
cd kubernetes
make quick-release
```

### k8s + harbor + kubesphere集群使用教程
#### 1、安装
```

```

### k8s + harbor + kubePi + harness集群使用教程
#### 1、安装、配置
```

```