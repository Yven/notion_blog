package structure

import (
	"fmt"
	"strings"
)

type BlockType string

const (
	BolckTypeBookmark         BlockType = "bookmark"
	BlockTypeBreadcrumb       BlockType = "breadcrumb"
	BolckTypeBulletedListItem BlockType = "bulleted_list_item"
	BolckTypeCallout          BlockType = "callout"
	BolckTypeChildDatabase    BlockType = "child_database"
	BolckTypeChildPage        BlockType = "child_page"
	BolckTypeCode             BlockType = "code"
	BlockTypeColumn           BlockType = "column"
	BlockTypeColumnList       BlockType = "column_list"
	BlockTypeDivider          BlockType = "divider"
	BolckTypeEmbed            BlockType = "embed"
	BolckTypeEquation         BlockType = "equation"
	BolckTypeFile             BlockType = "file"
	BolckTypeHeading1         BlockType = "heading_1"
	BolckTypeHeading2         BlockType = "heading_2"
	BolckTypeHeading3         BlockType = "heading_3"
	BolckTypeImage            BlockType = "image"
	BolckTypeLinkPreview      BlockType = "link_preview"
	BolckTypeNumberedListItem BlockType = "numbered_list_item"
	BolckTypeParagraph        BlockType = "paragraph"
	BolckTypePdf              BlockType = "pdf"
	BolckTypeQuote            BlockType = "quote"
	BolckTypeSyncedBlock      BlockType = "synced_block"
	BolckTypeTable            BlockType = "table"
	BolckTypeTableOfContents  BlockType = "table_of_contents"
	BolckTypeTableRow         BlockType = "table_row"
	BolckTypeToDo             BlockType = "to_do"
	BolckTypeToggle           BlockType = "toggle"
	BlcokTypeUnsupported      BlockType = "unsupported"
	BolckTypeVideo            BlockType = "video"
)

type ChildDatabase struct {
	Title string `json:"title"`
}

type ChildPage struct {
	Title string `json:"title"`
}

type BlockObject struct {
	Type        BlockType `json:"type,omitempty"`
	HasChildren *bool     `json:"has_children,omitempty"`
	Children    []*Object `json:"children,omitempty"`

	Bookmark         *Bookmark         `json:"bookmark,omitempty"`
	BulletedListItem *BulletedListItem `json:"bulleted_list_item,omitempty"`
	Callout          *Callout          `json:"callout,omitempty"`
	ChildDatabase    *ChildDatabase    `json:"child_database,omitempty"`
	ChildPage        *ChildPage        `json:"child_page,omitempty"`
	Code             *Code             `json:"code,omitempty"`
	Embed            *Embed            `json:"embed,omitempty"`
	Equation         *Equation         `json:"equation,omitempty"`
	File             *File             `json:"file,omitempty"`
	Heading1         *Heading1         `json:"heading_1,omitempty"`
	Heading2         *Heading2         `json:"heading_2,omitempty"`
	Heading3         *Heading3         `json:"heading_3,omitempty"`
	Image            *Image            `json:"image,omitempty"`
	LinkPreview      *LinkPreview      `json:"link_preview,omitempty"`
	NumberedListItem *NumberedListItem `json:"numbered_list_item,omitempty"`
	Paragraph        *Paragraph        `json:"paragraph,omitempty"`
	Pdf              *Pdf              `json:"pdf,omitempty"`
	Quote            *Quote            `json:"quote,omitempty"`
	SyncedBlock      *SyncedBlock      `json:"synced_block,omitempty"`
	Table            *Table            `json:"table,omitempty"`
	TableOfContents  *TableOfContents  `json:"table_of_contents,omitempty"`
	TableRow         *TableRow         `json:"table_row,omitempty"`
	ToDo             *ToDo             `json:"to_do,omitempty"`
	Toggle           *Toggle           `json:"toggle,omitempty"`
	Video            *Video            `json:"video,omitempty"`
}

func (obj *BlockObject) Output() string {
	entity := obj.GetDataObj()

	var content string
	switch obj.Type {
	case BolckTypeBulletedListItem:
		content = "* {{content}}"
	case BolckTypeCallout:
		content = "<aside>{{status}}{{content}}</aside>"
	case BolckTypeCode:
		content = "```{{status}}\n{{content}}\n```"
	case BolckTypeEquation:
		content = "$${{content}}$$"
	case BolckTypeEmbed, BolckTypeBookmark, BolckTypeLinkPreview, BolckTypeFile, BolckTypePdf, BolckTypeVideo:
		content = "[{{caption}}]({{file}})"
	case BolckTypeImage:
		// content = "![{{caption}}]({{file}})"
		content = `<img alt="{{caption}}" src="{{file}}" />`
	case BolckTypeHeading1:
		content = "# {{content}}"
		if len(entity.(ContentStatus).GetStatus()) != 0 {
			content = `<details><summary style="font-size: 2em;font-weight: bold;margin-bottom: 10px;">{{content}}</summary>{{child}}</details>`
		}
	case BolckTypeHeading2:
		content = "## {{content}}"
		if len(entity.(ContentStatus).GetStatus()) != 0 {
			content = `<details><summary style="font-size: 1.5em;font-weight: bold;margin-bottom: 8px;">{{content}}</summary>{{child}}</details>`
		}
	case BolckTypeHeading3:
		content = "### {{content}}"
		if len(entity.(ContentStatus).GetStatus()) != 0 {
			content = `<details><summary style="font-size: 1.2em;font-weight: bold;margin-bottom: 6px;">{{content}}</summary>{{child}}</details>`
		}
	case BolckTypeNumberedListItem:
		content = "1. {{content}}"
	case BolckTypeQuote:
		content = "> {{content}}"
	case BolckTypeToDo:
		content = "- [{{status}}] {{content}}"
	case BolckTypeToggle:
		content = "<details><summary>{{content}}</summary>{{child}}</details>"
	case BlockTypeDivider:
		content = "---"
	case BlockTypeColumnList:
		content = "<div style=\"display:flex;\">{{child}}\n</div>"
	case BlockTypeColumn:
		content = "<div style=\"flex:1;\">{{child}}\n</div>"
	case BolckTypeTableRow:
		content = "|{{content}}|"
	case BolckTypeTableOfContents:
		content = "[TOC]"
	case BolckTypeParagraph:
		content = "{{content}}"
	case BolckTypeTable, BolckTypeChildDatabase, BolckTypeChildPage, BolckTypeSyncedBlock, BlockTypeBreadcrumb, BlcokTypeUnsupported:
		content = ""
	default:
		content = ""
	}

	if en, ok := entity.(ContentTableRow); ok {
		content = strings.Replace(content, "{{content}}", en.GetTableRow(), 1)
	}
	if en, ok := entity.(Content); ok {
		content = strings.Replace(content, "{{content}}", en.GetContent(), 1)
		if obj.Type == BolckTypeCallout {
			content = strings.Replace(content, "\n", "\n<br />\n", -1)
		}
		if obj.Type == BolckTypeQuote {
			content = strings.Replace(content, "\n", "\n>\n> ", -1)
		}
	}
	if en, ok := entity.(ContentStatus); ok {
		content = strings.Replace(content, "{{status}}", en.GetStatus(), 1)
	}
	if en, ok := entity.(ContentCaption); ok {
		content = strings.Replace(content, "{{caption}}", en.GetCaption(), 1)
	}
	if en, ok := entity.(ContentFile); ok {
		content = strings.Replace(content, "{{file}}", en.GetFile(), 1)
	}
	if *obj.HasChildren {
		if strings.Contains(content, "{{child}}") {
			content = strings.Replace(content, "{{child}}", obj.GetChild(), 1)
		} else {
			content = fmt.Sprintf("%s%s", content, obj.GetChild())
		}
	}

	return content
}

func (obj *BlockObject) GetDataObj() interface{} {
	var entity interface{}
	switch obj.Type {
	case BolckTypeBulletedListItem:
		entity = obj.BulletedListItem
	case BolckTypeCallout:
		entity = obj.Callout
	case BolckTypeCode:
		entity = obj.Code
	case BolckTypeEquation:
		entity = obj.Equation
	case BolckTypeEmbed:
		entity = obj.Embed
	case BolckTypeBookmark:
		entity = obj.Bookmark
	case BolckTypeLinkPreview:
		entity = obj.LinkPreview
	case BolckTypeFile:
		entity = obj.File
	case BolckTypePdf:
		entity = obj.Pdf
	case BolckTypeVideo:
		entity = obj.Video
	case BolckTypeHeading1:
		entity = obj.Heading1
	case BolckTypeHeading2:
		entity = obj.Heading2
	case BolckTypeHeading3:
		entity = obj.Heading3
	case BolckTypeImage:
		entity = obj.Image
	case BolckTypeNumberedListItem:
		entity = obj.NumberedListItem
	case BolckTypeParagraph:
		entity = obj.Paragraph
	case BolckTypeQuote:
		entity = obj.Quote
	case BolckTypeToDo:
		entity = obj.ToDo
	case BolckTypeToggle:
		entity = obj.Toggle
	case BolckTypeTable:
		entity = obj.Table
	case BolckTypeTableOfContents:
		entity = obj.TableOfContents
	case BolckTypeTableRow:
		entity = obj.TableRow
	case BlockTypeDivider, BolckTypeChildDatabase, BolckTypeChildPage, BlockTypeColumn, BlockTypeColumnList, BolckTypeSyncedBlock, BlockTypeBreadcrumb, BlcokTypeUnsupported:
		entity = nil
	default:
		entity = nil
	}

	return entity
}

func (obj *BlockObject) AppendChild(other []*Object) {
	obj.Children = other
}

func (obj *BlockObject) GetChild() string {
	childContent := ""
	if obj.Children == nil {
		return childContent
	}
	for k, ch := range obj.Children {
		if obj.Type == BolckTypeParagraph {
			childContent = fmt.Sprintf("%s\n<br />\n&emsp;%s", childContent, strings.Replace(ch.Output(), "<br />\n", "<br />\n&emsp;", -1))
		} else if obj.Type == BolckTypeQuote {
			childContent = fmt.Sprintf("%s\n>\t%s", childContent, strings.Replace(ch.Output(), "\n", "\n>\t", -1))
		} else if obj.Type == BlockTypeColumn || obj.Type == BlockTypeColumnList {
			childContent = fmt.Sprintf("%s\n%s", childContent, ch.Output())
		} else if obj.Type == BolckTypeTable {
			var breakLine string
			if k == 1 {
				breakLine = fmt.Sprintf("\n%s|\n", strings.Repeat("|-", obj.GetDataObj().(*Table).TableWidth))
			} else {
				breakLine = "\n"
			}
			childContent = fmt.Sprintf("%s%s%s", childContent, breakLine, ch.Output())
		} else if obj.Type == BolckTypeToggle || obj.Type == BolckTypeHeading1 || obj.Type == BolckTypeHeading2 || obj.Type == BolckTypeHeading3 {
			childContent = fmt.Sprintf("%s%s", childContent, ch.Output())
		} else {
			childContent = fmt.Sprintf("%s\n\t%s", childContent, strings.Replace(ch.Output(), "\n", "\n\t", -1))
		}
	}

	return childContent
}

func (obj *BlockObject) GetBreakLink() string {
	var content string
	switch obj.Type {
	case BolckTypeBulletedListItem, BolckTypeNumberedListItem, BolckTypeToDo:
		content = "\n"
	case BolckTypeQuote, BolckTypeCallout, BolckTypeParagraph, BolckTypeEquation, BolckTypeCode, BolckTypeEmbed, BolckTypeBookmark, BolckTypeLinkPreview, BolckTypeFile, BolckTypePdf, BolckTypeVideo, BolckTypeHeading1, BolckTypeHeading2, BolckTypeHeading3, BolckTypeImage, BolckTypeToggle, BlockTypeDivider, BolckTypeChildDatabase, BolckTypeChildPage, BolckTypeTable, BolckTypeTableOfContents, BolckTypeTableRow, BlockTypeColumn, BlockTypeColumnList, BolckTypeSyncedBlock, BlockTypeBreadcrumb, BlcokTypeUnsupported:
		content = "\n\n"
	default:
		content = "\n\n"
	}

	return content
}
