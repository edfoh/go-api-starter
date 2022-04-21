package http

type errorResponse struct {
	Errors []*errorPayload `json:"errors"`
}

type errorPayload struct {
	Message string `json:"message"`
}
