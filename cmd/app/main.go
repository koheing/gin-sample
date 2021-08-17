package main

import (
	"github.com/nor-ko-hi-jp/gorilla-tutorial/internals/routers"
)

func main() {

	router := routers.NewRouter()

	router.Run()
}
