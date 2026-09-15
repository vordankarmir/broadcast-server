/*
Copyright © 2026 Henry Grigoryan henrygrigoryan5@gmail.com
*/
package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vordankarmir/broadcast-server/cmd"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	cmd.Execute()
}
