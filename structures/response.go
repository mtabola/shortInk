package structures

type Response struct {
	Message interface{} `json:"Message"`
	Error   interface{} `json:"Error"`
}

func (r *Response) GetResponse(msg interface{}, err interface{}) {
	r.Message = msg
	r.Error = err
}