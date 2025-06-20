package client

import (
	"os"
	"testing"

	"github.com/Yven/notion_blog/filter"
	"github.com/joho/godotenv"
)

func TestNewClient(t *testing.T) {
	godotenv.Load(".env")
	key := os.Getenv("NOTION_KEY")
	DbId := os.Getenv("NOTION_DB_ID")

	notion := NewClient(key)
	list, err := notion.NewDb(DbId).Query(QueryDatabase{
		Filter: filter.Status("Status").Equal("waiting"),
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

		// 更新文章状态
		err = notion.NewPage(pageItem.Id).Update(UpdatePage{
			Properties: filter.Status("Status").Set("name", "publish"),
		})
		if err != nil {
			t.Log(err)
		}

		break
	}
}

func TestPages_Retrieve(t *testing.T) {
	godotenv.Load(".env")
	key := os.Getenv("NOTION_KEY")
	pageId := "47652d31-e88e-45a5-aa0c-91fce09062d4"

	notion := NewClient(key)
	err := notion.NewPage(pageId).Update(UpdatePage{
		Properties: filter.Status("Status").Set("name", "publish"),
	})
	if err != nil {
		t.Log(err)
	}
}
