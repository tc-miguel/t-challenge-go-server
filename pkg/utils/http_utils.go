package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

const defaultHttpSizeResults = 50

func ParseToBytesReader(object any) *bytes.Reader {
	data, err := json.Marshal(object)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewReader(data)
}

func GetStrQueryParamValue(context *gin.Context, paramName string) string {
	queryParams := context.Request.URL.Query()
	return queryParams.Get(paramName)
}

func GetIntQueryParamOrDefault(context *gin.Context, paramName string, defaultValue int) int {
	paramValue := GetStrQueryParamValue(context, paramName)
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
