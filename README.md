# Lite Yun 分布式版

## 简介
本版本是[lite-yun-golang](https://github.com/SBOrg666/lite-yun-golang)的功能升级版，可以同时记录并监控多台服务器的信息，需要搭配[lite-yun-RESTful](https://github.com/SBOrg666/lite-yun-RESTful)共同使用。

---
## 使用说明
本项目只提供了web前端，将来自多个lite-yun-RESTful的数据进行整合并显示，将添加的服务器信息进行存储。  
首先需要将RESTful的可执行程序部署在需要监控的服务器上，然后将本项目部署至想要充当web服务器的机器上面（也可以是本地）。

默认的登录帐号是admin@liteyun.com，密码是lite_yun_admin，如果需要修改，请使用sqlite客户端修改ACCOUNT.sqlite数据库文件。