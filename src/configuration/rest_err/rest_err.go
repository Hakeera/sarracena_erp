package rest_err

type RestErr struct {
	Message string `json:"message"`
	Err string 
	Code int
	Causes []Causes
}

type Causes struct {
	Field string
	Message string
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
