package parser

import "strings"

func CleanJSON(
	content string,
) string {

	content = strings.TrimSpace(content)

	content = strings.TrimPrefix(
		content,
		"```json",
	)

	content = strings.TrimPrefix(
		content,
		"```",
	)

	content = strings.TrimSuffix(
		content,
		"```",
	)

	return strings.TrimSpace(content)
}
