package response

type Response struct {
	Status   string `json:"status"`
	Code     string `json:"code"`
	Message  string `json:"message"`
	TracerID string `json:"tracer_id,omitempty"`
}

type ErrorResponse struct {
	Response
	Errors []string `json:"errors"`
}

type SuccessResponse[T any] struct {
	Response
	Data interface{} `json:"data"`
}
