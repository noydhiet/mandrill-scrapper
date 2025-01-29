package storage

type Storage interface {
	Store(model string, data map[string]interface{}) error
	StoreDB(sql string, data map[string]interface{}) error
}
