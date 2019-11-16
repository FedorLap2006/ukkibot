package core

import (
	Discord "github.com/bwmarrin/discordgo"
	"errors"
	"time"
)

type Core struct {
	Cfg *CoreCfg
	Client *Discord.Session
	rtCrc chan bool
	running bool
	llEventHandlers map[string]EventHandler
	hlEventHandlers map[string]EventHandler
	Payload interface{}
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
		if !p.Enabled {
			continue
		}
		if !p.Runtime {
			p.Handler(c)
		}
	}
	for k, eh := range c.Cfg.Events {
		c.RegEventHandler(k,eh)
	}
}

func (c *Core) runtimeThread() {
	for {
		if !c.running {
			return
		}
		time.Sleep(255 * time.Millisecond)
		for _, p := range c.Cfg.Plugins {
			if !p.Enabled {
				continue
			}
			if p.Runtime {
				if p.FinishChnl == nil {
					p.FinishChnl = make(chan bool, 1)
				}
				select {
					case _ = <-p.FinishChnl:
						go func() {
							p.Handler(c)
							p.FinishChnl <- true
						}()
					default:
						continue
				}
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

	go c.runtimeThread()
	c.running = true

	return nil

}

func(c *Core) Stop() error {
	err := c.Client.Close()
	c.running = false
	return err
}