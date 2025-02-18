package i18n_test

import (
	"testing"

	"github.com/mohammadne/takhir/internal/api/http/i18n"
	"go.uber.org/zap"
)

func TestReader(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("error while creating development zap logger %v", err)
	}

	i18n, err := i18n.New(logger)
	if err != nil {
		t.Errorf("error while creating i18n %v", err)
	}

	tests := []struct {
		description string
		locale      string
		key         string
		expected    string
	}{
		// {
		// 	description: "testing non exsisting locale",
		// 	locale:      "fr",
		// 	key:         "non.existing.locale",
		// 	expected:    "non existing locale",
		// },
		// {
		// 	description: "testing non exsisting key in a valid locale",
		// 	locale:      "en",
		// 	key:         "non.existing.key",
		// 	expected:    "non existing key",
		// },
		{
			description: "testing an exsisting key in a valid locale",
			locale:      "en",
			key:         "categories.list.success",
			expected:    "Categories have been retrieved successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			translation := i18n.Translate(tt.key, tt.locale)
			if translation != tt.expected {
				t.Errorf("Translate(%q, %q) = %q; want %q", tt.key, tt.locale, translation, tt.expected)
			}
		})
	}
}
