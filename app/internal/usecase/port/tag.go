package port

import (
	"github.com/takokun778/2022/internal/domain/model/tag"
	"github.com/takokun778/2022/internal/usecase"
)

type NoticeTagInput struct {
	usecase.Input
	Owner tag.Owner
	Repo  tag.Repo
}

type NoticeTagOutput struct {
	usecase.Output
}

type NoticeTag interface {
	usecase.Usecase[NoticeTagInput, NoticeTagOutput]
}
