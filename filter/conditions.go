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

func (meta *MetaData) Equal(value any) *MetaData {
	return meta.ConditionCommon(Equal, value)
}
func (meta *MetaData) NotEqual(value any) *MetaData {
	return meta.ConditionCommon(NotEqual, value)
}
func (meta *MetaData) Empty(value any) *MetaData {
	return meta.ConditionCommon(Empty, value)
}
func (meta *MetaData) NotEmpty(value any) *MetaData {
	return meta.ConditionCommon(NotEmpty, value)
}
func (meta *MetaData) Contain(value any) *MetaData {
	return meta.ConditionCommon(Contain, value)
}
func (meta *MetaData) NotContain(value any) *MetaData {
	return meta.ConditionCommon(NotContain, value)
}

func (meta *MetaData) ConditionCommon(conditions FilterConditionType, value any) *MetaData {
	meta.Condition = &MetaDataValue{
		Field: conditions,
		Value: value,
	}

	return meta
}
