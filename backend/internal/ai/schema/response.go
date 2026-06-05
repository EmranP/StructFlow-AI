package schema

type Template struct {
	Directories []string `json:"directories"`
	Files       []string `json:"files"`
}

type ProjectStructures struct {
	Simple     Template `json:"simple"`
	Medium     Template `json:"medium"`
	Enterprise Template `json:"enterprise"`
}
