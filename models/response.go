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
	Token   string `json:"token"`
	IsAdmin bool   `json:"is_admin"`
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

func ResponseToken(code int, token string, is_admin bool) ResponseTokenData {
	return ResponseTokenData{
		code,
		token,
		is_admin,
	}
}

func ResponseImage(code int, images []string) ResponseImageData {
	return ResponseImageData{
		Code:   code,
		Images: images,
	}
}
