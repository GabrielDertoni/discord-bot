package discord

type Hello struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

type Ready struct {
	Version   int      `json:"v"`
	User      *User    `json:"user"`
	Guilds    []*Guild `json:"guilds"`
	SessionId string   `json:"session_id"`
}
