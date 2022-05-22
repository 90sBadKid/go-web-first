package result

const (
	StatusOk = 2000 //成功

	StatusClientError = 4000 //服务器系统错误

	StatusServerError = 5000 //服务器系统错误

	StatusUnknownError = 6000 //未知错误
)

var statusText = map[int]string{
	StatusOk: "成功",

	StatusClientError: "服务器系统错误",

	StatusServerError: "服务器系统错误",

	StatusUnknownError: "未知错误",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
