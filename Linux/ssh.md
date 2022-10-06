## SSH登录
### 基本用法
#### 远程登录服务器：
```
ssh user@hostname
user: 用户名
hostname: IP地址或域名
```
第一次登录时会提示：
```
The authenticity of host '123.57.47.211 (123.57.47.211)' can't be established.
ECDSA key fingerprint is SHA256:iy237yysfCe013/l+kpDGfEG9xxHxm0dnxnAbJTPpG8.
Are you sure you want to continue connecting (yes/no/[fingerprint])?
```
输入yes，然后回车即可。
这样会将该服务器的信息记录在`~/.ssh/known_hosts`文件中。

然后输入密码即可登录到远程服务器中。

默认登录端口号为22。如果想登录某一特定端口：
```
ssh user@hostname -p 22
```
#### 配置文件
创建文件 `~/.ssh/config`。

然后在文件中输入：
```
Host myserver1
    HostName IP地址或域名
    User 用户名

Host myserver2
    HostName IP地址或域名
    User 用户名
```
之后再使用服务器时，可以直接使用别名myserver1、myserver2。

#### 密钥登录
创建密钥：

`ssh-keygen`
然后一直回车即可。

执行结束后，~/.ssh/目录下会多两个文件：

`id_rsa`：私钥
`id_rsa.pub`：公钥
之后想免密码登录哪个服务器，就将公钥传给哪个服务器即可。

例如，想免密登录myserver服务器。则将公钥中的内容，复制到myserver中的`~/.ssh/authorized_keys`文件里即可。

也可以使用如下命令一键添加公钥：
```
ssh-copy-id myserver
```
#### 执行命令
命令格式：
```
ssh user@hostname command
```
例如：
```
ssh user@hostname ls -a
```
或者
```
# 单引号中的$i可以求值
ssh myserver 'for ((i = 0; i < 10; i ++ )) do echo $i; done'
```
或者
```
# 双引号中的$i不可以求值
ssh myserver "for ((i = 0; i < 10; i ++ )) do echo $i; done"
```
## SCP传文件
### 基本用法
命令格式：
```
scp source destination
```
将source路径下的文件复制到destination中

一次复制多个文件：
```
scp source1 source2 destination
```
复制文件夹：
```
scp -r ~/tmp myserver:/home/acs/
```
将本地家目录中的tmp文件夹复制到myserver服务器中的/home/acs/目录下。
```
scp -r ~/tmp myserver:homework/
```
将本地家目录中的tmp文件夹复制到myserver服务器中的~/homework/目录下。
```
scp -r myserver:homework .
```
将myserver服务器中的~/homework/文件夹复制到本地的当前路径下。

指定服务器的端口号：
```
scp -P 22 source1 source2 destination
```
注意： scp的-r -P等参数尽量加在source和destination之前。

使用scp配置其他服务器的vim和tmux
```
scp ~/.vimrc ~/.tmux.conf myserver:
```
