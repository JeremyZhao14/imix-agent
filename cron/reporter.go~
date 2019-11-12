package cron

import (
	"fmt"
	"github.com/imix-agent/g"
	"github.com/imix-agent/common/model"
	"log"
	"time"
)

func ReportAgentStatus() {
	// 判断hbs配置是否正常，正常则上报agent状态
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		// 根据配置的interval间隔上报信息
        //开启新的线程
		go reportAgentStatus(time.Duration(g.Config().Heartbeat.Interval) * time.Second)
	}
}

func reportAgentStatus(interval time.Duration) {
	for {
		// 获取hostname, 出错则错误赋值给hostname
		hostname, err := g.Hostname()
		if err != nil {
			hostname = fmt.Sprintf("error:%s", err.Error())
		}
		// 请求发送信息
		req := model.AgentReportRequest{
			Hostname:      hostname,
			IP:            g.IP(),
			AgentVersion:  g.VERSION,
			PluginVersion: g.GetCurrPluginVersion(),
		}

    //response,int code ,json
		var resp model.SimpleRpcResponse
		// 调用rpc接口
        //在var.go中定义HbsClient
		err = g.HbsClient.Call("Agent.ReportStatus", req, &resp)
		if err != nil || resp.Code != 0 {
			log.Println("call Agent.ReportStatus fail:", err, "Request:", req, "Response:", resp)
		}

		time.Sleep(interval)
	}
}
