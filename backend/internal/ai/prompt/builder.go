package prompt

import (
	"fmt"

	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/domain"
)

func BuildProjectPrompt(
	project *domain.Project,
) string {

	return fmt.Sprintf(`
Generate project structure.

Title:
%s

Project-type:
%s

Stack:
%s

Architecture:
%s

Features:
%s

AdditionalInfo:
%s

Generate:
1. Simple
2. Medium
3. Enterprise

Return JSON only.

Example:

{
  "simple": {
    "directories": [],
    "files": []
  },
  "medium": {
    "directories": [],
    "files": []
  },
  "enterprise": {
    "directories": [],
    "files": []
  }
}

No markdown.
No explanations.
No comments.
Only JSON.
`,
		project.Title,
		project.ProjectType,
		project.Stack,
		project.Architecture,
		project.Features,
		project.AdditionalInfo,
	)
}
