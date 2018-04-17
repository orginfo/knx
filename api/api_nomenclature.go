package api

///////////////////////////////////////////////////////////////////////////////
// APINomenclature
type APINomenclature struct {
	NomenclatureType            *APINomenclatureType `json:"nomenclature_type,omitempty"`
	ID                          int                  `json:"id,omitempty"`
	Name                        string               `json:"name,omitempty"`
	VendorCode                  string               `json:"vendor_code,omitempty"`
	MeasureUnit                 string               `json:"measure_unit,omitempty"`
	Material                    string               `json:"material,omitempty"`
	Thickness                   float32              `json:"thickness,omitempty"`
	Color                       *APIColor            `json:"color,omitempty"`
	Size                        float32              `json:"size,omitempty"`
	Price                       int                  `json:"price,omitempty"`
	Division                    []float32            `json:"division,omitempty"`
	DivisionServiceNomenclature *APINomenclature     `json:"division_service_nomencla,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request:
//a) GET /nomenclature
//   GET /nomenclature/<id>
//b) GET /nomenclature_types/<id>/nomenclature
//
/*Будут выведены только те поля, которые перечислены в nomenclature_type.use_fields*/
//
// Answer:
//	{
//		nomenclature_type: { /*Только для варианта a)*/
//			id 				int
//			name 			string
//			color_scheme_id int
//			use_fields: [<field_name>, <field_name>, ...]
//		}
//		id          int
//		name        string
//		vendor_code string
//		mesure_unit string
//		material    string
//		thickness   double
//		color: {
//			id 		int
//			name 	string
//			value 	int
//		}
//		size        double
//		price       int
//		division: [<value double>, <value double>, ...]
//		division_service_nomenclature: {
//			id   int
//			name string
//			vendor_code string
//			mesure_unit string
//			price       int
//		}
//	}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /nomenclature
//
func GetNomenclatures(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclatures"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /nomenclature/<id>
//
func GetNomenclature(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclature"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /nomenclature_types/<id>/nomenclature
//
func GetNomenclatureOfNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclatureOfNomenclatureType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /component_types/<id>/part_types/<id>/nomenclature
//
func GetNomenclatureForPartType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclatureForPartType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: /param_types/<id>/values/<value>/nomenclature
//
func GetNomenclatureForValueOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclatureForValueOfParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /nomenclature_types/<id>/nomenclature?name=<value>[?vendor_code=<value>][mesure_unit=<value>]
//	[?material=<value>][?thickness=<value>][?color_id=<value>][?size=<value>][division=<value>,<value>,...][division_service_nomenclature_id=<value>]
//
func PutNomenclature(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutNomenclature"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /nomenclature/<id>[?name=<value>][?vendor_code=<value>][mesure_unit=<value>]
//	[?material=<value>][?thickness=<value>][?color_id=<value>][?size=<value>][division=<value>,<value>,...][division_service_nomenclature_id=<value>]
//
func PostNomenclature(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostNomenclature"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /nomenclature/<id>
//
func DeleteNomenclature(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteNomenclature"
	return
}
