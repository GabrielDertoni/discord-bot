package discord

import (
	"regexp"
	"strings"
)

type IMessageHandler interface {
	Match(text string) bool
	Handle(message *Message)
}

type RegexpMessageHandler struct {
	Regexp   *regexp.Regexp
	HandleFn func(*Message)
}

func NewRegexpMessageHandler(expStr string, handle func(*Message)) *RegexpMessageHandler {
	return &RegexpMessageHandler{
		Regexp:   regexp.MustCompile(expStr),
		HandleFn: handle,
	}
}

func (handler RegexpMessageHandler) Handle(message *Message) {
	handler.HandleFn(message)
}

func (handler RegexpMessageHandler) Match(text string) bool {
	return handler.Regexp.MatchString(text)
}

type CommandMessageHandler struct {
	Command  string
	HandleFn func(*Message)
}

func NewCommandMessageHandler(command string, handle func(*Message)) *CommandMessageHandler {
	return &CommandMessageHandler{
		Command:  command,
		HandleFn: handle,
	}
}

func (handler CommandMessageHandler) Match(text string) bool {
	return strings.HasPrefix(text, handler.Command)
}

func (handler CommandMessageHandler) Handle(message *Message) {
	handler.HandleFn(message)
}

type GuildExcluseMessageHandler struct {
	GuildId string
	Handler IMessageHandler
}

func NewGuildExcluseMessageHandler(guildId string, handler IMessageHandler) *GuildExcluseMessageHandler {
	return &GuildExcluseMessageHandler{
		GuildId: guildId,
		Handler: handler,
	}
}

func (handler GuildExcluseMessageHandler) Match(text string) bool {
	return handler.Handler.Match(text)
}

func (handler GuildExcluseMessageHandler) Handle(message *Message) {
	if message.GuildId == handler.GuildId {
		handler.Handler.Handle(message);
	}
}
