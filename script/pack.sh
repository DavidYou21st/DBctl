#!/bin/bash

function SysTem() {
a=`uname  -a`

if [[ $a =~ "Darwin" ]];then
    os_name="darwin"
elif [[ $a =~ "MINGW" || $a =~ "MSYS" ]];then
    os_name="windows"
else
    os_name="linux"
fi
}

SysTem
echo "当前操作系统:"${os_name}

export GO111MODULE=on
export GOPROXY=https://goproxy.cn

#SetEnv
bash_dir=`pwd`
release_dir=./build/release

if [[ $a =~ "Darwin" ]];then
    upx=upx
  else
    upx=${bash_dir}/scripts/upx/${os_name}/upx
fi
echo ${upx}

DATE=`date +%Y-%m-%d-%H%M%S`
echo $DATE

rm -rf ${release_dir}
mkdir -p ${release_dir}

echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>进入项目目录<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
cd  ${bash_dir}

echo "开始编译···"
GOOS=linux GOARCH=amd64 go build -o ${release_dir}/linux-$DATE -ldflags "-s -w"
GOOS=windows GOARCH=amd64 go build -o ${release_dir}/windows-$DATE -ldflags "-s -w"
GOOS=darwin GOARCH=amd64 go build -o ${release_dir}/darwin-$DATE -ldflags "-s -w"

echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>复制项目文件<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"

cd  ${bash_dir}/${release_dir}


#压缩优化
#工具安装：从 github.com/upx/upx 下载upx安装文件
echo "压制执行文件···"
cd  ${bash_dir}/${release_dir}
${upx} --brute linux-${DATE}
${upx} --brute windows-${DATE}
${upx} --brute darwin-${DATE}
echo "压制完成"

#打包文件
echo "创建压缩包"
cd  ${bash_dir}/${release_dir}
tar -cf ./mctl.$DATE.tar.gz *

cd  ${bash_dir}