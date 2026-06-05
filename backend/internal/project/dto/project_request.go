package dto

type ProjectRequest struct {
	Title string `json:"title" validate:"required"`

	ProjectType string `json:"projectType" validate:"required"`

	Stack string `json:"stack" validate:"required"`

	Architecture string `json:"architecture"`

	Features string `json:"features" validate:"required"`

	AdditionalInfo string `json:"additionalInfo"`
}
