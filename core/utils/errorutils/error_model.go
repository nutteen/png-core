package errorutils

type ErrorResponseModel struct {
	Error ErrorModel `json:"error"`
}

type ErrorModel struct {
	Code string `json:"code"`
	Message string `json:"message"`
	DetailMessage string `json:"detailMessage"`
}

type ErrorCode struct {
	Code string
	Message string
}

func NewErrorCode(code string, message string) ErrorCode {
	return ErrorCode{
		Code: code,
		Message: message,
	}
}

func newShortErrorModel(errorCode ErrorCode) ErrorModel {
	return ErrorModel{
		Code: errorCode.Code,
		Message: errorCode.Message,
	}
}

func newErrorModel(errorCode ErrorCode, detailMessage string) ErrorModel {
	return ErrorModel{
		Code: errorCode.Code,
		Message: errorCode.Message,
		DetailMessage: detailMessage,
	}
}

func NewErrorResponseModel(errorCode ErrorCode, detailMessage string) *ErrorResponseModel {
	return &ErrorResponseModel{
		Error: newErrorModel(errorCode, detailMessage),
	}
}

func NewShortErrorResponseModel(errorCode ErrorCode) *ErrorResponseModel {
	return &ErrorResponseModel{
		Error: newShortErrorModel(errorCode),
	}
}