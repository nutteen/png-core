package errorutils

func NewErrorCode(code string, errorType string, title string, status int) *ErrorCode {
	return &ErrorCode{
		Code: code,
		Type: errorType,
		Title: title,
		Status: status,
	}
}

// ErrorCode predefines code, type and title to be used to construct ErrorResponseModel
type ErrorCode struct {
	Code string
	Type string
	Title string
	Status int
}

func (e ErrorCode) NewErrorResponseModel(errorDetail string) *ErrorResponseModel {
	return &ErrorResponseModel{
		Error: ErrorModel{
			Code: e.Code,
			Type: e.Type,
			Title: e.Title,
			Status: e.Status,
			Detail: errorDetail,
		},
	}
}

type ErrorResponseModel struct {
	Error ErrorModel `json:"error"`
}

// ErrorModel implements https://www.rfc-editor.org/rfc/rfc7807
type ErrorModel struct {
	// an error code that is numeric presentation of the error type
	Code string `json:"code"`
	// error type
	Type string `json:"type"`
	// a short, human-readable summary of the problem type
	Title string `json:"title"`
	// http status code
	Status int `json:"status"`
	// a human-readable explanation specific to this occurrence of the problem
	Detail string `json:"detail"`
	// A URI reference that identifies the specific resource that cause the problem
	Instance string `json:"Instance"`
}

type ValidationErrorResponseModel struct {
	Error ValidationErrorModel `json:"error"`
}

type ValidationErrorModel struct {
	// an error code that is numeric presentation of the error type
	Code string `json:"code"`
	// error type
	Type string `json:"type"`
	// a short, human-readable summary of the problem type
	Title string `json:"title"`
	// http status code
	Status int `json:"status"`
	// a human-readable explanation specific to this occurrence of the problem
	Detail string `json:"detail"`
	// A URI reference that identifies the specific resource that cause the problem
	Instance string `json:"Instance"`
	// Extension members
	InvalidParams []ValidationInvalidParam `json:"invalidParams"`
}

type ValidationInvalidParam struct {
	// Name of the invalid attribute
	Name string
	// Invalid reason of the attribute
	Reason string
}

func NewValidationErrorCode(code string, errorType string, title string, status int) *ValidationErrorCode {
	return &ValidationErrorCode{
		Code: code,
		Type: errorType,
		Title: title,
		Status: status,
	}
}

// ValidationErrorCode predefines code, type and title to be used to construct ValidationErrorResponseModel
type ValidationErrorCode struct {
	Code string
	Type string
	Title string
	Status int
}

func (e ValidationErrorCode) NewErrorResponseModel(errorDetail string, invalidParams []ValidationInvalidParam) *ValidationErrorResponseModel {
	return &ValidationErrorResponseModel{
		Error: ValidationErrorModel{
			Code: e.Code,
			Type: e.Type,
			Title: e.Title,
			Status: e.Status,
			Detail: errorDetail,
			InvalidParams: invalidParams,
		},
	}
}