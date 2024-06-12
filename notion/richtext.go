package notion

import (
	"fmt"
)

type RichTextType string
type MentionType string
type TemplateMentionType string
type TemplateMentionDateType string
type TemplateMentionUserType string

const (
	TypeText     RichTextType = "text"
	TypeMention  RichTextType = "mention"
	TypeEquation RichTextType = "equation"

	MentionTypeDatabase        MentionType = "database"
	MentionTypeDate            MentionType = "date"
	MentionTypeLinkPreview     MentionType = "link_preview"
	MentionTypePage            MentionType = "page"
	MentionTypeTemplateMention MentionType = "template_mention"
	MentionTypeUser            MentionType = "user"

	TemplateMentionDate TemplateMentionType = "template_mention_date"
	TemplateMentionUser TemplateMentionType = "template_mention_user"

	TemplateMentionDateToday TemplateMentionDateType = "today"
	TemplateMentionUserMe    TemplateMentionUserType = "me"
)

type Page struct {
	Id string `json:"id"`
}
type Database struct {
	Id string `json:"id"`
}

type Annotation struct {
	Bold          bool            `json:"bold"`
	Italic        bool            `json:"italic"`
	Strikethrough bool            `json:"strikethrough"`
	Underline     bool            `json:"underline"`
	Code          bool            `json:"code"`
	Color         AnnotationColor `json:"color"`
}

func (anno Annotation) Build(content string) string {
	if anno.Bold {
		content = fmt.Sprintf("**%s**", content)
	}
	if anno.Italic {
		content = fmt.Sprintf("*%s*", content)
	}
	if anno.Strikethrough {
		content = fmt.Sprintf("~~%s~~", content)
	}
	if anno.Underline {
		content = fmt.Sprintf("<u>%s</u>", content)
	}
	if anno.Code {
		content = fmt.Sprintf("`%s`", content)
	}
	if anno.Color != Default {
		content = fmt.Sprintf("<font %s>%s</font>", anno.Color.Output(), content)
	}

	return content
}

type RichText struct {
	Type RichTextType `json:"type"`
	Text *struct {
		Content string `json:"content"`
		Link    struct {
			Url string `json:"url"`
		} `json:"link"`
	} `json:"text,omitempty"`
	Mention *struct {
		Type        MentionType `json:"type"`
		Database    Database    `json:"database,omitempty"`
		Date        Date        `json:"date,omitempty"`
		Page        Page        `json:"date,omitempty"`
		User        User        `json:"date,omitempty"`
		LinkPreview struct {
			Url string `json:"url"`
		} `json:"date,omitempty"`
		TemplateMention struct {
			Type                TemplateMentionType     `json:"type"`
			TemplateMentionDate TemplateMentionDateType `json:"template_mention_date,omitempty"`
			TemplateMentionUser TemplateMentionUserType `json:"template_mention_user,omitempty"`
		} `json:"date,omitempty"`
	} `json:"mention,omitempty"`
	Equation    *Equation  `json:"equation,omitempty"`
	Annotations Annotation `json:"annotations"`
	PlainText   string     `json:"plain_text"`
	Href        string     `json:"href"`
}

func (richText RichText) Output() string {
	var res string
	switch richText.Type {
	case TypeText:
		if len(richText.Text.Link.Url) != 0 {
			res = fmt.Sprintf("[%s](%s)", richText.Text.Content, richText.Text.Link.Url)
		} else {
			res = richText.Text.Content
		}
		res = richText.Annotations.Build(res)
	case TypeMention:
		res = ""
	case TypeEquation:
		res = fmt.Sprintf("$%s$", richText.Equation.GetContent())
	}

	return res
}
func (richText RichText) NormalOutput() string {
	var res string
	switch richText.Type {
	case TypeText:
		res = richText.Text.Content
	case TypeMention:
		res = ""
	case TypeEquation:
		res = richText.Equation.GetContent()
	}

	return res
}
