package model


// 响应对象 对象
type Resp struct {
	ErrNo     	int64 `json:"err_no"`      
	ErrMsg		string `json:"err_msg"`
	Data  		interface{} `json:"data"`
}