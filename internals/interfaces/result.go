package interfaces

type Mapper interface {
	ToDomain(entity interface{}) interface{}
	ToEntity(domain interface{}) interface{}
}

type Result struct {
	Value  interface{}
	Mapper Mapper
}
