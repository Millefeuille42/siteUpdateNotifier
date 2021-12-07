package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"siteUpdateNotifier/utils"
	"time"
)

var gBot *discordgo.Session
var ownerID = "268431730967314435" //Please change this when using my bot
var botToken = "ODk3NDc4NDM3ODgxNjU5NDUz.YWWP7Q.7icXumiIziN7H6nMcA7Nb_l0kvo"
var gPrefix = os.Getenv("SEGBOT_PREFIX")
var channel *discordgo.Channel

// startBot Starts discord bot
func startBot() *discordgo.Session {
	discordBot, err := discordgo.New("Bot " + botToken)
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

func ReadHTTPResponse(response *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return []byte(""), err
	}
	return body, nil
}

func main() {
	gBot = startBot()
	oldData := ""

	for {
		res, err := http.Get("https://talizmo.bigcartel.com/")
		if err != nil {
			_, _ = gBot.ChannelMessageSend(channel.ID, err.Error())
			continue
		}
		data, err := ReadHTTPResponse(res)
		if err != nil {
			_, _ = gBot.ChannelMessageSend(channel.ID, err.Error())
			continue
		}
		fmt.Print("Checking... ")
		newData := string(data)
		if oldData != newData {
			for i := 0; i < 15; i++ {
				_, err = gBot.ChannelMessageSend(channel.ID, "UwU Y a un nouvo truc sur le site !\n	https://talizmo.bigcartel.com/")
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				time.Sleep(time.Second * 3)
			}
		}
		fmt.Println("Announced!")
		oldData = newData
		time.Sleep(time.Minute*1 + (time.Duration(rand.Int()%30))*time.Second)
	}
}
