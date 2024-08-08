package util

import (
	"notification-api/domain/model"
	"strings"
)

func ProcessTemplate(template string, parameters []model.Parameter) string {
	processedTemplate := template
	for _, param := range parameters {
		processedTemplate = strings.ReplaceAll(processedTemplate, "`"+param.Name+"`", param.Value)
	}
	return processedTemplate
}
