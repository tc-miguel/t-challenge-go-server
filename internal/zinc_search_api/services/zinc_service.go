package zinc_search

import (
	"encoding/base64"
	config "example/hello/internal/config"
	zincModels "example/hello/internal/zinc_search_api/models"
	utils "example/hello/pkg/utils"
	"io"
	"net/http"
)

func getHttpBase() string {
	return config.GetStringByKey("zinc.api.url")
}

func getClientUsername() string {
	return config.GetStringByKey("zinc.api.username")
}

func getClientPassword() string {
	return config.GetStringByKey("zinc.api.credentials")
}

func basicAuth() string {
	auth := getClientUsername() + ":" + getClientPassword()
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth())
	return nil
}

func getClient() *http.Client {
	return &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
}
func AddDocument(indexName string, body io.Reader) (*http.Response, error) {
	url := getHttpBase() + "/" + indexName + "/_doc"
	return postCall(url, body)
}

func SearchByIndex(indexName string, searchRequest zincModels.SearchRequest) (zincModels.SearchResponse, error) {
	url := getHttpBase() + "/" + indexName + "/_search"
	body := utils.ParseToBytesReader(searchRequest)
	response, error := postCall(url, body)
	if error != nil {
		return zincModels.SearchResponse{}, error
	}
	defer response.Body.Close()

	data, error := io.ReadAll(response.Body)
	if error != nil {
		return zincModels.SearchResponse{}, error
	}
	return zincModels.UnmarshalSearchResponse(data)
}

func postCall(url string, body io.Reader) (*http.Response, error) {
	client := getClient()
	request, requestError := http.NewRequest("POST", url, body)
	if requestError != nil {
		return nil, requestError
	}
	request.Header.Add("Authorization", "Basic "+basicAuth())
	return client.Do(request)
}
