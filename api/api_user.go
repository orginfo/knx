package api

import (
	"database/sql"
	"knx/db"
)

///////////////////////////////////////////////////////////////////////////////
// APIUser
type APIUser struct {
	ID       int64  `json:"id,omitempty"`
	Login    string `json:"login,omitempty"`
	Name     string `json:"name,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Position string `json:"position,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /users
//
func GetUsers(request []string, params map[string][]string) (answer Answer) {
	var err error
	var res []APIUser
	defer answer.make(&err, &res)

	// Select data from table [user]
	var rows *sql.Rows
	rows, err = db.DB.Query("SELECT id, login, name, phone, position, comment FROM user")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u APIUser
		err = rows.Scan(&u.ID, &u.Login, &u.Name, &u.Phone, &u.Position, &u.Comment)
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
// Request: GET /users/<login>
//
func GetUser(request []string, params map[string][]string) (answer Answer) {
	var err error
	var res APIUser
	defer answer.make(&err, &res)

	// Select data from [user]
	var row *sql.Row
	row = db.DB.QueryRow("SELECT id, name, phone, position, comment FROM user WHERE login=?", request[1])
	err = row.Scan(&res.ID, &res.Name, &res.Phone, &res.Position, &res.Comment)
	// If no rows, just return empty result
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		return
	}
	answer.ID = res.ID
	res.Login = request[1]

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /users?login=<login>[?name=<Value>][?phone=<value>][?position=<value>][?comment=<value>]
//
func PutUser(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	// Parse user request parameters
	var rp RequestParams = RequestParams{
		"login":    {Optional: false, Type: String},
		"name":     {Optional: true, Type: String},
		"phone":    {Optional: true, Type: String},
		"position": {Optional: true, Type: String},
		"comment":  {Optional: true, Type: String},
	}

	err = rp.Parse(params)
	if err != nil {
		answer.Code = BadRequest
		return
	}

	// Insert into [user]
	sqlText, sqlParams := rp.MakeSQLInsert("user", []string{"login", "name", "phone", "position", "comment"})
	var res sql.Result
	res, err = db.DB.Exec(sqlText, sqlParams...)
	if err != nil {
		return
	}

	answer.ID, err = res.LastInsertId()

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /users/<login>[?name=<Value>][?phone=<value>][?position=<value>][?comment=<value>]
//
func PostUser(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	login := request[1]

	// Parse user request parameters
	var rp RequestParams = RequestParams{
		"name":     {Optional: true, Type: String},
		"phone":    {Optional: true, Type: String},
		"position": {Optional: true, Type: String},
		"comment":  {Optional: true, Type: String},
	}

	err = rp.Parse(params)
	if err != nil {
		answer.Code = BadRequest
		return
	}

	// Get user id
	var row *sql.Row
	row = db.DB.QueryRow("SELECT id FROM user WHERE login=?", login)
	err = row.Scan(&answer.ID)
	if err != nil {
		return
	}

	// Update [user]
	sqlText, sqlParams := rp.MakeSQLUpdate("user", []string{"name", "phone", "position", "comment"}, answer.ID)
	if len(sqlParams) > 0 {
		_, err = db.DB.Exec(sqlText, sqlParams...)
		if err != nil {
			return
		}
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /users/<login>
//
func DeleteUser(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	// Delete from [user]
	_, err = db.DB.Exec("DELETE FROM user WHERE login=?", request[1])
	return
}
