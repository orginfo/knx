package api

///////////////////////////////////////////////////////////////////////////////
// APINomenclatureType
type APINomenclatureType struct {
	ID            int      `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	ColorSchemeID int      `json:"color_scheme_id,omitempty"`
	UseFields     []string `json:"use_fields,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /nomenclature_types
//
func GetNomenclatureTypes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclatureTypes"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /nomenclature_types/<id>
//
func GetNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclatureType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /nomenclature_types?name=<value>[?color_scheme_id=<value>][?use_fields=<field_name>,<field_name>,...]
//
func PutNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutNomenclatureType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /nomenclature_types/<id>?name=<value>[?color_scheme_id=<value>][?use_fields=<field_name>,<field_name>,...]
//
func PostNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostNomenclatureType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /nomenclature_types/<id>
//
func DeleteNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteNomenclatureType"
	return
}
