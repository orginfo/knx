package api

///////////////////////////////////////////////////////////////////////////////
// APIPrice
type APIPrice struct {
	Date      string `json:"date,omitempty"`
	Price     int    `json:"price,omitempty"`
	CostPrice int    `json:"cost_price,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /nomenclature/<id>/price
// Answer:
// [
//	{
//		date  		string
//		price 		int /*Цены в копейках*/
//		cost_price 	int
//	}
// ]
//
// Request: GET /nomenclature/<id>/price?date=<value>
// Answer:
// {
//      price 		int /*Цены в копейках*/
//	    cost_price 	int
// }

///////////////////////////////////////////////////////////////////////////////
// Request: GET /nomenclature/<id>/price
//
func GetPrice(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetPrice"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /nomenclature/<id>/price?date=<value>[?price=<value>][?cost_price=<value>]

//
func PutPrice(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutPrice"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /nomenclature/<id>/price?date=<value>[?price=<value>][?cost_price=<value>]
//
func PostPrice(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostPrice"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /nomenclature/<id>/price[?date=<value>]
//
func DeletePrice(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeletePrice"
	return
}
