# chat
这是一个关于用户管理和登陆的 Go 项目，用于管理用户信息、用户登陆和退出操作、简单的通信功能
# 基本介绍
1.该项目基于gin框架，基于目前主流的前后端开发
# 使用说明
使用 GoLand 等编辑工具，进入项目根目录，修改config/app.yaml 中 mysql 相关配置。
## 克隆项目
git clone https://github.com/aaaaaaliang/golang.git
##  使用 go mod 并安装go依赖包
go mod tidy
## 运行项目 
### 1.安装 swagger
go get -u github.com/swaggo/swag/cmd/swag  
### 2.生成api文档  
swag init
### 3.运行testGorm.go里文件生成数据库表  
go run testGorm.go
### 4.go run main.go  

## 实现聊天
启动完项目 浏览器访问http://localhost:8080  
做完注册 登陆操作以后 找到添加好友 即可开始聊天

# 主要功能
创建用户，包括用户名、密码、邮箱、电话和身份信息
根据用户 ID 删除用户
更新用户信息，包括用户名、密码、邮箱、电话等
用户登录和退出，记录用户 IP 地址
添加好友 上传文件
websocket的简单实现
# websocket流程分析
<img width="769" alt="websocket流程图" src="https://github.com/aaaaaaliang/chat/assets/117182742/ec7ed5b8-ecc2-4da7-8aed-4185dc546822">
