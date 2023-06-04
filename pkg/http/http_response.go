package http

// TODO: Move under types/http_types.go
type Status string

const (
	StatusSuccess Status = "success"
	StatusFail    Status = "fail"
	StatusError   Status = "error"
)

type BaseResponse struct {
	Status Status `json:"status"`
}

type SuccessResponse struct {
	BaseResponse
	Data any `json:"data"`
}

type FailResponse struct {
	BaseResponse
	Data any `json:"data"`
}

type ErrorResponse struct {
	BaseResponse
	Message string `json:"message"`
}

func NewSuccessResponse(data any) SuccessResponse {
	return SuccessResponse{
		BaseResponse: BaseResponse{
			Status: StatusSuccess,
		},
		Data: data,
	}
}

func NewFailResponse(data any) FailResponse {
	return FailResponse{
		BaseResponse: BaseResponse{
			Status: StatusFail,
		},
		Data: data,
	}
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		BaseResponse: BaseResponse{
			Status: StatusError,
		},
		Message: err.Error(),
	}
}
