package g

import (
	"github.com/toolkits/net"
	"log"
	"math"
	"net/rpc"
	"sync"
	"time"
)

// 简单封装rpc.Cilent
type SingleConnRpcClient struct {
	sync.Mutex
	rpcClient *rpc.Client
	RpcServer string
	Timeout   time.Duration
}

// 关闭rpc
func (this *SingleConnRpcClient) close() {
	if this.rpcClient != nil {
		this.rpcClient.Close()
		this.rpcClient = nil
	}
}

// 保证rpc存在,为空则重新创建, 如果server宕机, 死循环????
func (this *SingleConnRpcClient) insureConn() {
	if this.rpcClient != nil {
		return
	}

	var err error
	var retry int = 1

	for {
		if this.rpcClient != nil {
			return
		}
		// 根据timeout和server地址去连接rpc的server
		this.rpcClient, err = net.JsonRpcClient("tcp", this.RpcServer, this.Timeout)
		if err == nil {
			return
		}

		log.Printf("dial %s fail: %v", this.RpcServer, err)

		if retry > 6 {
			retry = 1
		}

		time.Sleep(time.Duration(math.Pow(2.0, float64(retry))) * time.Second)

		retry++
	}
}

// rpc client调用hbs函数
func (this *SingleConnRpcClient) Call(method string, args interface{}, reply interface{}) error {
	// 加锁保证一个agent只与server有一个连接,保证性能
	this.Lock()
	defer this.Unlock()
    // 保证rpc连接可用
	this.insureConn()

	timeout := time.Duration(50 * time.Second)
	done := make(chan error)

	go func() {
		err := this.rpcClient.Call(method, args, reply)
		done <- err
	}()
	// 超时控制
	select {
	case <-time.After(timeout):
		log.Printf("[WARN] rpc call timeout %v => %v", this.rpcClient, this.RpcServer)
		this.close()
	case err := <-done:
		if err != nil {
			this.close()
			return err
		}
	}

	return nil
}
