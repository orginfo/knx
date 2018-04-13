package api

///////////////////////////////////////////////////////////////////////////////
// APIParamType
type APIParamType struct {
	ID          int    `json:"id,omitempty"`
	UserName    string `json:"user_name,omitempty"`
	CodeName    string `json:"code_name,omitempty"`
	Description string `json:"description,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// APIParamValue
type APIParamValue struct {
	Value float64 `json:"value,omitempty"`
	Name  string  `json:"name,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /param_types
//
func GetParamTypes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetParamTypes"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /param_types/<id>
//
func GetParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /param_types/<id>[?name=<Value>][?description=<Value>]
//
func PostParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /param_types/<id>/part_types
//
func GetPartTypesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetPartTypesOfParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /param_types/<id>/part_types?<id>,<id>,...
//
func PutPartTypesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutPartTypesOfParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /param_types/<id>/part_types?<id>,<id>,...
//
func PostPartTypesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostPartTypesOfParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /param_types/<id>/part_types?<id>,<id>,...
//
func DeletePartTypesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeletePartTypesOfParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /param_types/<id>/values
//
func GetValuesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetValuesOfParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /param_types/<id>/values[?value=<Value>(<Name>)&value=<Value>(<Name>)&...]
//
func PutValuesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutValuesOfParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /param_types/<id>/values[?value=<Value>(<Name>)&value=<Value>(<Name>)&...]
//
func PostValuesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostValuesOfParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /param_types/<id>/values[?value=<Value>&value=<Value>&...]
//
func DeleteValuesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteParamTypesValues"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /param_types/<id>/values/<value>/nomenclature?<id>,<id>,...
//
func PutNomenclatureForValueOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutNomenclatureForValueOfParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /param_types/<id>/values/<value>/nomenclature?<id>,<id>,...
//
func PostNomenclatureForValueOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostNomenclatureForValueOfParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /param_types/<id>/values/<value>/nomenclature?<id>,<id>,...
//
func DeleteNomenclatureForValueOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteNomenclatureForValueOfParamType"
	return
}
