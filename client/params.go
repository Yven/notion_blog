package client

import (
	"net/url"

	"fmt"
)

type ListParams interface {
	Next(start string)
}

type DefaultNext struct {
	StartCursor string `json:"start_cursor,omitempty"`
}

func (d *DefaultNext) Next(start string) {
	d.StartCursor = start
}

// ---------------------------------------------- //

type BaseQuery struct {
	StartCursor string
	PageSize    int
}

func (q *BaseQuery) ToUrlValues() url.Values {
	if q.PageSize == 0 {
		q.PageSize = 100
	}

	query := url.Values{}
	if q.PageSize != 0 {
		query.Add("page_size", fmt.Sprint(q.PageSize))
	}
	if q.StartCursor != "" {
		query.Add("start_cursor", q.StartCursor)
	}

	return query
}

func (q *BaseQuery) Next(start string) {
	q.StartCursor = start
}
