<h1 align="center">IrisAdminApi</h1>

<div align="center">
    <a href="https://travis-ci.org/snowlyg/IrisAdminApi"><img src="https://travis-ci.org/snowlyg/IrisAdminApi.svg?branch=blog" alt="Build Status"></a>
    <a href="https://codecov.io/gh/snowlyg/IrisAdminApi"><img src="https://codecov.io/gh/snowlyg/IrisAdminApi/branch/master/graph/badge.svg" alt="Code Coverage"></a>
    <a href="https://goreportcard.com/report/github.com/snowlyg/IrisAdminApi"><img src="https://goreportcard.com/badge/github.com/snowlyg/IrisAdminApi" alt="Go Report Card"></a>
    <a href="https://godoc.org/github.com/snowlyg/IrisAdminApi"><img src="https://godoc.org/github.com/snowlyg/IrisAdminApi?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/snowlyg/IrisAdminApi/blob/master/LICENSE"><img src="https://img.shields.io/github/license/snowlyg/IrisAdminApi" alt="Licenses"></a>
    <h5 align="center">IrisAdminApi</h5>
</div>

> 简单项目仅供学习，欢迎指点！
>
#### 演示地址
主分支：
[http://irisadminapi.snowlyg.com](http://irisadminapi.snowlyg.com)

blog 分支:
[http://www.snowlyg.com](http://www.snowlyg.com) 

#### IRIS V12 中文文档
[IRIS V12 中文文档](https://www.snowlyg.com/chapter/1)


账号/密码 ： username/123456

###### `Iris-go` 学习交流 QQ 群 ：`676717248`
<a target="_blank" href="//shang.qq.com/wpa/qunwpa?idkey=cc99ccf86be594e790eacc91193789746af7df4a88e84fe949e61e5c6d63537c"><img border="0" src="http://pub.idqqimg.com/wpa/images/group.png" alt="Iris-go" title="Iris-go"></a>


#### 项目介绍
- 数据支持 `mysql`，`sqlite3` 配置; `sqlite3` 需要下载 `gcc`。  [gcc 下载地址](http://mingw-w64.org/doku.php/download)
- 使用了 [https://github.com/snowlyg/gotransformer](https://github.com/snowlyg/gotransformer) 转换数据，返回数据格式化，excel 导入数据转换，xml 文件生产数据转换等 
- 使用 `casbin` 做权限控制, `server/config/rbac_model.conf` 为相关配置。系统会根据路由名称生成对应路由权限，并配置到管理员角色。

 
---

#### 项目开发过程详解

1.[Iris-go 项目登陆 API 构建细节实现过程](https://www.snowlyg.com/#/detail/1)

2.[iris + casbin 从陌生到学会使用的过程](https://www.snowlyg.com/#/detail/2)

---

#### 更新日志

[UPDATE](UPDATE.MD)
---

#### 问题总结

[ERRORS](ERRORS.MD)


1. 运行项目
- 从地址 [https://github.com/snowlyg/IrisAdminApi/releases/tag/v0.3](https://github.com/snowlyg/IrisAdminApi/releases/tag/v0.3) ，选择下载编译好的系统文件,

- 运行项目 

```shell script
# 安装
./main install
# 卸载
./main uninstall
# 启动
./main start
# 停止
./main stop
# 查看版本
./main version
# 清除数据，并重新填充
./main seeder
# 修改 seeder/data/perms.yaml 后同步权限,
./main sync_perms
# 查看权限信息
./main perms
```


#### 如果需要重新编译项目

- 拉取项目

```shell script
git clone https://github.com/snowlyg/IrisAdminApi.git

# github 克隆太慢可以用 gitee 地址:

git clone https://gitee.com/snowlyg/IrisAdminApi.git

```

- 安装项目依赖

>加载依赖管理包 (解决国内下载依赖太慢问题)
>使用国内七牛云的 go module 镜像。
>
>参考 https://github.com/goproxy/goproxy.cn。
>
>阿里： https://mirrors.aliyun.com/goproxy/
>
>官方： https://goproxy.io/
>
>中国：https://goproxy.cn
>
>其他：https://gocenter.io

##### golang 1.13+ 可以直接执行：
```shell script
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

> 修改配置文件 `application.yml` ,配置文件需要放置在运行程序的同级目录

- 运行项目,
>推荐使用 air 或者 gowatch 等热编译工具,直接使用 `go run main.go `  方法运行，可能会出现配置文件无法加载的问题
>如果想使用 `go run main.go` 命令运行, 设置环境变量 `TRAVIS_BUILD_DIR=~/go/src/github.com/snowlyg/IrisAdminApi`

```shell script

# 安装工具 air     
go get -u github.com/cosmtrek/air

# 在 server 目录执行,可以通过 .air.conf 配置 air 工具
air
```

---
##### 单元测试 
> http test

```shell script
# 设置环境变量 TRAVIS_BUILD_DIR=~/go/src/github.com/snowlyg/IrisAdminApi
# 否则无法加载配置文件

# 所有测试
 go test -v ./... -tags test
 
#单个方法
 go test -run TestUserCreate -v  -tags test

# 测试
 gotest  
 
```
---

##### 接口文档
自动生成文档 (访问过的接口就会自动成功)
>访问文档，从浏览器直接打开 `http://localhost:8085/apiDoc`
---

#### 登录项目
- http://localhost:80

#### 感谢 

[JetBrains](https://www.jetbrains.com/?from=IrisAdminApi) 对本项目的支持。

