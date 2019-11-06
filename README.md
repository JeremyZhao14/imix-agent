imix-agent
===

## Installation
## 有待重新整理，到时直接打包好成为一个整体


## 安装golang、git 2
- 下载 git2
```bash
yum remove git
tar zxvf v2.2.1.tar.gz
cd git-2.2.1
make configure
./configure --prefix=/usr/local/git --with-iconv=/usr/local/libiconv
make all doc
make install install-doc install-html
echo "export PATH=$PATH:/usr/local/git/bin" >> /etc/bashrc
source /etc/bashrc
''' 
 
- 下载复制 golang
```bash
wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz --no-check-certificate
tar zxvf go1.8.3.linux-amd64.tar.gz
mv go /usr/local/
vim /etc/profile
'''
export PATH=$PATH:/usr/local/go/bin
export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export GOPATH= /root/go/ 
- 生效环境变量
 source /etc/profile



```bash
# set $GOPATH and $GOROOT
mkdir -p $GOPATH/src/github.com/imix-agent
cd $GOPATH/src/github.com/imix-agent
git clone https://JeremyZhao14@github.com/JeremyZhao14/imix-agent.git

go get ./...
./control build
./control start

 
```
## Configuration

- heartbeat: heartbeat server rpc address
- transfer: transfer rpc address
- ignore: the metrics should ignore



