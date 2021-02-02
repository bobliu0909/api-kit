package v1

type ResponseCode int

const (
	ResponseSuccessfullyCode        = ResponseCode(0)
	ErrClientRequestResolveCode     = ResponseCode(-1000)
	ErrServerResourceNotFoundCode   = ResponseCode(-2000)
	ErrServerInternalFailedCode     = ResponseCode(-2001)
	/*
	ValidateRequestOKCode           = ResponseCode(1)
	ErrClientRequestValidateCode    = ResponseCode(-1001)
	*/
)

var (
	RequestBodyInvalidMsg = "request body invalid"
	ResourceNotFoundMsg = "resource not found"
	ServiceInternalErrorMsg = "service internal error"
)

type IResponse interface {
	WriteDetail(code int, message string, err error)
	WriteData(data interface{})
}

type Response struct {
	Detail struct {
		Code ResponseCode `json:"code"`
		Message string `json:"message"`
		Error string `json:"error,omitempty"`
	} `json:"detail"`
	Data interface{} `json:"data,omitempty"`
}

func (resp *Response) WriteDetail(code ResponseCode, message string, err error) {
	resp.Detail.Code = code
	resp.Detail.Message = message
	if err != nil {
		resp.Detail.Error = err.Error()
	}
}

func (resp *Response) WriteData(data interface{}) {
	resp.Data = data
}

func ErrorResponse(code ResponseCode, message string, err error) *Response {
	resp := &Response{}
	resp.WriteDetail(code, message, err)
	return resp
}

func DataResponse(code ResponseCode, message string, data interface{}) *Response {
	resp := &Response{}
	resp.WriteDetail(code, message, nil)
	resp.WriteData(data)
	return resp
}