package discord

type Snowflake string

type Message struct {
	Id               string              `json:"id"`
	ChannelId        string              `json:"channel_id"`
	GuildId          string              `json:"guild_id"`
	Author           *User               `json:"author"`
	Member           *Member             `json:"member"`
	Content          string              `json:"content"`
	Timestamp        Timestamp           `json:"timestamp"`
	EditedTimestamp  Timestamp           `json:"edited_timestamp"`
	TTS              bool                `json:"tts"`
	MentionEveryone  bool                `json:"mention_everyone"`
	Mentions         []*User             `json:"mentions"`
	MentionRoles     []*Role             `json:"mention_roles"`
	MentionChannels  []*ChannelMention   `json:"mention_channels"`
	Attachments      []*Attachment       `json:"attachment"`
	Embeds           []*Embed            `json:"embeds"`
	Reactions        []*Reaction         `json:"reactions"`
	Pinned           bool                `json:"pinned"`
	Type             MessageType         `json:"type"`
	Activity         *MessageActivity    `json:"activity"`
	Application      *MessageApplication `json:"application"`
	MessageReference *MessageReference   `json:"message_reference"`
	Flags            MessageFlag         `json:"flags"`
}

type MessageReference struct {
	MessageId string `json:"message_id"`
	ChannelId string `json:"channel_id"`
	GuildId   string `json:"guild_id"`
}

type MessageApplication struct {
	Id          string `json:"id"`
	CoverImage  string `json:"cover_image"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
}

type MessageActivity struct {
	Type    MessageActivityType `json:"type"`
	PartyId string              `json:"party_id"`
}

type Reaction struct {
	Count int   `json:"count"`
	Me    bool  `json:"me"`
	Emoji Emoji `json:"emoji"`
}

type Emoji struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Roles         []*Role `json:"roles"`
	User          *User   `json:"user"`
	RequireColons bool    `json:"require_colons"`
	Managed       bool    `json:"managed"`
	Animated      bool    `json:"animated"`
	Available     bool    `json:"available"`
}

type Embed struct {
	Title       string     `json:"title"`
	Type        EmbedType  `json:"type"`
	Description string     `json:"description"`
	URL         string     `json:"url"`
	Timestamp   Timestamp  `json:"timestamp"`
	Color       int        `json:"color"`
	Image       EmbedImage `json:"image"`
	// NOTE: implement footer, image, thumbnail, video, provider, author, field
}

type EmbedImage struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type Attachment struct {
	Id       string `json:"attachment"`
	Filename string `json:"filename"`
	Size     int    `json:"size"` // In bytes
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"` // height if image
	Width    int    `json:"width"`  // width if image
}

type ChannelMention struct {
	Id      string      `json:"id"`
	GuildId string      `json:"guild_id"`
	Type    ChannelType `json:"type"`
	Name    string      `json:"name"`
}

type Role struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Color       int    `json:"color"`
	Hoist       bool   `json:"hoist"`
	Position    int    `json:"position"`
	Permissions string `json:"permissions"`
	Managed     bool   `json:"managed"`
	Mentionable bool   `json:"mentionable"`
}

type User struct {
	Id            string   `json:"id"`
	Username      string   `json:"username"`
	Discriminator string   `json:"discriminator"`
	Avatar        string   `json:"avatar"`
	Bot           bool     `json:"bot"`
	System        bool     `json:"system"`
	MFAEnabled    bool     `json:"mfa_enabled"`
	Locale        string   `json:"locale"`
	Verified      bool     `json:"verified"`
	Email         string   `json:"email"`
	Flags         UserFlag `json:"flags"`
	// PremiumType
	PublicFlags UserFlag `json:"public_flags"`
}

type Member struct {
	User     *User     `json:"user"`
	Nick     string    `json:"nick"`
	Roles    []string  `json:"roles"`
	JoinedAt Timestamp `json:"joined_at"`
	// PremiumSince
	Deaf bool `json:"deaf"`
	Mute bool `json:"mute"`
}

type UserFlag uint32

const (
	UserFlagNone                      UserFlag = 0
	UserFlagDiscordEmployee                    = 1
	UserFlagPartnerServerOwner                 = 1 << 1
	UserFlagHypeSquadEvents                    = 1 << 2
	UserFlagBugHunterLevel1                    = 1 << 3
	UserFlagHouseBravery                       = 1 << 6
	UserFlagHouseBrilliance                    = 1 << 7
	UserFlagHouseBalance                       = 1 << 8
	UserFlagEarlySupporter                     = 1 << 9
	UserFlagTeamUser                           = 1 << 10
	UserFlagSystem                             = 1 << 12
	UserFlagBugHunterLevel2                    = 1 << 14
	UserFlagVerifiedBot                        = 1 << 16
	UserFlagEarlyVerifiedBotDeveloper          = 1 << 17
)

type EmbedType string

const (
	Rich             EmbedType = "rich"
	EmbedTypeImage             = "image"
	EmbedTypeVideo             = "video"
	EmbedTypeGIFV              = "gifv"
	EmbedTypeArticle           = "article"
	EmbedTypeLink              = "link"
)

type MessageType uint16

const (
	MessageTypeDefault MessageType = iota
	MessageTypeRecipientAdd
	MessageTypeRecipientRemove
	MessageTypeCall
	MessageTypeChannelNameChange
	MessageTypeChannelIconChange
	MessageTypeChannelPinnedMessage
	MessageTypeGuildMemberJoin
	MessageTypeUserPremiumGuildSubscription
	MessageTypeUserPremiumGuildSubscriptionTier1
	MessageTypeUserPremiumGuildSubscriptionTier2
	MessageTypeUserPremiumGuildSubscriptionTier3
	MessageTypeChannelFollowAdd
	MessageTypeGuildDiscoveryDisqualified
	MessageTypeGuildDiscoveryRequalified
)

type MessageActivityType uint16

const (
	Join MessageActivityType = iota
	Spectate
	Listen
	JoinRequest
)

type MessageFlag uint32

const (
	Crossposted          MessageFlag = 1 << 0
	IsCrosspost                      = 1 << 1
	SupressEmbeds                    = 1 << 2
	SourceMessageDeleted             = 1 << 3
	Urgent                           = 1 << 4
)
