#登录服务器

用于用户登录的服务端。  

用户登录成功后返回 Message服务器信息，后续用户直接和Message服务器交互。  
登录成功的用户通过 Route 服务器来获取Message服务器信息。  

####服务流程
- 启动后从配置文件里读取 Cluster 服务器信息
- 向 Cluster 服务器注册自身
- 从 Cluster 服务器获取 Message, DB 服务器信息
- 启动定时任务，定时注册自身和获取 Message, DB 服务器信息
- 监听指定端口开始提供服务

##Todo  
- [ ] 启动后定时向 Cluster 服务器注册自身
- [ ] 启动后从 Cluster 服务器获取 Message, DB 服务器信息
- [ ] 定时向 Cluster 服务器注册自身
- [ ] 定时获取 Message, DB 服务器信息 
- [ ] 用户登录  