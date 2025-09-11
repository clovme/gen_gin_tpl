package enums

type Enums struct {
	ID   int64
	Key  string
	Name string
	Desc string
}

type Enum[T any] interface {
	ID() int64
	Key() string
	Name() string
	Desc() string
	Int() int
	Is(v T) bool
}
