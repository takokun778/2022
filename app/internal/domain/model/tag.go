package model

import "github.com/takokun778/2022/internal/domain/model/tag"

type Tag struct {
	Repo tag.Repo
	Name tag.Name
}

func TakeTags(from, target []Tag) []Tag {
	result := from

	for _, i := range target {
		list := make([]Tag, 0)

		for _, j := range result {
			if i != j {
				list = append(list, j)
			}
		}

		result = list
	}

	return result
}
