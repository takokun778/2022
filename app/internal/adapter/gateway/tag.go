package gateway

import (
	"context"
	"fmt"

	"github.com/takokun778/2022/internal/domain/model"
	"github.com/takokun778/2022/internal/domain/model/tag"
	"github.com/takokun778/2022/internal/domain/repository"
)

var _ repository.Tag = (*Tag)(nil)

type Tag struct {
	rdb *RDB
}

type TagEntity struct {
	Name string `bun:"name,pk"`
	Repo string `bun:"repo,pk"`
}

func NewTag(
	rdb *RDB,
) *Tag {
	return &Tag{
		rdb: rdb,
	}
}

func (t *Tag) SaveAll(ctx context.Context, models []model.Tag) error {
	fmt.Printf("Saving %d tags\n", len(models))
	return nil
}

func (t *Tag) FindAll(ctx context.Context, repo tag.Repo) ([]model.Tag, error) {
	return []model.Tag{}, nil
}
