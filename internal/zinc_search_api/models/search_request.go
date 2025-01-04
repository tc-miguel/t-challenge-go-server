// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    searchRequest, err := UnmarshalSearchRequest(bytes)
//    bytes, err = searchRequest.Marshal()

package zinc_search

import "encoding/json"

func UnmarshalSearchRequest(data []byte) (SearchRequest, error) {
	var r SearchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SearchRequest struct {
	SearchType string `json:"search_type"`
	Query      Query  `json:"query"`
	From       int    `json:"from"`
	MaxResults int    `json:"max_results"`
}

type Query struct {
	Term  string `json:"term"`
	Field string `json:"field"`
}
