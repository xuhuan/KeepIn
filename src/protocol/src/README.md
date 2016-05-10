#所有通讯协议
规定了所有约定的通讯协议

协议采用 protobuf v3 beta 2 ，统一包名为 protocol
生成最后存放的路径为
```
github.com/xuhuan/keepin/protocol
```
协议除了通用的基础协议存放在 Base 里面，其他按照服务器分

##Cluster
该服务器为集群服务器信息服务器，存放 (Login|Message|Route|Push|Db|File) 服务器的基本信息和负载信息
####1.0
请求协议


```
act:(get) //请求类型
data:
```

服务器类型 (Login|Message|Route|Push|Db|File)
服务器IP
服务器端口
