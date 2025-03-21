package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// DiscordUtils contains utility functions for Discord operations
type DiscordUtils struct {
	Session *discordgo.Session
}

// NewDiscordUtils creates a new DiscordUtils instance
func NewDiscordUtils(session *discordgo.Session) *DiscordUtils {
	return &DiscordUtils{
		Session: session,
	}
}

// SendMessage sends a message to a channel
func (d *DiscordUtils) SendMessage(channelID, content string) (*discordgo.Message, error) {
	if content == "" {
		return nil, errors.New("message content cannot be empty")
	}

	return d.Session.ChannelMessageSend(channelID, content)
}

// SendReply sends a reply to a message
func (d *DiscordUtils) SendReply(channelID, content string, reference *discordgo.MessageReference) (*discordgo.Message, error) {
	if content == "" {
		return nil, errors.New("message content cannot be empty")
	}

	return d.Session.ChannelMessageSendReply(channelID, content, reference)
}

// SendComplex sends a complex message with embeds, files, etc.
func (d *DiscordUtils) SendComplex(channelID string, data *discordgo.MessageSend) (*discordgo.Message, error) {
	return d.Session.ChannelMessageSendComplex(channelID, data)
}

// EditMessage edits a message
func (d *DiscordUtils) EditMessage(channelID, messageID, content string) (*discordgo.Message, error) {
	return d.Session.ChannelMessageEdit(channelID, messageID, content)
}

// DeleteMessage deletes a message
func (d *DiscordUtils) DeleteMessage(channelID, messageID string) error {
	return d.Session.ChannelMessageDelete(channelID, messageID)
}

// GetChannel gets a channel by ID
func (d *DiscordUtils) GetChannel(channelID string) (*discordgo.Channel, error) {
	return d.Session.State.Channel(channelID)
}

// GetGuild gets a guild by ID
func (d *DiscordUtils) GetGuild(guildID string) (*discordgo.Guild, error) {
	return d.Session.State.Guild(guildID)
}

// GetMember gets a guild member
func (d *DiscordUtils) GetMember(guildID, userID string) (*discordgo.Member, error) {
	return d.Session.GuildMember(guildID, userID)
}

// GetUser gets a user by ID
func (d *DiscordUtils) GetUser(userID string) (*discordgo.User, error) {
	return d.Session.User(userID)
}

// GetUserAvatar gets a user's avatar as bytes
func (d *DiscordUtils) GetUserAvatar(user *discordgo.User) ([]byte, error) {
	avatar, err := d.Session.UserAvatarDecode(user)
	if err != nil {
		return nil, fmt.Errorf("failed to decode user avatar: %w", err)
	}

	return avatar.Bytes(), nil
}

// GetRecentMessages gets recent messages from a channel
func (d *DiscordUtils) GetRecentMessages(channelID string, limit int, beforeID string) ([]*discordgo.Message, error) {
	return d.Session.ChannelMessages(channelID, limit, beforeID, "", "")
}

// FindRecentReply finds a recent reply to a specific message
func (d *DiscordUtils) FindRecentReply(channelID, messageID string, limit int) (*discordgo.Message, error) {
	messages, err := d.GetRecentMessages(channelID, limit, "")
	if err != nil {
		return nil, err
	}

	for _, msg := range messages {
		if msg.ReferencedMessage != nil && msg.ReferencedMessage.ID == messageID && msg.Author.ID == d.Session.State.User.ID {
			return msg, nil
		}
	}

	return nil, errors.New("reply not found")
}

// UpdateStatus updates the bot's status
func (d *DiscordUtils) UpdateStatus(status string) error {
	return d.Session.UpdateGameStatus(0, status)
}

// UpdateCustomStatus updates the bot's custom status
func (d *DiscordUtils) UpdateCustomStatus(status string) error {
	err := d.Session.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "online",
		Activities: []*discordgo.Activity{
			{
				Name: status,
				Type: discordgo.ActivityTypeCustom,
			},
		},
	})

	return err
}

// IsOwner checks if a user is a bot owner
func (d *DiscordUtils) IsOwner(userID string, ownerIDs []string) bool {
	for _, id := range ownerIDs {
		if id == userID {
			return true
		}
	}
	return false
}

// HasPermission checks if the bot has a specific permission in a channel
func (d *DiscordUtils) HasPermission(channelID string, permission int64) (bool, error) {
	perms, err := d.Session.State.UserChannelPermissions(d.Session.State.User.ID, channelID)
	if err != nil {
		return false, err
	}

	return (perms & permission) == permission, nil
}

// FormatTimestamp formats a Unix timestamp for Discord display
func (d *DiscordUtils) FormatTimestamp(timestamp int64, style string) string {
	switch style {
	case "R": // Relative time
		return fmt.Sprintf("<t:%d:R>", timestamp)
	case "F": // Full date and time
		return fmt.Sprintf("<t:%d:F>", timestamp)
	case "T": // Time only
		return fmt.Sprintf("<t:%d:T>", timestamp)
	case "D": // Date only
		return fmt.Sprintf("<t:%d:D>", timestamp)
	default: // Default style
		return fmt.Sprintf("<t:%d>", timestamp)
	}
}

// ParseMention extracts a user ID from a mention string
func (d *DiscordUtils) ParseMention(mention string) (string, error) {
	if !strings.HasPrefix(mention, "<@") || !strings.HasSuffix(mention, ">") {
		return "", errors.New("invalid mention format")
	}

	// Handle nickname mentions
	mention = strings.TrimPrefix(mention, "<@!")
	mention = strings.TrimPrefix(mention, "<@")
	mention = strings.TrimSuffix(mention, ">")

	return mention, nil
}

// IsBot checks if a user is a bot
func (d *DiscordUtils) IsBot(userID string) bool {
	user, err := d.GetUser(userID)
	if err != nil {
		return false
	}

	return user.Bot
}

// SoftFail sends an error message while maintaining the flow
func (d *DiscordUtils) SoftFail(channelID string, reference *discordgo.MessageReference, errMsg string) {
	d.SendReply(channelID, "I know what you are. "+errMsg, reference)
}

// HardFail sends a critical error message
func (d *DiscordUtils) HardFail(channelID string, reference *discordgo.MessageReference, err error) {
	errMsg := "Sorry, my creator must have fucked something up.\nPlease pierce him with a sanguine lance and drink his blood."
	if err != nil {
		fmt.Println("Error:", err)
	}
	d.SendReply(channelID, errMsg, reference)
}
