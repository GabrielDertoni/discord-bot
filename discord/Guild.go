package discord

type Guild struct {
	Id                    string
	Name                  string
	Icon                  string
	Spash                 string
	DiscoverySplash       string
	Owner                 bool
	OwnerId               string
	Permissions           string
	Region                string
	AFKChannelId          Snowflake
	AFKTimeout            int // in seconds
	WidgetEnabled         bool
	WidgetChannelId       Snowflake
	VerificationLevel     int // NOTE: should be its own type
	ExplicitContentFilter int // NOTE: should be its own type
	Roles                 []*Role
	Emojis                []*Emoji
	// Features
	// MFALevel
	ApplicationId            Snowflake
	SystemChannelId          Snowflake
	SystemChannelFlags       int // NOTE: should be its own type
	RulesChannelId           Snowflake
	JoinedAt                 Timestamp
	Large                    bool
	Unavailable              bool
	MemberCount              int
	VoiceStates              []*VoiceState
	Members                  []*Member
	Channels                 []*Channel
	Presences                []*PresenceUpdate
	MaxPresences             int
	MaxMembers               int
	VanityURLCode            string
	Description              string
	Banner                   string
	PremiumTier              int
	PremiumSubscriptionCount int
	PreferedLocale           string
	PublicUpdatesChannelId   Snowflake
	MaxVideoChannelUsers     int
	ApproximateMemberCount   int
	ApproximatePresenceCount int
}
