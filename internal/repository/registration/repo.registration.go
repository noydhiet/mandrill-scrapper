package registration

import (
	pkgStorage "github.com/noydhiet/mandrill-scrapper/internal/pkg/storage"
)

type RegistrationDB struct {
}

type Repository struct {
	storage pkgStorage.Storage
}

func NewRepository(storage pkgStorage.Storage) *Repository {
	return &Repository{
		storage: storage,
	}
}

func (s *Repository) StoreRegistrationDb(data map[string]interface{}) error {
	return nil
}
