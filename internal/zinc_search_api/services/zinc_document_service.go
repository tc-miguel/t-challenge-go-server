package zinc_search

import (
	zincModels "example/hello/internal/zinc_search_api/models"
	utils "example/hello/pkg/utils"
	"log"
)

const emailDocName = "email_doc"
const allDocumentsSearchType = "alldocuments"
const matchSearchType = "match"

func AddEmailDocuments(emailJsonList []zincModels.ZincEmailDocument) {
	for _, emailRecord := range emailJsonList {
		go AddEmailDocument(emailRecord)
	}
}

func AddEmailDocument(emailJson zincModels.ZincEmailDocument) bool {
	bodyReader := utils.ParseToBytesReader(emailJson)
	response, error := AddDocument(emailDocName, bodyReader)
	if error != nil {
		log.Println("Failed to create document email_doc in ZincSearch")
		return false
	}
	return response.StatusCode == 200
}

func validateSeachType(term string) string {
	if len(term) == 0 {
		return allDocumentsSearchType
	} else {
		return matchSearchType
	}
}

func SearchByDocument(term string, from int, maxResult int) (zincModels.SearchResponse, error) {
	searchRequest := zincModels.SearchRequest{
		SearchType: validateSeachType(term),
		From:       from,
		MaxResults: maxResult,
		Query: zincModels.Query{
			Term:  term,
			Field: "_all",
		},
	}
	return SearchByIndex(emailDocName, searchRequest)
}
