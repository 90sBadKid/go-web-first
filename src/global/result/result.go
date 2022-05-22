package result

type RestResult struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func newRestResult(status bool, code int, data interface{}, msg string) *RestResult {
	return &RestResult{
		Status:  status,
		Code:    code,
		Data:    data,
		Message: msg,
	}
}

func Successful() *RestResult {
	return newRestResult(true, StatusOk, nil, StatusText(StatusOk))
}

func SuccessfulData(data interface{}) *RestResult {
	return newRestResult(true, StatusOk, data, StatusText(StatusOk))
}

func Failure(message string) *RestResult {
	return newRestResult(false, StatusUnknownError, nil, message)
}
func FailureStatus(code int, message string) *RestResult {
	return newRestResult(false, code, nil, message)
}
func (e *RestResult) Error() string {
	return e.Message
}
