package controller

import (
	"context"
	"fmt"

	"github.com/takokun778/2022/internal/domain/model/tag"
	"github.com/takokun778/2022/internal/usecase/port"
)

type Tag struct {
	port.NoticeTag
}

func NewTag(
	noticeTag port.NoticeTag,
) *Tag {
	return &Tag{
		NoticeTag: noticeTag,
	}
}

type CmdReq struct {
	Owner string
	Repo  string
}

func (t *Tag) Cmd(ctx context.Context, req CmdReq) error {
	owner := tag.Owner(req.Owner)

	repo := tag.Repo(req.Repo)

	input := port.NoticeTagInput{
		Owner: owner,
		Repo:  repo,
	}

	if _, err := t.NoticeTag.Execute(ctx, input); err != nil {
		return fmt.Errorf("failed to notice tag execute: %w", err)
	}

	return nil
}
