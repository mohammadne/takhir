package entities

import (
	"net/url"
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

func QueryParamsToListOptions(queryParams url.Values) *ListOptions {
	pageSizeString := queryParams.Get("page_size")
	lastIdString := queryParams.Get("last_id")
	pageString := queryParams.Get("page")
	withCountsString := queryParams.Get("with_counts")
	searchString := queryParams.Get("search")

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
