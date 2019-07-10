package models

type ResponseError struct {
	ResponseErrorData ResponseErrorData `json:"data"`
}

type ResponseErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseTokenData struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}

func ResponseJSON(code int, message string) ResponseError {
	return ResponseError{
		ResponseErrorData: ResponseErrorData{
			Code:    code,
			Message: message,
		},
	}
}

func ResponseToken(code int, token string) ResponseTokenData {
	return ResponseTokenData{
		code,
		token,
	}
}
