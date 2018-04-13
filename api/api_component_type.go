package api

///////////////////////////////////////////////////////////////////////////////
// APIComponentType
type APIComponentType struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /component_types
//
func GetComponentTypes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetComponentTypes"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /component_types/<id>
//
func GetComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetComponentType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /component_types?name=<value>
//
func PutComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutComponentType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /component_types/<id>?name=<value>
//
func PostComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostComponentType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /component_types/<id>
//
func DeleteComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteComponentType"
	return
}
