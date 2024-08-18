package loclzr

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Localizer struct {
	localizers map[string]*i18n.Localizer
}

func New(bundle *i18n.Bundle, languages []language.Tag) *Localizer {
	localizers := make(map[string]*i18n.Localizer, len(languages))
	for i := range languages {
		localizers[languages[i].String()] = i18n.NewLocalizer(bundle, languages[i].String())
	}

	return &Localizer{
		localizers: localizers,
	}
}

func (l *Localizer) Localize(lang string, messageID string) string {
	localizeConfig := i18n.LocalizeConfig{
		MessageID: messageID,
	}

	localizer, ok := l.localizers[lang]
	if !ok {
		localizer = l.localizers[language.English.String()]
	}

	localizedMessage, err := localizer.Localize(&localizeConfig)
	if err != nil {
		localizer = l.localizers[language.English.String()]
		localizedMessage, err = localizer.Localize(&localizeConfig)
	}

	return localizedMessage
}
