# 批量添加互信的小工具


## 程序逻辑
- 读取本机生成的密钥,如果没有, 则自动生成id_rsa, id_rsa.pub
- 使用密码链接远端服务器,同步本机id_rsa.pub到远端服务器,构建互信.如果已经构建互信, 则自动跳过此步骤


## 使用方法
```shell
Usage of ./create_trust:
  -s string
    	输入登录服务器名称,多个服务器使用,分割 (不允许为空)
  -u string
    	输入登录用户名 (default "inro")
```
