module ukkibot

replace core => ./core

replace plugins => ./plugins

require (
	core v0.0.1
	github.com/bwmarrin/discordgo v0.20.1
	plugins v0.0.1
)
