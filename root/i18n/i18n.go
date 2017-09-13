package i18n

import (
	"github.com/ethereal-go/ethereal/root/config"
	"github.com/jinzhu/gorm"
)

type FillLocale interface {
	EstablishConnection(*gorm.DB) FillLocale
	Add(StorageLocale)
	Merge(map[string]map[string]string, StorageLocale) FillLocale
}

type Locale interface {
	Get(conf config.Configurable, key string) interface{}
}

type StorageLocale struct {
	Structure map[string]map[string]string
}
