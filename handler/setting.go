package handler

//ResType 前后端交互返回值通用结构
type ResType struct {
	Code int32
	Text string
}

//JSON 格式化返回
func JSON(Code int32, Text string) string {
	//rt := ResType{Code: Code, Text: Text}
	return "{Code:" + string(Code) + ",Text:" + Text + "}"
}
