package main

import (
	services "example/hello/internal/services"
	"log"
	"testing"
)

func TestProcessEmailFiles(t *testing.T) {
	log.Println("Start processing")
	success, error := services.ProcessEmailFiles()
	if error != nil {
		log.Fatal(error.Error())
	}
	log.Printf("Finish processing, succcess process: %t", success)
}
