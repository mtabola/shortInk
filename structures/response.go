package structures

type Response struct {
	Message interface{} `json:"Message"`
	Error   string      `json:"Error"`
}
