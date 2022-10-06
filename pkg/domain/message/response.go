package message

// custom code:
// 80 -> BAD_REQUEST
// 99 -> INTERNAL_SERVER_ERROR
// 0 -> SUCCESS

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    any    `json:"data"`
}
