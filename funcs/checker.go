package funcs

import (
	"fmt"
	"github.com/imix-agent/toolkits/nux"
	"github.com/imix-agent/toolkits/sys"
)
//nux实际上是funcs下的一些采集工具
//cmdg.o

func CheckCollector() {
    //返回键值对
	output := make(map[string]bool)
    
    //从字面看，获取当前进程状态、磁盘、监听端口
	_, procStatErr := nux.CurrentProcStat()
	_, listDiskErr := nux.ListDiskStats()
	ports, listeningPortsErr := nux.ListeningPorts()
	procs, psErr := nux.AllProcs()
    fmt.Println(ports)
    fmt.Println(procs)

    //du错误的处理
	_, duErr := sys.CmdOut("du", "--help")

	output["kernel  "] = len(KernelMetrics()) > 0
	output["df.bytes"] = len(DeviceMetrics()) > 0
	output["net.if  "] = len(CoreNetMetrics([]string{})) > 0
	output["loadavg "] = len(LoadAvgMetrics()) > 0
	output["cpustat "] = procStatErr == nil
	output["disk.io "] = listDiskErr == nil
	output["memory  "] = len(MemMetrics()) > 0
	output["netstat "] = len(NetstatMetrics()) > 0
	output["ss -s   "] = len(SocketStatSummaryMetrics()) > 0
	output["ss -tln "] = listeningPortsErr == nil && len(ports) > 0
	output["ps aux  "] = psErr == nil && len(procs) > 0
	output["du -bs  "] = duErr == nil

	for k, v := range output {
		status := "fail"
		if v {
			status = "ok"
		}
		fmt.Println(k, "...", status)
	}
}
/* 
[root@localhost bin]# imix-agent -check
ps aux   ... ok
kernel   ... ok
df.bytes ... ok
disk.io  ... ok
memory   ... ok
ss -s    ... ok
du -bs   ... ok
net.if   ... ok
loadavg  ... ok
cpustat  ... ok
netstat  ... ok
ss -tln  ... ok
    
*/
