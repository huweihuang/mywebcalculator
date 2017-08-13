#!/bin/sh

# 镜像名称
APP_NAME="mywebcalculator"
VERSION="1.0.0"

# 镜像版本号
IMAGE_NAME="${APP_NAME}:${VERSION}"

echo "准备创建镜像  $IMAGE_NAME"

# 删除原来镜像
docker rmi -f "$IMAGE_NAME"
# 编译新的镜像
docker build -t "$IMAGE_NAME" -f docker/Dockerfile ./
