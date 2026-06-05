package main

import (
	"log"
	"github.com/Promise111/Zero/cmd"
)

func main() {
	log.SetPrefix("Zero CLI project: ")
	log.SetFlags(1)
	log.Println("Starting Zero CLI project...")
	cmd.Execute()
}
