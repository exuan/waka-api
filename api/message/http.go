package message

const (
	ResFormat = `{"error":%d, "msg":"%s", "data":%s}`
)
const (
	OK       = 0
	ParamErr = iota + 10000
	CaptchaErr
	SaveFailed
)

var httpText = map[int]string{
	OK:         "",
	ParamErr:   "参数错误",
	CaptchaErr: "验证码错误",
	SaveFailed: "保存失败",
}

func HttpText(code int) string {
	if text, ok := httpText[code]; ok {
		return text
	}
	return ""
}
