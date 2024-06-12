package filter

type FilterConditionType string

const (
	Equal      FilterConditionType = "equals"
	NotEqual   FilterConditionType = "does_not_equal"
	Empty      FilterConditionType = "is_empty"
	NotEmpty   FilterConditionType = "is_not_empty"
	Contain    FilterConditionType = "contains"
	NotContain FilterConditionType = "does_not_contain"
)

func (meta *MetaData) Equal(value interface{}) *MetaData {
	return meta.ConditionCommon(Equal, value)
}
func (meta *MetaData) NotEqual(value interface{}) *MetaData {
	return meta.ConditionCommon(NotEqual, value)
}
func (meta *MetaData) Empty(value interface{}) *MetaData {
	return meta.ConditionCommon(Empty, value)
}
func (meta *MetaData) NotEmpty(value interface{}) *MetaData {
	return meta.ConditionCommon(NotEmpty, value)
}
func (meta *MetaData) Contain(value interface{}) *MetaData {
	return meta.ConditionCommon(Contain, value)
}
func (meta *MetaData) NotContain(value interface{}) *MetaData {
	return meta.ConditionCommon(NotContain, value)
}

func (meta *MetaData) ConditionCommon(conditions FilterConditionType, value interface{}) *MetaData {
	var objType string
	for k, _ := range *meta {
		if k != "property" {
			objType = k
		}
	}

	valueInside := (*meta)[objType].(MetaDataValue)
	valueInside[conditions] = value
	(*meta)[objType] = valueInside

	return meta
}
