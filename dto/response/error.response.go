package response

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
