package ethereal

import (
	"github.com/qor/i18n"
	"strings"
)

//
type ErrSeedI18N interface {
	fill()
}

type SchemaI18n struct {
	structure map[string]map[string]string
}

func graphQL() (graphQL SchemaI18n) {
	graphQL = SchemaI18n{
		structure: map[string]map[string]string{
			"en-US": map[string]string{
				"graphQL.User.Description": "List of users of your application.",
			},
			"ru-RU": map[string]string{
				"graphQL.User.Description": "Список пользователей вашего приложения.",
			},
		},
	}
	return
}

func (schema SchemaI18n) fill() {
	for locale, mapValues := range schema.structure {
		for key, value := range mapValues {
			app.I18n.SaveTranslation(&i18n.Translation{Key: key, Locale: locale, Value: value})
		}
	}
}

// Fill storage
func SeedI18N() {
	graphQL().fill()
}

func mapLanguage() map[string]string {
	return map[string]string{
		"en-US": "en-US",
		"ru-RU": "ru-RU",
	}
}

// temp functions determines type locale
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
	return ""
}
