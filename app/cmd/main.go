package main

import (
	"context"
	"log"

	"github.com/takokun778/2022/internal/adapter/controller"
	"github.com/takokun778/2022/internal/adapter/gateway"
	"github.com/takokun778/2022/internal/adapter/notifier"
	"github.com/takokun778/2022/internal/driver/config"
	"github.com/takokun778/2022/internal/driver/database"
	"github.com/takokun778/2022/internal/driver/slack"
	"github.com/takokun778/2022/internal/usecase/interactor"
)

func main() {
	config.Init()

	db := database.NewClient()

	rdb, err := db.Of(config.Get().DSN)
	if err != nil {
		log.Fatal(err)
	}

	tagGateway := gateway.NewTag(rdb)

	github := gateway.NewGithub()

	slack := slack.NewClient()

	ch, err := slack.Of(config.Get().SlackToken, config.Get().SlackChannelID)
	if err != nil {
		log.Fatal(err)
	}

	tagNotifier := notifier.NewTag(ch)

	tagNoticeInt := interactor.NewNoticeTag(github, tagGateway, tagNotifier)

	tagCtl := controller.NewTag(tagNoticeInt)

	if err := tagCtl.Cmd(context.Background(), controller.CmdReq{
		Owner: config.Get().GitHubOwner,
		Repo:  config.Get().GitHubRepo,
	}); err != nil {
		log.Fatal(err)
	}
}
