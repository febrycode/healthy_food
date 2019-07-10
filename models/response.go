package models

type ResponseError struct {
	ResponseErrorData ResponseErrorData `json:"data"`
}

type ResponseErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ResponseJSON(code int, message string) ResponseError {
	return ResponseError{
		ResponseErrorData: ResponseErrorData{
			Code:    code,
			Message: message,
		},
	}
}
