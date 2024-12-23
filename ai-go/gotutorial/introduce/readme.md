# 介绍(introduce)

### Go语言项目 为什么选择Go语言

#### 执行性能

缩短 API 的响应时长，解决批量请求访问超时的问题。

在 Uwork 的业务场景下，一次 API 批量请求，往往会涉及对另外接口服务的多次调用，而在之前的 **[PHP](https://haicoder.net/php/php-tutorial.html)** 实现模式下，要做到并行调用是非常困难的，串行处理却不能从根本上提高处理性能。

而 **[Go 语言](https://haicoder.net/golang/golang-tutorial.html)** 不一样，通过 **[Go 协程](https://haicoder.net/golang/golang-goroutine.html)** 可以方便的实现 API 的并行处理，达到处理效率的最大化。 依赖 Golang 的高性能 HTTP Server，提升系统吞吐能力，由 PHP 的数百级别提升到数千里甚至过万级别。

#### 开发效率

Go 语言使用起来简单、代码描述效率高、编码规范统一、上手快。 通过少量的代码，即可实现框架的标准化，并以统一的规范快速构建 API 业务逻辑。

能快速的构建各种通用组件和公共类库，进一步提升开发效率，实现特定场景下的功能量产。

#### Go语言应用领域

Golang 发布之后，很多公司特别是云计算公司开始用 Go 语言重构他们的基础架构，很多都是直接采用 Go 语言 进行了开发，最近热火朝天的 **[Docker](https://haicoder.net/docker/docker-course.html)** 就是采用 Go 语言开发的。

使用 Go 语言开发的开源项目非常多。早期的 Go 语言开源项目只是通过 Go 语言与传统项目进行 **[C 语言库](https://haicoder.net/c/c-tutorial.html)** 绑定实现，例如 Qt、Sqlite 等；后期的很多项目都使用 Go 语言进行重新原生实现，这个过程相对于其他语言要简单一些，这也促成了大量使用 Go 语言原生开发项目的出现。

#### 云计算基础设施领域

- docker
- kubernetes
- etcd
- consul
- cloudflare CDN
- 七牛云存储等。

#### 基础软件

- tidb
- influxdb
- cockroachdb

#### 微服务

- go-kit
- micro
- typhon

#### 互联网基础设施

- 以太坊
- hyperledger



### Go语言代表项目

#### Docker

Docker 是一种操作系统层面的虚拟化技术，可以在操作系统和应用程序之间进行隔离，也可以称之为容器。

Docker 可以在一台物理服务器上快速运行一个或多个实例。基于 lxc 的一个虚拟打包工具，能够实现 PAAS 平台的组建。

#### Go语言

Go 语言自己的早期源码使用 **[C 语言](https://haicoder.net/c/c-tutorial.html)** 和汇编语言写成。

从 Go 1.5 版本后，完全使用 Go 语言自身进行编写。Go 语言的源码对了解 Go 语言的底层调度有极大的参考意义，建议希望对 Go 语言有深入了解的读者读一读。

#### Kubernetes

Google 公司开发的构建于 Docker 之上的容器调度服务，用户可以通过 Kubernetes 集群进行云端容器集群管理。

#### etcd

一款分布式、可靠的 KV 存储系统，可以快速进行云配置。

#### beego

beego 是一个类似 **[Python](https://haicoder.net/python/python-tutorial.html)** 的 Tornado 框架，采用了 RESTFul 的设计思路，使用 Go 语言编写的一个极轻量级、高可伸缩性和高性能的 Web 应用框架。

#### martini

一款快速构建模块化的 Web 应用的 Web 框架。

#### codis

国产的优秀分布式 **[Redis](https://haicoder.net/redis/redis-tutorial.html)** 解决方案。

#### delve

Go 语言强大的调试器，被很多集成环境和编辑器整合。



### 使用Go语言公司

#### Facebook

Facebook 也在用 Go 语言，为此他们还专门在 Github 上建立了一个开源组织 facebookgo。

#### 腾讯

腾讯作为国内的大公司，还是敢于尝试的，尤其是 Docker 容器化这一块，他们在 15 年已经做了 docker 万台规模的实践。

#### 百度

目前所知的百度的使用是在运维这边，是百度运维的一个 BFE 项目，负责前端流量的接入。其次就是百度的消息系统。负责公司手百消息通讯系统服务器端开发及维护。

#### 京东

京东云消息推送系统、云存储，以及京东商城等都有使用 Go 语言做开发。

#### 小米

小米对 Golang 的支持，莫过于运维监控系统的开源。此外，小米互娱、小米商城、小米视频、小米生态链等团队都在使用 Golang。

#### 360

360 对 Golang 的使用也不少，一个是开源的日志搜索系统 Poseidon，托管在 Github 上。还有 360 的推送团队也在使用。

#### 七牛云

七牛云用了近 50 万行代码，来实现整个产品。

#### 美团

美团后台流量支撑程序。应用范围：支撑主站后台流量（排序，推荐，搜索等），提供负载均衡，cache，容错，按条件分流，统计运行指标（qps，latency）等功能。



### Go语言项目总结

Go 语言容易上手，解决了并发编程和写底层应用开发效率的痛点，Go 语言有 Google 这个世界一流的技术公司在后面。

Golang 的杀手级应用是 Docker，而 Docker 的生态圈在这几年完全爆棚了。
