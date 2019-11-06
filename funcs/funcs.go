package funcs

import (
	"github.com/open-falcon/agent/g"
	"github.com/open-falcon/common/model"
)

// 间隔internal时间执行fs中的函数
//拆分不同的采集函数集，方便通过不同goroutine运行
type FuncsAndInterval struct {
	Fs       []func() []*model.MetricValue
	Interval int
}

/*
type MetricValue struct {
	Endpoint  string      `json:"endpoint"`
	Metric    string      `json:"metric"`
	Value     interface{} `json:"value"`
	Step      int64       `json:"step"`
	Type      string      `json:"counterType"`
	Tags      string      `json:"tags"`
	Timestamp int64       `json:"timestamp"`
}
*/



var Mappers []FuncsAndInterval

// 根据调用指令类型和是否容易被挂起而分类(通过不同的goroutine去执行,避免相互之间的影响)
func BuildMappers() {
    //设置传输数据的时间间隔
	interval := g.Config().Transfer.Interval
    //一上来就是一个切片
    //每个元素是 FuncsAndInterval struct 
	Mappers = []FuncsAndInterval{
		FuncsAndInterval{
			Fs: []func() []*model.MetricValue{
				AgentMetrics,
				CpuMetrics,
				NetMetrics,
				KernelMetrics,
				LoadAvgMetrics,
				MemMetrics,
				DiskIOMetrics,
				IOStatsMetrics,
				NetstatMetrics,
				ProcMetrics,
				UdpMetrics,
			},
			Interval: interval,
		},
        ///agent/funcs/agent.go

		// 容易出问题
		FuncsAndInterval{
			Fs: []func() []*model.MetricValue{
				DeviceMetrics,
			},
			Interval: interval,
		},
		// 调用相同指令
		FuncsAndInterval{
			Fs: []func() []*model.MetricValue{
				PortMetrics,
				SocketStatSummaryMetrics,
			},
			Interval: interval,
		},
		FuncsAndInterval{
			Fs: []func() []*model.MetricValue{
				DuMetrics,
			},
			Interval: interval,
		},
		FuncsAndInterval{
			Fs: []func() []*model.MetricValue{
				UrlMetrics,
			},
			Interval: interval,
		},
	}
}
