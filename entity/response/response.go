package response

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

// NewResponse 创建rsponse
func NewResponse(code int, data interface{}, msg string) Response {
	return Response{
		Code:  code,
		Data:  data,
		Msg:   msg,
		Error: "",
	}
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}
