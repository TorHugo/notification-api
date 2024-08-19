package util

import (
	"notification-api/domain"
	"strings"
)

func ProcessTemplate(template string, parameters []domain.Parameter) string {
	processedTemplate := template
	for _, param := range parameters {
		processedTemplate = strings.ReplaceAll(processedTemplate, "`"+param.Name+"`", param.Value)
	}
	return processedTemplate
}
