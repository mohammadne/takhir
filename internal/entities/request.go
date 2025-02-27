package entities

import (
	"strconv"
)

type Language string

const (
	LanguageEnglish Language = "en"
	LanguagePersian Language = "fa"

	LanguageDefault Language = LanguageEnglish
)

func ToLanguage(rawLanguage string) Language {
	switch rawLanguage {
	case string(LanguageEnglish):
		return LanguageEnglish
	case string(LanguagePersian):
		return LanguagePersian
	default:
		return LanguageDefault
	}
}

//

type ListOptions struct {
	PageSize   int
	LastId     int
	Page       int
	WithCounts bool
	Search     string
}

func QueryParamsToListOptions(queryParams map[string]string) *ListOptions {
	pageSizeString := queryParams["page_size"]
	lastIdString := queryParams["last_id"]
	pageString := queryParams["page"]
	withCountsString := queryParams["with_counts"]
	searchString := queryParams["search"]

	pageSize, _ := strconv.Atoi(pageSizeString)
	lastId, _ := strconv.Atoi(lastIdString)
	page, _ := strconv.Atoi(pageString)
	var withCounts bool
	if withCountsString == "1" {
		withCounts = true
	}

	return &ListOptions{
		PageSize:   pageSize,
		LastId:     lastId,
		Page:       page,
		WithCounts: withCounts,
		Search:     searchString,
	}
}

func (p *ListOptions) DoPaginate() bool {
	if p == nil {
		return false
	}
	return p.PageSize > 0
}

func (p *ListOptions) DoSearch() bool {
	if p == nil {
		return false
	}
	return p.Search != ""
}
