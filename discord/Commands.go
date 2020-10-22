package discord

import "encoding/json"

type Status string

const (
	StatusOnline       Status = "online"
	StatusOffline             = "offline"
	StatusDoNotDisturb        = "dnd"
	StatusIdle                = "idle"
	StatusInvisible           = "invisible"
)

type Payload struct {
	Op       OpCode          `json:"op"`
	Data     json.RawMessage `json:"d"`
	Sequence int             `json:"s"`
	Event    string          `json:"t"`
}

type Identify struct {
	Token      string             `json:"token"`
	Intents    int                `json:"intents"`
	Properties IdentifyProperties `json:"properties"`
	Presence   PresenceUpdate
}

type IdentifyProperties struct {
	OS      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

type PresenceUpdate struct {
	Since      int
	Activities []Activity
	Status     Status
	AFK        bool
}

type Activity struct {
	Name       string              `json:"name"`
	Type       ActivityType        `json:"type"`
	URL        string              `json:"url"`
	CreatedAt  int                 `json:"created_at"`
	Timestamps *ActivityTimestamps `json:"timestamps"`
}

type ActivityTimestamps struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type ActivityType int

const (
	ActivityTypeGame ActivityType = iota
	ActivityTypeStreaming
	ActivityTypeListening
	ActivityTypeCustom
	ActivityTypeCompeting
)
