package manufacture

import (
	pkgStorage "github.com/noydhiet/mandrill-scrapper/internal/pkg/storage"
)

type ManufactureDB struct {
}

type Repository struct {
	storage pkgStorage.Storage
}

func NewRepository(storage pkgStorage.Storage) *Repository {
	return &Repository{
		storage: storage,
	}
}

func (s *Repository) StoreManufactureDb(data map[string]interface{}) error {
	return nil
}
