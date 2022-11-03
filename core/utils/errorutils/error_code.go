package errorutils

var (
	GENERIC_ERROR = NewErrorCode("G00001", "Generic exception")
	BAD_REQUEST_BODY_PARSE_ERROR = NewErrorCode("R00001", "Request can't be parsed")
	BAD_REQUEST_VALDATION_ERROR = NewErrorCode("R00002", "Request doesn't pass validation")
	NOT_FOUND_ERROR = NewErrorCode("G00003", "Data not found exception")
)