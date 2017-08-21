package ethereal

type ErrSeedI18N interface {
	fill()
}

// list structure
type EnUS struct{}
type RuRU struct{}

func (en EnUS) fill() {
	app.I18n.Default("Success").T("en-US", "StatusOK")
}

func (ru RuRU) fill() {
	app.I18n.Default("Успешно").T("ru-RU", "StatusOK")
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
