package responseCode

type Response struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

var (
	OK = Response{
		Code:        "00",
		Description: "OK.",
	}
	ErrorValidation = Response{
		Code:        "01",
		Description: "Error validation.",
	}
)
