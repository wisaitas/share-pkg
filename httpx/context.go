package httpx

type ErrorContext struct {
	FilePath     *string `json:"file_path"`
	ErrorMessage string  `json:"error_message"`
}
