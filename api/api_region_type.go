package api

import (
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
	answer.Code = OK
	var res []APIRegionType

	userNames := make(map[int]string)
	rows, err := db.DB.Query("SELECT id, name FROM tregion")
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		userNames[id] = name
	}
	err = rows.Err()

	for id, r := range calc.Regions {
		res = append(res, APIRegionType{ID: id, UserName: userNames[id], CodeName: r.Name})
	}
	answer.Result = res

	if err != nil {
		answer.Code = InternalServerError
		answer.Message = fmt.Sprintf("Ошибка SQL: %v", err)
	}
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /region_types/<id>
//
func GetRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Code = OK
	answer.Message = "GetRegionType"
	id, err := strconv.ParseInt(request[1], 10, 0)
	if err != nil || id < 0 || int(id) >= len(calc.Regions) {
		answer.Code = BadRequest
		answer.Message = fmt.Sprintf("Неверный ID '%s'", request[1])
		return
	}

	// Get code name from constant declaration
	res := APIRegionType{ID: int(id), CodeName: calc.Regions[id].Name}

	// Get user name from db
	rows, err := db.DB.Query("SELECT name FROM tregion WHERE id=?", int(id))
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&res.UserName)
	}
	err = rows.Err()

	if err != nil {
		answer.Code = InternalServerError
		answer.Message = fmt.Sprintf("Ошибка SQL: %v", err)
	} else {
		answer.Result = res
	}

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
