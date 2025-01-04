package email_processor

import (
	models "example/hello/internal/zinc_search_api/models"
	"fmt"
	"regexp"
	"strings"
)

func appendValues(principalSlide []string, values []string) []string {
	for i := 0; i < len(values); i++ {
		if len(values[i]) == 0 {
			continue
		}
		principalSlide = append(principalSlide, values[i])
	}
	return principalSlide
}

func buildFieldRegEx(prefix string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`^%s:\s*(.*)`, prefix))
}

func getSingleResult(values []string) string {
	return strings.TrimSpace(values[0])
}

func messageIdFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex: buildFieldRegEx(fieldName),
		AddValue: func(matches []string) {
			if len(item.MessageId) <= 0 {
				item.MessageId = getSingleResult(matches)
			}
		}}
}
func dateFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex: buildFieldRegEx(fieldName),
		AddValue: func(matches []string) {
			if len(item.Date) <= 0 {
				item.Date = getSingleResult(matches)
			}
		}}
}

func fromFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex: buildFieldRegEx(fieldName),
		AddValue: func(matches []string) {
			if len(item.From) <= 0 {
				item.From = getSingleResult(matches)
			}
		}}
}

func toFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex:      buildFieldRegEx(fieldName),
		SplitValue: true,
		AddValue: func(matches []string) {
			item.To = appendValues(item.To, matches)
		}}
}

func subjectFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex: buildFieldRegEx(fieldName),
		AddValue: func(matches []string) {
			if len(item.Subject) <= 0 {
				item.Subject = getSingleResult(matches)
			}
		}}
}

func mimeFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex: buildFieldRegEx(fieldName),
		AddValue: func(matches []string) {
			if len(item.MimeVersion) <= 0 {
				item.MimeVersion = getSingleResult(matches)
			}
		}}
}

func contentTypeFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex: buildFieldRegEx(fieldName),
		AddValue: func(matches []string) {
			if len(item.ContentType) <= 0 {
				item.ContentType = getSingleResult(matches)
			}
		}}
}

func contentTransferEncodingFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex: buildFieldRegEx(fieldName),
		AddValue: func(matches []string) {
			if len(item.ContentTransferEncoding) <= 0 {
				item.ContentTransferEncoding = getSingleResult(matches)
			}
		}}
}

func xFromFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex: buildFieldRegEx(fieldName),
		AddValue: func(matches []string) {
			if len(item.XFrom) <= 0 {
				item.XFrom = getSingleResult(matches)
			}
		}}
}

func xToFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex:      buildFieldRegEx(fieldName),
		SplitValue: true,
		AddValue: func(matches []string) {
			item.XTo = appendValues(item.XTo, matches)
		}}
}

func xCcFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex:      buildFieldRegEx(fieldName),
		SplitValue: true,
		AddValue: func(matches []string) {
			if len(item.XCc) <= 0 {
				item.XCc = getSingleResult(matches)
			}
		}}
}

func xBccFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex:      buildFieldRegEx(fieldName),
		SplitValue: true,
		AddValue: func(matches []string) {
			if len(item.XBcc) <= 0 {
				item.XBcc = getSingleResult(matches)
			}
		}}
}

func xFolderFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex: buildFieldRegEx(fieldName),
		AddValue: func(matches []string) {
			if len(item.XFolder) <= 0 {
				item.XFolder = getSingleResult(matches)
			}
		}}
}

func xOriginFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex: buildFieldRegEx(fieldName),
		AddValue: func(matches []string) {
			if len(item.XOrigin) <= 0 {
				item.XOrigin = getSingleResult(matches)
			}
		}}
}

func xFileNameFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex: buildFieldRegEx(fieldName),
		AddValue: func(matches []string) {
			if len(item.XFilename) <= 0 {
				item.XFilename = getSingleResult(matches)
			}
		}}
}

func ccFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex:      buildFieldRegEx(fieldName),
		SplitValue: true,
		AddValue: func(matches []string) {
			item.Cc = appendValues(item.Cc, matches)
		}}
}
func bccFieldConfig(fieldName string, item *models.ZincEmailDocument) models.FieldConfig {
	return models.FieldConfig{
		Regex:      buildFieldRegEx(fieldName),
		SplitValue: true,
		AddValue: func(matches []string) {
			item.Bcc = appendValues(item.Bcc, matches)
		}}
}

func GetFieldConfig(item *models.ZincEmailDocument) map[string]models.FieldConfig {
	// Mapa de configuración
	config := map[string]models.FieldConfig{
		"Message-ID":              messageIdFieldConfig("Message-ID", item),
		"Date":                    dateFieldConfig("Date", item),
		"From":                    fromFieldConfig("From", item),
		"To":                      toFieldConfig("To", item),
		"Subject":                 subjectFieldConfig("Subject", item),
		"Cc":                      ccFieldConfig("Cc", item),
		"Bcc":                     bccFieldConfig("Bcc", item),
		"MimeVersion":             mimeFieldConfig("MimeVersion", item),
		"ContentType":             contentTypeFieldConfig("ContentType", item),
		"ContentTransferEncoding": contentTransferEncodingFieldConfig("ContentTransferEncoding", item),
		"XFrom":                   xFromFieldConfig("XFrom", item),
		"XTo":                     xToFieldConfig("XTo", item),
		"XCc":                     xCcFieldConfig("XCc", item),
		"XBcc":                    xBccFieldConfig("XBcc", item),
		"XFolder":                 xFolderFieldConfig("XFolder", item),
		"XOrigin":                 xOriginFieldConfig("XOrigin", item),
		"XFilename":               xFileNameFieldConfig("XFilename", item),
	}
	return config
}

func changeFieldValue(configFieldMap map[string]models.FieldConfig, fieldKey string, line string) {
	lineValue := strings.Replace(line, fieldKey+":", "", 1)
	fieldConfig := configFieldMap[fieldKey]
	if fieldConfig.SplitValue {
		fieldConfig.AddValue(strings.Split(strings.ReplaceAll(strings.ReplaceAll(lineValue, "\t", ""), " ", ""), ","))
	} else {
		fieldConfig.AddValue([]string{lineValue})
	}
}
