package api

///////////////////////////////////////////////////////////////////////////////
// APIResultType
type APIResultType struct {
	ID          int    `json:"id,omitempty"`
	UserName    string `json:"user_name,omitempty"`
	CodeName    string `json:"code_name,omitempty"`
	Description string `json:"description,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /result_types
//
func GetResultTypes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetResultTypes"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /result_types/<id>
//
func GetResultType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetResultType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /result_types/<id>[?name=<Value>][?description=<Value>]
//
func PostResultType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostResultType"
	return
}
