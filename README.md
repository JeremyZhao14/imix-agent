imix-agent
===

## Installation
## 有待重新整理，到时直接打包好成为一个整体

## 安装golang


```bash
# set $GOPATH and $GOROOT
mkdir -p $GOPATH/src/github.com/imix-agent
cd $GOPATH/src/github.com/imix-agent
git clone https://github.com/JeremyZhao14/imix-agent.git

go get ./...
./control build
./control start

 
```
## Configuration

- heartbeat: heartbeat server rpc address
- transfer: transfer rpc address
- ignore: the metrics should ignore



