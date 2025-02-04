package handler

import (
	"context"

	modelPatent "github.com/noydhiet/mandrill-scrapper/internal/repository/patent"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type repoPatentInterface interface {
	StorePatent(ctx context.Context, data modelPatent.PatentDB) error
	FindPatent(ctx context.Context, filter bson.M) ([]modelPatent.PatentDB, error)
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
