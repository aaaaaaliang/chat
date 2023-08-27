# chat
这是一个关于用户管理和登陆的 Go 项目，用于管理用户信息、用户登陆和退出操作、简单的通信功能
# 基本介绍
1.该项目基于gin框架，基于目前主流的前后端开发
# 使用说明
1.安装 swagger
go get -u github.com/swaggo/swag/cmd/swag
2.生成api文档
swag init
# 主要功能
创建用户，包括用户名、密码、邮箱、电话和身份信息
根据用户 ID 删除用户
更新用户信息，包括用户名、密码、邮箱、电话等
用户登录和退出，记录用户 IP 地址
websocket的简单实现
<img width="769" alt="websocket流程图" src="https://github.com/aaaaaaliang/chat/assets/117182742/ec7ed5b8-ecc2-4da7-8aed-4185dc546822">
