package main

import (
	"fmt"

	"github.com/akshayaparida/todo_golang/config"
	"github.com/akshayaparida/todo_golang/db"
)

func main() {
	config.LoadEnv() 

	port := config.GetEnv("PORT")
	db.ConnectDB()
	fmt.Printf("App starting on port %s...\n", port)
	
	
}
