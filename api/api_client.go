package api

///////////////////////////////////////////////////////////////////////////////
// APIClient
type APIClient struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Comment string `json:"comment,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /clients
//
func GetClients(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetClients"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /clients/<id>
//
func GetClient(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetClient"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /clients?name=<value>[?comment=<value>]
//
func PutClient(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutClient"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /clients/<id>[?name=<value>][?comment=<value>]
//
func PostClient(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostClient"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /clients/<id>
//
func DeleteClient(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteClient"
	return
}
