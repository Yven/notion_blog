package filter

type LogicRelate string

const (
	AND LogicRelate = "and"
	OR  LogicRelate = "or"
)

func (metaData *MetaData) Or(value *MetaData) *MetaData {
	newMetaData := make(MetaData)
	newMetaData[string(OR)] = &[2]MetaData{*metaData, *value}
	return &newMetaData
}

func (metaData *MetaData) And(value *MetaData) *MetaData {
	newMetaData := make(MetaData)
	newMetaData[string(AND)] = &[2]MetaData{*metaData, *value}
	return &newMetaData
}
