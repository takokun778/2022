package external

import "context"

type Tag interface {
	Notice(context.Context, string) error
}
