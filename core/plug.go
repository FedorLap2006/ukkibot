package core

import (
	//Discord "github.com/bwmarrin/discordgo"
)

type PluginHandler func(c *Core)

type Plugin struct {
	Runtime bool
	PluginHandler
}