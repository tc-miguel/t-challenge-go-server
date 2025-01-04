package email_processor

import (
	services "example/hello/internal/services"
	utils "example/hello/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmails(context *gin.Context) {
	queryTerm := utils.GetStrQueryParamValue(context, "q")
	if len(queryTerm) == 0 {
		context.IndentedJSON(http.StatusBadRequest, "No search term found")
		return
	}
	page := utils.GetIntQueryParamOrDefault(context, "page", 0)
	size := utils.GetIntQueryParamOrDefault(context, "size", 50)
	emailDocumentList, error := services.GetEmails(page, size, queryTerm)
	if error != nil {
		context.IndentedJSON(http.StatusInternalServerError, error.Error())
	} else {
		context.IndentedJSON(http.StatusOK, emailDocumentList)
	}
}
