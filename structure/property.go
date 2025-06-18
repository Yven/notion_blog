package structure

import (
	"strconv"
)

type PropertyObject struct {
	Cover      string                 `json:"cover,omitempty"`
	Icon       string                 `json:"icon,omitempty"`
	Properties map[string]*Properties `json:"properties,omitempty"`
	Url        string                 `json:"url,omitempty"`
	PublicUrl  string                 `json:"public_url,omitempty"`
}

type PropertiesType string
type FormulaType string

const (
	PropertyTypeCheckbox       PropertiesType = "checkbox"
	PropertyTypeCreatedBy      PropertiesType = "created_by"
	PropertyTypeCreatedTime    PropertiesType = "created_time"
	PropertyTypeDate           PropertiesType = "date"
	PropertyTypeEmail          PropertiesType = "email"
	PropertyTypeFiles          PropertiesType = "files"
	PropertyTypeFormula        PropertiesType = "formula"
	PropertyTypeLastEditedBy   PropertiesType = "last_edited_by"
	PropertyTypeLastEditedTime PropertiesType = "last_edited_time"
	PropertyTypeMultiSelect    PropertiesType = "multi_select"
	PropertyTypeNumber         PropertiesType = "number"
	PropertyTypePeople         PropertiesType = "people"
	PropertyTypePhoneNumber    PropertiesType = "phone_number"
	PropertyTypeRelation       PropertiesType = "relation"
	PropertyTypeRollup         PropertiesType = "rollup"
	PropertyTypeRichText       PropertiesType = "rich_text"
	PropertyTypeSelect         PropertiesType = "select"
	PropertyTypeStatus         PropertiesType = "status"
	PropertyTypeTitle          PropertiesType = "title"
	PropertyTypeUrl            PropertiesType = "url"
	PropertyTypeUniqueId       PropertiesType = "unique_id"
	PropertyTypeVerification   PropertiesType = "verification"

	FormulaBoolean FormulaType = "boolean"
	FormulaDate    FormulaType = "date"
	FormulaNumber  FormulaType = "number"
	FormulaString  FormulaType = "string"
)

type Properties struct {
	Id   string         `json:"id,omitempty"`
	Type PropertiesType `json:"type,omitempty"`

	Checkbox       *bool         `json:"checkbox,omitempty"`
	CreatedBy      *CreatedBy    `json:"created_by,omitempty"`
	CreatedTime    *string       `json:"created_time,omitempty"`
	Date           *Date         `json:"date,omitempty"`
	Email          *string       `json:"email,omitempty"`
	Files          []*FileObject `json:"files,omitempty"`
	Formula        *Formula      `json:"formula,omitempty"`
	LastEditedBy   *LastEditedBy `json:"last_edited_by,omitempty"`
	LastEditedTime *string       `json:"last_edited_time,omitempty"`
	MultiSelect    []*Select     `json:"multi_select,omitempty"`
	Number         *Number       `json:"number,omitempty"`
	People         []*User       `json:"people,omitempty"`
	PhoneNumber    *string       `json:"phone_number,omitempty"`
	Relation       []*Relation   `json:"relation,omitempty"`
	Rollup         *Rollup       `json:"rollup,omitempty"`
	RichText       []*RichText   `json:"rich_text,omitempty"`
	Select         *Select       `json:"select,omitempty"`
	Status         *Status       `json:"status,omitempty"`
	Title          []*RichText   `json:"title,omitempty"`
	Url            *string       `json:"url,omitempty"`
	UniqueId       *UniqueId     `json:"unique_id,omitempty"`
	Verification   *Verification `json:"verification,omitempty"`
}

func (p *PropertyObject) Get(name string) any {
	if p.Properties[name] != nil {
		return p.Properties[name].GetValue()
	} else {
		return ""
	}
}

func (p *Properties) GetValue() any {
	// layout := "2006-01-02 15:04:05"
	switch p.Type {
	case PropertyTypeCheckbox:
		return *p.Checkbox
	case PropertyTypeCreatedBy:
		return p.CreatedBy.Name
	case PropertyTypeCreatedTime:
		return *p.CreatedTime
	case PropertyTypeDate:
		// return p.Date.Start.Format(layout) + "~" + p.Date.End.Format(layout)
		return p.Date.Start + "~" + p.Date.End
	case PropertyTypeEmail:
		return *p.Email
	case PropertyTypeFormula:
		return p.Formula.GetData()
	case PropertyTypeLastEditedBy:
		return p.LastEditedBy.Name
	case PropertyTypeLastEditedTime:
		return *p.LastEditedTime
	case PropertyTypeMultiSelect:
		var selectNames []string
		for _, slt := range p.MultiSelect {
			selectNames = append(selectNames, slt.Name)
		}
		return selectNames
	case PropertyTypeNumber:
		return p.Number.Format
	case PropertyTypePeople:
	case PropertyTypePhoneNumber:
		return *p.PhoneNumber
	case PropertyTypeRichText:
		var content string
		for _, rt := range p.RichText {
			content += rt.Output()
		}
		return content
	case PropertyTypeSelect:
		return p.Select.Name
	case PropertyTypeStatus:
		return p.Status.Name
	case PropertyTypeTitle:
		var content string
		for _, rt := range p.Title {
			content += rt.Output()
		}
		return content
	case PropertyTypeUrl:
		return *p.Url
	case PropertyTypeUniqueId:
		return p.UniqueId.Prefix + "_" + p.UniqueId.Number
	case PropertyTypeVerification:
		return p.Verification.State
	case PropertyTypeFiles, PropertyTypeRelation:
		return ""
	}

	return ""
}

type PropertyBaseInfo struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type CreatedBy struct {
	User
}
type Formula struct {
	Type    FormulaType `json:"type"`
	Boolean bool        `json:"boolean,omitempty"`
	Date    *Date       `json:"date,omitempty"`
	Number  int         `json:"number,omitempty"`
	String  string      `json:"string,omitempty"`
}

func (f Formula) GetData() string {
	switch f.Type {
	case FormulaBoolean:
		if f.Boolean {
			return "true"
		} else {
			return "false"
		}
	case FormulaDate:
		return f.Date.Start
		// return f.Date.Start.Format("2006-01-02 15:04:05")
	case FormulaNumber:
		return strconv.Itoa(f.Number)
	case FormulaString:
		return f.String
	}

	return ""
}

type LastEditedBy struct {
	User
}

//	type MultiSelect struct {
//		Options []*Select `json:"options"`
//	}
type Number struct {
	Format string `json:"format"`
}
type Relation struct {
	Id string `json:"id"`
}
type Rollup struct {
}
type Select struct {
	Id    string          `json:"id"`
	Name  string          `json:"name"`
	Color AnnotationColor `json:"color"`
}
type Status struct {
	Id    string          `json:"id"`
	Name  string          `json:"name"`
	Color AnnotationColor `json:"color"`
}
type UniqueId struct {
	Number string `json:"number"`
	Prefix string `json:"prefix,omitempty"`
}
type Verification struct {
	State      string `json:"state"`
	VerifiedBy *User  `json:"verified_by,omitempty"`
	Date       *Date  `json:"date,omitempty"`
}
