package patent

import (
	"context"
	"time"

	pkgStorage "github.com/noydhiet/mandrill-scrapper/internal/pkg/storage"
)

type PatentDB struct {
	CompanyName      string    `bson:"company_name"`
	PatentExpiryDate string    `bson:"patent_expiry_date"`
	PatentNumber     string    `bson:"patent_number"`
	Title            string    `bson:"title"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdateAt         time.Time `bson:"updated_at"`
}

type Repository struct {
	storage pkgStorage.Storage
}

func NewRepository(storage pkgStorage.Storage) *Repository {
	return &Repository{
		storage: storage,
	}
}

func (s *Repository) StorePatent(ctx context.Context, data PatentDB) error {
	if err := s.storage.Store(ctx, "patent", data); err != nil {
		return err
	}

	return nil
}

func (s *Repository) FindPatent(ctx context.Context) ([]PatentDB, error) {
	var data []PatentDB

	if err := s.storage.Find(ctx, "patent", nil, &data); err != nil {
		return nil, err
	}

	return data, nil
}
