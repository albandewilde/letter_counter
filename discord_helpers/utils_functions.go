package discord_helpers

import (
	"github.com/bwmarrin/discordgo"
)

func DiscordUserCompleteName(user *discordgo.User) string {
	return user.Username + "#" + user.Discriminator
}
