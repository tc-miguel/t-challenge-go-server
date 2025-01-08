package email_processor

import (
	models "example/hello/internal/models"
	zincModels "example/hello/internal/zinc_search_api/models"
	zincService "example/hello/internal/zinc_search_api/services"
	utils "example/hello/pkg/utils"
)

func GetEmails(page int, size int, term string) (models.PaginatorResponse, error) {
	from := utils.GetFrom(page, size)
	maxResults := utils.GetMaxResults(from, size)
	searchResponse, error := zincService.SearchByDocument(term, from, maxResults)
	if error != nil {
		return models.PaginatorResponse{}, error
	}
	data := emailDocumentCopyFrom(searchResponse)
	response := models.PaginatorResponse{
		Data:         data,
		Total:        searchResponse.Hits.Total.Value,
		RecordsCount: maxResults}
	return response, error
}

func emailDocumentCopyFrom(response zincModels.SearchResponse) []models.EmailDocument {
	var emailDocumentList = make([]models.EmailDocument, 0)
	for _, hit := range response.Hits.Hits {
		emailDocumentList = append(emailDocumentList, copy(hit.Source))
	}
	return emailDocumentList
}

func copy(data zincModels.Source) models.EmailDocument {
	return models.EmailDocument{
		MessageId:   data.MessageID,
		Date:        data.Date,
		From:        data.From,
		Subject:     data.Subject,
		MimeVersion: data.MIMEVersion,
		ContentType: data.ContentType,
		//ContentTransferEncoding :data.ContentTransferEncoding,
		Body:      data.Body,
		XFrom:     data.XFrom,
		XCc:       data.XCc,
		XBcc:      data.XBcc,
		XFolder:   data.XFolder,
		XOrigin:   data.XOrigin,
		XFilename: data.XFilename,
		To:        data.To,
		XTo:       data.XTo,
		Cc:        data.Cc,
		Bcc:       data.Bcc,
	}
}
