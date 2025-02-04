package patent

import (
	"time"

	pkgStorage "github.com/noydhiet/mandrill-scrapper/internal/pkg/storage"
)

type PatentDB struct {
	CompanyName      string    `db:"company_name"`
	PatentExpiryDate string    `db:"patent_expiry_date"`
	PatentNumber     string    `db:"patent_number"`
	Title            string    `db:"title"`
	CreatedAt        time.Time `db:"created_at"`
	UpdateAt         time.Time `db:"updated_at"`
}

type Repository struct {
	storage pkgStorage.Storage
}

func NewRepository(storage pkgStorage.Storage) *Repository {
	return &Repository{
		storage: storage,
	}
}

func (s *Repository) StorePatentDb(data map[string]interface{}) error {
	// sql := `
	// 	INSERT INTO patents (
	// 		company_name,
	// 		patent_expiry_date,
	// 		patent_number,
	// 		title,
	// 		created_at
	// 	) VALUES (
	// 		:company_name,
	// 		:patent_expiry_date,
	// 		:patent_number,
	// 		:title,
	// 		:created_at
	// 	)
	// `

	// if err := s.storage.StoreDB(sql, data); err != nil {
	// 	return errors.Wrap(err, "error inserting data patent")
	// }

	return nil
}
