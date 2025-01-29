package handler

type repoPatentInterface interface {
	StorePatentDb(data map[string]interface{}) error
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
