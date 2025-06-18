package client

import (
	"github.com/Yven/notion_blog/filter"
	"github.com/Yven/notion_blog/lib"
	"github.com/Yven/notion_blog/structure"
)

type Pages struct {
	Endpoints
}

type CreatePage struct {
	Parent struct {
		PageId     string `json:"page_id,omitempty"`
		DatabaseId string `json:"database_id,omitempty"`
	} `json:"parent,omitempty"`
	Properties *filter.MetaData    `json:"properties,omitempty"`
	Children   []*structure.Object `json:"children,omitempty"`
	Icon       string              `json:"icon,omitempty"`
	Cover      *structure.File     `json:"cover,omitempty"`

	DefaultBody
}

// TODO
func (page *Pages) Create(params CreatePage) error {
	_, err := page.Client.NewNotion(lib.Post, page.GetPath(true), &params)
	return err
}

func (page *Pages) Retrieve() (*structure.List, error) {
	return page.Client.NewNotion(lib.Get, page.GetPath(), &EmptyParams{})
}

type UpdatePage struct {
	Properties *filter.MetaData `json:"properties,omitempty"`
	InTrash    bool             `json:"in_trash,omitempty"`
	Icon       string           `json:"icon,omitempty"`
	Cover      *structure.File  `json:"cover,omitempty"`

	DefaultBody
}

func (page *Pages) Update(params UpdatePage) error {
	_, err := page.Client.NewNotion(lib.Patch, page.GetPath(), &params)
	return err
}
