package bot

import (
	"encoding/json"
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
)

const maxAdminsToQueryForNotification = 1000

// PostMessage posts a message to a specified channel.
func (b *Bot) PostMessage(channelID, format string, args ...interface{}) (*model.Post, error) {
	post := &model.Post{
		Message:   fmt.Sprintf(format, args...),
		UserId:    b.botUserID,
		ChannelId: channelID,
	}
	if err := b.pluginAPI.Post.CreatePost(post); err != nil {
		return nil, err
	}
	return post, nil
}

// PostMessageToThread posts a message to a specified channel and thread identified by rootPostID.
// If the rootPostID is blank, or the rootPost is deleted, it will create a standalone post. The
// returned post's RootID will be the correct rootID (save that if you want to continue the thread).
func (b *Bot) PostMessageToThread(channelID, rootPostID, format string, args ...interface{}) (*model.Post, error) {
	rootID := ""
	if rootPostID != "" {
		root, err := b.pluginAPI.Post.GetPost(rootPostID)
		if err == nil && root != nil && root.DeleteAt == 0 {
			rootID = root.Id
		}
	}

	post := &model.Post{
		Message:   fmt.Sprintf(format, args...),
		UserId:    b.botUserID,
		ChannelId: channelID,
		RootId:    rootID,
	}
	if err := b.pluginAPI.Post.CreatePost(post); err != nil {
		return nil, err
	}

	return post, nil
}

// PostMessageWithAttachments posts a message with slack attachments to channelID. Returns the post id if
// posting was successful. Often used to include post actions.
func (b *Bot) PostMessageWithAttachments(channelID string, attachments []*model.SlackAttachment, format string, args ...interface{}) (*model.Post, error) {
	post := &model.Post{
		Message:   fmt.Sprintf(format, args...),
		UserId:    b.botUserID,
		ChannelId: channelID,
	}
	model.ParseSlackAttachment(post, attachments)
	if err := b.pluginAPI.Post.CreatePost(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (b *Bot) PostCustomMessageWithAttachments(channelID, customType string, attachments []*model.SlackAttachment, format string, args ...interface{}) (*model.Post, error) {
	post := &model.Post{
		Message:   fmt.Sprintf(format, args...),
		UserId:    b.botUserID,
		ChannelId: channelID,
		Type:      customType,
	}
	model.ParseSlackAttachment(post, attachments)
	if err := b.pluginAPI.Post.CreatePost(post); err != nil {
		return nil, err
	}
	return post, nil
}

// DM sends a DM from the plugin bot to the specified user
func (b *Bot) DM(userID string, post *model.Post) error {
	channel, err := b.pluginAPI.Channel.GetDirect(userID, b.botUserID)
	if err != nil {
		b.pluginAPI.Log.Info("Couldn't get bot's DM channel", "user_id", userID)
		return err
	}
	post.ChannelId = channel.Id
	post.UserId = b.botUserID

	return b.pluginAPI.Post.CreatePost(post)
}

// EphemeralPost sends an ephemeral message to a user
func (b *Bot) EphemeralPost(userID, channelID string, post *model.Post) {
	post.UserId = b.botUserID
	post.ChannelId = channelID

	b.pluginAPI.Post.SendEphemeralPost(userID, post)
}

// EphemeralPostWithAttachments sends an ephemeral message to a user with Slack attachments.
func (b *Bot) EphemeralPostWithAttachments(userID, channelID, postID string, attachments []*model.SlackAttachment, format string, args ...interface{}) {
	post := &model.Post{
		Message:   fmt.Sprintf(format, args...),
		UserId:    b.botUserID,
		ChannelId: channelID,
		RootId:    postID,
	}

	model.ParseSlackAttachment(post, attachments)
	b.pluginAPI.Post.SendEphemeralPost(userID, post)
}

// PublishWebsocketEventToTeam sends a websocket event with payload to teamID
func (b *Bot) PublishWebsocketEventToTeam(event string, payload interface{}, teamID string) {
	payloadMap := b.makePayloadMap(payload)
	b.pluginAPI.Frontend.PublishWebSocketEvent(event, payloadMap, &model.WebsocketBroadcast{
		TeamId: teamID,
	})
}

// PublishWebsocketEventToChannel sends a websocket event with payload to channelID
func (b *Bot) PublishWebsocketEventToChannel(event string, payload interface{}, channelID string) {
	payloadMap := b.makePayloadMap(payload)
	b.pluginAPI.Frontend.PublishWebSocketEvent(event, payloadMap, &model.WebsocketBroadcast{
		ChannelId: channelID,
	})
}

// PublishWebsocketEventToUser sends a websocket event with payload to userID
func (b *Bot) PublishWebsocketEventToUser(event string, payload interface{}, userID string) {
	payloadMap := b.makePayloadMap(payload)
	b.pluginAPI.Frontend.PublishWebSocketEvent(event, payloadMap, &model.WebsocketBroadcast{
		UserId: userID,
	})
}

func (b *Bot) makePayloadMap(payload interface{}) map[string]interface{} {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		b.With(LogContext{
			"payload": payload,
		}).Errorf("could not marshall payload")
		payloadJSON = []byte("null")
	}
	return map[string]interface{}{"payload": string(payloadJSON)}
}

// DM posts a simple Direct Message to the specified user
func (b *Bot) dmAdmins(format string, args ...interface{}) error {
	for _, id := range b.configService.GetConfiguration().AllowedUserIDs {
		err := b.DM(id, &model.Post{
			Message: fmt.Sprintf(format, args...),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
