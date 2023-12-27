package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("hello world")
	godotenv.Load()
	portStr := os.Getenv("PORT")
	if portStr == ""{
		log.Fatal("No port found")
	}
	fmt.Println("Port:", portStr)

}