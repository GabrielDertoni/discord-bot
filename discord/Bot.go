package discord

import (
	"container/list"
	"io/ioutil"
	"log"
	"net/http"
)

// Bot struct allows users to use the Discord Bot API through an object.
type Bot struct {
	Token          string
	GatewayManager *GatewayManager
	intents        int
	online         bool
	listeners      map[Event]*list.List
}

// NewBot - Creates a new bot struct with a token.
func NewBot(token string, intents []Intent) *Bot {
	var intent int = 0
	for _, i := range intents {
		intent |= int(i)
	}
	bot := Bot{
		Token:     token,
		intents:   intent,
		listeners: make(map[Event]*list.List),
	}
	return &bot
}

func (bot *Bot) AddMessageHandler(handler IMessageHandler) {
	bot.AddEventListener(EventMessageCreate, func(data EventData) {
		message, err := data.ToMessage()
		if err == nil && !message.Author.Bot && handler.Match(message.Content) {
			handler.Handle(message)
		}
	})
}

// AddEventListener - Adds a new event callback to the specified event
func (bot *Bot) AddEventListener(eventName Event, callback func(EventData)) func() {
	val, exists := bot.listeners[eventName]
	if !exists {
		val = list.New()
		bot.listeners[eventName] = val
	}
	elem := val.PushBack(callback)
	return func() {
		val.Remove(elem)
	}
}

func (bot *Bot) TriggerEvent(eventName Event, data EventData) {
	list, exists := bot.listeners[eventName]
	if exists {
		for e := list.Front(); e != nil; e = e.Next() {
			callback := e.Value.(func(EventData))
			callback(data)
		}
	}
}

func (bot *Bot) Online() (err error) {
	if bot.GatewayManager == nil {
		bot.GatewayManager = NewGatewayManager()
	}
	err = bot.GatewayManager.OpenConnection(bot.Token, bot.intents)
	if err != nil {
		return
	}
	bot.online = true
	log.Println("=== BOT ONLINE ===")
	return nil
}

func (bot *Bot) Offline() {
	if bot.online {
		bot.GatewayManager.CloseConnection()
		bot.online = false
	}
}

func (bot *Bot) StartListening() {
	go func() {
		for {
			payload := <-bot.GatewayManager.Events
			go bot.TriggerEvent(Event(payload.Event), EventData(payload.Data))
		}
	}()
}

func (bot *Bot) SendMessage(message *MessageBody, channelID string) (sent *Message, err error) {
	return CreateMessage(message, channelID, bot.Token)
}

func (bot *Bot) SendText(text string, channelID string) (sent *Message, err error) {
	return CreateMessage(&MessageBody{
		Content: text,
	}, channelID, bot.Token)
}

func (bot *Bot) SendImageWithURL(imgURL string, channelID string) (sent *Message, err error) {
	return CreateMessage(&MessageBody{
		Embed: &Embed{
			Title: "InspiroBot Says",
			URL:   "https://inspirobot.me/",
			Image: EmbedImage{
				URL: imgURL,
			},
		},
	}, channelID, bot.Token)
}

func (bot *Bot) ReplyText(message *Message, text string) (sent *Message, err error) {
	return bot.SendText(text, message.ChannelId)
}

func (bot *Bot) ModifyAvatarImageFromURL(imgURL string, ext ImageFileExt) (err error) {
	resp, err := http.Get(imgURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	imgData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	dataURI := BytesToImageData(imgData, ext)
	ModifyCurrentUser(&ModifyCurrentUserBody{
		Avatar: dataURI,
	}, bot.Token)
	return
}
