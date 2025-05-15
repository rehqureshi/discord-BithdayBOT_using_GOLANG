package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rehqureshi/go-pingmod-Discord/config"
)

var BotID string

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return

	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Sprintln("Bot is running")

}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "hello bot" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "HELLO")
		return
	}
	for _, user := range m.Mentions {
		if user.ID == BotID {
			_, _ = s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> YEEH DAWG YOU MENTIONED ME ?")
			return
		}
	}

}
