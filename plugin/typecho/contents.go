package typecho

import (
	"time"

	"notion_blog/notion"
)

type ContentType string

const (
	ContentPublish ContentType = "publish"
	ContentWaiting ContentType = "waiting"
	ContentHidden  ContentType = "hidden"
	ContentPrivate ContentType = "private"
)

type Contents struct {
	Cid          uint   `gorm:"primaryKey;not null;autoIncrement"`
	Title        string `gorm:"size:150"`
	Slug         string `gorm:"size:150;unique"`
	Created      uint   `gorm:"default:0;index"`
	Modified     uint   `gorm:"default:0"`
	Text         string
	Order        uint   `gorm:"default:0"`
	AuthorId     uint   `gorm:"default:0;column:authorId"`
	Template     string `gorm:"size:32"`
	Type         string `gorm:"size:16;default:post"`
	Status       string `gorm:"size:16;default:publish"`
	Password     string `gorm:"size:32"`
	CommentsNum  uint   `gorm:"default:0;column:commentsNum"`
	AllowComment uint8  `gorm:"size:1;default:0;column:allowComment"`
	AllowPing    uint8  `gorm:"size:1;default:0;column:allowPing"`
	AllowFeed    uint8  `gorm:"size:1;default:0;column:allowFeed"`
	Parent       uint   `gorm:"default:0"`
}

func (c *Typecho) TransformContent(list *notion.List) *Contents {
	return &Contents{
		Title:        list.Property.Get("Name", notion.PropertyTypeTitle).(string),
		Slug:         list.Property.Get("Slug", notion.PropertyTypeRichText).(string),
		Created:      uint(list.Property.Get("Created time", notion.PropertyTypeRichText).(*time.Time).Unix()),
		Modified:     uint(list.Property.Get("Edited time", notion.PropertyTypeRichText).(*time.Time).Unix()),
		Text:         "<!--markdown-->\n" + list.Output(),
		AuthorId:     1,
		Type:         "post",
		Status:       string(ContentPublish),
		AllowComment: 0,
		AllowPing:    1,
		AllowFeed:    1,
	}
}
