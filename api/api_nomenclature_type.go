package api

import (
	"database/sql"
	"fmt"
	"knx/db"
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////
// APINomenclatureType
type APINomenclatureType struct {
	ID            int      `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	ColorSchemeID *int     `json:"color_scheme_id,omitempty"` // Can be nil if no color scheme for this type
	UseFields     []string `json:"use_fields,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /nomenclature_types
//
func GetNomenclatureTypes(request []string, params map[string][]string) (answer Answer) {
	var err error
	var res []APINomenclatureType
	defer answer.make(&err, &res)

	// Map of use_fileds
	useFields := make(map[int][]string)

	// Select use_fields
	var rows *sql.Rows
	rows, err = db.DB.Query("SELECT tnomenclature_id, field_name FROM cn_tnomenclature_usefield ORDER BY tnomenclature_id")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var field string
		err = rows.Scan(&id, &field)
		if err != nil {
			return
		}
		useFields[id] = append(useFields[id], field)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	rows.Close()

	// Select data from main table [tnomenclature]
	rows, err = db.DB.Query("SELECT id, name, color_scheme_id FROM tnomenclature")
	if err != nil {
		return
	}
	for rows.Next() {
		var t APINomenclatureType
		err = rows.Scan(&t.ID, &t.Name, &t.ColorSchemeID)
		if err != nil {
			return
		}
		t.UseFields = useFields[t.ID]
		res = append(res, t)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /nomenclature_types/<id>
//
func GetNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	var err error
	var res APINomenclatureType
	defer answer.make(&err, &res)

	var id int64
	id, err = strconv.ParseInt(request[1], 10, 0)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	// Select data from main table [tnomenclature]
	var row *sql.Row
	row = db.DB.QueryRow("SELECT name, color_scheme_id FROM tnomenclature WHERE id=?", int(id))
	err = row.Scan(&res.Name, &res.ColorSchemeID)
	// If no rows, just return empty result
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		return
	}
	res.ID = int(id)

	// Select use_fields
	var rows *sql.Rows
	rows, err = db.DB.Query("SELECT field_name FROM cn_tnomenclature_usefield WHERE tnomenclature_id=?", int(id))
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var field string
		err = rows.Scan(&field)
		if err != nil {
			return
		}
		res.UseFields = append(res.UseFields, field)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /nomenclature_types?name=<value>[?color_scheme_id=<value>][?use_fields=<field_name>,<field_name>,...]
//
func PutNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	// Parse user request parameters
	var rp RequestParams = RequestParams{
		"name":            {Optional: false, Type: String},
		"color_scheme_id": {Optional: true, Type: Int},
		"use_fields":      {Optional: true, Type: StringArray},
	}

	err = rp.Parse(params)
	if err != nil {
		answer.Code = BadRequest
		return
	}

	// Insert into [tnomenclature]
	sqlText, sqlParams := rp.MakeSQLInsert("tnomenclature", []string{"name", "color_scheme_id"})
	var res sql.Result
	res, err = db.DB.Exec(sqlText, sqlParams...)
	if err != nil {
		return
	}

	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return
	}
	answer.ID = int(id)

	// Insert into [cn_tnomenclature_usefield]
	pruf, _ := rp["use_fields"]
	if !pruf.Exists() {
		return
	}

	var ufs []string = rp["use_fields"].Value.StringArray
	for _, uf := range ufs {
		_, err = db.DB.Exec("INSERT INTO cn_tnomenclature_usefield(tnomenclature_id, field_name) VALUES(?,?)", answer.ID, uf)
		if err != nil {
			return
		}
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /nomenclature_types/<id>[?name=<value>][?color_scheme_id=<value>][?use_fields=<field_name>,<field_name>,...]
//
func PostNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	var id int64
	id, err = strconv.ParseInt(request[1], 10, 0)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	// Parse user request parameters
	var rp RequestParams = RequestParams{
		"name":            {Optional: true, Type: String},
		"color_scheme_id": {Optional: true, Type: Int},
		"use_fields":      {Optional: true, Type: StringArray},
	}

	err = rp.Parse(params)
	if err != nil {
		answer.Code = BadRequest
		return
	}

	// Update [tnomenclature]
	sqlText, sqlParams := rp.MakeSQLUpdate("tnomenclature", []string{"name", "color_scheme_id"}, int(id))
	if len(sqlParams) > 0 {
		_, err = db.DB.Exec(sqlText, sqlParams...)
		if err != nil {
			return
		}
	}

	// Update [cn_tnomenclature_usefield]
	pruf, _ := rp["use_fields"]
	if !pruf.Exists() {
		return
	}

	_, err = db.DB.Exec("DELETE FROM cn_tnomenclature_usefield WHERE tnomenclature_id=?", int(id))
	if err != nil {
		return
	}

	var ufs []string = rp["use_fields"].Value.StringArray
	for _, uf := range ufs {
		_, err = db.DB.Exec("INSERT INTO cn_tnomenclature_usefield(tnomenclature_id, field_name) VALUES(?,?)", int(id), uf)
		if err != nil {
			return
		}
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /nomenclature_types/<id>
//
func DeleteNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	var id int64
	id, err = strconv.ParseInt(request[1], 10, 0)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	// Delete from [tnomenclature]
	_, err = db.DB.Exec("DELETE FROM tnomenclature WHERE id=?", int(id))
	return
}
