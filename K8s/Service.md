# Service

Deployment 只是保证了支撑服务的微服务Pod的数量，但是没有解决如何访问这些服务的问题。一个Pod只是一个运行服务的实例，随时可能在一个节点上停止，在另一个节点以一个新的IP启动一个新的Pod，因此不能以确定的IP和端口号提供服务。

要稳定地提供服务需要服务发现和负载均衡能力。服务发现完成的工作，是针对客户端访问的服务，找到对应的后端服务实例。在K8S集群中，客户端需要访问的服务就是Service对象。每个Service会对应一个集群内部有效的虚拟IP，集群内部通过虚拟IP访问一个服务。

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305262213879.png)

## 特性

-   Service 通过 label 关联对应的 Pod
-   Servcie 生命周期不跟 Pod 绑定，不会因为 Pod 重创改变 IP
-   提供了负载均衡功能，自动转发流量到不同 Pod
-   可对集群外部提供访问端口
-   集群内部可通过服务名字访问

## 作用

1. 防止Pod失联**服务发现**
   因为Pod每次创建都对应一个IP地址，而这个IP地址是短暂的，每次随着Pod的更新都会变化，假设当我们的前端页面有多个Pod时候，同时后端也多个Pod，这个时候，他们之间的相互访问，就需要通过注册中心，拿到Pod的IP地址，然后去访问对应的Pod
2. 定义Pod访问策略**负载均衡**
   页面前端的Pod访问到后端的Pod，中间会通过Service一层，而Service在这里还能做负载均衡，负载均衡的策略有很多种实现策略如：随机、轮询、响应比

## Service常用类型

Service常用类型有三种

- ClusterIp：集群内部访问
- NodePort：暴露端口到节点，提供了集群外部访问的入口，端口范围固定 30000 ~ 32767
- LoadBalancer：对外访问应用使用，公有云
- Headless：适合数据库  clusterIp 设置为 None 就变成 Headless 了，不会再分配 IP

## 使用

### 创建Service

创建 一个 Service，通过标签`test-k8s`跟对应的 Pod 关联上

service.yaml：

```yml
apiVersion: v1
kind: Service
metadata:
  name: test-k8s
spec:
  selector:
    app: test-k8s
  type: ClusterIP
  ports:
    - port: 8080        # 本 Service 的端口
      targetPort: 8080  # 容器端口
```

应用配置 `kubectl apply -f service.yaml`  
查看服务 `kubectl get svc`
查看服务详情 `kubectl describe svc test-k8s`，可以发现 Endpoints 是各个 Pod 的 IP，也就是他会把流量转发到这些节点。

服务的默认类型是`ClusterIP`，只能在集群内部访问，我们可以进入到 Pod 里面访问：  
`kubectl exec -it pod-name -- bash`  
`curl http://test-k8s:8080`

如果要在集群外部访问，可以通过端口转发实现（只适合临时测试用）：  
`kubectl port-forward service/test-k8s 8888:8080`

### 对外暴露服务

上面我们是通过端口转发的方式可以在外面访问到集群里的服务，如果想要直接把集群服务暴露出来，我们可以使用`NodePort` 和 `Loadbalancer` 类型的 Service

```yml
apiVersion: v1
kind: Service
metadata:
  name: test-k8s
spec:
  selector:
    app: test-k8s
  # 默认 ClusterIP 集群内可访问，NodePort 节点可访问，LoadBalancer 负载均衡模式（需要负载均衡器才可用）
  type: NodePort
  ports:
    - port: 8080        # 本 Service 的端口
      targetPort: 8080  # 容器端口
      nodePort: 31000   # 节点端口，范围固定 30000 ~ 32767
```

应用配置 `kubectl apply -f service.yaml`  
在节点上，我们可以 `curl http://localhost:31000/hello/easydoc` 访问到应用  
并且是有负载均衡的，网页的信息可以看到被转发到了不同的 Pod

`Loadbalancer` 也可以对外提供服务，这需要一个负载均衡器的支持，因为它需要生成一个新的 IP 对外服务，否则状态就一直是 pendding，这个很少用了，后面我们会讲更高端的 Ingress 来代替它。

### 多端口配置

对于某些服务，你需要公开多个端口。 Kubernetes 允许你在 Service 对象上配置多个端口定义。 为服务使用多个端口时，必须提供所有端口名称，以使它们无歧义。

```yml
apiVersion: v1
kind: Service
metadata:
  name: test-k8s
spec:
  selector:
    app: test-k8s
  type: NodePort
  ports:
    - port: 8080        # 本 Service 的端口
      name: test-k8s    # 必须配置
      targetPort: 8080  # 容器端口
      nodePort: 31000   # 节点端口，范围固定 30000 ~ 32767
    - port: 8090
      name: test-other
      targetPort: 8090
      nodePort: 32000

```
