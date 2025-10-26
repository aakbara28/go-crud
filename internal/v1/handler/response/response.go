package response

type Response struct {
	RequestId   string      `json:"requestId"`
	RequestKey  string      `json:"requestKey"`
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data,omitempty"`
}

func HandleSuccess(reqId string, reqKey string, status int, message string, data ...interface{}) Response {
	res := Response{
		RequestId:  reqId,
		RequestKey: reqKey,
		Code:       status,
		Status:     "OK",
		Message:    message,
	}
	
	if len(data) > 0 {
		res.Data = data[0]
	}
	
	return res
}

func HandleError(reqId string, reqKey string, status int, err error) Response {
	return Response{
		RequestId:  reqId,
		RequestKey: reqKey,
		Code:       status,
		Status:     "NOK",
		Message:    err.Error(),
	}
}