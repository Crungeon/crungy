package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func DingDong(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ding" reply with "Dong!"
	if strings.ToLower(m.Content) == "ding" {
		s.ChannelMessageSend(m.ChannelID, "Dong!")
	}

	// If the message is "pong" reply with "ding!"
	if strings.ToLower(m.Content) == "dong" {
		s.ChannelMessageSend(m.ChannelID, "Ding?")
	}
}
