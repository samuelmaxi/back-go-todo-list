package main

import (
	"fmt"
	"os"

	"github.com/samuelmaxi/back-go-todo-list/cmd/api"
)

func main() {
	s := api.NewAPIServer("http://localhost")
	s.Run()
	a := os.Getenv("TESTE")
	fmt.Println("env: ", a)
}
