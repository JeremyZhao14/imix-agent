package funcs

import (
	"github.com/imix-agent/common/model"
	"github.com/imix-agent/toolkits/nux"
	"log"
)

// 执行shell指令采集socket信息
func SocketStatSummaryMetrics() (L []*model.MetricValue) {
	ssMap, err := nux.SocketStatSummary()
	if err != nil {
		log.Println(err)
		return
	}

	for k, v := range ssMap {
		L = append(L, GaugeValue("ss."+k, v))
	}

	return
}
