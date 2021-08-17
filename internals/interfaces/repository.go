package interfaces

type column string
type value string
type sign string

type Query struct {
	Column column
	Value  value
	Sign   sign
}

type Queries []Query

type ReadRepository interface {
	Find(id string) *Result
	FindAll() *Result
	QueryBy(queries Queries) *Result
}

type WriteRepository interface {
	Write(data interface{})
}

type UpdateRepository interface {
	Update(data interface{})
}

type DeleteRepository interface {
	Delete(data interface{})
}
