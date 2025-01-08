package email_processor

type EmailDocument struct {
	MessageId               string   `json:"messageId"`
	Date                    string   `json:"date"`
	From                    string   `json:"from"`
	Subject                 string   `json:"subject"`
	MimeVersion             string   `json:"mimeVersion"`
	ContentType             string   `json:"contentType"`
	ContentTransferEncoding string   `json:"contentTransferEncoder"`
	Body                    string   `json:"body"`
	XFrom                   string   `json:"xFrom"`
	XCc                     string   `json:"xCc"`
	XBcc                    string   `json:"xBcc"`
	XFolder                 string   `json:"xFolder"`
	XOrigin                 string   `json:"xOrigin"`
	XFilename               string   `json:"xFilename"`
	To                      []string `json:"to"`
	XTo                     []string `json:"xTo"`
	Cc                      []string `json:"cc"`
	Bcc                     []string `json:"bcc"`
	Error                   error
}
