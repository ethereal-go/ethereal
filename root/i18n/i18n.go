package i18n

import "github.com/ethereal-go/ethereal/root/config"

type FillLocale interface {
	EstablishConnection(configuration interface{}) FillLocale
	Add(StorageLocale)
	Merge(map[string]map[string]string, StorageLocale) FillLocale
}

type Locale interface {
	Get(conf config.Configurable, key string) interface{}
}

type StorageLocale struct {
	Structure map[string]map[string]string
}
