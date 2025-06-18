package filter

type SortDirection string
type SortTimestamp string

const (
	SortDirectionAscending  SortDirection = "ascending"
	SortDirectionDescending SortDirection = "descending"

	SortTimestampCreateTime     SortTimestamp = "created_time"
	SortTimestampLastEditedTime SortTimestamp = "last_edited_time"
)

type Sort struct {
	Property  string        `json:"property"`
	Timestamp SortTimestamp `json:"timestamp"`
	Direction SortDirection `json:"direction"`
}
