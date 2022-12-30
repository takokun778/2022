package gateway

import (
	"context"
	"fmt"

	"github.com/takokun778/2022/internal/domain/model"
	"github.com/takokun778/2022/internal/domain/model/tag"
	"github.com/takokun778/2022/internal/domain/repository"
	"github.com/uptrace/bun"
)

var _ repository.Tag = (*Tag)(nil)

type Tag struct {
	rdb *RDB
}

type TagEntity struct {
	bun.BaseModel `bun:"table:tags"`
	Name          string `bun:"name,pk"`
	Repo          string `bun:"repo,pk"`
}

func NewTag(
	rdb *RDB,
) *Tag {
	return &Tag{
		rdb: rdb,
	}
}

func (t *Tag) SaveAll(ctx context.Context, models []model.Tag) error {
	entities := make([]TagEntity, len(models))

	for i, model := range models {
		entities[i] = TagEntity{
			Name: model.Name.String(),
			Repo: model.Repo.String(),
		}
	}

	if _, err := t.rdb.DB.NewInsert().
		Model(&entities).
		Exec(ctx); err != nil {
		return fmt.Errorf("failed to save tags: %w", err)
	}

	return nil
}

func (t *Tag) FindAll(ctx context.Context, owner tag.Owner, repo tag.Repo) ([]model.Tag, error) {
	var entities []TagEntity

	err := t.rdb.DB.NewSelect().
		Model(&entities).
		Where("repo = ?", fmt.Sprintf("%s/%s", owner, repo)).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find tags: %w", err)
	}

	tags := make([]model.Tag, len(entities))

	for i, entity := range entities {
		tags[i] = model.Tag{
			Name: tag.Name(entity.Name),
			Repo: tag.Repo(entity.Repo),
		}
	}

	return tags, nil
}

func (t *Tag) CreateTable(ctx context.Context) error {
	if _, err := t.rdb.DB.NewCreateTable().
		Model((*TagEntity)(nil)).
		IfNotExists().
		Exec(ctx); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

func (t *Tag) CreateIndex(ctx context.Context) error {
	if _, err := t.rdb.DB.NewCreateIndex().
		Model((*TagEntity)(nil)).
		Index("name_idx").
		Column("name").
		Exec(ctx); err != nil {
		return fmt.Errorf("failed to create name index: %w", err)
	}

	if _, err := t.rdb.DB.NewCreateIndex().
		Model((*TagEntity)(nil)).
		Index("repo_idx").
		Column("repo").
		Exec(ctx); err != nil {
		return fmt.Errorf("failed to create repo index: %w", err)
	}

	return nil
}
