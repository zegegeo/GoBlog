#!/bin/bash

# 切换到项目根目录
cd "$(dirname "$0")"

# 启动 etcd (如果没有在后台运行)
echo "Starting etcd..."
etcd &

# 启动 User RPC 服务
echo "Starting User RPC service..."
cd service/user
go run user.go &
cd ../../

# 启动 Article RPC 服务
echo "Starting Article RPC service..."
cd service/article
go run article.go &
cd ../../

# 启动 Comment RPC 服务
echo "Starting Comment RPC service..."
cd service/comment
go run comment.go &
cd ../../

# 启动 API Gateway
echo "Starting API Gateway..."
cd api
go run gateway.go &
cd ..

# 等待所有服务启动
echo "正在启动所有服务..."
sleep 3
echo "所有服务已启动"

# 保持脚本运行
wait