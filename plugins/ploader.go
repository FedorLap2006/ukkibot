package plugins

import ( 
	"core"
)

type BotPlugin struct {
	Name string
	Commands map[string]core.Command
	Events map[string]core.EventHandler
	RuntimeThreads map[string]core.PluginHandler
	Enabled bool
}

type BotPluginCommandPayload struct {
	*BotPlugin
}

type BotPluginConfigPayload struct {
	Plugins []*BotPlugin
}

// var Commands map[string]core.Command = make(map[string]core.Command)
// var Events map[string]core.EventHandler = make(map[string]core.EventHandler)

var _plugins []*BotPlugin
var _rplugins []core.Plugin

func RegPlugin(bp *BotPlugin) {
	_plugins = append(_plugins, bp)
}


func PluginsLoader(c *core.Core) {
	c.Cfg.Payload = BotPluginConfigPayload{ Plugins: _plugins }
	for _, plug := range _plugins {
		if(!plug.Enabled) {
			continue
		}
		for name, inst := range plug.Commands {
			inst.Payload = BotPluginCommandPayload{ plug }
			c.Cfg.Commands[name] = inst
		}
		for name, inst := range plug.Events {
			c.Cfg.Events[name] = inst
		}
		for name, inst := range plug.RuntimeThreads {
			plg := core.Plugin{
				Runtime: true,
				Enabled: true,
				Name: plug.Name + "::" + name,
				Handler: inst,
			}
			c.Cfg.Plugins = append(c.Cfg.Plugins, plg)
		}
	}
}