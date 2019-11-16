package plugins

import (
	"core"
	Discord "github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

func init() {
	RegPlugin(&BotPlugin{
		Name: "Mod",
		Commands: map[string]core.Command {
			"ban": core.Command {
				Name: "ban",
				Usage: "ban @user|id reason",
				Handler: banCommand,
				Enabled: true,
			},
		},
		RuntimeThreads: map[string]core.PluginHandler {
			"checkban": modCheckBan,
		},
		//Events: map[string]core.EventHandler{},
	})
}

func modCheckBan(c *core.Core) {
	
}

func banCommand(s *Discord.Session, message *Discord.Message, args []string) {
	//core.BotLog("INFO","Bot::plugins::mod::ban", "yeah")
	if len(args) == 0 {
		return
	}
	useri := args[0]

	reid := regexp.MustCompile(`\d+`)

	if matched := string(reid.Find([]byte(useri))); len(matched) > 0 {
		if len(args) > 1 {
			reason := strings.Join(args[1:], " ")
			s.GuildBanCreateWithReason(message.GuildID, matched, reason, 0)
		} else {
			s.GuildBanCreate(message.GuildID, matched, 0)
		}
	}
}
