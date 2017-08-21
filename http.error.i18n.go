package ethereal

import "github.com/qor/i18n"

type ErrSeedI18N interface {
	fill()
}

// list structure
type EnUS struct{}
type RuRU struct{}

func (en EnUS) fill() {
	app.I18n.SaveTranslation(&i18n.Translation{Key: "StatusOK", Locale: "en-US", Value: "Success"})
}

func (ru RuRU) fill() {
	app.I18n.SaveTranslation(&i18n.Translation{Key: "StatusOK", Locale: "ru-RU", Value: "Успешно"})
}

func getTypesLanguage() []ErrSeedI18N {
	return []ErrSeedI18N{EnUS{}, RuRU{}}
}

// Fill storage
func SeedI18N() {
	for _, typeI18n := range getTypesLanguage() {
		typeI18n.fill()
	}
}
