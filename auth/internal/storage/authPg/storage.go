package authPg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func NewAuthStorage(
	host string,
	port string,
	dbName string,
	user string,
	password string,
	sslMode string,
) (authStorage *Storage, err error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbName, sslMode,
	)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return
	}

	return &Storage{db: pool}, nil
}

func (s *Storage) Close() {
	s.Close()
}
