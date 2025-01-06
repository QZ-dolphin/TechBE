#!/bin/bash

# 获取当前日期和时间，并格式化为 YYYY-MM-DD_HH-MM-SS
datetime=$(date +"%Y-%m-%d_%H-%M-%S")

# 定义变量a为 "revise" 加上当前日期时间
a="revise_$datetime"

# 添加所有更改到git
git add .
if [ $? -ne 0 ]; then
    echo "Failed to add changes. Exiting."
    exit 1
fi

# 使用变量a作为提交信息
git commit -m "$a"
if [ $? -ne 0 ]; then
    echo "Failed to commit changes. Exiting."
    exit 1
fi

# 推送更改到master分支
git push origin master
if [ $? -ne 0 ]; then
    echo "Failed to push changes. Exiting."
    exit 1
fi

echo "All git commands executed successfully."
