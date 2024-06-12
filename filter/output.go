package filter

type Filter interface {
	String() string
	StringIndent() string
}
