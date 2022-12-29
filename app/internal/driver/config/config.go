package config

import "os"

type Config struct {
	DSN            string
	GitHubOwner    string
	GitHubRepo     string
	SlackToken     string
	SlackChannelID string
}

var config Config //nolint:gochecknoglobals

func Init() {
	config = Config{
		DSN:            os.Getenv("DATABASE_URL"),
		GitHubOwner:    os.Getenv("GITHUB_OWNER"),
		GitHubRepo:     os.Getenv("GITHUB_REPOSITORY"),
		SlackToken:     os.Getenv("SLACK_TOKEN"),
		SlackChannelID: os.Getenv("SLACK_CHANNEL_ID"),
	}
}

func Get() Config {
	return config
}
