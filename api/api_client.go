package api

import (
	"database/sql"
	"fmt"
	"knx/db"
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////
// APIClient
type APIClient struct {
	ID      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Comment string `json:"comment,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /clients
//
func GetClients(request []string, params map[string][]string) (answer Answer) {
	var err error
	var res []APIClient
	defer answer.make(&err, &res)

	// Select data from table [client]
	var rows *sql.Rows
	rows, err = db.DB.Query("SELECT id, name, phone, comment FROM client")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u APIClient
		err = rows.Scan(&u.ID, &u.Name, &u.Phone, &u.Comment)
		if err != nil {
			return
		}
		res = append(res, u)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /clients/<id>
//
func GetClient(request []string, params map[string][]string) (answer Answer) {
	var err error
	var res APIClient
	defer answer.make(&err, &res)

	answer.ID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	// Select data from [client]
	var row *sql.Row
	row = db.DB.QueryRow("SELECT name, phone, comment FROM client WHERE id=?", answer.ID)
	err = row.Scan(&res.Name, &res.Phone, &res.Comment)
	if err != nil {
		return
	}
	res.ID = answer.ID
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /clients?name=<value>[?comment=<value>][?phone=<value>]
//
func PutClient(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	// Parse user request parameters
	var rp RequestParams = RequestParams{
		"name":    {Optional: false, Type: String},
		"phone":   {Optional: true, Type: String},
		"comment": {Optional: true, Type: String},
	}

	err = rp.Parse(params)
	if err != nil {
		answer.Code = BadRequest
		return
	}

	// Insert into [client]
	sqlText, sqlParams := rp.MakeSQLInsert("client", []string{"name", "phone", "comment"})
	var res sql.Result
	res, err = db.DB.Exec(sqlText, sqlParams...)
	if err != nil {
		return
	}

	answer.ID, err = res.LastInsertId()

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /clients/<id>[?name=<value>][?comment=<value>][?phone=<value>]
//
func PostClient(request []string, params map[string][]string) (answer Answer) {
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
		"name":    {Optional: true, Type: String},
		"phone":   {Optional: true, Type: String},
		"comment": {Optional: true, Type: String},
	}

	err = rp.Parse(params)
	if err != nil {
		answer.Code = BadRequest
		return
	}

	// Update [client]
	sqlText, sqlParams := rp.MakeSQLUpdate("client", []string{"name", "phone", "comment"}, answer.ID)
	if len(sqlParams) > 0 {
		_, err = db.DB.Exec(sqlText, sqlParams...)
		if err != nil {
			return
		}
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /clients/<id>
//
func DeleteClient(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	answer.ID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	// Delete from [client]
	_, err = db.DB.Exec("DELETE FROM client WHERE id=?", answer.ID)
	return
}
