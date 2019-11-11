package main

import (
	"core"
	Discord "github.com/bwmarrin/discordgo"
)

type BotPlugin struct {
	Name string
	Commands map[string]core.CommandHandler
	Events map[string]core.EventHandler
}


func (p BotPlugin) Init(c *Core) {
	for k, eh := range p.Events {
		c.RegEventHandler(k,eh)
	}
	for n, ch := range p.Commands {
		if _, ok := c.Cfg.Commands; !ok {
			c.Cfg.Commands[n] = ch
		}
	}
}