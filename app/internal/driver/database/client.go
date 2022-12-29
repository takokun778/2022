package database

import (
	"database/sql"
	"fmt"

	// postgres driver.
	_ "github.com/lib/pq"
	"github.com/takokun778/2022/internal/adapter/gateway"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var _ gateway.RDBFactory = (*Client)(nil)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c Client) Of(dsn string) (*gateway.RDB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &gateway.RDB{
		DB: bun.NewDB(db, pgdialect.New()),
	}, nil
}
