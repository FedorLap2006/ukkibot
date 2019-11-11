package core

import (
	Discord "github.com/bwmarrin/discordgo"
	"errors"
)

type Core struct {
	Cfg *CoreCfg
	Client *Discord.Session
	rtCrc <-chan bool
	llEventHandlers map[string]EventHandler
	hlEventHandlers map[string]EventHandler
}

type moduleInitHandler func(*Core)
var modulesInitHandlers []moduleInitHandler

func init() {
	modulesInitHandlers = []moduleInitHandler {
		init_events_Module,
	}
}

func (c *Core) Init() {
	c.Client, _ = Discord.New("Bot " + c.Cfg.Token)

	for _, m := range modulesInitHandlers {
		m(c)
	}
	for _, p := range c.Cfg.Plugins {
		if(!p.Runtime) p.Handler(c)
	}
	for k, eh := range c.Cfg.Events {
		c.RegEventHandler(k,eh)
	}
}

func (c *Core) runtimeThread(crc <-chan bool) {
	for {
		select {
		case _ := <-crc:
			return
		default:
			for _, p := range c.Cfg.Plugins {
				if(p.Runtime) p.Handler(c)
			}
		}
	}
}


func(c *Core) Run() error {
	if c.Cfg == nil {
		return errors.New("Cfg is empty")
	}
	if c.Client == nil {
		return errors.New("Core is not initialized")
	}

	err := c.Client.Open()

	if err != nil {
		return err
	}

	go runtimeThread(c.rtCrc)

	return nil

}

func(c *Core) Stop() error {
	c.rtCrc <- true
	return c.Client.Close()
}