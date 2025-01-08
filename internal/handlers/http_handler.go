package email_processor

import (
	services "example/hello/internal/services"
	utils "example/hello/pkg/utils"
	"net/http"
)

func GetEmails(responseWriter http.ResponseWriter, request *http.Request) {
	queryTerm := utils.GetStrQueryParamValue(request, "q")
	page := utils.GetIntQueryParamOrDefault(request, "page", 0)
	size := utils.GetIntQueryParamOrDefault(request, "size", 50)
	paginatorResponse, error := services.GetEmails(page, size, queryTerm)
	if error != nil {
		utils.HttpResponseServerError(responseWriter, error)
	} else {
		utils.HttpResponseOk(responseWriter, paginatorResponse)
	}
}
