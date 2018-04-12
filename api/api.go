// Реализация функций API
package api

import "fmt"

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

// make - defer функция, заполняющая ответ обработчика команды в конце каждого обработчика
func (a *Answer) make(err *error, result interface{}) {
	if (*err) != nil {
		if a.Code == 0 {
			a.Code = InternalServerError
		}
		a.Message = (*err).Error()
		a.Result = nil
	} else {
		a.Code = OK
		a.Message = ""
		a.Result = result
	}
}

///////////////////////////////////////////////////////////////////////////////
// NotImplemented - стандартный ответ на запросы, которые не должны обрабатываться,
// например, удаление типа забора
func NotImplemented(request []string, params map[string][]string) (answer Answer) {
	answer.Code = BadRequest
	answer.Message = fmt.Sprintf("Код ошибки: %d (%s). Функция не реализована.", BadRequest, APIErrorDesc[BadRequest])
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /region_types/<id>/param_types

// GetParamTypesOfRegionType -
func GetParamTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetParamTypesOfRegionType"
	return
}

// PutParamTypesOfRegionType -
func PutParamTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutParamTypesOfRegionType"
	return
}

// PostParamTypesOfRegionType -
func PostParamTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostRegionTypesParamTypes"
	return
}

// DeleteParamTypesOfRegionType - не может быть вызвана.
func DeleteParamTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteParamTypesOfRegionType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /region_types/<id>/component_types

// GetComponentTypesOfRegionType -
func GetComponentTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetComponentTypesOfRegionType"
	return
}

// PutComponentTypesOfRegionType -
func PutComponentTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutComponentTypesOfRegionType"
	return
}

// PostComponentTypesOfRegionType -
func PostComponentTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostComponentTypesOfRegionType"
	return
}

// DeleteComponentTypesOfRegionType -
func DeleteComponentTypesOfRegionType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteComponentTypesOfRegionType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /param_types

// GetParamTypes -
func GetParamTypes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetParamTypes"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /param_types/<id>

// GetParamType -
func GetParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetParamType"
	return
}

// PostParamType -
func PostParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /param_types/<id>/part_types

// GetPartTypesOfParamType -
func GetPartTypesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetPartTypesOfParamType"
	return
}

// PutPartTypesOfParamType -
func PutPartTypesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutPartTypesOfParamType"
	return
}

// PostPartTypesOfParamType -
func PostPartTypesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostPartTypesOfParamType"
	return
}

// DeletePartTypesOfParamType -
func DeletePartTypesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeletePartTypesOfParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /param_types/<id>/values

// GetValuesOfParamType -
func GetValuesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetValuesOfParamType"
	return
}

// PutValuesOfParamType -
func PutValuesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutValuesOfParamType"
	return
}

// PostValuesOfParamType -
func PostValuesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostValuesOfParamType"
	return
}

// DeleteValuesOfParamType -
func DeleteValuesOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteParamTypesValues"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /param_types/<id>/values/<id>/nomenclature

// GetNomenclatureForValueOfParamType -
func GetNomenclatureForValueOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclatureForValueOfParamType"
	return
}

// PutNomenclatureForValueOfParamType -
func PutNomenclatureForValueOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutNomenclatureForValueOfParamType"
	return
}

// PostNomenclatureForValueOfParamType -
func PostNomenclatureForValueOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostNomenclatureForValueOfParamType"
	return
}

// DeleteNomenclatureForValueOfParamType -
func DeleteNomenclatureForValueOfParamType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteNomenclatureForValueOfParamType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /result_types

// GetResultTypes -
func GetResultTypes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetResultTypes"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /result_types/<id>

// GetResultType -
func GetResultType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetResultType"
	return
}

// PostResultType -
func PostResultType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostResultType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /сalculation_types

// GetCalculationTypes -
func GetCalculationTypes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetCalculationTypes"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /сalculation_types/<id>

// GetCalculationType -
func GetCalculationType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetCalculationType"
	return
}

// PostCalculationType -
func PostCalculationType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostCalculationType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /users

// GetUsers -
func GetUsers(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetUsers"
	return
}

// PutUser -
func PutUser(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutUser"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /users/<id>

// GetUser -
func GetUser(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetUser"
	return
}

// PostUser -
func PostUser(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostUser"
	return
}

// DeleteUser -
func DeleteUser(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteUser"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /clients

// GetClients -
func GetClients(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetClients"
	return
}

// PutClient -
func PutClient(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutClient"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /clients/<id>

// GetClient -
func GetClient(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetClient"
	return
}

// PostClient -
func PostClient(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostClient"
	return
}

// DeleteClient - не может быть вызвана.
func DeleteClient(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteClient"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /clients/<id>/projects

// GetProjectsOfClient -
func GetProjectsOfClient(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetProjectsOfClient"
	return
}

// PutProjectOfClient -
func PutProjectOfClient(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutProjectOfClient"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /projects

// GetProjects -
func GetProjects(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetProjects"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /projects/<id>

// GetProject -
func GetProject(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetProject"
	return
}

// PostProject -
func PostProject(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostProject"
	return
}

// DeleteProject -
func DeleteProject(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteProject"
	return
}

///////////////////////////////////////////////////////////////////////////////////
// URL path: /projects/<id>/regions

// GetRegionsOfProject -
func GetRegionsOfProject(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetRegionsOfProject"
	return
}

// PutRegionOfProject -
func PutRegionOfProject(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutRegionOfProject"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /projects/<id>/regions/<id>/results

// GetResultsOfRegion -
func GetResultsOfRegion(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetResultsOfRegion"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /projects/<id>/regions/<id>/components

// GetComponentsOfRegion -
func GetComponentsOfRegion(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetComponentsOfRegion"
	return
}

// PutComponentOfRegion -
func PutComponentOfRegion(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutComponentOfRegion"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /projects/<id>/regions/<id>/components/<id>

// GetComponentOfRegion -
func GetComponentOfRegion(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetComponentOfRegion"
	return
}

// PostComponentOfRegion -
func PostComponentOfRegion(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostComponentOfRegion"
	return
}

// DeleteComponentOfRegion -
func DeleteComponentOfRegion(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteComponentOfRegion"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /projects/<id>/regions/<id>/components/<id>/parts

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /projects/<id>/regions/<id>/components/<id>/parts/<id>

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /component_types

// GetComponentTypes -
func GetComponentTypes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetComponentTypes"
	return
}

// PutComponentType -
func PutComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutComponentType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /component_types/<id>

// GetComponentType -
func GetComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetComponentType"
	return
}

// PostComponentType -
func PostComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostComponentType"
	return
}

// DeleteComponentType -
func DeleteComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteComponentType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /component_types/<id>/part_types

// GetPartTypesOfComponentType -
func GetPartTypesOfComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetPartTypesOfComponentType"
	return
}

// PutPartTypeOfComponentType -
func PutPartTypeOfComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutPartTypeOfComponentType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /component_types/<id>/part_types<id>

// GetPartTypeOfComponentType -
func GetPartTypeOfComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetPartTypeOfComponentType"
	return
}

// DeletePartTypeOfComponentType -
func DeletePartTypeOfComponentType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeletePartTypeOfComponentType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /component_types/<id>/part_types/<id>/nomenclature

// GetNomenclatureForPartType -
func GetNomenclatureForPartType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclatureForPartType"
	return
}

// PutNomenclatureForPartType -
func PutNomenclatureForPartType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutNomenclatureForPartType"
	return
}

// PostNomenclatureForPartType -
func PostNomenclatureForPartType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostNomenclatureForPartType"
	return
}

// DeleteNomenclatureForPartType -
func DeleteNomenclatureForPartType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteNomenclatureForPartType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /nomenclature_types

// GetNomenclatureTypes -
func GetNomenclatureTypes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclatureTypes"
	return
}

// PutNomenclatureType -
func PutNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutNomenclatureType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /nomenclature_types/<id>

// GetNomenclatureType -
func GetNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclatureType"
	return
}

// PostNomenclatureType -
func PostNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostNomenclatureType"
	return
}

// DeleteNomenclatureType -
func DeleteNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteNomenclatureType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /nomenclature_types/<id>/nomenclature

// GetNomenclatureOfNomenclatureType -
func GetNomenclatureOfNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclatureOfNomenclatureType"
	return
}

// PutNomenclatureOfNomenclatureType -
func PutNomenclatureOfNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutNomenclatureOfNomenclatureType"
	return
}

// DeleteNomenclatureOfNomenclatureType -
func DeleteNomenclatureOfNomenclatureType(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteNomenclatureOfNomenclatureType"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /nomenclature

// GetNomenclatures -
func GetNomenclatures(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclatures"
	return
}

// PutNomenclature -
func PutNomenclature(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutNomenclature"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /nomenclature/<id>

// GetNomenclature -
func GetNomenclature(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetNomenclature"
	return
}

// PostNomenclature -
func PostNomenclature(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostNomenclature"
	return
}

// DeleteNomenclature -
func DeleteNomenclature(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteNomenclature"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /nomenclature/<id>/price

// GetPriceOfNomenclature -
func GetPriceOfNomenclature(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetPriceOfNomenclature"
	return
}

// PutPriceOfNomenclature -
func PutPriceOfNomenclature(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutPriceOfNomenclature"
	return
}

// PostPriceOfNomenclature -
func PostPriceOfNomenclature(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostPriceOfNomenclature"
	return
}

// DeletePriceOfNomenclature -
func DeletePriceOfNomenclature(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeletePriceOfNomenclature"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /color_schemes

// GetColorSchemes -
func GetColorSchemes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetColorSchemes"
	return
}

// PutColorScheme -
func PutColorScheme(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutColorScheme"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Путь URL: /color_schemes/<id>

// GetColorScheme -
func GetColorScheme(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "GetColorScheme"
	return
}

// PostColorSchemes -
func PostColorSchemes(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostColorSchemes"
	return
}

// DeleteColorScheme -
func DeleteColorScheme(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteColorScheme"
	return
}
