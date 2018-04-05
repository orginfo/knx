package api

import (
	"fmt"
	"knx/db"
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////
//
// Request: GET /projects/<id>/regions/<id>
// Answer:
//		{
//			id             int
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
//					/*Секция может не содержать параметров, если для номенклатура секции от параметров не зависит*/
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
// Request: GET /projects/<id>/regions/<id>
func GetRegionOfProject(request []string, params map[string][]string) (answer Answer) {
	answer.Code = OK
	answer.Message = "GetRegionOfProject"
	projectID, err := strconv.ParseInt(request[1], 10, 0)
	if err != nil {
		answer.Code = BadRequest
		answer.Message = fmt.Sprintf("Неверный ID проекта '%s'", request[1])
		return
	}

	regionID, err := strconv.ParseInt(request[3], 10, 0)
	if err != nil {
		answer.Code = BadRequest
		answer.Message = fmt.Sprintf("Неверный ID участка '%s'", request[3])
		return
	}

	rows, err := db.DB.Query(
		`SELECT tregion_id, tregion.name
		FROM region LEFT JOIN tregion ON region.tregion_id = tregion.id
		WHERE project_id=? AND id=?`, int(projectID), int(regionID))
	defer rows.Close()
	for rows.Next() {
		var rt APIRegionType
		err = rows.Scan(&rt.ID, &rt.UserName)
	}
	err = rows.Err()

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /projects/<id>/regions/<id>[?param=<param_id>(<value>)&param=<param_id>(<value>)&...]
//
func PostRegionOfProject(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "PostRegionOfProject"
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /projects/<id>/regions/<id>
//
func DeleteRegionOfProject(request []string, params map[string][]string) (answer Answer) {
	answer.Message = "DeleteRegionOfProject"
	return
}
