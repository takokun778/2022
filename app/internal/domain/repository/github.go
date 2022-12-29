package repository

import (
	"context"

	"github.com/takokun778/2022/internal/domain/model"
	"github.com/takokun778/2022/internal/domain/model/tag"
)

type GitHub interface {
	FindAll(context.Context, tag.Owner, tag.Repo) ([]model.Tag, error)
}
