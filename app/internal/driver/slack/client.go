package slack

import (
	"github.com/slack-go/slack"
	"github.com/takokun778/2022/internal/adapter/notifier"
)

var _ notifier.ChannelFactory = (*Client)(nil)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c Client) Of(token string, id string) (*notifier.Channel, error) {
	return &notifier.Channel{
		ID:     id,
		Client: slack.New(token),
	}, nil
}
