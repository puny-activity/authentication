package loclzr

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Localizer struct {
	bundle       *i18n.Bundle
	engLocalizer *i18n.Localizer
}

func New(bundle *i18n.Bundle) *Localizer {
	return &Localizer{
		bundle:       bundle,
		engLocalizer: i18n.NewLocalizer(bundle, "en_US"),
	}
}

func (l *Localizer) Localize(lang string, messageID string) (string, error) {
	localizeConfig := i18n.LocalizeConfig{
		MessageID: messageID,
	}

	langLocalizer := i18n.NewLocalizer(l.bundle, lang)
	localizedMessage, err := langLocalizer.Localize(&localizeConfig)
	if err != nil {
		return "", fmt.Errorf("failed to localize message with id=%s to lang=%s", messageID, lang)
	}

	return localizedMessage, nil
}

func (l *Localizer) English(messageID string) string {
	localizeConfig := i18n.LocalizeConfig{
		MessageID: messageID,
	}

	engMessage := l.engLocalizer.MustLocalize(&localizeConfig)

	return engMessage
}
