package filter

type LogicRelate string

const (
	AND LogicRelate = "and"
	OR  LogicRelate = "or"
)

func (metaData *MetaData) Or(value *MetaData) *MetaData {
	return metaData.CommonLogic(OR, value)
}
func (metaData *MetaData) And(value *MetaData) *MetaData {
	return metaData.CommonLogic(AND, value)
}

func (metaData *MetaData) CommonLogic(logic LogicRelate, value *MetaData) *MetaData {
	if metaData.Related != nil {
		if metaData.Related.Logic != logic {
			return metaData
		} else {
			data := append(metaData.Related.Value, value)
			metaData.Related.Value = data
			return metaData
		}
	}

	metaData.Related = &Related{
		Logic: logic,
		Value: []*MetaData{value},
	}

	return metaData
}
