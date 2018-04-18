package api

///////////////////////////////////////////////////////////////////////////////
// APICalculationType

type APICalculationType struct {
	ID       int64  `json:"id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	CodeName string `json:"code_name,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /calculation_types
//
func GetCalculationTypes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetCalculationTypes"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /calculation_types/<id>
//
func GetCalculationType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetCalculationType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /calculation_types/<id>?name=<Value>
//
func PostCalculationType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostCalculationType"
	return
}
