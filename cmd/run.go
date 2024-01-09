package main

import (
	"github.com/kam2yar/user-service/internal"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("tmp/app.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panicln(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	internal.Bootstrap()
}
