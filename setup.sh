#!/usr/bin/env bash

set -x
#sudo apt update
#sudo apt upgrade -y
#自定义别名
echo "alias cls='clear'" >> ~/.bashrc
echo "alias apt='sudo apt'" >> ~/.bashrc
echo "alias apti='apt install -y'" >> ~/.bashrc
echo "alias npm='sudo npm'" >> ~/.bashrc
echo "alias cnpm='sudo cnpm'" >> ~/.bashrc
echo "alias yarn='sudo yarn'" >> ~/.bashrc
echo "alias pip='sudo pip'" >> ~/.bashrc
echo "alias l='ls -al'" >> ~/.bashrc
echo "alias sc=systemctl" >> ~/.bashrc
echo "alias jc=journalctl" >> ~/.bashrc
echo "alias kc=kubectl" >> ~/.bashrc
echo "alias kcp='kubectl get pod -A -owide'" >> ~/.bashrc
echo "alias kcd='kubectl get deploy -A'" >> ~/.bashrc
echo "alias kcn='kubectl get node --show-labels'" >> ~/.bashrc
echo "alias kcm='kubectl get cm -A -owide'" >> ~/.bashrc
source ~/.bashrc
#sudo apt install -y unity-tweak-tool
#输入法
sudo apt install -y fcitx
sudo apt install -y ibus
#常用开发工具
sudo apt install -y chromium-browser
sudo apt install -y wget curl git vim tree
sudo apt install -y build-essential cpp gcc g++ gdb cmake automake
sudo apt install -y default-jdk
#sudo apt install -y manpages manpages-dev manpages-posix manpages-posix-dev manpages-zh

##golang
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install golang-go

#nodejs相关
sudo apt install -y nodejs npm
sudo npm config set registry https://registry.npm.taobao.org
sudo npm i -g npm
sudo npm i -g cnpm --registry=https://registry.npm.taobao.org
sudo npm i -g yarn
yarn config set registry https://registry.npm.taobao.org
#python3相关
sudo apt install -y python3 python3-pip python3-setuptools 
sudo mkdir ~/.pip
sudo touch ~/.pip/pip.conf
sudo echo -e "[global]\nindex-url = https://pypi.tuna.tsinghua.edu.cn/simple" > ~/.pip/pip.conf

sudo pip3 install --upgrade pip
sudo pip3 install cheat tldr
sudo pip3 install pipenv
#sudo pip3 install numpy scipy matplotlib pandas scikit-learn
#重新安装pip
#https://www.imooc.com/article/31953?block_id=tuijian_wz
#sudo python3 -m pip uninstall pip && sudo apt install python3-pip --reinstall
#vscode
curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.gpg
sudo mv microsoft.gpg /etc/apt/trusted.gpg.d/microsoft.gpg
sudo sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main" > /etc/apt/sources.list.d/vscode.list'
sudo apt-get update
sudo apt-get install code
#zsh & oh my zsh
sudo apt install -y zsh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
#自定义别名
echo "alias cls='clear'" >> ~/.zshrc
echo "alias apt='sudo apt'" >> ~/.zshrc
echo "alias apti='apt install -y'" >> ~/.zshrc
echo "alias npm='sudo npm'" >> ~/.zshrc
echo "alias cnpm='sudo cnpm'" >> ~/.zshrc
echo "alias yarn='sudo yarn'" >> ~/.zshrc
echo "alias pip='sudo pip'" >> ~/.zshrc
echo "alias l='ls -al'" >> ~/.zshrc
echo "alias sc=systemctl" >> ~/.zshrc
echo "alias jc=journalctl" >> ~/.zshrc
echo "alias kc=kubectl" >> ~/.zshrc
echo "alias kcp='kubectl get pod -A -owide'" >> ~/.zshrc
echo "alias kcd='kubectl get deploy -A'" >> ~/.zshrc
echo "alias kcn='kubectl get node --show-labels'" >> ~/.zshrc
echo "alias kcm='kubectl get cm -A -owide'" >> ~/.zshrc
source ~/.zshrc
#超好用的vim配置
wget -qO- https://raw.github.com/ma6174/vim/master/setup.sh | sh -x
#手动安装deb包
#sudo apt-get install -f
#sudo dpkg -i deb文件名
#升级所有过期库
pip list --outdated | awk '{print "pip install --upgrade "$1}' > python-upgrade.sh
sh python-upgrade.sh

// docker 安装
sudo apt install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io

