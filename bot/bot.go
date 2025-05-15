package bot

import (
	"fmt"
	"strings"
	"time"

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
	loadBirthdays()
	goBot.AddHandler(messageHandler)
	startDailyBirthdayCheck(goBot, "1372543500985438272")

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Sprintln("Bot is running")

}
func startDailyBirthdayCheck(s *discordgo.Session, channelID string) {
	go func() {
		for {
			now := time.Now()
			next := now.AddDate(0, 0, 1)
			nextMidnight := time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			time.Sleep(time.Until(nextMidnight))

			today := time.Now().Format("01-02")
			for _, b := range birthday {
				t, err := time.Parse("2006-01-02", b.Date)
				if err != nil {
					continue
				}
				if t.Format("01-02") == today {
					msg := "🎉 Happy Birthday <@" + b.UserID + ">!"
					_, _ = s.ChannelMessageSend(channelID, msg)
				}
			}

		}
	}()
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
			_, _ = s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> YEAH DAWG YOU MENTIONED ME ?")
			return
		}
	}
	if strings.HasPrefix(m.Content, "!birthday ") {
		date := strings.TrimSpace(strings.TrimPrefix(m.Content, "!birthday "))
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Please use format YYYY-MM-DD")
			return
		}
		updated := false

		for i, b := range birthday {
			if b.UserID == m.Author.ID {
				birthday[i].Date = date
				updated = true
				break
			}
		}
		if !updated {
			birthday = append(birthday, Birthday{
				UserID:   m.Author.ID,
				Date:     date,
				Username: m.Author.Username,
			})
		}
		saveBirthdays()
		if parsedDate.Format("01-02") == time.Now().Format("01-02") {
			s.ChannelMessageSend("1372543500985438272", "🎉 Happy Birthday <@"+m.Author.ID+">! 🎂")
		}
		if updated {
			s.ChannelMessageSend(m.ChannelID, "Updated your birthday to "+date+", "+m.Author.Username)
		} else {
			s.ChannelMessageSend(m.ChannelID, "Got it! Birthday saved for "+m.Author.Username)
		}
		return
	}

}
