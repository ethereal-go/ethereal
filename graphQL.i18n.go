package ethereal

import (
	"fmt"
	"github.com/qor/i18n"
	"strings"
)

//
type ErrSeedI18N interface {
	fill()
}

// list structure
type EnUS struct{}
type RuRU struct{}

func (en EnUS) fill() {
	app.I18n.SaveTranslation(&i18n.Translation{Key: "graphQL.User.Description", Locale: "en-US", Value: "List of users of your application."})
}

func (ru RuRU) fill() {
	app.I18n.SaveTranslation(&i18n.Translation{Key: "graphQL.User.Description", Locale: "ru-RU", Value: "Список пользователей вашего приложения."})
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

func mapLanguage() map[string]string {
	return map[string]string{
		"en-US": "en-US",
		"ru-RU": "ru-RU",
	}
}

func parserLocale(header []string) string {
	splitSign := strings.Split(header[0], ";")
	for _, multiLocal := range splitSign {
		local := strings.Split(multiLocal, ",")
		for _, possible := range local {
			if _, isExist := mapLanguage()[possible]; isExist {
				return mapLanguage()[possible]
			}
		}
	}
}
