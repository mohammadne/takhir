package helpers

import (
	"net/url"
	"strconv"
)

type Pagination struct {
	PageSize   int
	LastId     int
	Page       int
	WithCounts bool
}

func (p *Pagination) Valid() bool {
	if p == nil {
		return false
	}
	return p.PageSize > 0
}

func (p *Pagination) ToQueryParams() string {
	params := url.Values{}
	params.Add("page_size", strconv.Itoa(p.PageSize))
	params.Add("last_id", strconv.Itoa(p.LastId))
	params.Add("page", strconv.Itoa(p.Page))
	params.Add("with_counts", strconv.FormatBool(p.WithCounts))
	return params.Encode()
}
