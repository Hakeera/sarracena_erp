package rest_err

type RestErr struct {
	Message string `json:"message"`
	Err string `json: "error"`
	Code int `json: "cpde"`
	Causes []Causes `json: "causes"`
}

type Causes struct {
	Field string `json: "field"`
	Message string `json: "message"`
}

func (r *RestErr) Error() sting {
	return r.Message
}

func NewRestError(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}


func NewBadRequestError(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad request",
		Code: http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad request",
		Code: http.StatusBadRequest,
		Causes: causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad request",
		Code: http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad request",
		Code: http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad request",
		Code: http.StatusForbidden,
	}
}