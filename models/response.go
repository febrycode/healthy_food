package models

type ResponseError struct {
	ResponseErrorData ResponseErrorData `json:"data"`
}

type ResponseErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseTokenData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type ResponseImageData struct {
	Code   int      `json:"code"`
	Images []string `json:"images"`
}

func ResponseJSON(code int, message string) ResponseError {
	return ResponseError{
		ResponseErrorData: ResponseErrorData{
			Code:    code,
			Message: message,
		},
	}
}

func ResponseToken(code int, message, token string) ResponseTokenData {
	return ResponseTokenData{
		code,
		message,
		token,
	}
}

func ResponseImage(code int, images []string) ResponseImageData {
	return ResponseImageData{
		Code:   code,
		Images: images,
	}
}
