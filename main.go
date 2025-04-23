package main

import (
	"fmt"

	"github.com/akshayaparida/todo_golang/config"
)

func main() {
	config.LoadEnv() 

	port := config.GetEnv("PORT")
	fmt.Printf("App starting on port %s...\n", port)

	
}
