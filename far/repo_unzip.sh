#!/bin/bash

set -ex

pwd

function read_dir(){
    for file in `ls $1` #注意此处这是两个反引号，表示运行系统命令
    do
      if [ -d $1"/"$file ] #如果是目录
      then
        read_dir $1"/"$file
      else
        echo $1"/"$file
        cd $1
        # 如果是压缩文件就解压
        ls $file | grep *.tar.gz | xargs -n1 tar -xvf
        cd -
      fi
    done
}

read_dir $1
