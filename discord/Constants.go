package discord

type OpCode int

const (
	OpCodeDispatch           OpCode = 0
	OpCodeHeartbeat                 = 1
	OpCodeIdentify                  = 2
	OpCodePresenceUpdate            = 3
	OpCodeVoiceStateUpdate          = 4
	OpCodeResume                    = 6
	OpCodeReconnect                 = 7
	OpCodeRequestGuilMembers        = 8
	OpCodeInvalidSession            = 9
	OpCodeHello                     = 10
	OpCodeHeartbeatACK              = 11
)

type Event string

const (
	EventHello                      Event = "Hello"
	EventReady                            = "Ready"
	EventResumed                          = "Resumed"
	EventReconnect                        = "Reconnect"
	EventInvalidSession                   = "Invalid Session"
	EventChannelCreate                    = "Channel Create"
	EventChannelUpdate                    = "Channel Update"
	EventChannelDelete                    = "Channel Delete"
	EventChannelPinsUpdate                = "Channel Pins Update"
	EventGuildCreate                      = "Guild Create"
	EventGuildUpdate                      = "Guild Update"
	EventGuildDelete                      = "Guild Delete"
	EventGuildBanAdd                      = "Guild Ban Add"
	EventGuildBanRemove                   = "Guild Ban Remove"
	EventGuildEmojisUpdate                = "Guild Emojis Update"
	EventGuildIntegrationsUpdate          = "Guild Integrations Update"
	EventGuildMemberAdd                   = "Guild Member Add"
	EventGuildMemberRemove                = "Guild Member Remove"
	EventGuildMemberUpdate                = "Guild Member Update"
	EventGuildMembersChunk                = "Guild Members Chunk"
	EventGuildRoleCreate                  = "Guild Role Create"
	EventGuildRoleUpdate                  = "Guild Role Update"
	EventGuildRoleDelete                  = "Guild Role Delete"
	EventInviteCreate                     = "Invite Create"
	EventInviteDelete                     = "Invite Delete"
	EventMessageCreate                    = "MESSAGE_CREATE"
	EventMessageUpdate                    = "Message Update"
	EventMessageDelete                    = "Message Delete"
	EventMessageDeleteBulk                = "Message Delete Bulk"
	EventMessageReactionAdd               = "Message Reaction Add"
	EventMessageReactionRemove            = "Message Reaction Remove"
	EventMessageReactionRemoveAll         = "Message Reaction Remove All"
	EventMessageReactionRemoveEmoji       = "Message Reaction Remove Emoji"
	EventPresenceUpdate                   = "Presence Update"
	EventTypingStart                      = "Typing Start"
	EventUserUpdate                       = "User Update"
	EventVoiceStateUpdate                 = "Voice State Update"
	EventVoiceServerUpdate                = "Voice Server Update"
	EventWebhooksUpdate                   = "Webhooks Update"
)

type Intent int

const (
	IntentGuilds                 Intent = 1
	IntentGuildMembers                  = 1 << 1
	IntentGuildBans                     = 1 << 2
	IntentGuildEmojis                   = 1 << 3
	IntentGuildIntegrations             = 1 << 4
	IntentGuildWebhooks                 = 1 << 5
	IntentGuildInvites                  = 1 << 6
	IntentGuildVoiceStates              = 1 << 7
	IntentGuildPresences                = 1 << 8
	IntentGuildMessages                 = 1 << 9
	IntentGuildMessageReactions         = 1 << 10
	IntentGuildMessageTyping            = 1 << 11
	IntentDirectMessages                = 1 << 12
	IntentDirectMessageReactions        = 1 << 13
	IntentDirectMessageTyping           = 1 << 14
)

const GatewayURL string = "wss://gateway.discord.gg/?v=8&encoding=json"
