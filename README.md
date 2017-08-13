# mywebcalculator

## 说明

mywebcalculator是一个基于beego框架构建的小程序，实现的功能为在网页上进行四则运算。该程序可以通过Docker镜像的方式运行。

## 安装

```shell
#将代码clone到$GOPATH/src下
cd $GOPATH/src
git clone https://github.com/huweihuang/mywebcalculator
#编译代码
cd mywebcalculator
go build
#执行构建镜像脚本
sh build.sh
#查看镜像
docker images
#运行容器
docker run -d -p 8080:8080 mywebcalculator:1.0.0
```

## 使用

打开浏览器，输入运行机器IP加端口，例如：10.8.216.25:8080

