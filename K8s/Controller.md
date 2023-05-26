# Controller

Controller是在集群上管理和运行容器的对象，Controller是实际存在的，Pod是虚拟机的

Pod是通过Controller实现应用的运维，比如弹性伸缩，滚动升级等
Pod 和 Controller之间是通过label标签来建立关系，同时Controller又被称为控制器工作负载

## 升级回滚和弹性伸缩

- 升级：  假设从版本为1.14 升级到 1.15 ，这就叫应用的升级【升级可以保证服务不中断】
- 回滚：从版本1.15 变成 1.14，这就叫应用的回滚
- 弹性伸缩：我们根据不同的业务场景，来改变Pod的数量对外提供服务，这就是弹性伸缩

```sh
# 查看历史版本
kubectl rollout history deployment web
# 回滚
kubectl rollout undo deployment web
kubectl rollout undo deployment web --to-revision=2
# 查看回滚状态
kubectl rollout status deployment web

# 弹性伸缩
kubectl scale deployment web --replicas=10 # 创建10个副本
```

## 工作负载分类

-   Deployment  
    适合无状态应用，所有pod等价，可替代
-   StatefulSet  
    有状态的应用，适合数据库这种类型。
-   DaemonSet  
    在每个节点上跑一个 Pod，可以用来做节点监控、节点日志收集等
-   Job & CronJob  
    Job 用来表达的是一次性的任务，而 CronJob 会根据其时间规划反复运行。

### Deployment部署应用

#### 直接命令运行

```sh
kubectl run testapp --image=ccr.ccs.tencentyun.com/k8s-tutorial/test-k8s:v1
```

#### Pod

```yml
apiVersion: v1 
kind: Pod 
metadata: 
	name: test-pod spec: # 定义容器，可以多个 
	containers: 
		- name: test-k8s # 容器名字 
		  image: ccr.ccs.tencentyun.com/k8s-tutorial/test-k8s:v1 # 镜像
```

#### Deployment

```yml
apiVersion: apps/v1 
kind: Deployment 
metadata: # 部署名字 
	name: test-k8s 
spec: 
	replicas: 2 
	# 用来查找关联的 Pod，所有标签都匹配才行 
	selector: 
		matchLabels: 
			app: test-k8s 
	# 定义 Pod 相关数据
	template: 
		metadata: 
			labels:
				app: test-k8s 
		spec: 
		# 定义容器，可以多个 
			containers:
			- name: test-k8s # 容器名字 
			  image: ccr.ccs.tencentyun.com/k8s-tutorial/test-k8s:v1 # 镜像
```

#### Deployment 通过 label 关联起来 Pods

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305262152129.png)

#### Pod 报错解决

如果你运行 `kubectl describe pod/pod-name` 发现 Events 中有下面这个错误

```sh
networkPlugin cni failed to set up pod "test-k8s-68bb74d654-mc6b9_default" network: open /run/flannel/subnet.env: no such file or directory
```

在每个节点创建文件`/run/flannel/subnet.env`写入以下内容，配置后等待一会就好了

```env
FLANNEL_NETWORK=10.244.0.0/16 
FLANNEL_SUBNET=10.244.0.1/24 
FLANNEL_MTU=1450 
FLANNEL_IPMASQ=true
```

#### 更多命令

```
# 查看全部
kubectl get all
# 重新部署
kubectl rollout restart deployment test-k8s
# 命令修改镜像，--record 表示把这个命令记录到操作历史中
kubectl set image deployment test-k8s test-k8s=ccr.ccs.tencentyun.com/k8s-tutorial/test-k8s:v2-with-error --record
# 暂停运行，暂停后，对 deployment 的修改不会立刻生效，恢复后才应用设置
kubectl rollout pause deployment test-k8s
# 恢复
kubectl rollout resume dep

loyment test-k8s
# 输出到文件
kubectl get deployment test-k8s -o yaml >> app2.yaml
# 删除全部资源
kubectl delete all --all
```

### StatefulSet

StatefulSet 是用来管理有状态的应用，例如数据库。  
前面我们部署的应用，都是不需要存储数据，不需要记住状态的，可以随意扩充副本，每个副本都是一样的，可替代的。  
而像数据库、Redis 这类有状态的，则不能随意扩充副本。  
StatefulSet 会固定每个 Pod 的名字

对于StatefulSet中的Pod，每个Pod挂载自己独立的存储，如果一个Pod出现故障，从其他节点启动一个同样名字的Pod，要挂载上原来Pod的存储继续以它的状态提供服务。

适合StatefulSet的业务包括数据库服务MySQL 和 PostgreSQL，集群化管理服务Zookeeper、etcd等有状态服务

StatefulSet的另一种典型应用场景是作为一种比普通容器更稳定可靠的模拟虚拟机的机制。传统的虚拟机正是一种有状态的宠物，运维人员需要不断地维护它，容器刚开始流行时，我们用容器来模拟虚拟机使用，所有状态都保存在容器里，而这已被证明是非常不安全、不可靠的。

使用StatefulSet，Pod仍然可以通过漂移到不同节点提供高可用，而存储也可以通过外挂的存储来提供
高可靠性，StatefulSet做的只是将确定的Pod与确定的存储关联起来保证状态的连续性。

#### 无状态应用

使用 deployment，部署的都是无状态的应用

- 认为Pod都是一样的
- 没有顺序要求
- 不考虑应用在哪个node上运行
- 能够进行随意伸缩和扩展

#### 有状态应用

上述的因素都需要考虑到

- 让每个Pod独立的
- 让每个Pod独立的，保持Pod启动顺序和唯一性
- 唯一的网络标识符，持久存储
- 有序，比如mysql中的主从

#### 部署

Mongodb

```yml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mongodb
spec:
  serviceName: mongodb
  replicas: 3
  selector:
    matchLabels:
      app: mongodb
  template:
    metadata:
      labels:
        app: mongodb
    spec:
      containers:
        - name: mongo
          image: mongo:4.4
          # IfNotPresent 仅本地没有镜像时才远程拉，Always 永远都是从远程拉，Never 永远只用本地镜像，本地没有则报错
          imagePullPolicy: IfNotPresent
---
apiVersion: v1
kind: Service
metadata:
  name: mongodb
spec:
  selector:
    app: mongodb
  type: ClusterIP
  # HeadLess
  clusterIP: None
  ports:
    - port: 27017
      targetPort: 27017
```

`kubectl apply -f mongo.yaml`

#### 特性

-   Service 的 `CLUSTER-IP` 是空的，Pod 名字也是固定的。
-   Pod 创建和销毁是有序的，创建是顺序的，销毁是逆序的。
-   Pod 重建不会改变名字，除了IP，所以不要用IP直连

访问时，如果直接使用 Service 名字连接，会随机转发请求  
要连接指定 Pod，可以这样`pod-name.service-name`  
运行一个临时 Pod 连接数据测试下  
`kubectl run mongodb-client --rm --tty -i --restart='Never' --image docker.io/bitnami/mongodb:4.4.10-debian-10-r20 --command -- bash`

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305262225466.png)

### DaemonSet

DaemonSet 即后台支撑型服务，主要是用来部署守护进程

长期伺服型和批处理型的核心在业务应用，可能有些节点运行多个同类业务的Pod，有些节点上又没有这类的Pod运行；而后台支撑型服务的核心关注点在K8S集群中的节点(物理机或虚拟机)，要保证每个节点上都有一个此类Pod运行。节点可能是所有集群节点，也可能是通过 nodeSelector选定的一些特定节点。典型的后台支撑型服务包括：存储、日志和监控等。在每个节点上支撑K8S集群运行的服务。

守护进程在我们每个节点上，运行的是同一个pod，新加入的节点也同样运行在同一个pod里面

### Job和CronJob

一次性任务 和 定时任务

- 一次性任务：一次性执行完就结束
- 定时任务：周期性执行

Job是K8S中用来控制批处理型任务的API对象。批处理业务与长期伺服业务的主要区别就是批处理业务的运行有头有尾，而长期伺服业务在用户不停止的情况下永远运行。Job管理的Pod根据用户的设置把任务成功完成就自动退出了。成功完成的标志根据不同的 spec.completions 策略而不同：单Pod型任务有一个Pod成功就标志完成；定数成功行任务保证有N个任务全部成功；工作队列性任务根据应用确定的全局成功而标志成功。

### Replication Controller

Replication Controller 简称 **RC**，是K8S中的复制控制器。RC是K8S集群中最早的保证Pod高可用的API对象。通过监控运行中的Pod来保证集群中运行指定数目的Pod副本。指定的数目可以是多个也可以是1个；少于指定数目，RC就会启动新的Pod副本；多于指定数目，RC就会杀死多余的Pod副本。

即使在指定数目为1的情况下，通过RC运行Pod也比直接运行Pod更明智，因为RC也可以发挥它高可用的能力，保证永远有一个Pod在运行。RC是K8S中较早期的技术概念，只适用于长期伺服型的业务类型，比如控制Pod提供高可用的Web服务。

#### Replica Set

Replica Set 检查 RS，也就是副本集。RS是新一代的RC，提供同样高可用能力，区别主要在于RS后来居上，能够支持更多种类的匹配模式。副本集对象一般不单独使用，而是作为Deployment的理想状态参数来使用
