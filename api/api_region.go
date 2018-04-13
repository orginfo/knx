package api

import (
	"fmt"
	"knx/calc"
	"knx/db"
	"strconv"
)

// ControlType - тип элемента управления в GUI
type ControlType int

const (
	CTCheckBox ControlType = iota + 1
	CTEditBox
	CTNumericEditBox
	CTChooseColorBtn
	CTComboBox
	CTEditableComboBox
)

///////////////////////////////////////////////////////////////////////////////
// APIParam
type APIParam struct {
	ParamType APIParamType    `json:"param_type,omitempty"`
	Value     float64         `json:"value,omitempty"`
	Control   ControlType     `json:"control,omitempty"` //Тип элемента интерфейса: поле для ввода, комбо-бокс, галочка,...
	Enabled   bool            `json:"enabled,omitempty"`
	Visible   bool            `json:"visible,omitempty"`
	ValueList []APIParamValue `json:"value_list,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// APIPartNomenclature
type APIPartNomenclature struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	List  bool   `json:"list,omitempty"`  // Показывает, есть ли список для выбора номенклатуры, или это значение не может поменяться
	Empty bool   `json:"empty,omitempty"` // Показывает, может ли не указывыть номенклатуру для этой части
}

///////////////////////////////////////////////////////////////////////////////
// APIPart
type APIPart struct {
	ID                int                 `json:"id,omitempty"`
	Name              string              `json:"name,omitempty"`
	CalculationTypeID int                 `json:"calculation_type_id,omitempty"`
	Nomenclature      APIPartNomenclature `json:"nomenclature,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// APIRegion
type APIRegion struct {
	ID          int            `json:"id,omitempty"`
	Description string         `json:"description,omitempty"`
	RegionType  *APIRegionType `json:"region_type,omitempty"`
	Params      []APIParam     `json:"params,omitempty"`
	Parts       []APIPart      `json:"parts,omitempty"`
	Components  []APIComponent `json:"components,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
//
// Request: GET /projects/<id>/regions/<id>
// Answer:
//		{
//			id             int
//          descrption     string
//			region_type   {
//			     id        int
//			     user_name string
//			     code_name string
//			}
//			params [
//				{
//					param_type   {id int, name string, description string}
//					value        float
//					control      int /*Тип элемента интерфейса: поле для ввода, комбо-бокс, галочка,...*/
//					/*1: checkbox
//					2: editbox для любого текста
//					3: editbox для числа
//					4: кнопка выбора цвета
//					5: combobox нередактируемый
//					6: combobox редактируемый с фильтрацией*/
//					enabled      bool
//					visible      bool
//					value_list   [
//						{
//							value float
//							name string
//						}
//					]
//				}
//			]
//			parts: [
//				{
//					id   int
//					name string
//					calculation_type_id int
//					nomenclature {
//						id    int
//						name  string
//						list  bool /*показывает, есть ли список для выбора номенклатуры, или это значение не может поменяться*/
//						empty bool /*показывает, может ли не указывыть номенклатуру для этой части*/
//					}
//				}
//			]
//			components [
//				{
//					id  int
//					component_type: {
//						id   int
//						name string
//					}
//
//					/*Для GUI:*/
//					/*Содержимое компонента разбито по секциям, каждая секция содержит параметры и части для выбора номенклатуры*/
//					/*Параметры одной секции влияют на выбор номенклатуры этой секции и, возможно, на выбор других параметров секции*/
//					/*Секция может не содержать параметров, если номенклатура секции от параметров не зависит*/
//					sections: [
//						{
//							params []int /*id параметров из массива params этого участка*/
//							parts  []int /*id частей из массива parts этого участка*/
//						}
//					]
//				}
//			]
//		}
//

///////////////////////////////////////////////////////////////////////////////
// Request: GET /projects/<id>/regions
//
func GetRegionsOfProject(request []string, params map[string][]string) (answer Answer) {
	var err error
	var res APIRegion
	defer answer.make(&err, &res)

	answer.Message = "GetRegionsOfProject"
	_, err = strconv.ParseInt(request[1], 10, 0)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID проекта '%s'", request[1])
		return
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /projects/<id>/regions/<id>
//
func GetRegion(request []string, params map[string][]string) (answer Answer) {
	var err error
	var res APIRegion
	defer answer.make(&err, &res)

	answer.Message = "GetRegion"
	projectID, err := strconv.ParseInt(request[1], 10, 0)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID проекта '%s'", request[1])
		return
	}

	regionID, err := strconv.ParseInt(request[3], 10, 0)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID участка '%s'", request[3])
		return
	}

	rows, err := db.DB.Query(
		`SELECT tregion_id, tregion.name, region.description
		FROM region LEFT JOIN tregion ON region.tregion_id = tregion.id
		WHERE project_id=? AND region.id=?`, int(projectID), int(regionID))

	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&res.RegionType.ID, &res.RegionType.UserName, &res.Description)
		if err != nil {
			return
		}

		// Get code name of region type
		if res.RegionType.ID < 0 || res.RegionType.ID >= len(calc.Regions) {
			err = fmt.Errorf("Неверный тип учатска '%d'", res.RegionType.ID)
			return
		}
		res.RegionType.CodeName = calc.Regions[res.RegionType.ID].Name

		//
	}
	err = rows.Err()

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /projects/<id>/regions?region_type=<Value>[?description=<value>][?params=<param_id>(<value>),<param_id>(<value>),...]
//
func PutRegion(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PutRegion"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /projects/<id>/regions/<id>[?description=<value>][?param=<param_id>(<value>)&param=<param_id>(<value>)&...]
//
func PostRegion(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostRegion"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /projects/<id>/regions/<id>
//
func DeleteRegion(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteRegion"
	return
}
