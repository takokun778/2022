package gateway

import (
	"github.com/uptrace/bun"
)

type RDB struct {
	*bun.DB
}

type RDBFactory interface {
	Of(dsn string) (*RDB, error)
}
