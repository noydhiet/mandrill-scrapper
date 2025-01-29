package lawsuit

import (
	pkgStorage "github.com/suryatresna/srg-radar-project/internal/pkg/storage"
)

type LawsuitDB struct {
}

type Repository struct {
	storage pkgStorage.Storage
}

func NewRepository(storage pkgStorage.Storage) *Repository {
	return &Repository{
		storage: storage,
	}
}

func (s *Repository) StoreLawsuitDb(data map[string]interface{}) error {
	return nil
}
