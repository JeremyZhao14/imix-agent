package plugins

// 插件信息: 路径、修改时间、运行周期(来自plugin插件)
type Plugin struct {
	FilePath string
	MTime    int64
	Cycle    int
}

// 插件map和调度器map
var (
	Plugins              = make(map[string]*Plugin)
	PluginsWithScheduler = make(map[string]*PluginScheduler)
)

// 删除不需要的plugin
func DelNoUsePlugins(newPlugins map[string]*Plugin) {
	for currKey, currPlugin := range Plugins {
		newPlugin, ok := newPlugins[currKey]
		if !ok || currPlugin.MTime != newPlugin.MTime {
			deletePlugin(currKey)
		}
	}
}

// 添加同步时增加的plugin
func AddNewPlugins(newPlugins map[string]*Plugin) {
	for fpath, newPlugin := range newPlugins {
		// 去除重复插件
		if _, ok := Plugins[fpath]; ok && newPlugin.MTime == Plugins[fpath].MTime {
			continue
		}
		// 为新添加的插件新建调度器
		Plugins[fpath] = newPlugin
		sch := NewPluginScheduler(newPlugin)
		PluginsWithScheduler[fpath] = sch
		// 启动plugin调度
		sch.Schedule()
	}
}

func ClearAllPlugins() {
	for k := range Plugins {
		deletePlugin(k)
	}
}

func deletePlugin(key string) {
	v, ok := PluginsWithScheduler[key]
	if ok {
		// 暂停调度plugin
		v.Stop()
		delete(PluginsWithScheduler, key)
	}
	delete(Plugins, key)
}
