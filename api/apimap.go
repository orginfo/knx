package api

// Коды ошибок API
type APIErrorCode int

const (
	OK                  = 200
	BadRequest          = 400
	InternalServerError = 500
)

// Описание ошибок API
var APIErrorDesc = [...]string{
	OK:                  "OK",
	BadRequest:          "Bad Request",           // некорректный запрос
	InternalServerError: "Internal server error", // внутренняя ошибка сервера
}

// Ответ. Используется как возвращаемое значение для Get/Put/Post/Delete команд
type Answer struct {
	Code    APIErrorCode
	Message string
	ID      int // id измененного(возвращаемого) элемента
	Result  interface{}
}

type HTTPCallbackFunc func([]string, map[string][]string) Answer

type HTTPCallbackSet struct {
	Get  HTTPCallbackFunc
	Put  HTTPCallbackFunc
	Post HTTPCallbackFunc
	Del  HTTPCallbackFunc
}

/*
/////////////////////////////////////////////////////////////////////////////////////////////////////////
// Карта всех вызовов функций для команд (get, put, post, del)

GET /region_types
GET /region_types/<id>
GET /region_types/<id>/param_types
GET /region_types/<id>/component_types

PUT /region_types/<id>?name=<Value>
PUT /region_types/<id>/param_types[?<param_id>,<param_id>,...]
PUT /region_types/<id>/component_types?<id>,<id>,...


POST /region_types/<id>?name=<Value>
POST /region_types/<id>/param_types[?<param_id>,<param_id>,...]
POST /region_types/<id>/component_types?<id>,<id>,...

--------------------------------------------------------------------------------------------------------

GET /param_types
GET /param_types/<id>/part_types
GET /param_types/<id>/values
GET /param_types/<id>/values/<value>/nomenclature

PUT /param_types
PUT /param_types/<id>/part_types?<id>,<id>,...
PUT /param_types/<id>/values[?<Value>(<Name>),<Value>(<Name>),...]
PUT /param_types/<id>/values/<value>/nomenclature?<id>,<id>,...

POST /param_types
POST /param_types/<id>/part_types?<id>,<id>,...
POST /param_types/<id>/values[?<Value>(<Name>),<Value>(<Name>),...]
POST /param_types/<id>/values/<value>/nomenclature?<id>,<id>,...

DELETE /param_types/<id>/values[?<Value>,<Value>,...]

--------------------------------------------------------------------------------------------------------

GET /result_types
GET /result_types/<id>

PUT /result_types/<id>?name=<Value>[?description=<Value>]

POST /result_types/<id>[?name=<Value>][?description=<Value>]

--------------------------------------------------------------------------------------------------------

GET /сalculation_types/

PUT /calculation_types/<id>?name=<value>

--------------------------------------------------------------------------------------------------------

GET /users
GET /users/<Login>

PUT /users?login=<Login>[?name=<Value>][?phone=<value>][?position=<value>][?comment=<value>]

POST /users/<login>[?name=<Value>][?phone=<value>][?position=<value>][?comment=<value>]

--------------------------------------------------------------------------------------------------------

GET /clients
GET /clients/<id>
GET /clients/<id>/projects

PUT /clients/<id>/projects?contract_date=<value>[?install_date=<value>][?comment=<value>]

--------------------------------------------------------------------------------------------------------
GET /projects
GET /projects/<id>
GET /projects/<id>/regions
GET /projects/<id>/regions/<nr>
GET /projects/<id>/regions/<nr>/results
GET /projects/<id>/regions/<nr>/components
GET /projects/<id>/regions/<nr>/components/<id>
GET /projects/<id>/regions/<nr>/components/<id>/parts
GET /projects/<id>/regions/<nr>/components/<id>/parts/<id>

PUT /projects/<id>/regions?region_type=<Value>[?params=<param_id>(<value>),<param_id>(<value>),...]
PUT /projects/<id>/regions/<nr>/results		//return error
PUT /projects/<id>/regions/<nr>/components?component_type=<id>,<id>,...

POST /projects/<id>[?contract_date=<value>][?install_date=<value>][?comment=<value>]
POST /projects/<id>/regions/<nr>[?params=<param_id>(<value>),<param_id>(<value>),...]
POST /projects/<id>/regions/<nr>/results	//return error
POST /projects/<id>/regions/<nr>/components?component_type=<id>,<id>,...

DELETE /projects/<id>/regions/<nr>

--------------------------------------------------------------------------------------------------------

GET /component_types
GET /component_types/<id>
GET /component_types/<id>/part_types
GET /component_types/<id>/part_types/<id>
GET /component_types/<id>/part_types/<id>/nomenclature

PUT /component_types?name=<value>
PUT /component_types/<id>/part_types?name=<value>[?calculation_type_id=<value>]
PUT /component_types/<id>/part_types/<id>/nomenclature?<id>,<id>,...

POST /component_types/<id>?name=<value>
POST /component_types/<id>/part_types/<id>[?name=<value>][?calculation_type_id=<value>]
POST /component_types/<id>/part_types/<id>/nomenclature?<id>,<id>,...

--------------------------------------------------------------------------------------------------------

GET /nomenclature_types
GET /nomenclature_types/<id>
GET /nomenclature_types/<id>/nomenclature
GET /nomenclature_types/<id>/nomenclature

PUT /nomenclature_types?name=<value>[?color_scheme_id=<value>][?use_fields=<field_name>,<field_name>,...]
PUT /nomenclature_types/<id>/nomenclature?name=<value>[?vendor_code=<value>][mesure_unit=<value>]
	[?material=<value>][?thickness=<value>][?color_id=<value>][?size=<value>][division=<value>,<value>,...][division_service_nomenclature_id=<value>]

POST /nomenclature_types/<id>?name=<value>[?color_scheme_id=<value>][?use_fields=<field_name>,<field_name>,...]
POST /nomenclature_types/<id>/nomenclature/<id>[?те же параметры, что и в PUT]

--------------------------------------------------------------------------------------------------------

GET /nomenclature
GET /nomenclature/<id>
GET /nomenclature/<id>/price
GET /nomenclature/<id>/price?date=<value>

PUT /nomenclature/<id>/price?date=<value>[?price=<value>][?cost_price=<value>]

POST /nomenclature/<id>/price?date=<value>[?price=<value>][?cost_price=<value>]

--------------------------------------------------------------------------------------------------------

GET /color_schemes
GET /color_schemes/<id>

PUT /color_schemes[?name=<value>][?colors=<value>[(<name>)],<value>[(<name>)],...]

POST /color_schemes/<id>[?name=<value>][?colors=<value>[(<name>)],<value>[(<name>)],...]

/////////////////////////////////////////////////////////////////////////////////////////////////////////
*/

var F map[string]HTTPCallbackSet = map[string]HTTPCallbackSet{
	"region_types":                    {GetRegionTypes, NotImplemented, NotImplemented, NotImplemented},
	"region_types<id>":                {GetRegionType, NotImplemented, PostRegionType, NotImplemented},
	"region_types<id>param_types":     {GetParamTypesOfRegionType, PutParamTypesOfRegionType, PostParamTypesOfRegionType, DeleteParamTypesOfRegionType},
	"region_types<id>component_types": {GetComponentTypesOfRegionType, PutComponentTypesOfRegionType, PostComponentTypesOfRegionType, DeleteComponentTypesOfRegionType},

	"param_types":                           {GetParamTypes, NotImplemented, NotImplemented, NotImplemented},
	"param_types<id>":                       {GetParamType, NotImplemented, PostParamType, NotImplemented},
	"param_types<id>part_types":             {GetPartTypesOfParamType, PutPartTypesOfParamType, PostPartTypesOfParamType, DeletePartTypesOfParamType},
	"param_types<id>values":                 {GetValuesOfParamType, PutValuesOfParamType, PostValuesOfParamType, DeleteValuesOfParamType},
	"param_types<id>values<id>nomenclature": {GetNomenclatureForValueOfParamType, PutNomenclatureForValueOfParamType, PostNomenclatureForValueOfParamType, DeleteNomenclatureForValueOfParamType},

	"result_types":     {GetResultTypes, NotImplemented, NotImplemented, NotImplemented},
	"result_types<id>": {GetResultType, NotImplemented, PostResultType, NotImplemented},

	"сalculation_types":     {GetCalculationTypes, NotImplemented, NotImplemented, NotImplemented},
	"сalculation_types<id>": {GetCalculationType, NotImplemented, PostCalculationType, NotImplemented},

	"users":     {GetUsers, PutUser, NotImplemented, NotImplemented},
	"users<id>": {GetUser, NotImplemented, PostUser, DeleteUser},

	"clients":             {GetClients, PutClient, NotImplemented, NotImplemented},
	"clients<id>":         {GetClient, NotImplemented, PostClient, NotImplemented},
	"clients<id>projects": {GetProjectsOfClient, PutProjectOfClient, NotImplemented, NotImplemented},

	"projects":                                       {GetProjects, NotImplemented, NotImplemented, NotImplemented},
	"projects<id>":                                   {GetProject, NotImplemented, PostProject, DeleteProject},
	"projects<id>regions":                            {GetRegionsOfProject, PutRegionOfProject, NotImplemented, NotImplemented},
	"projects<id>regions<id>":                        {GetRegionOfProject, NotImplemented, PostRegionOfProject, DeleteRegionOfProject},
	"projects<id>regions<id>results":                 {GetResultsOfRegion, NotImplemented, NotImplemented, NotImplemented},
	"projects<id>regions<id>components":              {GetComponentsOfRegion, PutComponentOfRegion, NotImplemented, NotImplemented},
	"projects<id>regions<id>components<id>":          {GetComponentOfRegion, NotImplemented, PostComponentOfRegion, DeleteComponentOfRegion},
	"projects<id>regions<id>components<id>parts":     {NotImplemented, NotImplemented, NotImplemented, NotImplemented},
	"projects<id>regions<id>components<id>parts<id>": {NotImplemented, NotImplemented, NotImplemented, NotImplemented},

	"component_types":                               {GetComponentTypes, PutComponentType, NotImplemented, NotImplemented},
	"component_types<id>":                           {GetComponentType, NotImplemented, PostComponentType, DeleteComponentType},
	"component_types<id>part_types":                 {GetPartTypesOfComponentType, PutPartTypeOfComponentType, NotImplemented, NotImplemented},
	"component_types<id>part_types<id>":             {GetPartTypeOfComponentType, NotImplemented, NotImplemented, DeletePartTypeOfComponentType},
	"component_types<id>part_types<id>nomenclature": {GetNomenclatureForPartType, PutNomenclatureForPartType, PostNomenclatureForPartType, DeleteNomenclatureForPartType},

	"nomenclature_types":                 {GetNomenclatureTypes, PutNomenclatureType, NotImplemented, NotImplemented},
	"nomenclature_types<id>":             {GetNomenclatureType, NotImplemented, PostNomenclatureType, DeleteNomenclatureType},
	"nomenclature_types<id>nomenclature": {GetNomenclatureOfNomenclatureType, PutNomenclatureOfNomenclatureType, NotImplemented, DeleteNomenclatureOfNomenclatureType},

	"nomenclature":          {GetNomenclatures, PutNomenclature, NotImplemented, NotImplemented},
	"nomenclature<id>":      {GetNomenclature, NotImplemented, PostNomenclature, DeleteNomenclature},
	"nomenclature<id>price": {GetPriceOfNomenclature, PutPriceOfNomenclature, PostPriceOfNomenclature, DeletePriceOfNomenclature},

	"color_schemes":     {GetColorSchemes, PutColorScheme, NotImplemented, NotImplemented},
	"color_schemes<id>": {GetColorScheme, NotImplemented, PostColorSchemes, DeleteColorScheme},
}
