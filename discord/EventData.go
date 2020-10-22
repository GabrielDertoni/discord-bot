package discord

import (
	"encoding/json"
)

type EventData json.RawMessage

func (data EventData) ToMessage() (message *Message, err error) {
	message = &Message{}
	err = json.Unmarshal([]byte(data), message)
	return
}

func (data EventData) ToChannel() (channel *Channel, err error) {
	channel = &Channel{}
	err = json.Unmarshal([]byte(data), channel)
	return
}
