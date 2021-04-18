package types

// IKeywords -- интефейс к ключевым словам языка
type IKeywords interface {
	// IsKey -- проверяет обраец с ключевым словом
	IsKey(sample, key string) bool
}
