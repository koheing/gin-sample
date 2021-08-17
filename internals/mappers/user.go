package mappers

import (
	"strings"

	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/domains"
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/entities"
)

type User struct{}

func (user User) ToDomain(entity interface{}) interface{} {
	u := entity.(*entities.User)
	return &domains.User{UserId: u.Id, Name: u.FamilyName + " " + u.FirstName}
}

func (user User) ToEntity(domain interface{}) interface{} {
	u := domain.(*domains.User)
	return &entities.User{Id: u.UserId, FirstName: strings.Split(u.Name, " ")[1], FamilyName: strings.Split(u.Name, " ")[0]}
}
