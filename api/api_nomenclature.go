package api

import (
	"database/sql"
	"fmt"
	"knx/db"
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////
// APINomenclature
type APINomenclature struct {
	NomenclatureType            *APINomenclatureType `json:"nomenclature_type,omitempty"`
	ID                          int64                `json:"id,omitempty"`
	Name                        string               `json:"name,omitempty"`
	VendorCode                  string               `json:"vendor_code,omitempty"`
	MeasureUnit                 string               `json:"measure_unit,omitempty"`
	Material                    string               `json:"material,omitempty"`
	Thickness                   float64              `json:"thickness,omitempty"`
	Color                       *APIColor            `json:"color,omitempty"`
	Size                        float64              `json:"size,omitempty"`
	Price                       int                  `json:"price,omitempty"`
	Division                    []float64            `json:"division,omitempty"`
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
	var err error
	defer answer.make(&err, nil)

	var NomenclatureTypeID int64
	NomenclatureTypeID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	// Parse user request parameters
	var rp RequestParams = RequestParams{
		"name":                             {Optional: false, Type: String},
		"vendor_code":                      {Optional: true, Type: String},
		"measure_unit":                     {Optional: true, Type: String},
		"material":                         {Optional: true, Type: String},
		"thickness":                        {Optional: true, Type: Float},
		"color_id":                         {Optional: true, Type: Int},
		"size":                             {Optional: true, Type: Float},
		"division":                         {Optional: true, Type: String},
		"division_service_nomenclature_id": {Optional: true, Type: Int},
	}

	err = rp.Parse(params)
	if err != nil {
		answer.Code = BadRequest
		return
	}

	// Add SQL-parameter [tnomenclature_id]
	rp["tnomenclature_id"] = RequestParam{Optional: false, Type: Int, Value: RequestParamValue{Type: Int, IntValue: NomenclatureTypeID}}

	// Insert into [tnomenclature]
	sqlText, sqlParams := rp.MakeSQLInsert("nomenclature", []string{"tnomenclature_id", "name", "vendor_code", "measure_unit",
		"material", "thickness", "color_id", "size", "division", "division_service_nomenclature_id"})
	var res sql.Result
	res, err = db.DB.Exec(sqlText, sqlParams...)
	if err != nil {
		return
	}

	answer.ID, err = res.LastInsertId()
	if err != nil {
		return
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /nomenclature/<id>[?name=<value>][?vendor_code=<value>][mesure_unit=<value>]
//	[?material=<value>][?thickness=<value>][?color_id=<value>][?size=<value>][division=<value>,<value>,...][division_service_nomenclature_id=<value>]
//
func PostNomenclature(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	answer.ID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	// Parse user request parameters
	var rp RequestParams = RequestParams{
		"name":                             {Optional: true, Type: String},
		"vendor_code":                      {Optional: true, Type: String},
		"measure_unit":                     {Optional: true, Type: String},
		"material":                         {Optional: true, Type: String},
		"thickness":                        {Optional: true, Type: Float},
		"color_id":                         {Optional: true, Type: Int},
		"size":                             {Optional: true, Type: Float},
		"division":                         {Optional: true, Type: String},
		"division_service_nomenclature_id": {Optional: true, Type: Int},
	}

	err = rp.Parse(params)
	if err != nil {
		answer.Code = BadRequest
		return
	}

	// Update [tnomenclature]
	sqlText, sqlParams := rp.MakeSQLUpdate("nomenclature", []string{"name", "vendor_code", "measure_unit",
		"material", "thickness", "color_id", "size", "division", "division_service_nomenclature_id"}, answer.ID)
	if len(sqlParams) > 0 {
		_, err = db.DB.Exec(sqlText, sqlParams...)
		if err != nil {
			return
		}
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /nomenclature/<id>
//
func DeleteNomenclature(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	answer.ID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	// Delete from [nomenclature]
	_, err = db.DB.Exec("DELETE FROM nomenclature WHERE id=?", answer.ID)
	return
}
