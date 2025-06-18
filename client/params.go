package client

import (
	"encoding/json"
	"net/url"

	"fmt"
)

type RequestParams interface {
	String() string
}

type EmptyParams struct{}

func (d *EmptyParams) String() string {
	return ""
}

type DefaultBody struct{}

func (d *DefaultBody) String() string {
	body, _ := json.Marshal(d)

	return string(body)
}

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

func (q *BaseQuery) String() string {
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

	return query.Encode()
}

func (q *BaseQuery) Next(start string) {
	q.StartCursor = start
}
