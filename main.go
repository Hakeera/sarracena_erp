package main

import (
    "fmt"
    "os"
    "log"

    "github.com/joho/godotenv"
)
func main() {
    err:= godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
    fmt.Println(os.Getenv("TEST"))
}


