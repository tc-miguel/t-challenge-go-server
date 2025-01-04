package main

import (
	zincModels "example/hello/internal/zinc_search_api/models"
	zincService "example/hello/internal/zinc_search_api/services"
	utils "example/hello/pkg/utils"
	"fmt"
	"log"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAddDocument(t *testing.T) {

	toValues := make([]string, 3)
	toValues = append(toValues, "ewalsh@edisonmission.com")
	toValues = append(toValues, "clong@enron.com")
	toValues = append(toValues, "csandhe@enron.com")
	body := zincModels.ZincEmailDocument{
		MessageId: "123456789",
		Date:      "Tue, 19 Dec 2000 07:48:00 -0800 (PST)",
		From:      "raislerk@sullcrom.com",
		Subject:   "H.R. 5660",
		Body:      "This e-mail is sent by a law firm and contains information that may be privileged and confidential. If you are not the intended recipient, please delete the e-mail and notify us immediately.",
		To:        toValues,
	}
	reader := utils.ParseToBytesReader(body)
	response, error := zincService.AddDocument("myindex", reader)
	if error != nil {
		log.Fatalln("No fue posible ejecutar test")
	}
	assert.Equal(t, response.StatusCode, 200)
}

func TestSearchByIndex(t *testing.T) {
	searchRequest := zincModels.SearchRequest{
		SearchType: "match",
		Query: zincModels.Query{
			Term:  "IMPROVEMENT",
			Field: "_all",
		},
		From:       0,
		MaxResults: 50,
	}
	response, error := zincService.SearchByIndex("email_doc", searchRequest)
	assert.Equal(t, error, nil)
	fmt.Print(response)
}
