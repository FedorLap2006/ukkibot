package core

import (
	Discord "github.com/bwmarrin/discordgo"
	strs "strings"
	//"regexp"
	//"log"
)

type CommandHandler func(s *Discord.Session, message *Discord.Message, args []string)

type Command struct {
	Name string
	Usage string
	Enabled bool
	Handler CommandHandler
	Aliases []string
	Payload interface{}
}


func peh_messageCreate(c *Core, event *Event) {
	_callHighLevelEventHandler(c,event)

	rch := event.Data.(*Discord.MessageCreate)

	if !strs.HasPrefix(rch.Content, c.Cfg.Prefix) {
		return
	}
	msg := strs.TrimPrefix(rch.Content,c.Cfg.Prefix)

	args := _msgCommandParse(msg)

	cmd := args[0]
	args = args[1:]
	for _, ci := range c.Cfg.Commands {
		if !ci.Enabled {
			continue
		}
		
		if ci.Name == cmd {
			ci.Handler(c.Client, rch.Message, args)
		} else {
			if ci.Aliases == nil {
				continue
			}
			for _, alias := range ci.Aliases {
				if alias == cmd {
					ci.Handler(c.Client, rch.Message, args)
				}
			}
		}
	}
}

func _msgCommandParse(msg string) []string {
	args := strs.Fields(msg)
	return args
}


func init__commands__Module(c *Core) {

}