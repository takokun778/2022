package notifier

import (
	"github.com/slack-go/slack"
)

type Channel struct {
	ID string
	*slack.Client
}

type ChannelFactory interface {
	Of(string, string) (*Channel, error)
}
