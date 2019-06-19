// Code generated by \"eventhandlers\"; DO NOT EDIT
// See events.go

package discordgo

// Following are all the event types.
// Event type values are used to match the events returned by Discord.
// EventTypes surrounded by __ are synthetic and are internal to DiscordGo.
const (
	channelCreateEventType            = "CHANNEL_CREATE"
	channelDeleteEventType            = "CHANNEL_DELETE"
	channelPinsUpdateEventType        = "CHANNEL_PINS_UPDATE"
	channelUpdateEventType            = "CHANNEL_UPDATE"
	connectEventType                  = "__CONNECT__"
	disconnectEventType               = "__DISCONNECT__"
	eventEventType                    = "__EVENT__"
	guildBanAddEventType              = "GUILD_BAN_ADD"
	guildBanRemoveEventType           = "GUILD_BAN_REMOVE"
	guildCreateEventType              = "GUILD_CREATE"
	guildDeleteEventType              = "GUILD_DELETE"
	guildEmojisUpdateEventType        = "GUILD_EMOJIS_UPDATE"
	guildIntegrationsUpdateEventType  = "GUILD_INTEGRATIONS_UPDATE"
	guildMemberAddEventType           = "GUILD_MEMBER_ADD"
	guildMemberRemoveEventType        = "GUILD_MEMBER_REMOVE"
	guildMemberUpdateEventType        = "GUILD_MEMBER_UPDATE"
	guildMembersChunkEventType        = "GUILD_MEMBERS_CHUNK"
	guildRoleCreateEventType          = "GUILD_ROLE_CREATE"
	guildRoleDeleteEventType          = "GUILD_ROLE_DELETE"
	guildRoleUpdateEventType          = "GUILD_ROLE_UPDATE"
	guildUpdateEventType              = "GUILD_UPDATE"
	messageAckEventType               = "MESSAGE_ACK"
	messageCreateEventType            = "MESSAGE_CREATE"
	messageDeleteEventType            = "MESSAGE_DELETE"
	messageDeleteBulkEventType        = "MESSAGE_DELETE_BULK"
	messageReactionAddEventType       = "MESSAGE_REACTION_ADD"
	messageReactionRemoveEventType    = "MESSAGE_REACTION_REMOVE"
	messageReactionRemoveAllEventType = "MESSAGE_REACTION_REMOVE_ALL"
	messageUpdateEventType            = "MESSAGE_UPDATE"
	presenceUpdateEventType           = "PRESENCE_UPDATE"
	presencesReplaceEventType         = "PRESENCES_REPLACE"
	rateLimitEventType                = "__RATE_LIMIT__"
	readyEventType                    = "READY"
	relationshipAddEventType          = "RELATIONSHIP_ADD"
	relationshipRemoveEventType       = "RELATIONSHIP_REMOVE"
	resumedEventType                  = "RESUMED"
	typingStartEventType              = "TYPING_START"
	userGuildSettingsUpdateEventType  = "USER_GUILD_SETTINGS_UPDATE"
	userNoteUpdateEventType           = "USER_NOTE_UPDATE"
	userSettingsUpdateEventType       = "USER_SETTINGS_UPDATE"
	userUpdateEventType               = "USER_UPDATE"
	voiceServerUpdateEventType        = "VOICE_SERVER_UPDATE"
	voiceStateUpdateEventType         = "VOICE_STATE_UPDATE"
	webhooksUpdateEventType           = "WEBHOOKS_UPDATE"
)

// channelCreateEventHandler is an event handler for ChannelCreate events.
type channelCreateEventHandler func(Sessioner, *ChannelCreate)

// Type returns the event type for ChannelCreate events.
func (eh channelCreateEventHandler) Type() string {
	return channelCreateEventType
}

// New returns a new instance of ChannelCreate.
func (eh channelCreateEventHandler) New() interface{} {
	return &ChannelCreate{}
}

// Handle is the handler for ChannelCreate events.
func (eh channelCreateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*ChannelCreate); ok {
		eh(s, t)
	}
}

// channelDeleteEventHandler is an event handler for ChannelDelete events.
type channelDeleteEventHandler func(Sessioner, *ChannelDelete)

// Type returns the event type for ChannelDelete events.
func (eh channelDeleteEventHandler) Type() string {
	return channelDeleteEventType
}

// New returns a new instance of ChannelDelete.
func (eh channelDeleteEventHandler) New() interface{} {
	return &ChannelDelete{}
}

// Handle is the handler for ChannelDelete events.
func (eh channelDeleteEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*ChannelDelete); ok {
		eh(s, t)
	}
}

// channelPinsUpdateEventHandler is an event handler for ChannelPinsUpdate events.
type channelPinsUpdateEventHandler func(Sessioner, *ChannelPinsUpdate)

// Type returns the event type for ChannelPinsUpdate events.
func (eh channelPinsUpdateEventHandler) Type() string {
	return channelPinsUpdateEventType
}

// New returns a new instance of ChannelPinsUpdate.
func (eh channelPinsUpdateEventHandler) New() interface{} {
	return &ChannelPinsUpdate{}
}

// Handle is the handler for ChannelPinsUpdate events.
func (eh channelPinsUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*ChannelPinsUpdate); ok {
		eh(s, t)
	}
}

// channelUpdateEventHandler is an event handler for ChannelUpdate events.
type channelUpdateEventHandler func(Sessioner, *ChannelUpdate)

// Type returns the event type for ChannelUpdate events.
func (eh channelUpdateEventHandler) Type() string {
	return channelUpdateEventType
}

// New returns a new instance of ChannelUpdate.
func (eh channelUpdateEventHandler) New() interface{} {
	return &ChannelUpdate{}
}

// Handle is the handler for ChannelUpdate events.
func (eh channelUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*ChannelUpdate); ok {
		eh(s, t)
	}
}

// connectEventHandler is an event handler for Connect events.
type connectEventHandler func(Sessioner, *Connect)

// Type returns the event type for Connect events.
func (eh connectEventHandler) Type() string {
	return connectEventType
}

// Handle is the handler for Connect events.
func (eh connectEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*Connect); ok {
		eh(s, t)
	}
}

// disconnectEventHandler is an event handler for Disconnect events.
type disconnectEventHandler func(Sessioner, *Disconnect)

// Type returns the event type for Disconnect events.
func (eh disconnectEventHandler) Type() string {
	return disconnectEventType
}

// Handle is the handler for Disconnect events.
func (eh disconnectEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*Disconnect); ok {
		eh(s, t)
	}
}

// eventEventHandler is an event handler for Event events.
type eventEventHandler func(Sessioner, *Event)

// Type returns the event type for Event events.
func (eh eventEventHandler) Type() string {
	return eventEventType
}

// Handle is the handler for Event events.
func (eh eventEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*Event); ok {
		eh(s, t)
	}
}

// guildBanAddEventHandler is an event handler for GuildBanAdd events.
type guildBanAddEventHandler func(Sessioner, *GuildBanAdd)

// Type returns the event type for GuildBanAdd events.
func (eh guildBanAddEventHandler) Type() string {
	return guildBanAddEventType
}

// New returns a new instance of GuildBanAdd.
func (eh guildBanAddEventHandler) New() interface{} {
	return &GuildBanAdd{}
}

// Handle is the handler for GuildBanAdd events.
func (eh guildBanAddEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildBanAdd); ok {
		eh(s, t)
	}
}

// guildBanRemoveEventHandler is an event handler for GuildBanRemove events.
type guildBanRemoveEventHandler func(Sessioner, *GuildBanRemove)

// Type returns the event type for GuildBanRemove events.
func (eh guildBanRemoveEventHandler) Type() string {
	return guildBanRemoveEventType
}

// New returns a new instance of GuildBanRemove.
func (eh guildBanRemoveEventHandler) New() interface{} {
	return &GuildBanRemove{}
}

// Handle is the handler for GuildBanRemove events.
func (eh guildBanRemoveEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildBanRemove); ok {
		eh(s, t)
	}
}

// guildCreateEventHandler is an event handler for GuildCreate events.
type guildCreateEventHandler func(Sessioner, *GuildCreate)

// Type returns the event type for GuildCreate events.
func (eh guildCreateEventHandler) Type() string {
	return guildCreateEventType
}

// New returns a new instance of GuildCreate.
func (eh guildCreateEventHandler) New() interface{} {
	return &GuildCreate{}
}

// Handle is the handler for GuildCreate events.
func (eh guildCreateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildCreate); ok {
		eh(s, t)
	}
}

// guildDeleteEventHandler is an event handler for GuildDelete events.
type guildDeleteEventHandler func(Sessioner, *GuildDelete)

// Type returns the event type for GuildDelete events.
func (eh guildDeleteEventHandler) Type() string {
	return guildDeleteEventType
}

// New returns a new instance of GuildDelete.
func (eh guildDeleteEventHandler) New() interface{} {
	return &GuildDelete{}
}

// Handle is the handler for GuildDelete events.
func (eh guildDeleteEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildDelete); ok {
		eh(s, t)
	}
}

// guildEmojisUpdateEventHandler is an event handler for GuildEmojisUpdate events.
type guildEmojisUpdateEventHandler func(Sessioner, *GuildEmojisUpdate)

// Type returns the event type for GuildEmojisUpdate events.
func (eh guildEmojisUpdateEventHandler) Type() string {
	return guildEmojisUpdateEventType
}

// New returns a new instance of GuildEmojisUpdate.
func (eh guildEmojisUpdateEventHandler) New() interface{} {
	return &GuildEmojisUpdate{}
}

// Handle is the handler for GuildEmojisUpdate events.
func (eh guildEmojisUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildEmojisUpdate); ok {
		eh(s, t)
	}
}

// guildIntegrationsUpdateEventHandler is an event handler for GuildIntegrationsUpdate events.
type guildIntegrationsUpdateEventHandler func(Sessioner, *GuildIntegrationsUpdate)

// Type returns the event type for GuildIntegrationsUpdate events.
func (eh guildIntegrationsUpdateEventHandler) Type() string {
	return guildIntegrationsUpdateEventType
}

// New returns a new instance of GuildIntegrationsUpdate.
func (eh guildIntegrationsUpdateEventHandler) New() interface{} {
	return &GuildIntegrationsUpdate{}
}

// Handle is the handler for GuildIntegrationsUpdate events.
func (eh guildIntegrationsUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildIntegrationsUpdate); ok {
		eh(s, t)
	}
}

// guildMemberAddEventHandler is an event handler for GuildMemberAdd events.
type guildMemberAddEventHandler func(Sessioner, *GuildMemberAdd)

// Type returns the event type for GuildMemberAdd events.
func (eh guildMemberAddEventHandler) Type() string {
	return guildMemberAddEventType
}

// New returns a new instance of GuildMemberAdd.
func (eh guildMemberAddEventHandler) New() interface{} {
	return &GuildMemberAdd{}
}

// Handle is the handler for GuildMemberAdd events.
func (eh guildMemberAddEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildMemberAdd); ok {
		eh(s, t)
	}
}

// guildMemberRemoveEventHandler is an event handler for GuildMemberRemove events.
type guildMemberRemoveEventHandler func(Sessioner, *GuildMemberRemove)

// Type returns the event type for GuildMemberRemove events.
func (eh guildMemberRemoveEventHandler) Type() string {
	return guildMemberRemoveEventType
}

// New returns a new instance of GuildMemberRemove.
func (eh guildMemberRemoveEventHandler) New() interface{} {
	return &GuildMemberRemove{}
}

// Handle is the handler for GuildMemberRemove events.
func (eh guildMemberRemoveEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildMemberRemove); ok {
		eh(s, t)
	}
}

// guildMemberUpdateEventHandler is an event handler for GuildMemberUpdate events.
type guildMemberUpdateEventHandler func(Sessioner, *GuildMemberUpdate)

// Type returns the event type for GuildMemberUpdate events.
func (eh guildMemberUpdateEventHandler) Type() string {
	return guildMemberUpdateEventType
}

// New returns a new instance of GuildMemberUpdate.
func (eh guildMemberUpdateEventHandler) New() interface{} {
	return &GuildMemberUpdate{}
}

// Handle is the handler for GuildMemberUpdate events.
func (eh guildMemberUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildMemberUpdate); ok {
		eh(s, t)
	}
}

// guildMembersChunkEventHandler is an event handler for GuildMembersChunk events.
type guildMembersChunkEventHandler func(Sessioner, *GuildMembersChunk)

// Type returns the event type for GuildMembersChunk events.
func (eh guildMembersChunkEventHandler) Type() string {
	return guildMembersChunkEventType
}

// New returns a new instance of GuildMembersChunk.
func (eh guildMembersChunkEventHandler) New() interface{} {
	return &GuildMembersChunk{}
}

// Handle is the handler for GuildMembersChunk events.
func (eh guildMembersChunkEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildMembersChunk); ok {
		eh(s, t)
	}
}

// guildRoleCreateEventHandler is an event handler for GuildRoleCreate events.
type guildRoleCreateEventHandler func(Sessioner, *GuildRoleCreate)

// Type returns the event type for GuildRoleCreate events.
func (eh guildRoleCreateEventHandler) Type() string {
	return guildRoleCreateEventType
}

// New returns a new instance of GuildRoleCreate.
func (eh guildRoleCreateEventHandler) New() interface{} {
	return &GuildRoleCreate{}
}

// Handle is the handler for GuildRoleCreate events.
func (eh guildRoleCreateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildRoleCreate); ok {
		eh(s, t)
	}
}

// guildRoleDeleteEventHandler is an event handler for GuildRoleDelete events.
type guildRoleDeleteEventHandler func(Sessioner, *GuildRoleDelete)

// Type returns the event type for GuildRoleDelete events.
func (eh guildRoleDeleteEventHandler) Type() string {
	return guildRoleDeleteEventType
}

// New returns a new instance of GuildRoleDelete.
func (eh guildRoleDeleteEventHandler) New() interface{} {
	return &GuildRoleDelete{}
}

// Handle is the handler for GuildRoleDelete events.
func (eh guildRoleDeleteEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildRoleDelete); ok {
		eh(s, t)
	}
}

// guildRoleUpdateEventHandler is an event handler for GuildRoleUpdate events.
type guildRoleUpdateEventHandler func(Sessioner, *GuildRoleUpdate)

// Type returns the event type for GuildRoleUpdate events.
func (eh guildRoleUpdateEventHandler) Type() string {
	return guildRoleUpdateEventType
}

// New returns a new instance of GuildRoleUpdate.
func (eh guildRoleUpdateEventHandler) New() interface{} {
	return &GuildRoleUpdate{}
}

// Handle is the handler for GuildRoleUpdate events.
func (eh guildRoleUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildRoleUpdate); ok {
		eh(s, t)
	}
}

// guildUpdateEventHandler is an event handler for GuildUpdate events.
type guildUpdateEventHandler func(Sessioner, *GuildUpdate)

// Type returns the event type for GuildUpdate events.
func (eh guildUpdateEventHandler) Type() string {
	return guildUpdateEventType
}

// New returns a new instance of GuildUpdate.
func (eh guildUpdateEventHandler) New() interface{} {
	return &GuildUpdate{}
}

// Handle is the handler for GuildUpdate events.
func (eh guildUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*GuildUpdate); ok {
		eh(s, t)
	}
}

// messageAckEventHandler is an event handler for MessageAck events.
type messageAckEventHandler func(Sessioner, *MessageAck)

// Type returns the event type for MessageAck events.
func (eh messageAckEventHandler) Type() string {
	return messageAckEventType
}

// New returns a new instance of MessageAck.
func (eh messageAckEventHandler) New() interface{} {
	return &MessageAck{}
}

// Handle is the handler for MessageAck events.
func (eh messageAckEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*MessageAck); ok {
		eh(s, t)
	}
}

// messageCreateEventHandler is an event handler for MessageCreate events.
type messageCreateEventHandler func(Sessioner, *MessageCreate)

// Type returns the event type for MessageCreate events.
func (eh messageCreateEventHandler) Type() string {
	return messageCreateEventType
}

// New returns a new instance of MessageCreate.
func (eh messageCreateEventHandler) New() interface{} {
	return &MessageCreate{}
}

// Handle is the handler for MessageCreate events.
func (eh messageCreateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*MessageCreate); ok {
		eh(s, t)
	}
}

// messageDeleteEventHandler is an event handler for MessageDelete events.
type messageDeleteEventHandler func(Sessioner, *MessageDelete)

// Type returns the event type for MessageDelete events.
func (eh messageDeleteEventHandler) Type() string {
	return messageDeleteEventType
}

// New returns a new instance of MessageDelete.
func (eh messageDeleteEventHandler) New() interface{} {
	return &MessageDelete{}
}

// Handle is the handler for MessageDelete events.
func (eh messageDeleteEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*MessageDelete); ok {
		eh(s, t)
	}
}

// messageDeleteBulkEventHandler is an event handler for MessageDeleteBulk events.
type messageDeleteBulkEventHandler func(Sessioner, *MessageDeleteBulk)

// Type returns the event type for MessageDeleteBulk events.
func (eh messageDeleteBulkEventHandler) Type() string {
	return messageDeleteBulkEventType
}

// New returns a new instance of MessageDeleteBulk.
func (eh messageDeleteBulkEventHandler) New() interface{} {
	return &MessageDeleteBulk{}
}

// Handle is the handler for MessageDeleteBulk events.
func (eh messageDeleteBulkEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*MessageDeleteBulk); ok {
		eh(s, t)
	}
}

// messageReactionAddEventHandler is an event handler for MessageReactionAdd events.
type messageReactionAddEventHandler func(Sessioner, *MessageReactionAdd)

// Type returns the event type for MessageReactionAdd events.
func (eh messageReactionAddEventHandler) Type() string {
	return messageReactionAddEventType
}

// New returns a new instance of MessageReactionAdd.
func (eh messageReactionAddEventHandler) New() interface{} {
	return &MessageReactionAdd{}
}

// Handle is the handler for MessageReactionAdd events.
func (eh messageReactionAddEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*MessageReactionAdd); ok {
		eh(s, t)
	}
}

// messageReactionRemoveEventHandler is an event handler for MessageReactionRemove events.
type messageReactionRemoveEventHandler func(Sessioner, *MessageReactionRemove)

// Type returns the event type for MessageReactionRemove events.
func (eh messageReactionRemoveEventHandler) Type() string {
	return messageReactionRemoveEventType
}

// New returns a new instance of MessageReactionRemove.
func (eh messageReactionRemoveEventHandler) New() interface{} {
	return &MessageReactionRemove{}
}

// Handle is the handler for MessageReactionRemove events.
func (eh messageReactionRemoveEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*MessageReactionRemove); ok {
		eh(s, t)
	}
}

// messageReactionRemoveAllEventHandler is an event handler for MessageReactionRemoveAll events.
type messageReactionRemoveAllEventHandler func(Sessioner, *MessageReactionRemoveAll)

// Type returns the event type for MessageReactionRemoveAll events.
func (eh messageReactionRemoveAllEventHandler) Type() string {
	return messageReactionRemoveAllEventType
}

// New returns a new instance of MessageReactionRemoveAll.
func (eh messageReactionRemoveAllEventHandler) New() interface{} {
	return &MessageReactionRemoveAll{}
}

// Handle is the handler for MessageReactionRemoveAll events.
func (eh messageReactionRemoveAllEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*MessageReactionRemoveAll); ok {
		eh(s, t)
	}
}

// messageUpdateEventHandler is an event handler for MessageUpdate events.
type messageUpdateEventHandler func(Sessioner, *MessageUpdate)

// Type returns the event type for MessageUpdate events.
func (eh messageUpdateEventHandler) Type() string {
	return messageUpdateEventType
}

// New returns a new instance of MessageUpdate.
func (eh messageUpdateEventHandler) New() interface{} {
	return &MessageUpdate{}
}

// Handle is the handler for MessageUpdate events.
func (eh messageUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*MessageUpdate); ok {
		eh(s, t)
	}
}

// presenceUpdateEventHandler is an event handler for PresenceUpdate events.
type presenceUpdateEventHandler func(Sessioner, *PresenceUpdate)

// Type returns the event type for PresenceUpdate events.
func (eh presenceUpdateEventHandler) Type() string {
	return presenceUpdateEventType
}

// New returns a new instance of PresenceUpdate.
func (eh presenceUpdateEventHandler) New() interface{} {
	return &PresenceUpdate{}
}

// Handle is the handler for PresenceUpdate events.
func (eh presenceUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*PresenceUpdate); ok {
		eh(s, t)
	}
}

// presencesReplaceEventHandler is an event handler for PresencesReplace events.
type presencesReplaceEventHandler func(Sessioner, *PresencesReplace)

// Type returns the event type for PresencesReplace events.
func (eh presencesReplaceEventHandler) Type() string {
	return presencesReplaceEventType
}

// New returns a new instance of PresencesReplace.
func (eh presencesReplaceEventHandler) New() interface{} {
	return &PresencesReplace{}
}

// Handle is the handler for PresencesReplace events.
func (eh presencesReplaceEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*PresencesReplace); ok {
		eh(s, t)
	}
}

// rateLimitEventHandler is an event handler for RateLimit events.
type rateLimitEventHandler func(Sessioner, *RateLimit)

// Type returns the event type for RateLimit events.
func (eh rateLimitEventHandler) Type() string {
	return rateLimitEventType
}

// Handle is the handler for RateLimit events.
func (eh rateLimitEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*RateLimit); ok {
		eh(s, t)
	}
}

// readyEventHandler is an event handler for Ready events.
type readyEventHandler func(Sessioner, *Ready)

// Type returns the event type for Ready events.
func (eh readyEventHandler) Type() string {
	return readyEventType
}

// New returns a new instance of Ready.
func (eh readyEventHandler) New() interface{} {
	return &Ready{}
}

// Handle is the handler for Ready events.
func (eh readyEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*Ready); ok {
		eh(s, t)
	}
}

// relationshipAddEventHandler is an event handler for RelationshipAdd events.
type relationshipAddEventHandler func(Sessioner, *RelationshipAdd)

// Type returns the event type for RelationshipAdd events.
func (eh relationshipAddEventHandler) Type() string {
	return relationshipAddEventType
}

// New returns a new instance of RelationshipAdd.
func (eh relationshipAddEventHandler) New() interface{} {
	return &RelationshipAdd{}
}

// Handle is the handler for RelationshipAdd events.
func (eh relationshipAddEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*RelationshipAdd); ok {
		eh(s, t)
	}
}

// relationshipRemoveEventHandler is an event handler for RelationshipRemove events.
type relationshipRemoveEventHandler func(Sessioner, *RelationshipRemove)

// Type returns the event type for RelationshipRemove events.
func (eh relationshipRemoveEventHandler) Type() string {
	return relationshipRemoveEventType
}

// New returns a new instance of RelationshipRemove.
func (eh relationshipRemoveEventHandler) New() interface{} {
	return &RelationshipRemove{}
}

// Handle is the handler for RelationshipRemove events.
func (eh relationshipRemoveEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*RelationshipRemove); ok {
		eh(s, t)
	}
}

// resumedEventHandler is an event handler for Resumed events.
type resumedEventHandler func(Sessioner, *Resumed)

// Type returns the event type for Resumed events.
func (eh resumedEventHandler) Type() string {
	return resumedEventType
}

// New returns a new instance of Resumed.
func (eh resumedEventHandler) New() interface{} {
	return &Resumed{}
}

// Handle is the handler for Resumed events.
func (eh resumedEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*Resumed); ok {
		eh(s, t)
	}
}

// typingStartEventHandler is an event handler for TypingStart events.
type typingStartEventHandler func(Sessioner, *TypingStart)

// Type returns the event type for TypingStart events.
func (eh typingStartEventHandler) Type() string {
	return typingStartEventType
}

// New returns a new instance of TypingStart.
func (eh typingStartEventHandler) New() interface{} {
	return &TypingStart{}
}

// Handle is the handler for TypingStart events.
func (eh typingStartEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*TypingStart); ok {
		eh(s, t)
	}
}

// userGuildSettingsUpdateEventHandler is an event handler for UserGuildSettingsUpdate events.
type userGuildSettingsUpdateEventHandler func(Sessioner, *UserGuildSettingsUpdate)

// Type returns the event type for UserGuildSettingsUpdate events.
func (eh userGuildSettingsUpdateEventHandler) Type() string {
	return userGuildSettingsUpdateEventType
}

// New returns a new instance of UserGuildSettingsUpdate.
func (eh userGuildSettingsUpdateEventHandler) New() interface{} {
	return &UserGuildSettingsUpdate{}
}

// Handle is the handler for UserGuildSettingsUpdate events.
func (eh userGuildSettingsUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*UserGuildSettingsUpdate); ok {
		eh(s, t)
	}
}

// userNoteUpdateEventHandler is an event handler for UserNoteUpdate events.
type userNoteUpdateEventHandler func(Sessioner, *UserNoteUpdate)

// Type returns the event type for UserNoteUpdate events.
func (eh userNoteUpdateEventHandler) Type() string {
	return userNoteUpdateEventType
}

// New returns a new instance of UserNoteUpdate.
func (eh userNoteUpdateEventHandler) New() interface{} {
	return &UserNoteUpdate{}
}

// Handle is the handler for UserNoteUpdate events.
func (eh userNoteUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*UserNoteUpdate); ok {
		eh(s, t)
	}
}

// userSettingsUpdateEventHandler is an event handler for UserSettingsUpdate events.
type userSettingsUpdateEventHandler func(Sessioner, *UserSettingsUpdate)

// Type returns the event type for UserSettingsUpdate events.
func (eh userSettingsUpdateEventHandler) Type() string {
	return userSettingsUpdateEventType
}

// New returns a new instance of UserSettingsUpdate.
func (eh userSettingsUpdateEventHandler) New() interface{} {
	return &UserSettingsUpdate{}
}

// Handle is the handler for UserSettingsUpdate events.
func (eh userSettingsUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*UserSettingsUpdate); ok {
		eh(s, t)
	}
}

// userUpdateEventHandler is an event handler for UserUpdate events.
type userUpdateEventHandler func(Sessioner, *UserUpdate)

// Type returns the event type for UserUpdate events.
func (eh userUpdateEventHandler) Type() string {
	return userUpdateEventType
}

// New returns a new instance of UserUpdate.
func (eh userUpdateEventHandler) New() interface{} {
	return &UserUpdate{}
}

// Handle is the handler for UserUpdate events.
func (eh userUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*UserUpdate); ok {
		eh(s, t)
	}
}

// voiceServerUpdateEventHandler is an event handler for VoiceServerUpdate events.
type voiceServerUpdateEventHandler func(Sessioner, *VoiceServerUpdate)

// Type returns the event type for VoiceServerUpdate events.
func (eh voiceServerUpdateEventHandler) Type() string {
	return voiceServerUpdateEventType
}

// New returns a new instance of VoiceServerUpdate.
func (eh voiceServerUpdateEventHandler) New() interface{} {
	return &VoiceServerUpdate{}
}

// Handle is the handler for VoiceServerUpdate events.
func (eh voiceServerUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*VoiceServerUpdate); ok {
		eh(s, t)
	}
}

// voiceStateUpdateEventHandler is an event handler for VoiceStateUpdate events.
type voiceStateUpdateEventHandler func(Sessioner, *VoiceStateUpdate)

// Type returns the event type for VoiceStateUpdate events.
func (eh voiceStateUpdateEventHandler) Type() string {
	return voiceStateUpdateEventType
}

// New returns a new instance of VoiceStateUpdate.
func (eh voiceStateUpdateEventHandler) New() interface{} {
	return &VoiceStateUpdate{}
}

// Handle is the handler for VoiceStateUpdate events.
func (eh voiceStateUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*VoiceStateUpdate); ok {
		eh(s, t)
	}
}

// webhooksUpdateEventHandler is an event handler for WebhooksUpdate events.
type webhooksUpdateEventHandler func(Sessioner, *WebhooksUpdate)

// Type returns the event type for WebhooksUpdate events.
func (eh webhooksUpdateEventHandler) Type() string {
	return webhooksUpdateEventType
}

// New returns a new instance of WebhooksUpdate.
func (eh webhooksUpdateEventHandler) New() interface{} {
	return &WebhooksUpdate{}
}

// Handle is the handler for WebhooksUpdate events.
func (eh webhooksUpdateEventHandler) Handle(s Sessioner, i interface{}) {
	if t, ok := i.(*WebhooksUpdate); ok {
		eh(s, t)
	}
}

func handlerForInterface(handler interface{}) EventHandler {
	switch v := handler.(type) {
	case func(Sessioner, interface{}):
		return interfaceEventHandler(v)
	case func(Sessioner, *ChannelCreate):
		return channelCreateEventHandler(v)
	case func(Sessioner, *ChannelDelete):
		return channelDeleteEventHandler(v)
	case func(Sessioner, *ChannelPinsUpdate):
		return channelPinsUpdateEventHandler(v)
	case func(Sessioner, *ChannelUpdate):
		return channelUpdateEventHandler(v)
	case func(Sessioner, *Connect):
		return connectEventHandler(v)
	case func(Sessioner, *Disconnect):
		return disconnectEventHandler(v)
	case func(Sessioner, *Event):
		return eventEventHandler(v)
	case func(Sessioner, *GuildBanAdd):
		return guildBanAddEventHandler(v)
	case func(Sessioner, *GuildBanRemove):
		return guildBanRemoveEventHandler(v)
	case func(Sessioner, *GuildCreate):
		return guildCreateEventHandler(v)
	case func(Sessioner, *GuildDelete):
		return guildDeleteEventHandler(v)
	case func(Sessioner, *GuildEmojisUpdate):
		return guildEmojisUpdateEventHandler(v)
	case func(Sessioner, *GuildIntegrationsUpdate):
		return guildIntegrationsUpdateEventHandler(v)
	case func(Sessioner, *GuildMemberAdd):
		return guildMemberAddEventHandler(v)
	case func(Sessioner, *GuildMemberRemove):
		return guildMemberRemoveEventHandler(v)
	case func(Sessioner, *GuildMemberUpdate):
		return guildMemberUpdateEventHandler(v)
	case func(Sessioner, *GuildMembersChunk):
		return guildMembersChunkEventHandler(v)
	case func(Sessioner, *GuildRoleCreate):
		return guildRoleCreateEventHandler(v)
	case func(Sessioner, *GuildRoleDelete):
		return guildRoleDeleteEventHandler(v)
	case func(Sessioner, *GuildRoleUpdate):
		return guildRoleUpdateEventHandler(v)
	case func(Sessioner, *GuildUpdate):
		return guildUpdateEventHandler(v)
	case func(Sessioner, *MessageAck):
		return messageAckEventHandler(v)
	case func(Sessioner, *MessageCreate):
		return messageCreateEventHandler(v)
	case func(Sessioner, *MessageDelete):
		return messageDeleteEventHandler(v)
	case func(Sessioner, *MessageDeleteBulk):
		return messageDeleteBulkEventHandler(v)
	case func(Sessioner, *MessageReactionAdd):
		return messageReactionAddEventHandler(v)
	case func(Sessioner, *MessageReactionRemove):
		return messageReactionRemoveEventHandler(v)
	case func(Sessioner, *MessageReactionRemoveAll):
		return messageReactionRemoveAllEventHandler(v)
	case func(Sessioner, *MessageUpdate):
		return messageUpdateEventHandler(v)
	case func(Sessioner, *PresenceUpdate):
		return presenceUpdateEventHandler(v)
	case func(Sessioner, *PresencesReplace):
		return presencesReplaceEventHandler(v)
	case func(Sessioner, *RateLimit):
		return rateLimitEventHandler(v)
	case func(Sessioner, *Ready):
		return readyEventHandler(v)
	case func(Sessioner, *RelationshipAdd):
		return relationshipAddEventHandler(v)
	case func(Sessioner, *RelationshipRemove):
		return relationshipRemoveEventHandler(v)
	case func(Sessioner, *Resumed):
		return resumedEventHandler(v)
	case func(Sessioner, *TypingStart):
		return typingStartEventHandler(v)
	case func(Sessioner, *UserGuildSettingsUpdate):
		return userGuildSettingsUpdateEventHandler(v)
	case func(Sessioner, *UserNoteUpdate):
		return userNoteUpdateEventHandler(v)
	case func(Sessioner, *UserSettingsUpdate):
		return userSettingsUpdateEventHandler(v)
	case func(Sessioner, *UserUpdate):
		return userUpdateEventHandler(v)
	case func(Sessioner, *VoiceServerUpdate):
		return voiceServerUpdateEventHandler(v)
	case func(Sessioner, *VoiceStateUpdate):
		return voiceStateUpdateEventHandler(v)
	case func(Sessioner, *WebhooksUpdate):
		return webhooksUpdateEventHandler(v)
	}

	return nil
}

func init() {
	registerInterfaceProvider(channelCreateEventHandler(nil))
	registerInterfaceProvider(channelDeleteEventHandler(nil))
	registerInterfaceProvider(channelPinsUpdateEventHandler(nil))
	registerInterfaceProvider(channelUpdateEventHandler(nil))
	registerInterfaceProvider(guildBanAddEventHandler(nil))
	registerInterfaceProvider(guildBanRemoveEventHandler(nil))
	registerInterfaceProvider(guildCreateEventHandler(nil))
	registerInterfaceProvider(guildDeleteEventHandler(nil))
	registerInterfaceProvider(guildEmojisUpdateEventHandler(nil))
	registerInterfaceProvider(guildIntegrationsUpdateEventHandler(nil))
	registerInterfaceProvider(guildMemberAddEventHandler(nil))
	registerInterfaceProvider(guildMemberRemoveEventHandler(nil))
	registerInterfaceProvider(guildMemberUpdateEventHandler(nil))
	registerInterfaceProvider(guildMembersChunkEventHandler(nil))
	registerInterfaceProvider(guildRoleCreateEventHandler(nil))
	registerInterfaceProvider(guildRoleDeleteEventHandler(nil))
	registerInterfaceProvider(guildRoleUpdateEventHandler(nil))
	registerInterfaceProvider(guildUpdateEventHandler(nil))
	registerInterfaceProvider(messageAckEventHandler(nil))
	registerInterfaceProvider(messageCreateEventHandler(nil))
	registerInterfaceProvider(messageDeleteEventHandler(nil))
	registerInterfaceProvider(messageDeleteBulkEventHandler(nil))
	registerInterfaceProvider(messageReactionAddEventHandler(nil))
	registerInterfaceProvider(messageReactionRemoveEventHandler(nil))
	registerInterfaceProvider(messageReactionRemoveAllEventHandler(nil))
	registerInterfaceProvider(messageUpdateEventHandler(nil))
	registerInterfaceProvider(presenceUpdateEventHandler(nil))
	registerInterfaceProvider(presencesReplaceEventHandler(nil))
	registerInterfaceProvider(readyEventHandler(nil))
	registerInterfaceProvider(relationshipAddEventHandler(nil))
	registerInterfaceProvider(relationshipRemoveEventHandler(nil))
	registerInterfaceProvider(resumedEventHandler(nil))
	registerInterfaceProvider(typingStartEventHandler(nil))
	registerInterfaceProvider(userGuildSettingsUpdateEventHandler(nil))
	registerInterfaceProvider(userNoteUpdateEventHandler(nil))
	registerInterfaceProvider(userSettingsUpdateEventHandler(nil))
	registerInterfaceProvider(userUpdateEventHandler(nil))
	registerInterfaceProvider(voiceServerUpdateEventHandler(nil))
	registerInterfaceProvider(voiceStateUpdateEventHandler(nil))
	registerInterfaceProvider(webhooksUpdateEventHandler(nil))
}
