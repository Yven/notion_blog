package client

import (
	"encoding/json"

	"github.com/Yven/notion_blog/filter"
	"github.com/Yven/notion_blog/lib"
	"github.com/Yven/notion_blog/structure"
)

type Databases struct {
	Endpoints
}

// TODO
// func (db *Databases) Create(params RequestParams) error {
// 	_, err := db.Client.NewNotion(lib.Post, db.GetPath(true), params)
// 	return err
// }

type QueryDatabase struct {
	StartCursor string           `json:"start_cursor,omitempty"`
	PageSize    int              `json:"page_size,omitempty"`
	Filter      *filter.MetaData `json:"filter,omitempty"`
	Sort        []*filter.Sort   `json:"sorts,omitempty"`

	DefaultNext
}

func (d *QueryDatabase) String() string {
	if d.PageSize == 0 {
		d.PageSize = 100
	}

	body, _ := json.Marshal(d)

	return string(body)
}

func (db *Databases) Query(params QueryDatabase) (*structure.List, error) {
	return db.Client.NewNotion(lib.Post, db.GetPath("query"), &params)
}

func (db *Databases) Retrieve() (*structure.List, error) {
	return db.Client.NewNotion(lib.Get, db.GetPath(), &EmptyParams{})
}

// TODO
// func (db *Databases) Update(params RequestParams) error {
// 	_, err := db.Client.NewNotion(lib.Patch, db.GetPath(), params)
// 	return err
// }
