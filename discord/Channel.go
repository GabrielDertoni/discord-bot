package discord

type Channel struct {
	Id                  Snowflake    `json:"id"`
	Type                ChannelType  `json:"type"`
	GuildId             Snowflake    `json:"guild_id"`
	Position            int          `json:"position"`
	PermisionOverwrites []*Overwrite `json:"permision_overwrites"`
	Name                string       `json:"name"`
	Topic               string       `json:"topic"`
	NSFW                bool         `json:"nsfw"`
	LastMessageId       Snowflake    `json:"last_message_id"`
	Bitrate             int          `json:"bitrate"`
	UserLimit           int          `json:"user_limit"`
	RateLimitPerUser    int          `json:"rateLimit_per_user"`
	Recipients          []*User      `json:"recipients"`
	Icon                string       `json:"icon"`
	OwnerId             Snowflake    `json:"owner_id"`
	ApplicationId       Snowflake    `json:"application_id"`
	ParentId            Snowflake    `json:"parent_id"`
	LastPinTimestamp    Timestamp    `json:"last_pin_timestamp"`
}

type ChannelType int

const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDirectMessage
	ChannelTypeGuildVoice
	ChannelTypeGroupDirectMessage
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	ChannelTypeGuildStore
)

type Overwrite struct {
	Id    Snowflake     `json:"id"`
	Type  OverwriteType `json:"type"`
	Allow string        `json:"allow"`
	Deny  string        `json:"deny"`
}

type OverwriteType uint8

const (
	OverwriteTypeRole   OverwriteType = 0
	OverwriteTypeMember               = 1
)

type AllowedMentionType string

const (
	RoleMention     AllowedMentionType = "roles"
	UserMention                        = "users"
	EveryoneMention                    = "everyone"
)
