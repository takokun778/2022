package gateway

import (
	"context"
	"fmt"

	"github.com/google/go-github/v47/github"
	"github.com/takokun778/2022/internal/domain/model"
	"github.com/takokun778/2022/internal/domain/model/tag"
	"github.com/takokun778/2022/internal/domain/repository"
)

var _ repository.GitHub = (*GitHub)(nil)

const (
	maxPage   = 100
	startPage = 1
)

type GitHub struct {
	client *github.Client
}

func NewGithub() *GitHub {
	return &GitHub{
		client: github.NewClient(nil),
	}
}

func (g *GitHub) FindAll(ctx context.Context, owner tag.Owner, repo tag.Repo) ([]model.Tag, error) {
	var res []model.Tag

	opts := &github.ListOptions{
		PerPage: maxPage,
		Page:    startPage,
	}

	for {
		tags, _, err := g.client.Repositories.ListTags(ctx, owner.String(), repo.String(), opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list tags: %w", err)
		}

		for _, t := range tags {
			res = append(res, model.Tag{
				Name: tag.Name(*t.Name),
				Repo: tag.Repo(fmt.Sprintf("%s/%s", owner, repo)),
			})
		}

		if len(tags) < maxPage {
			break
		}

		opts.Page++
	}

	return res, nil
}
