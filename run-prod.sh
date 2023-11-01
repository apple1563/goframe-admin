#!/bin/bash

# Go 可执行文件名
name="goframe-admin"
# 日志文件名
logFile="output_$(date '+%Y%m%d_%H%M%S').log"
configFile="config.prod.yaml"

# 执行发布流程
#echo "Fetching latest code..."
#git fetch --all
#git reset --hard origin/master

# 停止正在运行的程序
echo "停止正在运行的程序..."
pkill -f $name
sleep 2  # 等待程序停止

# 启动 Go 程序并将输出重定向到日志文件
echo "启动 Go 程序..."
nohup ./$name --gf.gcfg.file=$configFile > $logFile 2>&1 &

# 打印进程 ID
echo "Go 程序已在后台运行，进程 ID：$!"
