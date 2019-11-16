package core

import (
	//Discord "github.com/bwmarrin/discordgo"
)

type PluginHandler func(c *Core)

type Plugin struct {
	Runtime bool
	Enabled bool
	Name string
	Handler PluginHandler
	Payload interface{}
	FinishChnl chan bool
}