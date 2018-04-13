package api

///////////////////////////////////////////////////////////////////////////////
// APIUser
type APIUser struct {
	ID       int    `json:"id,omitempty"`
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
	answer.Message = "GetUsers"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /users/<login>
//
func GetUser(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetUser"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /users?login=<login>[?name=<Value>][?phone=<value>][?position=<value>][?comment=<value>]
//
func PutUser(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutUser"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /users/<login>[?name=<Value>][?phone=<value>][?position=<value>][?comment=<value>]
//
func PostUser(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostUser"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /users/<login>
//
func DeleteUser(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteUser"
	return
}
