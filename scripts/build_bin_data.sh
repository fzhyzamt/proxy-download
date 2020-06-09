#!/usr/bin/env bash

# 此文件用于重新打包资源文件，如果未修改则不需要执行

if command -v go-bindata > /dev/null 2>&1; then
    cd ../src
    go-bindata ../asset/
else
    echo 'You must install go-bindata, see: "https://github.com/jteeuwen/go-bindata"'
    exit 1
fi