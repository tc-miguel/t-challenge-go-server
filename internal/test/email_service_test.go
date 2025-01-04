package main

import (
	services "example/hello/internal/services"
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetEmails(t *testing.T) {
	emailDocuments, error := services.GetEmails(0, 20, "Partner")
	assert.Equal(t, error, nil)
	fmt.Print(emailDocuments)
}
