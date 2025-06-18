package structure

type DbAdapter interface {
	Writer(list *List) error
	Delete()
	Close()
}

type ListWriter interface {
	Writer(list *List) error
}
