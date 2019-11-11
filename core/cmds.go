package core

import (
	Discord "github.com/bwmarrin/discordgo"
	strs "strings"
	"regexp"
	//"log"
)

type CommandHandler func(message *Discord.Message, args []string)


func __REH__messageCreate(c *Core, event *Event) {
	_callHighLevelEventHandler(c,event)

	rch := event.Data.(*Discord.MessageCreate)

	if !strs.HasPrefix(rch.Content, c.Cfg.Prefix) {
		return
	}
	msg := strs.TrimPrefix(rch.Content,c.Cfg.Prefix)

	args := _msgCommandParse(msg)

	cmd := args[0]
	args = args[1:]

	if ch, ok := c.Cfg.Commands[cmd]; ok {
		ch(rch.Message, args)
	}
}

func _msgCommandParse(msg string) []string {
	re := regexp.MustCompile("\\s")
	args := re.Split(msg, -1)
	return args
}


func init__commands__Module(c *Core) {

}