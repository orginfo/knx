package api

import (
	"database/sql"
	"fmt"
	"knx/db"
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////
// APIProject
type APIProject struct {
	ID           int64      `json:"id,omitempty"`
	Nr           string     `json:"nr,omitempty"`
	ContractDate *string    `json:"contract_date,omitempty"`
	InstallDate  *string    `json:"install_date,omitempty"`
	Address      string     `json:"address,omitempty"`
	Comment      string     `json:"comment,omitempty"`
	User         *APIUser   `json:"user,omitempty"`
	Client       *APIClient `json:"client,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
//
// Request: GET /clients/<id>/projects
//          GET /projects
//          GET /projects/<project_id>
// Answer:
//	[
//		{
//			id             int
//          nr             string
//			contract_date  string
//			install_date   string
//          address        string
//			comment        string
//			user: {
//			    id         int
//				name       string
//				phone      string
//				position   string
//				comment    string
//			}
//			client: {
//				id         int
//				name       string
//              phone      string
//				comment    string
//			}
//		}
//	]

///////////////////////////////////////////////////////////////////////////////
// Request: GET /clients/<id>/projects
//
func GetProjectsOfClient(request []string, params map[string][]string) (answer Answer) {
	var err error
	var res []APIProject
	defer answer.make(&err, &res)

	answer.ID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT p.id, p.nr, p.contract_date, p.install_date, p.address, p.comment,
		u.id, u.name, u.phone, u.position, u.comment,
		c.id, c.name, c.phone, c.comment
		FROM project p INNER JOIN user u ON p.user_id = u.id INNER JOIN client c ON p.client_id = c.id AND p.client_id=?`, answer.ID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u APIUser
		var c APIClient
		p := APIProject{User: &u, Client: &c}
		err = rows.Scan(&p.ID, &p.Nr, &p.ContractDate, &p.InstallDate, &p.Address, &p.Comment,
			&u.ID, &u.Name, &u.Phone, &u.Position, &u.Comment,
			&c.ID, &c.Name, &c.Phone, &c.Comment)
		if err != nil {
			return
		}
		res = append(res, p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /projects
//
func GetProjects(request []string, params map[string][]string) (answer Answer) {
	var err error
	var res []APIProject
	defer answer.make(&err, &res)

	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT p.id, p.nr, p.contract_date, p.install_date, p.address, p.comment,
		u.id, u.name, u.phone, u.position, u.comment,
		c.id, c.name, c.phone, c.comment
		FROM project p INNER JOIN user u ON p.user_id = u.id INNER JOIN client c ON p.client_id = c.id`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u APIUser
		var c APIClient
		p := APIProject{User: &u, Client: &c}
		err = rows.Scan(&p.ID, &p.Nr, &p.ContractDate, &p.InstallDate, &p.Address, &p.Comment,
			&u.ID, &u.Name, &u.Phone, &u.Position, &u.Comment,
			&c.ID, &c.Name, &c.Phone, &c.Comment)
		if err != nil {
			return
		}
		res = append(res, p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /projects/<project_id>
//
func GetProject(request []string, params map[string][]string) (answer Answer) {
	var err error
	var u APIUser
	var c APIClient
	res := APIProject{User: &u, Client: &c}
	defer answer.make(&err, &res)

	answer.ID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	var row *sql.Row
	row = db.DB.QueryRow(`SELECT p.nr, p.contract_date, p.install_date, p.address, p.comment,
		u.id, u.name, u.phone, u.position, u.comment,
		c.id, c.name, c.phone, c.comment
		FROM project p INNER JOIN user u ON p.user_id = u.id INNER JOIN client c ON p.client_id = c.id
		WHERE p.id=?`, answer.ID)
	err = row.Scan(&res.Nr, &res.ContractDate, &res.InstallDate, &res.Address, &res.Comment,
		&u.ID, &u.Name, &u.Phone, &u.Position, &u.Comment,
		&c.ID, &c.Name, &c.Phone, &c.Comment)
	if err != nil {
		return
	}
	res.ID = answer.ID
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /clients/<id>/projects?contract_date=<value>[?install_date=<value>][?comment=<value>][?address=<value>][?nr=<value>]
//
func PutProject(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	var clientID int64
	var userID int64 = 1 // Use 1 for default user, TODO: later make sessions and get user id from session, which opened by login-api
	clientID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	// Parse user request parameters
	var rp RequestParams = RequestParams{
		"contract_date": {Optional: false, Type: String},
		"install_date":  {Optional: true, Type: String},
		"comment":       {Optional: true, Type: String},
		"address":       {Optional: true, Type: String},
		"nr":            {Optional: true, Type: String},
	}

	err = rp.Parse(params)
	if err != nil {
		answer.Code = BadRequest
		return
	}

	// Add 2 parameters for SQL-query: user_id and client_id
	rp["user_id"] = RequestParam{Type: Int, Value: RequestParamValue{Type: Int, IntValue: userID}}
	rp["client_id"] = RequestParam{Type: Int, Value: RequestParamValue{Type: Int, IntValue: clientID}}

	// Insert into [project]
	sqlText, sqlParams := rp.MakeSQLInsert("project", []string{"contract_date", "install_date", "comment", "address", "nr", "user_id", "client_id"})
	var res sql.Result
	res, err = db.DB.Exec(sqlText, sqlParams...)
	if err != nil {
		return
	}

	answer.ID, err = res.LastInsertId()

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /project/<project_id>[?contract_date=<value>][?install_date=<value>][?comment=<value>][address=<value>][?nr=<value>]
//
func PostProject(request []string, params map[string][]string) (answer Answer) {
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
		"contract_date": {Optional: true, Type: String},
		"install_date":  {Optional: true, Type: String},
		"comment":       {Optional: true, Type: String},
		"address":       {Optional: true, Type: String},
		"nr":            {Optional: true, Type: String},
	}

	err = rp.Parse(params)
	if err != nil {
		answer.Code = BadRequest
		return
	}

	// Update [project]
	sqlText, sqlParams := rp.MakeSQLUpdate("project", []string{"contract_date", "install_date", "comment", "address", "nr"}, answer.ID)
	if len(sqlParams) > 0 {
		_, err = db.DB.Exec(sqlText, sqlParams...)
		if err != nil {
			return
		}
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /projects/<id>
//
func DeleteProject(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	answer.ID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	// Delete from [project]
	_, err = db.DB.Exec("DELETE FROM project WHERE id=?", answer.ID)
	return
}
