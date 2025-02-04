package handler

import (
	"context"

	modelPatent "github.com/noydhiet/mandrill-scrapper/internal/repository/patent"
)

type repoPatentInterface interface {
	StorePatent(ctx context.Context, data modelPatent.PatentDB) error
	FindPatent(ctx context.Context) ([]modelPatent.PatentDB, error)
}

type repoLawsuitInterface interface {
	StoreLawsuitDb(data map[string]interface{}) error
}

type repoManufacturerInterface interface {
	StoreManufactureDb(data map[string]interface{}) error
}

type repoRecallInterface interface {
	StoreRecallDb(data map[string]interface{}) error
}

type repoRegistrationInterface interface {
	StoreRegistrationDb(data map[string]interface{}) error
}
