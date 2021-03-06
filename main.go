package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/imix-agent/cron"
	"github.com/imix-agent/funcs"
	"github.com/imix-agent/g"
	"github.com/imix-agent/http"
)
//flag包：用于命令行参数解析
//fmt：格式化
//os: 系统命令

//cron:定时任务
//funcs: 信息采集包
//g:全局结构与变量
//http:agent简易网页


func main() {
    //参数、默认、说明文字
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	flag.Parse()

    //如果version的内容为true
    //在const.go
	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

    //如果check的内容为true,检查collector
    //checker.go 查看一些命令是否有返回值    
	if *check {
		funcs.CheckCollector()
		os.Exit(0)
	}

    //解析配置文件（根据路径）
    //在cfg.go
	g.ParseConfig(*cfg)

//初始化
    //以当前目录作为根目录
	g.InitRootDir()
    //获取内网地址
	g.InitLocalIps()
    //获得heartbeat配置
	g.InitRpcClients()


    //获取采集数据的列表
	funcs.BuildMappers()
// 初始化历史数据,只有cpu和disk需要历史数据,用于获取当前时刻的参数
	go cron.InitDataHistory()

	// 上报本机agent状态
	cron.ReportAgentStatus()
	// 同步插件
	cron.SyncMinePlugins()
	// 同步监控端口、路径、进程和URL，同步内置metric,包括端口、目录和进程信息
	cron.SyncBuiltinMetrics()
	// 后门调试agent,允许执行shell指令的ip列表
    //同步可信IP列表
    //请求获取远程访问执行shell命令的IP白名单，在通过http/run.go调用shell命令是会判断请求IP是否可信
	cron.SyncTrustableIps()
	// 开始数据次采集
	cron.Collect()
	// 启动dashboard server
	go http.Start()



//很多时候我们需要让main函数不退出，让它在后台一直执行，
	select {}

}
