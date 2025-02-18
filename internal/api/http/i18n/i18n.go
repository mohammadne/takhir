package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mohammadne/takhir/internal/entities"
	"go.uber.org/zap"
)

//go:embed languages/*.json
var languages embed.FS

type I18N interface {
	Translate(key string, language entities.Language) string
}

type i18n struct {
	logger   *zap.Logger
	messages map[entities.Language]map[string]any
}

func New(logger *zap.Logger) (I18N, error) {
	i := &i18n{
		logger:   logger,
		messages: make(map[entities.Language]map[string]any),
	}

	files, err := languages.ReadDir("languages")
	if err != nil {
		return nil, fmt.Errorf("error reading languages directory, %v", err)
	}

	for _, file := range files {
		name := file.Name()
		if len(name) < 6 || name[len(name)-5:] != ".json" {
			continue
		}

		languageRaw := name[:len(name)-5]
		language := entities.ToLanguage(languageRaw)
		if languageRaw != string(language) {
			return nil, fmt.Errorf("invalid file language %s, %v", name, err)
		}

		data, err := languages.ReadFile("languages/" + name)
		if err != nil {
			return nil, fmt.Errorf("error reading language file %s, %v", name, err)
		}

		var messages map[string]any
		if err := json.Unmarshal(data, &messages); err != nil {
			return nil, fmt.Errorf("error parsing language file %s, %v", name, err)
		}

		i.messages[language] = messages
	}

	return i, nil
}

func (i *i18n) Translate(key string, language entities.Language) string {
	if translations, exists := i.messages[language]; exists {
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

	i.logger.Error("local not found", zap.String("locale", string(language)))
	return translateFromKey(key)
}

func translateFromKey(key string) string {
	return strings.ReplaceAll(key, ".", " ")
}
