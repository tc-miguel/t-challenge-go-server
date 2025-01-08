package utils

import (
	"bytes"
	"encoding/json"
	utils "example/hello/pkg/utils/models"
	"log"
	"net/http"
	"strconv"
)

const defaultHttpSizeResults = 50

func ParseToBytesReader(object any) *bytes.Reader {
	data, err := json.Marshal(object)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewReader(data)
}

func GetStrQueryParamValue(request *http.Request, paramName string) string {
	return request.URL.Query().Get(paramName)
}

func GetIntQueryParamOrDefault(request *http.Request, paramName string, defaultValue int) int {
	paramValue := GetStrQueryParamValue(request, paramName)
	if len(paramValue) == 0 {
		return defaultValue
	}
	paramInt, error := strconv.Atoi(paramValue)
	if error == nil {
		return paramInt
	} else {
		return -1
	}
}

func GetFrom(page int, size int) int {
	if page < 0 {
		page = 0
	}
	if size < 0 {
		size = defaultHttpSizeResults
	}
	return page * size
}

func GetMaxResults(from int, size int) int {
	return from + size
}

func HttpResponseServerError(responseWriter http.ResponseWriter, error error) {
	httpErrorResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error", error.Error())
}

func HttpResponseBadRequest(responseWriter http.ResponseWriter, message string) {
	httpErrorResponse(responseWriter, http.StatusBadRequest, "Bad Request", message)
}

func HttpResponseOk(responseWriter http.ResponseWriter, body any) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(body)
}

func httpErrorResponse(responseWriter http.ResponseWriter, statusCode int, error string, message string) {
	errorResponse := utils.ErrorResponse{
		Error:   error,
		Message: message,
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	json.NewEncoder(responseWriter).Encode(errorResponse)
}
