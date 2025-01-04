package main

import (
	config "example/hello/internal/config"
	"fmt"
	"testing"
)

func TestGetStringByKey(t *testing.T) {
	zincUrl := config.GetStringByKey("zinc.api.url")
	fmt.Printf("property: %s", zincUrl)
}
