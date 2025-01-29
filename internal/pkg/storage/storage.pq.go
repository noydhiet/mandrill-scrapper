package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type StorageDB struct {
	dbconn *sqlx.DB
}

func NewStorageDB(dsn string) (Storage, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &StorageDB{
		dbconn: db,
	}, nil
}

func (s *StorageDB) Store(model string, data map[string]interface{}) error {
	return errors.New("model not found")
}

func (s *StorageDB) StoreDB(sql string, data map[string]interface{}) error {
	_, err := s.dbconn.NamedExec(sql, data)
	if err != nil {
		return errors.Wrap(err, "error inserting data")
	}
	return nil
}
