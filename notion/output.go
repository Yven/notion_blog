package notion

import (
	"fmt"
	"log"
)

type OutputType string

const (
	Markdown OutputType = "markdown"
	Html     OutputType = "html"
)

// 设置文章输出格式
func (list *List) SetOutputType(outputType string) {
	switch outputType {
	case "md":
		list.OutputType = Markdown
	case "html":
		list.OutputType = Html
	default:
		log.Panicln("输出格式设置错误")
	}
}

func (list *List) Output() string {
	if list.Type != "block" {
		return ""
	}

	var content string
	switch list.OutputType {
	case Markdown:
		content = list.ToMarkdown()
	case Html:
		content = list.ToHtml()
	default:
		log.Panicf("输出格式错误：%s", list.OutputType)
	}

	return content
}

func (list *List) ToMarkdown() string {
	var content string
	for _, obj := range list.GetContent() {
		content = fmt.Sprintf("%s%s%s", content, obj.GetBreakLink(), obj.Output())
	}

	return content
}

func (list *List) ToHtml() string {
	return ""
}
