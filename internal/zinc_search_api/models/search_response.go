// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    searchResponse, err := UnmarshalSearchResponse(bytes)
//    bytes, err = searchResponse.Marshal()

package zinc_search

import (
	"encoding/json"
	"time"
)

func UnmarshalSearchResponse(data []byte) (SearchResponse, error) {
	var r SearchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SearchResponse struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
}

type Hits struct {
	Total    Total   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hit   `json:"hits"`
}

type Hit struct {
	Index     Index     `json:"_index"`
	Type      Type      `json:"_type"`
	ID        string    `json:"_id"`
	Score     float64   `json:"_score"`
	Timestamp time.Time `json:"@timestamp"`
	Source    Source    `json:"_source"`
}

type Source struct {
	Timestamp              time.Time `json:"@timestamp"`
	Bcc                    []string  `json:"Bcc"`
	Body                   string    `json:"Body"`
	Cc                     []string  `json:"Cc"`
	ContentTransferEncoder string    `json:"ContentTransferEncoder"`
	ContentType            string    `json:"ContentType"`
	Date                   string    `json:"Date"`
	//Error                  interface{} `json:"Error"`
	From        string   `json:"From"`
	MessageID   string   `json:"MessageId"`
	MIMEVersion string   `json:"MimeVersion"`
	Subject     string   `json:"Subject"`
	To          []string `json:"To"`
	XFrom       string   `json:"XFrom"`
	XBcc        string   `json:"xBcc"`
	XCc         string   `json:"xCc"`
	XFilename   string   `json:"xFilename"`
	XFolder     string   `json:"xFolder"`
	XOrigin     string   `json:"xOrigin"`
	XTo         []string `json:"xTo"`
}

type Total struct {
	Value int `json:"value"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Index string

const (
	EmailDoc Index = "email_doc"
)

type Type string

const (
	Doc Type = "_doc"
)
