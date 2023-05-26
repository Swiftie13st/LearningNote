# Ingress

[Ingress | Kubernetes](https://kubernetes.io/zh-cn/docs/concepts/services-networking/ingress/)

原来我们需要将端口号对外暴露，通过 ip + 端口号就可以进行访问

原来是使用Service中的NodePort来实现

- 在每个节点上都会启动端口
- 在访问的时候通过任何节点，通过ip + 端口号就能实现访问

但是NodePort还存在一些缺陷

- 因为端口不能重复，所以每个端口只能使用一次，一个端口对应一个应用
- 实际访问中都是用域名，根据不同域名跳转到不同端口服务中

ngress 为外部访问集群提供了一个 **统一** 入口，避免了对外暴露集群端口；  
功能类似 Nginx，可以根据域名、路径把请求转发到不同的 Service。  
可以配置 https

## 跟 LoadBalancer 有什么区别？

LoadBalancer 需要对外暴露端口，不安全；  
无法根据域名、路径转发流量到不同 Service，多个 Service 则需要开多个 LoadBalancer；  
功能单一，无法配置 https

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305262246313.png)

## Ingress和Pod关系

pod 和 ingress 是通过service进行关联的，而ingress作为统一入口，由service关联一组pod中

- 首先service就是关联我们的pod
- 然后ingress作为入口，首先需要到service，然后发现一组pod
- 发现pod后，就可以做负载均衡等操作

**工作流程**：

![](https://raw.githubusercontent.com/Swiftie13st/Figurebed/main/img/202305262248729.png)
