package core

import (
	//Discord "github.com/bwmarrin/discordgo"
)

type CoreCfg struct {
	Token string
	Prefix string
	Plugins []Plugin
	Events map[string]EventHandler
	Commands map[string]Command
	Payload interface{}
}


