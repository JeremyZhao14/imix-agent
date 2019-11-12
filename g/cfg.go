package g

import (
	"encoding/json"
	"log"
	"os"
	"sync"
//	"github.com/imix-agent/godump"
	"github.com/imix-agent/toolkits/file"
)

type PluginConfig struct {
	Enabled bool   `json:"enabled"`
	Dir     string `json:"dir"`
	Git     string `json:"git"`
	LogDir  string `json:"logs"`
}

type HeartbeatConfig struct {
	Enabled  bool   `json:"enabled"`
	Addr     string `json:"addr"`
	Interval int    `json:"interval"`
	Timeout  int    `json:"timeout"`
}

type TransferConfig struct {
	Enabled  bool     `json:"enabled"`
	Addrs    []string `json:"addrs"`
	Interval int      `json:"interval"`
	Timeout  int      `json:"timeout"`
}

type HttpConfig struct {
	Enabled  bool   `json:"enabled"`
	Listen   string `json:"listen"`
	Backdoor bool   `json:"backdoor"`
}

type CollectorConfig struct {
	IfacePrefix []string `json:"ifacePrefix"`
}

// hostname 和IP要agent从机器获取
type GlobalConfig struct {
	Debug         bool             `json:"debug"`
	Hostname      string           `json:"hostname"`
	IP            string           `json:"ip"`
	Plugin        *PluginConfig    `json:"plugin"`
	Heartbeat     *HeartbeatConfig `json:"heartbeat"`
	Transfer      *TransferConfig  `json:"transfer"`
	Http          *HttpConfig      `json:"http"`
	Collector     *CollectorConfig `json:"collector"`
	IgnoreMetrics map[string]bool  `json:"ignore"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func Hostname() (string, error) {
	hostname := Config().Hostname
	if hostname != "" {
		return hostname, nil
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Println("ERROR: os.Hostname() fail", err)
	}
	return hostname, err
}

func IP() string {
	ip := Config().IP
	if ip != "" {
		// use ip in configuration
		return ip
	}

	if len(LocalIps) > 0 {
		ip = LocalIps[0]
	}

	return ip
}


//解析配置文件
func ParseConfig(cfg string) {
    //检查文件名是否为空
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}


	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg
    //获取配置文件中内容并删除首尾多余的空白字符
	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

    //GlobalConfig 定义在上方,全局配置，实际上就是cfg.json的内容
	var c GlobalConfig
    //反序列化，json转为数列
    //func Unmarshal(data []byte, v interface{}) error 注意参数的类型
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c
//    godump.Dump(config)

	log.Println("read config file:", cfg, "successfully")
}
