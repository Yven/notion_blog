package client

import (
	"testing"

	"github.com/Yven/notion_blog/filter"
)

func TestNewClient(t *testing.T) {
	key := "secret_dwk0kAwUXTGXXXevAH6cJ2fvwh8pgHhLwuQLvo5hylc"
	DbId := "1bf0172cf38b46ad9d6b1eb43d9b334d"

	notion := NewClient(key)
	list, err := notion.NewDb(DbId).Query(QueryDatabase{
		Filter: filter.Status("Status").Equal("edit"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(list)
	t.Log(len(list.Results))

	for _, pageItem := range list.GetContent() {
		page, err := notion.NewBlock(pageItem.Id).Children(pageItem, BaseQuery{})
		if err != nil {
			t.Fatal(err)
		}
		t.Log(page.Property.Get("Name"))
		t.Log(page.Property.Get("Created time"))
		t.Log(page.ToMarkdown())

		break
	}
}
