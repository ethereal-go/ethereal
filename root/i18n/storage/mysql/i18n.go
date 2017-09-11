package mysql

import (
	localeI18n "github.com/ethereal-go/ethereal/root/i18n"
	"github.com/jinzhu/gorm"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"strings"
)

type LocaleStorageMysql struct {
	DB   *gorm.DB
	I18n *i18n.I18n
}

func (l LocaleStorageMysql) EstablishConnection(config interface{}) {
	user  := config.(map[string]string)["user"]
	password  := config.(map[string]string)["password"]
	dbName  := config.(map[string]string)["dbname"]

	db, _ := gorm.Open("mysql", user + ":" + password+ "@/" + dbName + "?charset=utf8&parseTime=True&loc=Local")
	l.I18n = i18n.New(database.New(db))
}

func (l LocaleStorageMysql) Add(i18n localeI18n.StorageLocale) {
	for locale, mapValues := range i18n.Structure {
		for key, value := range mapValues {
			l.I18n.SaveTranslation(&i18n.Translation{Key: key, Locale: locale, Value: value})
		}
	}
}

//
//func I18nGraphQL() (graphQL graphQLI18n) {
//	graphQL = graphQLI18n{
//		structure: map[string]map[string]string{
//			"en-US": map[string]string{
//				"graphQL.User.Description":  "List of users of your application.",
//				"graphQL.Role.Description":  "List of roles of your application.",
//				"graphQL.UserType.id":       "ID your user.",
//				"graphQL.UserType.email":    "Email your user.",
//				"graphQL.UserType.name":     "Name your user.",
//				"graphQL.UserType.password": "Hashed password your user.",
//				"graphQL.UserType.role":     "Concrete role your user.",
//
//				"graphQL.RoleType.id":           "Id role.",
//				"graphQL.RoleType.name":         "Name role.",
//				"graphQL.RoleType.display_name": "Display name role.",
//				"graphQL.RoleType.description":  "Description role.",
//
//				"jwtToken.ValidationErrorMalformed": "That's not even a token",
//				"jwtToken.ValidationErrorExpired":   "Timing is everything",
//				"jwtToken.ErrorBase":                "Couldn't handle this token ",
//			},
//			"ru-RU": map[string]string{
//				"graphQL.User.Description":  "Список пользователей вашего приложения.",
//				"graphQL.Role.Description":  "Список ролей вашего приложения.",
//				"graphQL.UserType.id":       "ID вашего пользователя.",
//				"graphQL.UserType.email":    "Email вашего пользователя.",
//				"graphQL.UserType.name":     "Имя вашего пользователя.",
//				"graphQL.UserType.password": "Захэшированный пароль",
//				"graphQL.UserType.role":     "Роль которой принадлежить пользователь",
//
//				"graphQL.RoleType.id":           "Id Роли.",
//				"graphQL.RoleType.name":         "Имя роли.",
//				"graphQL.RoleType.display_name": "Имя роли для отображения.",
//				"graphQL.RoleType.description":  "Подробное описание роли.",
//
//				"jwtToken.ValidationErrorMalformed": "Это не похоже на токен",
//				"jwtToken.ValidationErrorExpired":   "Время истекло",
//				"jwtToken.ErrorBase":                "Не удалось обрабоать этот токен ",
//			},
//		},
//	}
//	return
//}

/**
/ Function merge structure graph i!8n
*/
func (l LocaleStorageMysql) Merge(merge map[string]map[string]string, storage localeI18n.StorageLocale) localeI18n.FillLocale {
	for locale, mapValues := range merge {
		for key, value := range mapValues {
			if _, exist := storage.Structure[locale]; !exist {
				storage.Structure[locale] = map[string]string{
					key: value,
				}
			} else {
				storage.Structure[locale][key] = value
			}
		}
	}
	return l
}

// temp functions determines type locale from header http request
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
	return ""
}
