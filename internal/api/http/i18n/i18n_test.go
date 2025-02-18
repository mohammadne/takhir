package i18n_test

import (
	"testing"

	"github.com/mohammadne/zanbil/internal/api/http/i18n"
	"github.com/mohammadne/zanbil/internal/entities"
	"go.uber.org/zap"
)

func TestReader(t *testing.T) {
	i18n, err := i18n.New(zap.NewNop())
	if err != nil {
		t.Errorf("error while creating i18n %v", err)
	}

	tests := []struct {
		description string
		language    entities.Language
		key         string
		expected    string
	}{
		{
			description: "testing non exsisting locale",
			language:    "fr",
			key:         "non.existing.locale",
			expected:    "non existing locale",
		},
		{
			description: "testing non exsisting key in a valid locale",
			language:    entities.LanguageEnglish,
			key:         "non.existing.key",
			expected:    "non existing key",
		},
		{
			description: "testing an exsisting key in a valid locale",
			language:    entities.LanguageEnglish,
			key:         "categories.list.success",
			expected:    "Categories have been retrieved successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			translation := i18n.Translate(tt.key, tt.language)
			if translation != tt.expected {
				t.Errorf("Translate(%q, %q) = %q; want %q", tt.key, tt.language, translation, tt.expected)
			}
		})
	}
}
