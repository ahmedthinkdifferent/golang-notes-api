package http

type AppError struct {
	Errors     map[string]any
	StatusCode int
	Error      string
}

func (receiver *AppError) HasErrors() bool {
	return len(receiver.Errors) > 0 || receiver.Error != ""
}
