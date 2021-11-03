package plugin

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Plugin struct {
	Name string
	Enabled bool
	Args []string
	Short string
}

type PluginManager struct {
	Plugins []Plugin
}

func NewManager() PluginManager {
	pluginManager := PluginManager{}
	data, err := os.ReadFile("plugins.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(data, &pluginManager)
	if err != nil {
		println(err.Error())
	}
	return pluginManager
}
func (p PluginManager) GetRunnablePlugins() []PluginRunner{
	runnablePlugins := make([]PluginRunner, 0)
	for _, plugin := range p.Plugins {
	   if plugin.Enabled {
		   currentPlugin := mapPluginsRun[plugin.Name]
		   currentPlugin.SetRunArgs(plugin.Args)
		   runnablePlugins = append(runnablePlugins, *currentPlugin)
	   }
	}
	return runnablePlugins
}