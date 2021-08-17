package repositories

import (
	"fmt"

	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/entities"
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/interfaces"
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/mappers"
)

type UserRepository struct {
}

func (repository UserRepository) Find(id string) *interfaces.Result {
	client, context := NewFirestoreClient()
	defer client.Close()
	ref := client.Collection("users").Doc(id)
	snapshot, err := ref.Get(context)
	if err != nil {
		fmt.Println(err)
		return &interfaces.Result{Value: err}
	}
	var user entities.User
	if mappedError := snapshot.DataTo(&user); mappedError != nil {
		fmt.Println(mappedError)
		return &interfaces.Result{Value: mappedError}
	}
	return &interfaces.Result{Value: &user, Mapper: mappers.User{}}
}

func (repository UserRepository) FindAll() *interfaces.Result {
	return &interfaces.Result{Value: ""}
}

func (repository UserRepository) QueryBy(queries interfaces.Queries) *interfaces.Result {
	return &interfaces.Result{}
}
