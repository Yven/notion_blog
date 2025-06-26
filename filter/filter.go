package filter

import (
	"encoding/json"
)

type FilterObjectType string

const (
	ObjectTypeCheckbox    FilterObjectType = "checkbox"
	ObjectTypeDate        FilterObjectType = "date"
	ObjectTypeFiles       FilterObjectType = "files"
	ObjectTypeFormula     FilterObjectType = "formula"
	ObjectTypeMultiSelect FilterObjectType = "multi_select"
	ObjectTypeNumber      FilterObjectType = "number"
	ObjectTypePeople      FilterObjectType = "people"
	ObjectTypePhoneNumber FilterObjectType = "phone_number"
	ObjectTypeRelation    FilterObjectType = "relation"
	ObjectTypeRichText    FilterObjectType = "rich_text"
	ObjectTypeSelect      FilterObjectType = "select"
	ObjectTypeStatus      FilterObjectType = "status"
	ObjectTypeTimestamp   FilterObjectType = "timestamp"
	ObjectTypeID          FilterObjectType = "ID"
)

type MetaData struct {
	Property  string
	Type      FilterObjectType
	Condition *MetaDataValue
	Related   *Related
	Value     any
}

type MetaDataValue struct {
	Field FilterConditionType
	Value any
}

type Related struct {
	Logic LogicRelate
	Value []*MetaData
}

func Checkbox(name string) *MetaData {
	return ObjectCommon(ObjectTypeCheckbox, name)
}
func Date(name string) *MetaData {
	return ObjectCommon(ObjectTypeDate, name)
}
func Files(name string) *MetaData {
	return ObjectCommon(ObjectTypeFiles, name)
}
func Formula(name string) *MetaData {
	return ObjectCommon(ObjectTypeFormula, name)
}
func MultiSelect(name string) *MetaData {
	return ObjectCommon(ObjectTypeMultiSelect, name)
}
func Number(name string) *MetaData {
	return ObjectCommon(ObjectTypeNumber, name)
}
func People(name string) *MetaData {
	return ObjectCommon(ObjectTypePeople, name)
}
func PhoneNumber(name string) *MetaData {
	return ObjectCommon(ObjectTypePhoneNumber, name)
}
func Relation(name string) *MetaData {
	return ObjectCommon(ObjectTypeRelation, name)
}
func RichText(name string) *MetaData {
	return ObjectCommon(ObjectTypeRichText, name)
}
func Select(name string) *MetaData {
	return ObjectCommon(ObjectTypeSelect, name)
}
func Status(name string) *MetaData {
	return ObjectCommon(ObjectTypeStatus, name)
}
func Timestamp(name string) *MetaData {
	return ObjectCommon(ObjectTypeTimestamp, name)
}
func ID(name string) *MetaData {
	return ObjectCommon(ObjectTypeID, name)
}

func ObjectCommon(objType FilterObjectType, name string) *MetaData {
	metaData := MetaData{
		Property: name,
		Type:     objType,
	}

	return &metaData
}

func (meta *MetaData) Set(key string, value any) *MetaData {
	meta.Value = map[string]any{key: value}

	return meta
}

type filterBody map[string]any

func (meta *MetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(meta.makeFilter())
}

func (meta *MetaData) makeFilter() *filterBody {
	if meta.Value != nil {
		list := make(filterBody)
		if meta.Related != nil && meta.Related.Logic == AND {
			for _, val := range meta.Related.Value {
				list[val.Property] = map[string]any{
					string(val.Type): val.Value,
				}
			}
		}

		list[meta.Property] = map[string]any{
			string(meta.Type): meta.Value,
		}

		return &list
	}

	if meta.Related == nil || len(meta.Related.Value) == 0 {
		return meta.makefilterBody()
	}

	related := make(filterBody)
	logic := string(meta.Related.Logic)
	related[logic] = []*filterBody{}
	related[logic] = append(related[logic].([]*filterBody), meta.makefilterBody())
	for _, data := range meta.Related.Value {
		related[logic] = append(related[logic].([]*filterBody), data.makeFilter())
	}

	return &related
}

func (meta *MetaData) makefilterBody() *filterBody {
	return &filterBody{
		"property": meta.Property,
		string(meta.Type): map[string]any{
			string(meta.Condition.Field): meta.Condition.Value,
		},
	}
}
