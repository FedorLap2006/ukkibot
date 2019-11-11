package core

import (
	//Discord "github.com/bwmarrin/discordgo"
	"log"
)

func BotLog(msgtype, module, message string) {
	switch msgtype {
	case "ERR":
		log.Fatalf("<%s> [%s]: (%s)\n", msgtype, module,message)
	case "ERROR":
		log.Fatalf("<%s> [%s]: (%s)\n", msgtype, module,message)
	case "WARN":
		log.Panicf("<%s> [%s]: (%s)\n", msgtype, module,message)
	case "WARNING":
		log.Panicf("<%s> [%s]: (%s)\n", msgtype, module,message)
	default:
		log.Printf("<%s> [%s]: (%s)\n", msgtype, module,message)
	}
	
}
