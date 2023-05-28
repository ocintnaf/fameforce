package http

// HttpErrorResponse defines the structure of an error HTTP response.
type HttpErrorResponse struct {
	Error string `json:"error,omitempty"`
}

// HttpSuccessResponse defines the structure of a success HTTP response.
type HttpSuccessResponse struct {
	Data interface{} `json:"data,omitempty"`
}

// HttpResponse defines the structure of a HTTP response.
type HttpResponse struct {
	HttpSuccessResponse
	HttpErrorResponse
}

// NewHttpErrorResponse creates a new HttpErrorResponse.
func NewHttpErrorResponse(err error) HttpResponse {
	return HttpResponse{
		HttpErrorResponse: HttpErrorResponse{
			Error: err.Error(),
		},
	}
}

// NewHttpSuccessResponse creates a new HttpSuccessResponse.
func NewHttpSuccessResponse(data interface{}) HttpResponse {
	return HttpResponse{
		HttpSuccessResponse: HttpSuccessResponse{
			Data: data,
		},
	}
}

// NewHttpResponse creates a new HttpResponse.
func NewHttpResponse(data interface{}, err error) HttpResponse {
	if err != nil {
		return NewHttpErrorResponse(err)
	}

	return NewHttpSuccessResponse(data)
}
