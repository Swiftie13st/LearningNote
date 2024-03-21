# K8s

[Kubernetes 文档 | Kubernetes](https://kubernetes.io/zh-cn/docs/home/)

K8s是kubernetes的简称，其本质是一个开源的容器编排系统，主要用于管理容器化的应用，其目标是让部署容器化的应用简单并且高效（powerful）,Kubernetes提供了应用部署，规划，更新，维护的一种机制。  
说简单点：k8s就是一个编排容器的系统，一个可以管理容器应用全生命周期的工具，从创建应用，应用的部署，应用提供服务，扩容缩容应用，应用更新，都非常的方便，而且还可以做到故障自愈，所以，k8s是一个非常强大的容器编排系统。

它使用 Label 和 Pod 的概念来将容器划分为逻辑单元。Pods 是同地协作（co-located）容器的集合，这些容器被共同部署和调度，形成了一个服务，这是 Kubernetes 和其他两个框架的主要区别。相比于基于相似度的容器调度方式（就像 Swarm 和Mesos），这个方法简化了对集群的管理。
通过 Pods 这一抽象的概念，解决 Container 之间的依赖于通信问题。Pods，Services，Deployments 是独立部署的部分，可以通过 Selector 提供更多的灵活性。内置服务注册表和负载平衡。


**主要特性：**

-   高可用，不宕机，自动灾难恢复
-   灰度更新，不影响业务正常运转
-   一键回滚到历史版本
-   方便的伸缩扩展（应用伸缩，机器加减）、提供负载均衡
-   有一个完善的生态

## 部署方案

传统部署 -> 虚拟化部署时代 -> 容器部署时代

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305262129440.png)

### 传统部署方式：

应用直接在物理机上部署，机器资源分配不好控制，出现Bug时，可能机器的大部分资源被某个应用占用，导致其他应用无法正常运行，无法做到应用隔离。

程序员/运维工程师手动操作部署应用，直接将应用部署在目标机器上，由于资源不隔离，容易出现资源争抢、依赖冲突等各方面问题。

### 虚拟机部署

虚拟化功能允许应用程序在VM之间隔离，并提供安全级别，因为一一个应用程序的信息不能被另一应用程序自由地访问。因为虚拟化可以轻松地添加或更新应用程序、降低硬件成本等等，所以虚拟化可以更好地利用物理服务器中的资源，并可以实现更好的可伸缩性。每个VM是一台完整的计算机，在虚拟化硬件之上运行所有组件，包括其自己的操作系统。即在单个物理机上运行多个虚拟机，每个虚拟机都是完整独立的系统，性能损耗大。

利用 OpenStask / VMware 等虚拟化技术，将一台目标机器虚拟化为多个虚拟机器，按照需求将应用部署到不同的虚拟机中，对虚拟机进行动态的水平扩容等管理操作。  
  
==相对传统部署自动化、资源隔离的能力提升了，带来的问题是虚拟化的逻辑过重，导致效率不高，且耗费资源较多==。

### 容器部署

容器类似于VM，但是它们具有轻量级的隔离属性，可以在应用程序之间共享操作系统，因此，容器被认为是轻量级的。容器与VM类似，具有自己的文件系统、CPU、内存、进程空间等。由于它们与基础架构分离，因此可以跨云和OS分发进行移植。

可以理解为轻量级的虚拟化，完美弥补虚拟化技术过重的问题，且由于直接共享主机硬件资源，只是通过系统提供的命名空间等技术实现资源隔离，损耗更小，且效率更高。

==所有容器共享主机的系统，轻量级的虚拟机，性能损耗小，资源隔离，CPU和内存可按需分配==

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202403211039562.png)

## K8s功能

Kubernetes 是一个轻便的和可扩展的开源平台，用于管理容器化应用和服务。通过Kubernetes 能够进行应用的自动化部署和扩缩容。在Kubernetes 中，会将组成应用的容器组合成一个逻辑单元以更易管理和发现。

kubernetes的本质是一组服务器集群，它可以在集群的每个节点上运行特定的程序，来对节点中的容器进行管理。目的是实现资源管理的自动化，主要提供了如下的主要功能：  
  
- **自我修复**：一旦某一个容器崩溃，能够在1秒中左右迅速启动新的容器，自动重启
- **弹性伸缩**：可以根据需要，自动对集群中正在运行的容器数量进行调整
- **服务发现**：服务可以通过自动发现的形式找到它所依赖的服务
- **负载均衡**：如果一个服务起动了多个容器，能够自动实现请求的负载均衡
- **版本回退**：如果发现新发布的程序版本有问题，可以立即回退到原来的版本
- **存储编排**：可以根据容器自身的需求自动创建存储卷，抽象虚拟磁盘
- **自动部署和回滚**：通过配置文件，自动部署，滚动更新
- **机密和配置管理**：动态配置，统一管理
- **批处理**：批量任务


### 自动装箱

基于容器对应用运行环境的资源配置要求自动部署应用容器

### 自我修复(自愈能力)

当容器失败时，会对容器进行重启
当所部署的Node节点有问题时，会对容器进行重新部署和重新调度
当容器未通过监控检查时，会关闭此容器直到容器正常运行时，才会对外提供服务

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305262138775.png)

### 水平扩展

通过简单的命令、用户UI 界面或基于CPU 等资源使用情况，对应用容器进行规模扩大或规模剪裁
当我们有大量的请求来临时，我们可以增加副本数量，从而达到水平扩展的效果
当黄色应用过度忙碌，会来扩展一个应用：

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305262139497.png)

### 服务发现

用户不需使用额外的服务发现机制，就能够基于Kubernetes 自身能力实现服务发现和负载均衡
对外提供统一的入口，让它来做节点的调度和负载均衡， 相当于微服务里面的网关

### 滚动更新

可以根据应用的变化，对应用容器运行的应用，进行一次性或批量式更新
添加应用的时候，不是加进去就马上可以进行使用，而是需要判断这个添加进去的应用是否能够正常使用

### 版本回退

可以根据应用部署情况，对应用容器运行的应用，进行历史版本即时回退
类似于Git中的回滚

### 密钥和配置管理

在不需要重新构建镜像的情况下，可以部署和更新密钥和应用配置，类似热部署。

### 存储编排

自动实现存储系统挂载及应用，特别对有状态应用实现数据持久化非常重要
存储系统可以来自于本地目录、网络存储(NFS、Gluster、Ceph 等)、公共云存储服务

### 批处理

提供一次性任务，定时任务；满足批量数据处理和分析的场景


## K8s集群架构、组件

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305262142679.png)

`kube-apiserver` API 服务器，公开了 Kubernetes API  
`etcd` 键值数据库，可以作为保存 Kubernetes 所有集群数据的后台数据库  
`kube-scheduler` 调度 Pod 到哪个节点运行  
`kube-controller` 集群控制器  
`cloud-controller` 与云服务商交互

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202403211135284.png)

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202403211500814.png)

### Master

主节点，控制平台，不需要很高性能，不跑任务，通常一个就行了，也可以开多个主节点来提高集群可用度。


#### kube-apiserver

API 服务器是 Kubernetes [控制平面](https://kubernetes.io/zh-cn/docs/reference/glossary/?all=true#term-control-plane)的组件， 该组件负责公开了 Kubernetes API，负责处理接受请求的工作。 API 服务器是 Kubernetes 控制平面的前端。  
  
Kubernetes API 服务器的主要实现是 [kube-apiserver](https://kubernetes.io/zh-cn/docs/reference/command-line-tools-reference/kube-apiserver/)。 kube-apiserver 设计上考虑了水平扩缩，也就是说，它可通过部署多个实例来进行扩缩。 你可以运行 kube-apiserver 的多个实例，并在这些实例之间平衡流量。

#### kube-controller-manager

[kube-controller-manager](https://kubernetes.io/zh-cn/docs/reference/command-line-tools-reference/kube-controller-manager/) 是[控制平面](https://kubernetes.io/zh-cn/docs/reference/glossary/?all=true#term-control-plane)的组件， 负责运行[控制器](https://kubernetes.io/zh-cn/docs/concepts/architecture/controller/)进程。  
  
从逻辑上讲， 每个[控制器](https://kubernetes.io/zh-cn/docs/concepts/architecture/controller/)都是一个单独的进程， 但是为了降低复杂性，它们都被编译到同一个可执行文件，并在同一个进程中运行。  

这些控制器包括：

- 节点控制器（Node Controller）：负责在节点出现故障时进行通知和响应
- 任务控制器（Job Controller）：监测代表一次性任务的 Job 对象，然后创建 Pods 来运行这些任务直至完成
- 端点分片控制器（EndpointSlice controller）：填充端点分片（EndpointSlice）对象（以提供 Service 和 Pod 之间的链接）。
- 服务账号控制器（ServiceAccount controller）：为新的命名空间创建默认的服务账号（ServiceAccount）。

#### cloud-controller-manager

嵌入了特定于云平台的控制逻辑。 云控制器管理器（Cloud Controller Manager）允许你将你的集群连接到云提供商的 API 之上， 并将与该云平台交互的组件同与你的集群交互的组件分离开来。  
  
cloud-controller-manager 仅运行特定于云平台的控制器。 因此如果你在自己的环境中运行 Kubernetes，或者在本地计算机中运行学习环境， 所部署的集群不需要有云控制器管理器。  

与 kube-controller-manager 类似，cloud-controller-manager 将若干逻辑上独立的控制回路组合到同一个可执行文件中， 供你以同一进程的方式运行。 你可以对其执行水平扩容（运行不止一个副本）以提升性能或者增强容错能力。

#### kube-scheduler

scheduler 负责资源的调度，按照预定的调度策略将 Pod 调度到相应的机器上；

#### etcd

[etcd文档](https://etcd.io/docs/)

一致且高度可用的键值存储，用作 Kubernetes 的所有集群数据的后台数据库。  

如果你的 Kubernetes 集群使用 etcd 作为其后台数据库， 请确保你针对这些数据有一份 [备份](https://kubernetes.io/zh-cn/docs/tasks/administer-cluster/configure-upgrade-etcd/#backing-up-an-etcd-cluster)计划。  
早期数据存放在内存，现在已经是持久化存储的了。

### worker / Node

工作节点，可以是虚拟机或物理计算机，任务都在这里跑，机器性能需要好点；通常都有很多个，可以不断加机器扩大集群；每个工作节点由主节点管理

#### kubelet

kubelet 负责维护容器的生命周期，同时也负责 Volume（CVI）和网络（CNI）的管理；

#### kube-proxy

kube-proxy 负责为 Service 提供 cluster 内部的服务发现和负载均衡；

#### container runtime

Container runtime 负责镜像管理以及 Pod 和容器的真正运行（CRI）；

Kubernetes 支持许多容器运行环境，例如 [containerd](https://containerd.io/docs/)、 [CRI-O](https://cri-o.io/#what-is-cri-o) 以及 [Kubernetes CRI (容器运行环境接口)](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md) 的其他任何实现。

### 附加组件

- `kube-dns` 负责为整个集群提供 DNS 服务
- `Ingress Controller` 为服务提供外网入口
- `Prometheus` 提供资源监控
- `Dashboard` 提供 GUI
- `Federation` 提供跨可用区的集群
- `Fluentd-elasticsearch` 提供集群日志采集、存储与查询


### Pod

**Pod** 是可以在 Kubernetes 中创建和管理的、最小的可部署的计算单元。
是一组（一个或多个）容器； 这些容器共享存储、网络、以及怎样运行这些容器的声明。 Pod 中的内容总是并置（colocated）的并且一同调度，在共享的上下文中运行。 Pod 所建模的是特定于应用的 “逻辑主机”，其中包含一个或多个应用容器， 这些容器相对紧密地耦合在一起。 在非云环境中，在相同的物理机或虚拟机上运行的应用类似于在同一逻辑主机上运行的云应用。

一个 Pod 可以包含一个或多个容器，每个 Pod 有自己的虚拟IP。一个工作节点可以有多个 pod，主节点会考量负载自动调度 pod 到哪个节点运行。

- 创建容器使用docker，一个docker对应一个容器，一个容器运行一个应用进程
- Pod是多进程设计，运用多个应用程序，也就是一个Pod里面有多个容器，而一个容器里面运行一个应用程序

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305262146394.png)

- 最小部署的单元
- Pod里面是由一个或多个容器组成【一组容器的集合】
- 一个pod中的容器是共享网络命名空间
- Pod是短暂的
- 每个Pod包含一个或多个紧密相关的用户业务容器

#### 创建Pod流程

- 首先创建一个pod，然后创建一个API Server 和 Etcd【把创建出来的信息存储在etcd中】
- 然后创建 Scheduler，监控API Server是否有新的Pod，如果有的话，会通过调度算法，把pod调度某个node上
- 在node节点，会通过 `kubelet -- apiserver ` 读取etcd 拿到分配在当前node节点上的pod，然后通过docker创建容器

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305262206956.png)


### 组件调用关系

以部署一个Nginx服务来说明Kubernetes各个组件的调用关系

1. 首先，K8s环境启动后，master和node都会将自身的信息存储到etcd数据库中
2. 一个Nginx服务的安装请求会首先发送到master节点的apiServer组件
3. apiServer组件会调用scheduler组件来决定应该把这个服务安装到哪个node节点上。它会从etcd中读取各个node的信息，按一定的算法进行选择，并将选择结果告知apiServer
4. apiServer调用controller-manger去调度相应node节点来安装Nginx服务
5. kuberlet接收到指令后，会通知docker/container，然后由docker/container来启动一个Nginx的pod
6. 至此一个Nginx服务就运行成功了。访问该服务时，需通过kube-proxy来对pod产生访问的代理


## k8s的组件有哪些，作用分别是什么？

k8s主要由master节点和node节点构成。master节点负责管理集群，node节点是容器应用真正运行的地方。  
master节点包含的组件有：kube-api-server、kube-controller-manager、kube-scheduler、etcd。  
node节点包含的组件有：kubelet、kube-proxy、container-runtime。

kube-api-server：以下简称api-server，api-server是k8s最重要的核心组件之一，它是k8s集群管理的统一访问入口，提供了RESTful API接口, 实现了认证、授权和准入控制等安全功能；api-server还是其他组件之间的数据交互和通信的枢纽，其他组件彼此之间并不会直接通信，其他组件对资源对象的增、删、改、查和监听操作都是交由api-server处理后，api-server再提交给etcd数据库做持久化存储，只有api-server才能直接操作etcd数据库，其他组件都不能直接操作etcd数据库，其他组件都是通过api-server间接的读取，写入数据到etcd。

kube-controller-manager：以下简称controller-manager，controller-manager是k8s中各种控制器的的管理者，是k8s集群内部的管理控制中心，也是k8s自动化功能的核心；controller-manager内部包含replication controller、node controller、deployment controller、endpoint controller等各种资源对象的控制器，每种控制器都负责一种特定资源的控制流程，而controller-manager正是这些controller的核心管理者。

kube-scheduler：以下简称scheduler，scheduler负责集群资源调度，其作用是将待调度的pod通过一系列复杂的调度算法计算出最合适的node节点，然后将pod绑定到目标节点上。shceduler会根据pod的信息，全部节点信息列表，过滤掉不符合要求的节点，过滤出一批候选节点，然后给候选节点打分，选分最高的就是最佳节点，scheduler就会把目标pod安置到该节点。

Etcd：etcd是一个分布式的键值对存储数据库，主要是用于保存k8s集群状态数据，比如，pod，service等资源对象的信息；etcd可以是单个也可以有多个，多个就是etcd数据库集群，etcd通常部署奇数个实例，在大规模集群中，etcd有5个或7个节点就足够了；另外说明一点，etcd本质上可以不与master节点部署在一起，只要master节点能通过网络连接etcd数据库即可。

kubelet：每个node节点上都有一个kubelet服务进程，kubelet作为连接master和各node之间的桥梁，负责维护pod和容器的生命周期，当监听到master下发到本节点的任务时，比如创建、更新、终止pod等任务，kubelet 即通过控制docker来创建、更新、销毁容器；  
每个kubelet进程都会在api-server上注册本节点自身的信息，用于定期向master汇报本节点资源的使用情况。

kube-proxy：kube-proxy运行在node节点上，在Node节点上实现Pod网络代理，维护网络规则和四层负载均衡工作，kube-proxy会监听api-server中从而获取service和endpoint的变化情况，创建并维护路由规则以提供服务IP和负载均衡功能。简单理解此进程是Service的透明代理兼负载均衡器，其核心功能是将到某个Service的访问请求转发到后端的多个Pod实例上。

container-runtime：容器运行时环境，即运行容器所需要的一系列程序，目前k8s支持的容器运行时有很多，如docker、rkt或其他，比较受欢迎的是docker，但是新版的k8s已经宣布弃用docker。

## K8s核心概念

### Pod

- Pod是K8s中最小的单元
- 一组容器的集合
- 共享网络【一个Pod中的所有容器共享同一网络】
- 生命周期是短暂的（服务器重启后，就找不到了）

### Volume

- 声明在Pod容器中可访问的文件目录
- 可以被挂载到Pod中一个或多个容器指定路径下
- 支持多种后端存储抽象【本地存储、分布式存储、云存储】

### Controller

- 确保预期的pod副本数量【ReplicaSet】
- 无状态应用部署【Deployment】
  - 无状态就是指，不需要依赖于网络或者ip
- 有状态应用部署【StatefulSet】
  - 有状态需要特定的条件
- 确保所有的node运行同一个pod 【DaemonSet】
- 一次性任务和定时任务【Job和CronJob】

### Deployment

- 定义一组Pod副本数目，版本等
- 通过控制器【Controller】维持Pod数目【自动回复失败的Pod】
- 通过控制器以指定的策略控制版本【滚动升级、回滚等】

### Service

- 定义一组pod的访问规则
- Pod的负载均衡，提供一个或多个Pod的稳定访问地址
- 支持多种方式【ClusterIP、NodePort、LoadBalancer】
- 可以用来组合pod，同时对外提供服务

### Label

label：标签，用于对象资源查询，筛选

### Namespace

命名空间，逻辑隔离

如果一个集群中部署了多个应用，所有应用都在一起，就不太好管理，也可以导致名字冲突等。  
我们可以使用 namespace 把应用划分到不同的命名空间，跟代码里的 namespace 是一个概念，只是为了划分空间。

- 一个集群内部的逻辑隔离机制【鉴权、资源】
- 每个资源都属于一个namespace
- 同一个namespace所有资源不能重复
- 不同namespace可以资源名重复

### API

我们通过Kubernetes的API来操作整个集群

同时我们可以通过 kubectl 、ui、curl 最终发送 http + json/yaml 方式的请求给API Server，然后控制整个K8S集群，K8S中所有的资源对象都可以采用 yaml 或 json 格式的文件定义或描述

如下：使用yaml部署一个nginx的pod
