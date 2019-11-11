package main

import (
	"core"
	"os"
	"os/signal"
	"syscall"
	"log"
	Discord "github.com/bwmarrin/discordgo"
)

var gcfg core.BotCfg

const botToken = os.Getenv("UKKIBOT_TOKEN")


func main() {
	gcfg.Prefix = ">~"
	gcfg.Token = botToken
	gcfg.Events = map[string]core.EventHandler {}
	gcfg.Plugins = []core.Plugin {}
	gcfg.Commands = map[string]core.CommandHandler {

	}

	bot := core.Core{Cfg: &gcfg}
	

	bot.Init()



	err := bot.Run()

	if err != nil {
		core.BotLog("ERR", "Bot::Run", err.Error())
	}

	core.BotLog("INFO", "Bot::Worker", "Bot is running")
	
	sc := make(chan os.Signal, 1)
	ech := make(chan bool, 1)


	go func() {
		for {
			select {
			case _ := <- ech:
				return
			default:
				botThread(bot)
			}	
		}
	}()



	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	ech <- true

	err = bot.Stop()

	if err != nil {
		core.BotLog("ERR", "Bot::Stop", err.Error())
	}
}