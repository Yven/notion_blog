package client

import (
	"github.com/Yven/notion_blog/lib"
	"github.com/Yven/notion_blog/structure"
)

type Blocks struct {
	Endpoints
}

type AppendBlock struct {
	Children []*structure.Object `json:"children"`
	After    string              `json:"after,omitempty"`
	DefaultBody
}

// TODO
func (b *Blocks) Append(params AppendBlock) error {
	_, err := b.Client.NewNotion(lib.Patch, b.GetPath("children"), &params)
	return err
}

func (b *Blocks) Retrieve() (*structure.List, error) {
	return b.Client.NewNotion(lib.Get, b.GetPath(), &EmptyParams{})
}

func (b *Blocks) Children(parent *structure.Object, params BaseQuery) (*structure.List, error) {
	data, err := b.Client.NewNotion(lib.Get, b.GetPath("children"), &params)
	if err != nil {
		return nil, err
	}
	data.Property = parent

	for k, obj := range data.GetContent() {
		if *obj.HasChildren {
			child, err := b.Client.NewBlock(obj.Id).Children(obj, BaseQuery{})
			if err != nil {
				return nil, err
			}

			data.Results[k].AppendChild(child.GetContent())
		}
	}

	return data, nil
}

// TODO
// func (b *Blocks) Update(params RequestParams) error {
// 	_, err := b.Client.NewNotion(lib.Patch, b.GetPath(), params)
// 	return err
// }

func (b *Blocks) Delete() error {
	_, err := b.Client.NewNotion(lib.Delete, b.GetPath(), nil)
	return err
}
