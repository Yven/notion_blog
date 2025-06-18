package structure

import (
	"fmt"
	"strings"
)

type FileType string

const (
	FileTypeFile     FileType = "file"
	FileTypeExternal FileType = "external"
)

/**
 * ------------------------------------------------------
 */

type Content interface {
	GetContent() string
	BuildColor(content string) string
}

type TextBaseInfo struct {
	RichText []RichText      `json:"rich_text"`
	Color    AnnotationColor `json:"color"`
}

func (obj *TextBaseInfo) GetContent() string {
	var textContent string
	for _, text := range obj.RichText {
		textContent += text.Output()
	}
	if len(obj.RichText) == 0 {
		textContent = "<br />"
	}

	return obj.BuildColor(textContent)
}

func (obj *TextBaseInfo) BuildColor(content string) string {
	if obj.Color == Default {
		return content
	} else {
		return fmt.Sprintf("<font %s>%s</font>", obj.Color.Output(), content)
	}
}

type ContentCaption interface {
	GetCaption() string
}

type Caption struct {
	Caption []RichText `json:"caption"`
}

func (obj *Caption) GetCaption() string {
	var textContent string
	for _, text := range obj.Caption {
		textContent += text.NormalOutput()
	}

	return textContent
}

type ContentStatus interface {
	GetStatus() string
}

type ContentFile interface {
	GetFile() string
}

type ContentTableRow interface {
	GetTableRow() string
}

func (table *TableRow) GetTableRow() string {
	var content []string
	for _, row := range table.Cells {
		textContent := ""
		for _, txtEle := range row {
			textContent += txtEle.Output()
		}
		content = append(content, textContent)
	}

	return strings.Join(content, "|")
}

/**
 * ------------------------------------------------------
 */

type Bookmark struct {
	Caption
	Url string `json:"url"`
}

type BulletedListItem struct {
	TextBaseInfo
}

type Callout struct {
	TextBaseInfo
	Icon struct {
		Emoji      string `json:"emoji,omitempty"`
		Url        string `json:"url,omitempty"`
		ExpiryTime string `json:"expiry_time,omitempty"`
	} `json:"icon"`
}

func (obj *Callout) GetStatus() string {
	if len(obj.Icon.Emoji) != 0 {
		return obj.Icon.Emoji
	} else if len(obj.Icon.Url) != 0 {
		return obj.Icon.Url
		// return lib.DownloadFile(obj.Icon.Url)
	} else {
		return ""
	}
}

type Code struct {
	Caption
	RichText []RichText `json:"rich_text"`
	Language string     `json:"language"`
}

func (obj *Code) GetStatus() string {
	return obj.Language
}

func (obj *Code) GetContent() string {
	var textContent string
	for _, text := range obj.RichText {
		textContent += text.Output()
	}

	return textContent
}

func (obj *Code) BuildColor(content string) string {
	return content
}

type Embed struct {
	Url string `json:"url"`
}

type Equation struct {
	Expression string `json:"expression"`
}

func (e *Equation) GetContent() string {
	return e.Expression
}
func (e *Equation) BuildColor(content string) string {
	return e.Expression
}

type FileObject struct {
	Caption
	Type FileType `json:"type"`
	Name string   `json:"name,omitempty"`
	File struct {
		Url        string `json:"url"`
		ExpiryTime string `json:"expiry_time"`
	} `json:"file,omitempty"`
	External struct {
		Url string `json:"url"`
	} `json:"external,omitempty"`
}

func (file *FileObject) GetFile() string {
	switch file.Type {
	case FileTypeFile:
		return file.File.Url
		// fileStream, err := lib.DownloadFile(file.File.Url)
		// if err != nil {
		// }
	case FileTypeExternal:
		return file.External.Url
	default:
		return ""
	}
}

type File struct {
	Caption
	Type FileType   `json:"type"`
	Name string     `json:"name"`
	file FileObject `json:"file"`
}

type Heading1 struct {
	TextBaseInfo
	IsToggleable bool `json:"is_toggleable"`
}

func (head Heading1) GetStatus() string {
	if head.IsToggleable {
		return "true"
	} else {
		return ""
	}
}

type Heading2 struct{ Heading1 }
type Heading3 struct{ Heading1 }

type Image struct {
	FileObject
}

type LinkPreview struct {
	Url string `json:"url"`
}

type NumberedListItem struct {
	TextBaseInfo
}

type Paragraph struct {
	TextBaseInfo
}

type Pdf struct {
	Caption
	FileObject
}

type Quote struct {
	TextBaseInfo
}

type SyncedBlock struct {
	SyncedFrom struct {
		BlockId string `json:"block_id,omitempty"`
	} `json:"synced_from"`
}

type Table struct {
	TableWidth      int  `json:"table_width"`
	HasColumnHeader bool `json:"has_column_header"`
	HasRowHeader    bool `json:"has_row_header"`
}

type TableRow struct {
	Cells [][]RichText `json:"cells"`
}

type TableOfContents struct {
	Color AnnotationColor `json:"color"`
}

type ToDo struct {
	TextBaseInfo
	Checked bool `json:"checked,omitempty"`
}

func (obj *ToDo) GetStatus() string {
	if obj.Checked {
		return "x"
	} else {
		return " "
	}
}

type Toggle struct {
	TextBaseInfo
}

type Video struct {
	FileObject
}
