package main

import (
	"core"
	"os"
	"os/signal"
	"syscall"
	"plugins"
	//"fmt"
	//"log"
	Discord "github.com/bwmarrin/discordgo"
)

var gcfg core.CoreCfg

var botToken string = os.Getenv("UKKIBOT_TOKEN")

// type ConfigPayload struct {
// 	BotPlugins map[string]BotPlugin
// }

func main() {
	gcfg.Prefix = "$"
	gcfg.Token = botToken
	gcfg.Events = map[string]core.EventHandler {}
	gcfg.Plugins = []core.Plugin {
		core.Plugin {
			Handler: plugins.PluginsLoader,
			Name: "PluginsLoader",
			Runtime: false,
			Enabled: true,
		},
	}
	gcfg.Commands = map[string]core.Command {
		"help": core.Command {
			Handler: func(client *Discord.Session, message *Discord.Message, args []string) {
				if len(args) != 0 {
					cmd, ok := gcfg.Commands[args[0]]
					if ok {
						client.ChannelMessageSendEmbed(message.ChannelID,&Discord.MessageEmbed{
							Title: "Command Help: " + cmd.Name, 
							Fields: []*Discord.MessageEmbedField{
								&Discord.MessageEmbedField{
									Name: "Usage",
									Value: cmd.Usage,
								},
							},
							
						})
					}
				} else {
					fields := []*Discord.MessageEmbedField{}
					for _, plug := range (gcfg.Payload.(plugins.BotPluginConfigPayload)).Plugins {
						chelp := ""
						for _, c := range plug.Commands {
							chelp += "**`" + c.Name + "`**" + " - " + c.Usage + "\n"
						}
						field := &Discord.MessageEmbedField {
							Name: plug.Name,
							Value: chelp,
						}
						fields = append(fields, field)
					}

					client.ChannelMessageSendEmbed(message.ChannelID,&Discord.MessageEmbed{Title: "Help", Fields: fields} )
				}
			},
			Name: "help",
			Usage: "help for info",
			Enabled: true,
		},
		"ping": core.Command {
			Handler: func(client *Discord.Session, message *Discord.Message, args []string) {
				client.ChannelMessageSendEmbed(message.ChannelID,&Discord.MessageEmbed{Title: "Help", Description: "pong!"} )
			},
			Name: "ping",
			Usage: "ping pong",
			Enabled: true,
		},
	}
	// gcfg.Payload = ConfigPayload{
	// 	BotPlugins: []BotPlugin {
	// 		Name: "mod"
	// 		Commands: 
	// 	}
	// }

	bot := core.Core{Cfg: &gcfg}

	bot.Init()


	err := bot.Run()

	if err != nil {
		core.BotLog("ERR", "Bot::Run", err.Error())
	}

	core.BotLog("INFO", "Bot::Worker", "Bot is running")
	
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc
  err = bot.Stop()

	if err != nil {
		core.BotLog("ERR", "Bot::Stop", err.Error())
	}
	core.BotLog("INFO", "Bot::Stop", "Bot is stopped . _.")
  os.Exit(0)
}