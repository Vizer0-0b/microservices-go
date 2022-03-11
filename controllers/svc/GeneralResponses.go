package svc


const (
	OK		= 20000 // Inherr Standard OK
	ERROR	= 50000 // Inherr Standard ERROR
)
type Resp struct {
	Code int	`json:"code"`
	Data interface{} `json:"data"`
}
type PageResp struct {
	Code 		int			`json:"code"`
	Data 		interface{} `json:"data"`
	PageIndex 	int			`json:"page_index"`
	PageSize	int			`json:"page_size"`
	PageCount	int			`json:"page_count"`
}

func RespOK(code int, v interface{}) Resp{
	return Resp{OK,v}
}

