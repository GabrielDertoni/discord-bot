package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/GabrielDertoni/discord-bot/discord"
	"github.com/joho/godotenv"
)

type CreatedMessageData struct {
	Author  MessageAuthorData `json:"author"`
	Content string            `json:"content"`
}

type MessageAuthorData struct {
	Bot bool `json:"bot"`
}

var dertoniCounter int

func getHandleDertoni(bot *discord.Bot) func(*discord.Message) {
	return func(message *discord.Message) {
		channel := message.ChannelId
		if dertoniCounter%4 == 0 {
			file, err := os.Open("./assets/obamameme.png")
			if err != nil {
				return
			}
			defer file.Close()
			reader := bufio.NewReader(file)

			_, err = bot.SendMessage(&discord.MessageBody{
				Files: []*discord.File{
					&discord.File{
						Name:        "obamameme.png",
						ContentType: "image/png",
						Reader:      reader,
					},
				},
			}, channel)

			if err != nil {
				return
			}
		} else {
			bot.ReplyText(message, "Dertoni eh foda!")
		}
		dertoniCounter++
	}
}


func main() {
	godotenv.Load()
	token := os.Getenv("DISCORD_TOKEN")

	bot := discord.NewBot(token, []discord.Intent{discord.IntentGuildMessages})

	dertoniCounter = 1

	bot.AddMessageHandler(discord.NewCommandMessageHandler(".oi", func(message *discord.Message) {
		// Procura no google por resultados da pesquisa.
		_, err := bot.ReplyText(message, "Olá pessoa, como vai você?")
		if err != nil {
			fmt.Println("Message error:", err)
		}
	}))

	bot.AddMessageHandler(discord.NewCommandMessageHandler(".tchau", func(message *discord.Message) {
		bot.ReplyText(message, "Tchau, bom te ver!")
	}))

	bot.AddMessageHandler(discord.NewCommandMessageHandler(".quote", func(message *discord.Message) {
		resp, err := http.Get("https://inspirobot.me/api?generate=true")
		if err != nil {
			return
		}
		defer resp.Body.Close()
		imgURL, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		bot.SendImageWithURL(string(imgURL), message.ChannelId)
	}))

	bot.AddMessageHandler(discord.NewCommandMessageHandler(".pic", func(message *discord.Message) {
		if message.Author.Id != "383012204405981216" {
			bot.ReplyText(message, "Você não tem autorização para fazer essa mudança.")
			return
		}
		_, err := bot.ReplyText(message, "Mudando imagem de perfil...")
		if err != nil {
			return
		}
		err = bot.ModifyAvatarImageFromURL("https://thispersondoesnotexist.com/image", discord.ImageFileExtPNG)
		if err != nil {
			log.Println("An error occurred while trying to change bot avatar image:", err)
		}
		bot.ReplyText(message, "Prontinho!")
	}))

	servezaoHandler := discord.NewRegexpMessageHandler("(?i)Dertoni", getHandleDertoni(bot))
	bot.AddMessageHandler(discord.NewGuildExcluseMessageHandler("361126565725077504", servezaoHandler))

	err := bot.Online()
	if err != nil {
		log.Fatal(err)
	}
	defer bot.Offline()
	bot.StartListening()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	<-interrupt
	fmt.Println("User exit")
	return
}
