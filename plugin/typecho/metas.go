package typecho

import (
	"notion_blog/notion"
)

type Metaype string

const (
	Tag      Metaype = "tag"
	Category Metaype = "category"
)

type Metas struct {
	Mid         uint   `gorm:"primaryKey;not null;autoIncrement"`
	Name        string `gorm:"size:150"`
	Slug        string `gorm:"size:150"`
	Type        string `gorm:"size:32;not null"`
	Description string `gorm:"size:150"`
	Count       uint   `gorm:"default:0"`
	Order       uint   `gorm:"default:0"`
	Parent      uint   `gorm:"default:0"`
}

func (c *Typecho) TransformMeta(list *notion.List) []*Metas {
	tagList := list.Property.Get("Tag", notion.PropertyTypeMultiSelect).([]string)
	category := list.Property.Get("Category", notion.PropertyTypeSelect).(string)

	var metaList []*Metas
	for _, tag := range tagList {
		metaList = append(metaList, &Metas{
			Name:  tag,
			Slug:  tag,
			Type:  string(Tag),
			Count: 0,
		})
	}
	metaList = append(metaList, &Metas{
		Name:  category,
		Slug:  category,
		Type:  string(Category),
		Count: 1,
	})

	return metaList
}
