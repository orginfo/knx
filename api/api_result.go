package api

///////////////////////////////////////////////////////////////////////////////
// APIResult
type APIResult struct {
	ResultType   *APIResultType   `json:"result_type,omitempty"`
	Value        string           `json:"value,omitempty"`
	Nomenclature *APINomenclature `json:"nomenclature,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// APIRegionResults
type APIProjectResults struct {
	Region  *APIRegion  `json:"region,omitempty"`
	Results []APIResult `json:"results,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request:
// GET /projects/<id>/results
//
// Answer:
//[
//		region {
//			id             int
//			description    string
//			region_type   {
//			     id        int
//			     user_name string
//			     code_name string
//			}
//		},
//		results:
//		[
//				{
//					result_type  {id int, name string, description string}
//					value        string
//					nomenclature {id int, ...}
//				}
//		]
//]

///////////////////////////////////////////////////////////////////////////////
// Request: GET /projects/<id>/regions/<id>/results
//
func GetResultsOfRegion(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetResultsOfRegion"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /projects/<id>/results
//
func GetResultsOfProject(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetResultsOfProject"
	return
}
