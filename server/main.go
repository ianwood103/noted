package main

import (
	"fmt"
	"net/http"
	"noted/helper"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("start server")

	routes := httprouter.New()

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}