package repository

import (
	"context"

	"github.com/takokun778/2022/internal/domain/model"
	"github.com/takokun778/2022/internal/domain/model/tag"
)

type Tag interface {
	SaveAll(context.Context, []model.Tag) error
	FindAll(context.Context, tag.Owner, tag.Repo) ([]model.Tag, error)
}
