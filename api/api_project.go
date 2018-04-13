package api

///////////////////////////////////////////////////////////////////////////////
// APIProject
type APIProject struct {
	ID           int        `json:"id,omitempty"`
	ContractDate string     `json:"contract_date,omitempty"`
	InstallDate  string     `json:"install_date,omitempty"`
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
//			contract_date  string
//			install_date   string
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
//				comment    string
//			}
//		}
//	]

///////////////////////////////////////////////////////////////////////////////
// Request: GET /clients/<id>/projects
//
func GetProjectsOfClient(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetProjectsOfClient"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /projects
//
func GetProjects(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetProjects"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /projects/<project_id>
//
func GetProject(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetProject"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /clients/<id>/projects?contract_date=<value>[?install_date=<value>][?comment=<value>]
//
func PutProject(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutProject"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /project/<project_id>[?contract_date=<value>][?install_date=<value>][?comment=<value>]
//
func PostProject(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostProject"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /projects/<id>
//
func DeleteProject(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteProject"
	return
}
