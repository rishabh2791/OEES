package value_objects

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}
