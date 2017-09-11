package i18n

type Locale interface {
	Fill()
	Merge(map[string]map[string]string) Locale
}


