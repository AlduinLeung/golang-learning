package Utils

type ResponseData struct {
	Code    int
	Message string
	Data    interface{} // 空接口
}

func Response(Code int, Message string, Data interface{}) ResponseData {
	res := ResponseData{}
	res.Code = Code
	res.Message = Message
	res.Data = Data
	return res
}
