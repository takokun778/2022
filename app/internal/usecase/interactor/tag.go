package interactor

import (
	"context"
	"fmt"
	"log"

	"github.com/takokun778/2022/internal/domain/external"
	"github.com/takokun778/2022/internal/domain/model"
	"github.com/takokun778/2022/internal/domain/repository"
	"github.com/takokun778/2022/internal/usecase/port"
)

type NoticeTag struct {
	github        repository.GitHub
	tagRepository repository.Tag
	tagExternal   external.Tag
}

var _ port.NoticeTag = (*NoticeTag)(nil)

func NewNoticeTag(
	github repository.GitHub,
	tagRepository repository.Tag,
	tagExternal external.Tag,
) *NoticeTag {
	return &NoticeTag{
		github:        github,
		tagRepository: tagRepository,
		tagExternal:   tagExternal,
	}
}

func (nt *NoticeTag) Execute(ctx context.Context, input port.NoticeTagInput) (port.NoticeTagOutput, error) {
	dst, err := nt.github.FindAll(ctx, input.Owner, input.Repo)
	if err != nil {
		return port.NoticeTagOutput{}, fmt.Errorf("failed to list tags: %w", err)
	}

	src, err := nt.tagRepository.FindAll(ctx, input.Owner, input.Repo)
	if err != nil {
		return port.NoticeTagOutput{}, fmt.Errorf("failed to list tags: %w", err)
	}

	diff := model.TakeTags(dst, src)

	if len(diff) == 0 {
		log.Printf("no tags changed\n")

		return port.NoticeTagOutput{}, nil
	}

	if err := nt.tagRepository.SaveAll(ctx, diff); err != nil {
		return port.NoticeTagOutput{}, fmt.Errorf("failed to save tags: %w", err)
	}

	for _, tag := range diff {
		url := fmt.Sprintf("https://github.com/%s/%s/releases/tag/%s", input.Owner, input.Repo, tag.Name)

		msg := fmt.Sprintf("released %s\n\n%s", tag.Name, url)

		if err := nt.tagExternal.Notice(ctx, msg); err != nil {
			return port.NoticeTagOutput{}, fmt.Errorf("failed to notice tag: %w", err)
		}
	}

	return port.NoticeTagOutput{}, nil
}
