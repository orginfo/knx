package api

///////////////////////////////////////////////////////////////////////////////
// APIColor
type APIColor struct {
	Name  string `json:"name,omitempty"`
	Value int    `json:"value,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// APIColorScheme
type APIColorScheme struct {
	ID     int64      `json:"id,omitempty"`
	Name   string     `json:"name,omitempty"`
	Colors []APIColor `json:"colors,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /color_schemes
//
func GetColorSchemes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetColorSchemes"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /color_schemes/<id>
//
func GetColorScheme(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetColorScheme"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /color_schemes[?name=<value>][?color=<value>[(<name>)]&color=<value>[(<name>)]&...]
//
func PutColorScheme(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutColorScheme"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /color_schemes/<id>[?name=<value>][?color=<value>[(<name>)]&color=<value>[(<name>)]&...]
//
func PostColorSchemes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostColorSchemes"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /color_schemes/<id>
//
func DeleteColorScheme(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteColorScheme"
	return
}
