package zinc_search

type ZincEmailDocument struct {
	MessageId               string   `json:"MessageId"`
	Date                    string   `json:"Date"`
	From                    string   `json:"From"`
	Subject                 string   `json:"Subject"`
	MimeVersion             string   `json:"MimeVersion"`
	ContentType             string   `json:"ContentType"`
	ContentTransferEncoding string   `json:"ContentTransferEncoder"`
	Body                    string   `json:"Body"`
	XFrom                   string   `json:"XFrom"`
	XCc                     string   `json:"xCc"`
	XBcc                    string   `json:"xBcc"`
	XFolder                 string   `json:"xFolder"`
	XOrigin                 string   `json:"xOrigin"`
	XFilename               string   `json:"xFilename"`
	To                      []string `json:"To"`
	XTo                     []string `json:"xTo"`
	Cc                      []string `json:"Cc"`
	Bcc                     []string `json:"Bcc"`
	Error                   error
}
