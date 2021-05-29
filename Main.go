package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"os/signal"

	"github.com/GabrielDertoni/discord-bot/discord"
	"github.com/joho/godotenv"
)

type CodeCreateResponseData struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type CodeDetailsResponseData struct {
	Stdout      string `json:"stdout"`
	Stderr      string `json:"stderr"`
	BuildStderr string `json:"build_stderr"`
}

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
					{
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

	bot.AddMessageHandler(discord.NewRegexpMessageHandler(".*", func(message *discord.Message) {
		code := ParseCodeBlock(message.Content)
		if code != nil {
			apiURL := "http://api.paiza.io/runners"
			params := url.Values{}
			params.Add("api_key", "guest")
			params.Add("longpoll", "true")
			params.Add("source_code", code.Code)
			params.Add("language", string(code.Lang))

			resp, err := http.Post(apiURL + "/create?" + params.Encode(), "text/plain;charset=UTF-8", nil)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return
			}
			result := CodeCreateResponseData{}
			err = json.Unmarshal(body, &result)
			if err != nil {
				return
			}
			detailsParams := url.Values{}
			detailsParams.Add("api_key", "guest")
			detailsParams.Add("id", result.Id)
			detailsResp, err := http.Get(apiURL + "/get_details?" + detailsParams.Encode())
			if err != nil {
				return
			}
			defer detailsResp.Body.Close()
			detailsBody, err := ioutil.ReadAll(detailsResp.Body)
			if err != nil {
				return
			}
			details := CodeDetailsResponseData{}
			err = json.Unmarshal(detailsBody, &details)
			if err != nil {
				return
			}
			if details.BuildStderr != "" {
				stderr := fmt.Sprintf("```\n%s```", details.BuildStderr)
				bot.ReplyText(message, stderr)
			} else if details.Stderr != "" {
				stderr := fmt.Sprintf("```\n%s```", details.Stderr)
				bot.ReplyText(message, stderr)
			} else if details.Stdout != "" {
				stdout := fmt.Sprintf("```\n%s```", details.Stdout)
				bot.ReplyText(message, stdout)
			} else {
				bot.ReplyText(message, "Nenhuma saída produzida pelo código...")
			}
		}
	}))

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
