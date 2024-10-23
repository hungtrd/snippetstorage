package main

import (
	"fmt"

	"github.com/hungtrd/snippetstorage/internal/server"
)

func main() {
	server := server.NewServer()

	fmt.Printf("Server is starting at :%s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
