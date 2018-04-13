package api

import (
	"database/sql"
	"fmt"
	"knx/calc"
	"knx/db"
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////
// APIRegionType
type APIRegionType struct {
	ID       int    `json:"id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	CodeName string `json:"code_name,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /region_types
//
func GetRegionTypes(request []string, params map[string][]string) (answer Answer) {
	var err error
	var res []APIRegionType
	defer answer.make(&err, &res)

	userNames := make(map[int]string)
	var rows *sql.Rows
	rows, err = db.DB.Query("SELECT id, name FROM tregion")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return
		}
		userNames[id] = name
	}
	err = rows.Err()
	if err != nil {
		return
	}
	for id, r := range calc.Regions {
		res = append(res, APIRegionType{ID: id, UserName: userNames[id], CodeName: r.Name})
	}
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /region_types/<id>
//
func GetRegionType(request []string, params map[string][]string) (answer Answer) {
	var err error
	var res APIRegionType
	defer answer.make(&err, &res)
	var id int64
	id, err = strconv.ParseInt(request[1], 10, 0)
	if err != nil || id < 0 || int(id) >= len(calc.Regions) {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	// Get code name from constant declaration
	res = APIRegionType{ID: int(id), CodeName: calc.Regions[id].Name}

	// Get user name from db
	var rows *sql.Rows
	rows, err = db.DB.Query("SELECT name FROM tregion WHERE id=?", int(id))
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&res.UserName)
		if err != nil {
			return
		}
	}
	err = rows.Err()
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /region_types/<id>?name=<Value>
//
func PostRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Code = OK
	answer.Message = "PostRegionType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /region_types/<id>/param_types
//
func GetParamTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetParamTypesOfRegionType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /region_types/<id>/param_types[?<param_id>,<param_id>,...]
//
func PutParamTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutParamTypesOfRegionType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /region_types/<id>/param_types[?<param_id>,<param_id>,...]
//
func PostParamTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostRegionTypesParamTypes"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /region_types/<id>/param_types[?<param_id>,<param_id>,...]
//
func DeleteParamTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteParamTypesOfRegionType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /region_types/<id>/component_types
//
func GetComponentTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetComponentTypesOfRegionType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /region_types/<id>/component_types[?<id>,<id>,...]
//
func PutComponentTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutComponentTypesOfRegionType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /region_types/<id>/component_types[?<id>,<id>,...]
//
func PostComponentTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostComponentTypesOfRegionType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /region_types/<id>/component_types[?<id>,<id>,...]
//
func DeleteComponentTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteComponentTypesOfRegionType"
	return
}
