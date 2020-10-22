package discord

type VoiceState struct {
	GuildId    Snowflake `json:"guild_id"`
	ChannelId  Snowflake `json:"channel_id"`
	UserId     Snowflake `json:"user_id"`
	Member     *Member   `json:"member"`
	SessionId  string    `json:"session_id"`
	Deaf       bool      `json:"deaf"`
	Mute       bool      `json:"mute"`
	SelfDeaf   bool      `json:"self_deaf"`
	SelfMute   bool      `json:"self_mute"`
	SelfStream bool      `json:"self_stream"`
	SelfVideo  bool      `json:"self_video"`
	Supress    bool      `json:"supress"`
}

type VoiceRegion struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	VIP        bool   `json:"vip"`
	Optimal    bool   `json:"optimal"`
	Deprecated bool   `json:"deprecated"`
	Custom     bool   `json:"custom"`
}
