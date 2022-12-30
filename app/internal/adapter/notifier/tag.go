package notifier

import (
	"context"
	"fmt"

	"github.com/slack-go/slack"
	"github.com/takokun778/2022/internal/domain/external"
)

var _ external.Tag = (*Tag)(nil)

type Tag struct {
	channel *Channel
}

func NewTag(
	channel *Channel,
) *Tag {
	return &Tag{
		channel: channel,
	}
}

func (t *Tag) Notice(ctx context.Context, msg string) error {
	if _, _, err := t.channel.PostMessage(t.channel.ID, slack.MsgOptionText(msg, true)); err != nil {
		return fmt.Errorf("failed to post message: %w", err)
	}

	return nil
}
