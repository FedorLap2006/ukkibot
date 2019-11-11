package core

import (
	Discord "github.com/bwmarrin/discordgo"	
	//"log"
)

type Event struct {
	Client *Discord.Session
	Type string
	Data interface{}
}

type EventHandler func(*Core, *Event)

const (
	EVENT_MESSAGE_CREATE = "messageCreate"
	EVENT_MESSAGE_DELETE = "messageDelete"
)

// var lleventHandlers map[string]EventHandler
// var hlEventHandlers map[string]EventHandler

func (c *Core) RegEventHandler(event string, eh EventHandler) {
	if _, ok := _reservedEventHandlers[event]; ok {
		c.hlEventHandlers[event] = eh
	} else {
		c.llEventHandlers[event] = eh
	}
}

func _call_llEventHandler(c *Core,event string, data interface{}) {
	if eh, ok := c.llEventHandlers["messageCreate"]; ok {
		eh(c, &Event{ Client: c.Client, Type: event, Data: data})
	}
}

func (c *Core) llBaseEventHandler(client *Discord.Session, event interface{}) {
	switch event.(type) {
	case *Discord.MessageCreate:
		_call_llEventHandler(c,EVENT_MESSAGE_CREATE,event.(*Discord.MessageCreate))
		// c.llEventHandlers["messageCreate"](c, &Event{ Client: client, Type: EVENT_MESSAGE_CREATE, Data: event.(*Discord.MessageCreate)})
	case *Discord.MessageUpdate:
		_call_llEventHandler(c,EVENT_MESSAGE_DELETE,event.(*Discord.MessageDelete))
	}
}

/////////////////// RESERVED EVENT HANDLERS ///////////////////////////

func _callHighLevelEventHandler(c *Core, event *Event) {
	if eh, ok := c.hlEventHandlers[event.Type]; ok {
		eh(c,event)
	}
}

///////////////////////////////////////////////////////////////////////

var _reservedEventHandlers = map[string]EventHandler {
	EVENT_MESSAGE_CREATE: peh_messageCreate,
}

func init_events_Module(c *Core) {
	c.Client.AddHandler(c.llBaseEventHandler)
	c.llEventHandlers = make(map[string]EventHandler)
	c.hlEventHandlers = make(map[string]EventHandler)
	for k,v := range _reservedEventHandlers {
		c.llEventHandlers[k] = v
	}
}