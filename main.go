package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net/http"
	"os"
	"siteUpdateNotifier/utils"
	"time"
)

var gBot *discordgo.Session
var ownerID = "268431730967314435" //Please change this when using my bot
var gPrefix = os.Getenv("SEGBOT_PREFIX")
var channel *discordgo.Channel

// startBot Starts discord bot
func startBot() *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	utils.CheckError(err)
	err = discordBot.Open()
	utils.CheckError(err)
	fmt.Println("Discord bot created")
	if os.Getenv("SEGBOT_IN_PROD") == "" {
		channel, err = discordBot.UserChannelCreate(ownerID)
		if err != nil {
			return nil
		}
		hostname, _ := os.Hostname()
		_, _ = discordBot.ChannelMessageSend(channel.ID, "Bot up - "+
			time.Now().Format(time.Stamp)+" - "+hostname)
	}
	if gPrefix == "" {
		gPrefix = "!"
	}
	utils.SetUpCloseHandler(discordBot)

	return discordBot
}

func main() {
	gBot = startBot()

	for {
		http.Get()
		time.Sleep(time.Second * 3)
	}
}
