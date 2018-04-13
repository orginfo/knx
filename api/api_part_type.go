package api

///////////////////////////////////////////////////////////////////////////////
// APIPartType
type APIPartType struct {
	ID                int               `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	CalculationTypeID int               `json:"calculation_type_id,omitempty"`
	ComponentType     *APIComponentType `json:"component_type,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /component_types/<id>/part_types
//
func GetPartTypesOfComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetPartTypesOfComponentType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /component_types/<id>/part_types/<id>
//
func GetPartType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetPartType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /component_types/<id>/part_types?name=<value>[?calculation_type_id=<value>]
//
func PutPartType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutPartType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /component_types/<id>/part_types/<id>[?name=<value>][?calculation_type_id=<value>]
//
func PostPartType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostPartType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /component_types/<id>/part_types/<id>
//
func DeletePartType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeletePartType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /component_types/<id>/part_types/<id>/nomenclature?<id>,<id>,...
//
func PutNomenclatureForPartType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutNomenclatureForPartType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /component_types/<id>/part_types/<id>/nomenclature?<id>,<id>,...
//
func PostNomenclatureForPartType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostNomenclatureForPartType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /component_types/<id>/part_types/<id>/nomenclature?<id>,<id>,...
//
func DeleteNomenclatureForPartType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteNomenclatureForPartType"
	return
}
