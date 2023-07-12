package store

import (
	"context"
	"database/sql"
	"github.com/dbubel/manifold/pkg/logging"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"time"
)

type Postgres struct {
	DbReader *sqlx.DB
	log      *logging.Logger
}

func NewPostgresStore(dsn string, log *logging.Logger) (*Postgres, error) {
	var s Postgres
	var dbReader *sqlx.DB
	var errReader error

	dbReader, errReader = sqlx.Connect("pgx", dsn)
	if errReader != nil {
		return &s, errReader
	}

	dbReader.SetMaxOpenConns(32)

	s.DbReader = dbReader
	s.log = log

	if err := s.HealthCheck(context.WithTimeout(context.Background(), time.Second)); err != nil {
		return &s, err
	}

	return &s, nil
}

func (s *Postgres) HealthCheck(ctx context.Context, fn context.CancelFunc) error {
	defer fn()
	return s.DbReader.PingContext(ctx)
}

func (s *Postgres) Insert(ctx context.Context, topic string, data []byte) (string, error) {
	var shardID string
	tx, err := s.DbReader.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return shardID, err
	}

	_, err = tx.Exec("update kinesis_shards set in_use = true,last_updated=now() where shard_id=$1", shardID)

	if err != nil {
		tx.Rollback()
		return shardID, err
	}

	tx.Commit()
	return shardID, nil
}
