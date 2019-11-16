package plugins

import (
	"core"
	Discord "github.com/bwmarrin/discordgo"
	"fmt"
	"time"
//	"math"
)

func init() {
	RegPlugin(&BotPlugin{
		Name: "Info",
		Commands: map[string]core.Command {
			"botinfo": core.Command {
				Name: "botinfo",
				Usage: "bot info",
				Handler: botInfoCommand,
				Enabled: true,
				Aliases: []string{"bot"},
			},
			"serverinfo": core.Command {
				Name: "serverinfo",
				Usage: "serverinfo",
				Handler: serverInfoCommand,
				Enabled: true,
				Aliases: []string{"servinfo", "sinfo", "server"},
			},
		},
		Events: map[string]core.EventHandler{},
	})
	//client.ChannelMessageSendEmbed(message.ChannelID,&Discord.MessageEmbed{Title: "Help", Description: chelp} )
}

func serverInfoCommand(s *Discord.Session, message *Discord.Message, args []string) {
	guild, _ := s.Guild(message.GuildID)
	guildIcon := guild.IconURL()
	//author, _ := s.GuildMember(message.GuildID, message.Author.ID)
	owner, _ := s.GuildMember(message.GuildID, guild.OwnerID)
	bots := []*Discord.Member{}
	users := []*Discord.Member{}
	for _, member := range guild.Members {
		if member.User.Bot {
			bots = append(bots, member)
		} else {
			users = append(users, member)
		}
	}

	s.ChannelMessageSendEmbed(message.ChannelID, &Discord.MessageEmbed{
		Title: "Server Info",
		Author : &Discord.MessageEmbedAuthor{	
			Name: guild.Name,
			IconURL: guildIcon,
		},
		Thumbnail: &Discord.MessageEmbedThumbnail {
			URL: guildIcon,
		},
		Fields: []*Discord.MessageEmbedField {
			&Discord.MessageEmbedField{
				Name: "Name",
				Value: guild.Name,
			},
			&Discord.MessageEmbedField{
				Name: "Owner",
				Value: owner.Mention(),
			},
			&Discord.MessageEmbedField{
				Name: "Members",
				Value: fmt.Sprintf("**Total Count**: %v\n**Users Count**: %v\n**Bots Count**: %v",guild.MemberCount, len(users), len(bots)),
			},
			&Discord.MessageEmbedField{
				Name: "Region",
				Value: guild.Region,
			},
		},
	})
}

func botInfoCommand(s *Discord.Session, message *Discord.Message, args []string) {
	// checkPing := func(message *Discord.Message) uint64 {
	// 	timestamp, err := message.Timestamp.Parse()
	// 	if err != nil {
	// 		core.BotLog("WARN", "Bot::Plugins::Info", err.Error())
	// 		return 0
	// 	}

	// 	ping := uint64(time.Since(timestamp))
	// 	return ping
//	}
	pingts, _ := message.Timestamp.Parse()
//	fmt.Printf("%v", -time.Since(pingts))
	//ursize := float64(memstats.Sys) / 1024.0 / 1024.0
	author, _ := s.GuildMember(message.GuildID, message.Author.ID)
	s.ChannelMessageSendEmbed(message.ChannelID, &Discord.MessageEmbed{ 
		Title: "Bot Info", 
		Author: &Discord.MessageEmbedAuthor{
			Name: author.Nick,
			IconURL: author.User.AvatarURL(""),
		},
		Fields: []*Discord.MessageEmbedField {
			&Discord.MessageEmbedField {
				Name: "Ping",
				Value: fmt.Sprintf("%v", -time.Since(pingts)),
				Inline: true,
			},
			&Discord.MessageEmbedField {
				Name: "Developer",
				Value: "<@!553557567591284737>",
				Inline: true,
			},
			&Discord.MessageEmbedField {
				Name: "Version",
				Value: "1.0.0",
			},
		},
	})
}
