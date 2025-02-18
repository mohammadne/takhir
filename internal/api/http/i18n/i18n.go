package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

//go:embed locales/*.json
var locales embed.FS

type I18N interface {
	Translate(message, locale string) string
}

type i18n struct {
	logger   *zap.Logger
	messages map[string]map[string]any
}

func New(logger *zap.Logger) (I18N, error) {
	i := &i18n{
		logger:   logger,
		messages: make(map[string]map[string]any),
	}

	files, err := locales.ReadDir("locales")
	if err != nil {
		return nil, fmt.Errorf("error reading locales directory, %v", err)
	}

	for _, file := range files {
		name := file.Name()
		if len(name) < 6 || name[len(name)-5:] != ".json" {
			continue
		}

		locale := name[:len(name)-5]
		data, err := locales.ReadFile("locales/" + name)
		if err != nil {
			return nil, fmt.Errorf("error reading locale file %s, %v", name, err)
		}

		var messages map[string]any
		if err := json.Unmarshal(data, &messages); err != nil {
			return nil, fmt.Errorf("error parsing locale file %s, %v", name, err)
		}

		i.messages[locale] = messages
	}

	return i, nil
}

func (i *i18n) Translate(key, locale string) string {
	if translations, exists := i.messages[locale]; exists {
		keyParts := strings.Split(key, ".")
		var current any = translations

		for index := 0; index <= len(keyParts); index++ {
			lastElement := index == len(keyParts)

			switch current1 := current.(type) {
			case string:
				if !lastElement {
					break
				}
				return current1

			case map[string]any:
				if lastElement {
					break
				}

				if translation, found := current1[keyParts[index]]; found {
					current = translation
					continue
				} else {
					break
				}
			}
		}

		i.logger.Error("key not found", zap.String("key", key))
		return translateFromKey(key)
	}

	i.logger.Error("local not found", zap.String("locale", locale))
	return translateFromKey(key)
}

func translateFromKey(key string) string {
	return strings.ReplaceAll(key, ".", " ")
}
