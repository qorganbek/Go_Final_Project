package main

import (
	"fmt"
	"github.com/ZhanserikKalmukhambet/Go_Final_Project/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	fmt.Println("Hello ")
}
