package errorutils

var (
	UNKNOWN_ERROR = NewErrorCode("G99999", "UNKNOWN_ERROR", "Unknown exception", 500)
	BAD_REQUEST_BODY_PARSE_ERROR = NewErrorCode("G00001", "REQUEST_PARSE_ERROR", "Request parse exception", 400)
	BAD_REQUEST_VALDATION_ERROR = NewValidationErrorCode("G00002", "REQUEST_VALIDATION_ERROR", "Request validation exception", 400)
	NOT_FOUND_ERROR = NewErrorCode("G00003", "DATA_NOT_FOUND_ERROR","Data not found exception", 404)
)