package api

///////////////////////////////////////////////////////////////////////////////
// APIComponentSection
type APIComponentSection struct {
	Params []int64 `json:"params,omitempty"` // ID параметров из массива params этого участка
	Parts  []int64 `json:"parts,omitempty"`  // ID частей из массива parts этого участка
}

///////////////////////////////////////////////////////////////////////////////
// APIComponent
type APIComponent struct {
	ID            int64                 `json:"id,omitempty"`
	ComponentType *APIComponentType     `json:"component_type,omitempty"`
	PartTypes     []APIPartType         `json:"part_types,omitempty"`
	Sections      []APIComponentSection `json:"sections,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /projects/<id>/regions/<id>/components
//
func GetComponentsOfRegion(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetComponentsOfRegion"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /projects/<id>/regions/<id>/components/<id>
//
func GetComponent(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetComponent"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /projects/<id>/regions/<id>/components?component_type=<id>,<id>,...
//
func PutComponent(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutComponent"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /projects/<id>/regions/<id>/components?component_type=<id>,<id>,...
//
func PostComponent(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostComponent"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /projects/<id>/regions/<id>/components?component_type=<id>,<id>,...
//
func DeleteComponent(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteComponent"
	return
}
