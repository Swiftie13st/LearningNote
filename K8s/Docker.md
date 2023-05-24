# Docker

**一句话概括容器：容器就是将软件打包成标准化单元，以用于开发、交付和部署。**

-   **容器镜像是轻量的、可执行的独立软件包** ，包含软件运行所需的所有内容：代码、运行时环境、系统工具、系统库和设置。
-   **容器化软件适用于基于 Linux 和 Windows 的应用，在任何环境中都能够始终如一地运行。**
-   **容器赋予了软件独立性**，使其免受外在环境差异（例如，开发和预演环境的差异）的影响，从而有助于减少团队间在相同基础设施上运行不同软件时的冲突

## docker 和虚拟机的区别

Docker 和虚拟机（Virtual Machine，VM）都是虚拟化技术的实现方式，但是它们有以下区别：

1.  架构不同：虚拟机在物理主机上运行一个完整的操作系统，每个虚拟机都有自己的内核、系统库、应用程序等，每个虚拟机之间互相独立。而 Docker 不需要运行整个操作系统，而是在主机操作系统上使用 Docker 引擎运行容器，所有容器共享同一个主机操作系统内核。
2.  资源占用不同：由于每个虚拟机都需要自己的操作系统和应用程序，所以虚拟机需要占用较多的计算资源、内存和存储空间。而 Docker 容器与宿主机共享资源，运行时只需占用少量的内存和磁盘空间。
3.  启动时间不同：启动虚拟机需要加载完整的操作系统，需要较长的时间，而 Docker 容器仅需要启动应用程序即可，启动时间很短。
4.  网络连接方式不同：虚拟机通常使用虚拟网络适配器与主机网络通信，而 Docker 容器可以通过主机网络或者 Docker 网络进行通信。
5. 与虚拟机相比，Docker隔离性更弱，Docker属于进程之间的隔离，虚拟机可实现系统级别隔离;Docker更轻量，Docker的架构可以共用一个内核与共享应用程序库，所占内存极小;同样的硬件环境，Docker运行的镜像数远多于虚拟机数量，对系统的利用率非常高;

| 特性   | 虚拟机                                                                       | Docker                                               |
| ------ | ---------------------------------------------------------------------------- | ---------------------------------------------------- |
| 跨平台 | 通常只能在桌面级系统运行，例如 Windows/Mac，无法在不带图形界面的服务器上运行 | 支持的系统非常多，各类 windows 和 Linux 都支持       |
| 性能   | 性能损耗大，内存占用高，因为是把整个完整系统都虚拟出来了                     | 性能好，只虚拟软件所需运行环境，最大化减少没用的配置 |
| 自动化 | 需要手动安装所有东西                                                         | 一个命令就可以自动部署好所需环境                     |
| 稳定性 | 稳定性不高，不同系统差异大                                                   | 稳定性好，不同系统都一样部署方式                     |
## 安装

[Install Docker Engine | Docker Documentation](https://docs.docker.com/engine/install/#server)
[Download Docker Desktop | Docker](https://www.docker.com/products/docker-desktop/)

### 镜像加速源

Docker 中国官方镜像 https://registry.docker-cn.com
DaoCloud 镜像站 http://f1361db2.m.daocloud.io
Azure 中国镜像 https://dockerhub.azk8s.cn
科大镜像站 https://docker.mirrors.ustc.edu.cn
阿里云 https://ud6340vz.mirror.aliyuncs.com
七牛云 https://reg-mirror.qiniu.com
网易云 https://hub-mirror.c.163.com
腾讯云 https://mirror.ccs.tencentyun.com

```
"registry-mirrors": ["https://registry.docker-cn.com"]
```

### 将当前用户添加到docker用户组

为了避免每次使用`docker`命令都需要加上`sudo`权限，可以将当前用户加入安装中自动创建的`docker`用户组(可以参考[官方文档](https://docs.docker.com/engine/install/linux-postinstall/))：
```
sudo usermod -aG docker $USER
```
```
newgrp docker
```

```
scp /var/lib/acwing/docker/images/docker_lesson_1_0.tar server_name:  # 将镜像上传到自己租的云端服务器
ssh server_name  # 登录自己的云端服务器

docker load -i docker_lesson_1_0.tar  # 将镜像加载到本地
docker run -p 20000:22 --name my_docker_server -itd docker_lesson:1.0  # 创建并运行docker_lesson:1.0镜像

docker attach my_docker_server  # 进入创建的docker容器
passwd  # 设置root密码
```

## 镜像（images）

镜像：镜像是一种轻量级、可执行的独立软件包，它包含运行某个软件所需的所有内容，我们把应用程序和配置依赖打包好形成一个可交付的运行环境(包括代码、运行时需要的库、环境变量和配置文件等)，这个打包好的运行环境就是image镜像文件。  

### 常用命令

`docker pull ubuntu:20.04`：拉取一个镜像
`docker images`：列出本地所有镜像
`docker image rm ubuntu:20.04` 或 `docker rmi ubuntu:20.04`：删除镜像ubuntu:20.04
`docker [container] commit CONTAINER IMAGE_NAME:TAG`：创建某个container的镜像
`docker save -o ubuntu_20_04.tar ubuntu:20.04`：将镜像ubuntu:20.04导出到本地文件ubuntu_20_04.tar中
`docker load -i ubuntu_20_04.tar`：将镜像ubuntu:20.04从本地文件ubuntu_20_04.tar中加载出来

### 制作自己的镜像

#### 编写Dockerfile

[Dockerfile reference | Docker Documentation](https://docs.docker.com/engine/reference/builder/#run)
```dockerfile
FROM node:11
MAINTAINER user

# 复制代码
ADD . /app

# 设置容器启动后的默认运行目录
WORKDIR /app

# 运行命令，安装依赖
# RUN 命令可以有多个，但是可以用 && 连接多个命令来减少层级。
# 例如 RUN npm install && cd /app && mkdir logs
RUN npm install --registry=https://registry.npm.taobao.org

# CMD 指令只能一个，是容器启动后执行的命令，算是程序的入口。
# 如果还需要运行其他命令可以用 && 连接，也可以写成一个shell脚本去执行。
# 例如 CMD cd /app && ./start.sh
CMD node app.js
```

#### 编译

[docker build | Docker Documentation](https://docs.docker.com/engine/reference/commandline/build/)

```sh
docker build -t test:v1 .
```

`-t` 设置镜像名字和版本号

#### 运行

[Docker run reference | Docker Documentation](https://docs.docker.com/engine/reference/run/)

```sh
docker run -p 8080:8080 --name test-hello test:v1
```

`-p` 映射容器内端口到宿主机  
`--name` 容器名字  
`-d` 后台运行  

### Dockerfile的基本指令有哪些？

FROM 指定基础镜像（必须为第一个指令，因为需要指定使用哪个基础镜像来构建镜像）；  
MAINTAINER 设置镜像作者相关信息，如作者名字，日期，邮件，联系方式等；  
COPY 复制文件到镜像；  
ADD 复制文件到镜像（ADD与COPY的区别在于，ADD会自动解压tar、zip、tgz、xz等归档文件，而COPY不会，同时ADD指令还可以接一个url下载文件地址，一般建议使用COPY复制文件即可，文件在宿主机上是什么样子复制到镜像里面就是什么样子这样比较好）；  
ENV 设置环境变量；  
EXPOSE 暴露容器进程的端口，仅仅是提示别人容器使用的哪个端口，没有过多作用；  
VOLUME 数据卷持久化，挂载一个目录；  
WORKDIR 设置工作目录，如果目录不在，则会自动创建目录；  
RUN 在容器中运行命令，RUN指令会创建新的镜像层，RUN指令经常被用于安装软件包；  
CMD 指定容器启动时默认运行哪些命令，如果有多个CMD，则只有最后一个生效，另外，CMD指令可以被docker run之后的参数替换；  
ENTRYOINT 指定容器启动时运行哪些命令，如果有多个ENTRYOINT，则只有最后一个生效，另外，如果Dockerfile中同时存在CMD和ENTRYOINT，那么CMD或docker run之后的参数将被当做参数传递给ENTRYOINT；

## 容器(container)

docker container，容器，一个系统级别的服务，拥有自己的ip和系统目录结构；运行容器前需要本地存在对应的镜像，如果本地不存在该镜像则就去镜像仓库下载。

### 常用命令

`docker [container] create -it ubuntu:20.04`：利用镜像ubuntu:20.04创建一个容器。
`docker ps -a`：查看本地的所有容器
`docker [container] start CONTAINER`：启动容器
`docker [container] stop CONTAINER`：停止容器
`docker [container] restart CONTAINER`：重启容器
`docker [contaienr] run -itd ubuntu:20.04`：创建并启动一个容器
`docker [container] attach CONTAINER`：进入容器
先按`Ctrl-p`，再按`Ctrl-q`可以挂起容器
`docker [container] exec CONTAINER COMMAND`：在容器中执行命令
`docker [container] rm CONTAINER`：删除容器
`docker container prune`：删除所有已停止的容器
`docker export -o xxx.tar CONTAINER`：将容器CONTAINER导出到本地文件xxx.tar中
`docker import xxx.tar image_name:tag`：将本地文件xxx.tar导入成镜像，并将镜像命名为image_name:tag
- `docker export/import`与`docker save/load`的区别：
	- `export/import`会丢弃历史记录和元数据信息，仅保存容器当时的快照状态
	- `save/load`会保存完整记录，体积更大

`docker top CONTAINER`：查看某个容器内的所有进程
`docker stats`：查看所有容器的统计信息，包括CPU、内存、存储、网络等信息
`docker cp xxx CONTAINER:xxx` 或 `docker cp CONTAINER:xxx xxx`：在本地和容器间复制文件
`docker rename CONTAINER1 CONTAINER2`：重命名容器
`docker update CONTAINER --memory 500MB`：修改容器限制

### 目录挂载

-   使用 Docker 运行后，我们改了项目代码不会立刻生效，需要重新`build`和`run`，很是麻烦。
-   容器里面产生的数据，例如 log 文件，数据库备份文件，容器删除后就丢失了。

挂载方式：

-   `bind mount` 直接把宿主机目录映射到容器内，适合挂代码目录和配置文件。可挂到多个容器上。方式用绝对路径：`-v D:/code:/app`
-   `volume` 由容器创建和管理，创建在宿主机，所以删除容器不会丢失，官方推荐，更高效，Linux 文件系统，适合存储数据库数据。可挂到多个容器上。方式只需名字：`-v db-data:/app`
-   `tmpfs mount` 适合存储临时文件，存宿主机内存中。不可多容器共享。

```sh
docker run -p 8080:8080 --name test-hello -v D:/code:/app -d test:v1
```

### 多容器通信

[docker network | Docker Documentation](https://docs.docker.com/engine/reference/commandline/network/)

项目往往都不是独立运行的，需要数据库、缓存这些东西配合运作。要想多容器之间互通，从 Web 容器访问 Redis 容器，我们只需要把他们放到同个网络中就可以了。

- 创建一个名为`test-net`的网络：

```sh
docker network create test-net
```

- 运行 Redis 在 `test-net` 网络中，别名`redis`

```sh
docker run -d --name redis --network test-net --network-alias redis redis:latest
```
 
- 修改代码中访问`redis`的地址为网络别名

- 运行 Web 项目，使用同个网络

```sh
docker run -p 8080:8080 --name test -v D:/test:/app --network test-net -d test:v1
```

## Docker-Compose

[Docker Compose overview | Docker Documentation](https://docs.docker.com/compose/)

Compose 是用于定义和运行多容器 Docker 应用程序的工具。通过 Compose，可以使用 YML 文件来配置应用程序需要的所有服务。然后，使用一个命令，就可以从 YML 文件配置中创建并启动所有服务。

Compose 使用的三个步骤：

-   使用 Dockerfile 定义应用程序的环境。
-   使用 docker-compose.yml 定义构成应用程序的服务，这样它们可以在隔离环境中一起运行。
-   最后，执行 docker-compose up 命令来启动并运行整个应用程序。

### 使用

要把项目依赖的多个服务集合到一起，我们需要编写一个`docker-compose.yml`文件，描述依赖哪些服务

```yml
version: "3.7"

services:
  app:
    build: ./
    ports:
      - 80:8080
    volumes:
      - ./:/app
    environment:
      - TZ=Asia/Shanghai
  redis:
    image: redis:5.0.13
    volumes:
      - redis:/data
    environment:
      - TZ=Asia/Shanghai

volumes:
  redis:

```

> 容器默认时间不是北京时间，增加 TZ=Asia/Shanghai 可以改为北京时间

在`docker-compose.yml` 文件所在目录，执行：`docker-compose up`就可以跑起来了。

在后台运行只需要加一个 -d 参数`docker-compose up -d`  
查看运行状态：`docker-compose ps`  
停止运行：`docker-compose stop`  
重启：`docker-compose restart`  
重启单个服务：`docker-compose restart service-name`  
进入容器命令行：`docker-compose exec service-name sh`  
查看容器运行log：`docker-compose logs [service-name]`
