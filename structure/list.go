package structure

type List struct {
	Object         string    `json:"object"`
	Results        []*Object `json:"results"`
	NextCursor     string    `json:"next_cursor"`
	HasMore        bool      `json:"has_more"`
	Type           string    `json:"type"`
	PageOrDatabase any       `json:"page_or_database"`
	RequestId      string    `json:"request_id"`

	Property *Object

	OutputType OutputType
}

func (list *List) GetContent() []*Object {
	return list.Results
}

func (list *List) AppendContent(objList []*Object) {
	list.Results = append(list.Results, objList...)
}
