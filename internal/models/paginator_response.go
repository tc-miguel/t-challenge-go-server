package email_processor

type PaginatorResponse struct {
	Total        int `json:"total"`
	RecordsCount int `json:"recordsCount"`
	Data         any `json:"data"`
}
