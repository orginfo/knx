package api

import (
	"bytes"
	"database/sql"
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
	ParamType *APIParamType   `json:"param_type,omitempty"`
	Value     float64         `json:"value,omitempty"`
	Control   ControlType     `json:"control,omitempty"` //Тип элемента интерфейса: поле для ввода, комбо-бокс, галочка,...
	Enabled   bool            `json:"enabled,omitempty"`
	Visible   bool            `json:"visible,omitempty"`
	ValueList []APIParamValue `json:"value_list,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// APIPart
type APIPart struct {
	APIPartType
	Value     *APINomenclature   `json:"value,omitempty"`
	ValueList []*APINomenclature `json:"value_list,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// APIRegion
type APIRegion struct {
	ID          int64                  `json:"id,omitempty"`
	Description string                 `json:"description,omitempty"`
	RegionType  *APIRegionType         `json:"region_type,omitempty"`
	Params      map[int64]APIParam     `json:"params,omitempty"`
	Parts       map[int64]APIPart      `json:"parts,omitempty"`
	Components  map[int64]APIComponent `json:"components,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
//
// Request: GET /projects/<id>/regions
// Answer:
//      {
//          id             int
//          description     string
//          region_type   {
//               id        int
//               user_name string
//               code_name string
//          }
//      }
//
// Request: GET /projects/<id>/regions/<id>
// Answer:
//		{
//			id             int
//          description     string
//			region_type   {
//			     id        int
//			     user_name string
//			     code_name string
//			}
//			params map(int)
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
//			parts: map(int)
//				{
//					id   int
//					name string
//					calculation_type_id int
//					value {
//						id    int
//						name  string
//					}
//                  value_list	[
//						{
//							id int
//							name string
//						}
//					]
//				}
//			components map(int)
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
	var res []APIRegion
	defer answer.make(&err, &res)

	var projectID int64
	projectID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID проекта '%s'", request[1])
		return
	}

	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT r.id, r.description, t.id, t.name
		FROM region r INNER JOIN tregion t ON r.tregion_id = t.id WHERE r.project_id=? ORDER BY r.nr`, projectID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var r APIRegion
		r.RegionType = new(APIRegionType)
		err = rows.Scan(&r.ID, &r.Description, &r.RegionType.ID, &r.RegionType.UserName)
		if err != nil {
			return
		}
		// Get code name of region type
		if r.RegionType.ID < 0 || r.RegionType.ID >= int64(len(calc.Regions)) {
			err = fmt.Errorf("Неверный тип '%d' участка '%d'", r.RegionType.ID, r.ID)
			return
		}
		r.RegionType.CodeName = calc.Regions[r.RegionType.ID].Name
		res = append(res, r)
	}
	err = rows.Err()
	if err != nil {
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
	res.RegionType = new(APIRegionType)
	defer answer.make(&err, &res)

	var projectID int64
	projectID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID проекта '%s'", request[1])
		return
	}

	answer.ID, err = strconv.ParseInt(request[3], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID участка '%s'", request[3])
		return
	}
	res.ID = answer.ID

	var row *sql.Row
	row = db.DB.QueryRow(
		`SELECT r.tregion_id, t.name, r.description
		FROM region r LEFT JOIN tregion t ON r.tregion_id = t.id
		WHERE r.project_id=? AND r.id=?`, projectID, answer.ID)

	err = row.Scan(&res.RegionType.ID, &res.RegionType.UserName, &res.Description)
	if err != nil {
		return
	}

	// Get code name of region type
	if res.RegionType.ID < 0 || res.RegionType.ID >= int64(len(calc.Regions)) {
		err = fmt.Errorf("Неверный тип '%d' участка '%d'", res.RegionType.ID, answer.ID)
		return
	}
	res.RegionType.CodeName = calc.Regions[res.RegionType.ID].Name

	// Get name and description for all the parameters
	var mapParamTypes map[int64]APIParamType = make(map[int64]APIParamType)
	func() {
		var rows *sql.Rows
		rows, err = db.DB.Query(`SELECT id, name, description FROM tparam`)
		if err != nil {
			return
		}
		defer rows.Close()
		for rows.Next() {
			var tp APIParamType
			err = rows.Scan(&tp.ID, &tp.UserName, &tp.Description)
			if err != nil {
				return
			}
			mapParamTypes[tp.ID] = tp
		}
		err = rows.Err()
		return
	}()
	if err != nil {
		return
	}

	// Get all the params and parts of the region
	var resParams map[int64]db.DBParamValue
	var resParts map[int64]db.DBPartNomenclatureValue
	resParams, resParts, err = db.GetParamPartValues(answer.ID, map[int64]float64{}, map[int64]int64{})
	if err != nil {
		return
	}
	if len(resParams) > 0 {
		res.Params = make(map[int64]APIParam)
	}
	for tparamID, paramValue := range resParams {
		var param APIParam
		tp := mapParamTypes[tparamID]
		param.ParamType = &tp
		param.Value = paramValue.Value
		for pvalue, pname := range paramValue.ValueList {
			param.ValueList = append(param.ValueList, APIParamValue{Value: pvalue, Name: pname})
		}
		param.SetGUIFields()
		res.Params[tparamID] = param
	}

	var mapNomenclature map[int64]string // Nomenclature names loaded from DB
	var listNomenclature bytes.Buffer    // list of ID to filter SQL nomenclature

	if len(resParts) > 0 {
		res.Parts = make(map[int64]APIPart)
		mapNomenclature = make(map[int64]string)
	}
	for tpartID, partValue := range resParts {
		var part APIPart
		part.ID = tpartID

		if partValue.ID != nil {
			part.Value = new(APINomenclature)
			part.Value.ID = *partValue.ID
		}

		for _, nomenclatureID := range partValue.IDList {
			if nomenclatureID == nil {
				part.ValueList = append(part.ValueList, nil)
			} else {
				part.ValueList = append(part.ValueList, &APINomenclature{ID: *nomenclatureID})
				listNomenclature.WriteString(fmt.Sprintf("%n,", nomenclatureID))
			}
		}
		// Add "0" to close last comma
		listNomenclature.WriteString("0")

		// Get nomenclature info from DB
		func() {
			var rows *sql.Rows
			rows, err = db.DB.Query(fmt.Sprintf(`SELECT id, name FROM nomenclature WHERE id IN (%s)`, listNomenclature.String()))
			if err != nil {
				return
			}
			defer rows.Close()
			for rows.Next() {
				var ID int64
				var name string
				err = rows.Scan(&ID, &name)
				if err != nil {
					return
				}
				mapNomenclature[ID] = name
			}
			err = rows.Err()
			return
		}()
		if err != nil {
			return
		}

		// Set up nomenclature info
		if part.Value != nil {
			part.Value.Name = mapNomenclature[part.Value.ID]
		}

		for i := range part.ValueList {
			if part.ValueList[i] != nil {
				part.ValueList[i].Name = mapNomenclature[part.ValueList[i].ID]
			}
		}

		res.Parts[tpartID] = part
	}

	// Select components and parts from DB
	if len(res.Parts) > 0 {
		res.Components = make(map[int64]APIComponent)
	}
	func() {
		var rows *sql.Rows
		rows, err = db.DB.Query(`SELECT c.id, t.id, t.name, tp.id, tp.name, tp.tcalculation_id
			FROM component c INNER JOIN tcomponent t ON c.tcomponent_id = t.id
			LEFT JOIN part p ON p.component_id = c.id
			LEFT JOIN tpart tp ON p.tpart_id = tp.id
			WHERE c.region_id = ?
			ORDER BY c.id, p.id`, answer.ID)
		if err != nil {
			return
		}
		defer rows.Close()
		for rows.Next() {
			var ct APIComponentType
			var partID *int64
			var partName *string
			var partCalculationTypeID *int64
			var c APIComponent = APIComponent{ComponentType: &ct}
			err = rows.Scan(&c.ID, &ct.ID, &ct.Name, &partID, &partName, &partCalculationTypeID)
			if err != nil {
				return
			}

			if _, ok := res.Components[c.ID]; !ok {
				res.Components[c.ID] = c
			}

			if partID != nil {
				p := APIPartType{ID: *partID, Name: *partName, CalculationTypeID: *partCalculationTypeID}
				if p0, ok := res.Parts[p.ID]; ok {
					p0.Name = p.Name
					p0.CalculationTypeID = p.CalculationTypeID
					res.Parts[p.ID] = p0
				}
				c0 := res.Components[c.ID]
				c0.PartTypes = append(c0.PartTypes, p)
				res.Components[c.ID] = c0
			}
		}
		err = rows.Err()
		return
	}()
	if err != nil {
		return
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: PUT /projects/<id>/regions?region_type=<Value>[?description=<value>][?nr=<value>]
//
func PutRegion(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	var projectID int64
	projectID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID '%s'", request[1])
		return
	}

	// Parse user request parameters
	var rp RequestParams = RequestParams{
		"region_type": {Optional: false, Type: Int},
		"description": {Optional: true, Type: String},
		"nr":          {Optional: true, Type: Int},
	}

	err = rp.Parse(params)
	if err != nil {
		answer.Code = BadRequest
		return
	}

	// Add SQL-parameters: tregion_id, project_id
	rp["tregion_id"] = rp["region_type"]
	rp["project_id"] = RequestParam{Type: Int, Value: RequestParamValue{Type: Int, IntValue: projectID}}
	regionTypeID := calc.RegionTypeID(rp["tregion_id"].Value.IntValue)

	// Check for region type RTProject - there should be only one region with such a type
	if regionTypeID == calc.RTProject {
		var row *sql.Row
		var count int
		row = db.DB.QueryRow(`SELECT count(*) FROM region WHERE project_id=? AND tregion_id=?`, projectID, regionTypeID)
		err = row.Scan(&count)
		if err != nil {
			return
		}
		if count > 0 {
			err = fmt.Errorf("В проекте [%d] уже добавлен основной участок. Допускается только один участок такого типа", projectID)
			return
		}
	}

	// Insert into [region]
	sqlText, sqlParams := rp.MakeSQLInsert("region", []string{"tregion_id", "project_id", "description", "nr"})
	var res sql.Result
	res, err = db.DB.Exec(sqlText, sqlParams...)
	if err != nil {
		return
	}

	answer.ID, err = res.LastInsertId()
	if err != nil {
		return
	}

	// Add params with default values for all the params of this region type
	res, err = db.DB.Exec(`INSERT INTO param(region_id, tparam_id, value)
		SELECT ?, p.tparam_id, ifnull(min(v.value), 0)
		FROM cn_tparam_tregion p
		LEFT JOIN tparamvalue v ON p.tparam_id = v.tparam_id
		WHERE tregion_id=?
		GROUP BY p.tparam_id`, answer.ID, regionTypeID)
	if err != nil {
		return
	}

	// Add all the possible components for this type of region
	res, err = db.DB.Exec(`INSERT INTO component(region_id, tcomponent_id)
		SELECT ?, cn.tcomponent_id FROM cn_tregion_tcomponent cn WHERE cn.tregion_id = ?`, answer.ID, regionTypeID)
	if err != nil {
		return
	}

	// Add all the parts of components of the region, set default value of nomenclature for each part (AS SELECT min from possible values)
	res, err = db.DB.Exec(`INSERT INTO part(tpart_id, component_id, nomenclature_id)  
		SELECT p.id, c.id, min(cn.nomenclature_id)
		FROM component c
		INNER JOIN tpart p ON c.tcomponent_id = p.tcomponent_id
		LEFT JOIN cn_tpart_nomenclature cn ON p.id = cn.tpart_id
		WHERE c.region_id = ?
		GROUP BY p.id, c.id`, answer.ID)
	if err != nil {
		return
	}

	// Correct dependent param and part values
	var resParams map[int64]db.DBParamValue
	var resParts map[int64]db.DBPartNomenclatureValue
	resParams, resParts, err = db.GetParamPartValues(answer.ID, map[int64]float64{}, map[int64]int64{})
	if err != nil {
		return
	}
	err = db.WriteParamPartValues(answer.ID, resParams, resParts)
	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: POST /projects/<id>/regions/<id>[?description=<value>][?nr=<value>][?param=<param_id>(<value>)&param=<param_id>(<value>)&...]
//
func PostRegion(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	_, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID проекта '%s'", request[1])
		return
	}

	answer.ID, err = strconv.ParseInt(request[3], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID участка '%s'", request[1])
		return
	}

	// Parse user request parameters
	var rp RequestParams = RequestParams{
		"description": {Optional: true, Type: String},
		"nr":          {Optional: true, Type: Int},
		"param":       {Optional: true, Type: IntFloatMap},
	}

	err = rp.Parse(params)
	if err != nil {
		answer.Code = BadRequest
		return
	}

	// Update [region]
	sqlText, sqlParams := rp.MakeSQLUpdate("region", []string{"description", "nr"}, answer.ID)
	if len(sqlParams) > 0 {
		_, err = db.DB.Exec(sqlText, sqlParams...)
		if err != nil {
			return
		}
	}

	// Update region params and dependent parts
	pp := rp["param"]
	if !pp.Exists() {
		return
	}

	// Get user set parameters as map with key [paramID] and value [paramValue]
	regionParams := pp.Value.IntFloatMap

	// Correct dependent param and part values
	var resParams map[int64]db.DBParamValue
	var resParts map[int64]db.DBPartNomenclatureValue
	resParams, resParts, err = db.GetParamPartValues(answer.ID, regionParams, map[int64]int64{})
	if err != nil {
		return
	}
	err = db.WriteParamPartValues(answer.ID, resParams, resParts)

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: DELETE /projects/<id>/regions/<id>
//
func DeleteRegion(request []string, params map[string][]string) (answer Answer) {
	var err error
	defer answer.make(&err, nil)

	var projectID int64
	projectID, err = strconv.ParseInt(request[1], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID проекта '%s'", request[1])
		return
	}

	answer.ID, err = strconv.ParseInt(request[3], 10, 64)
	if err != nil {
		answer.Code = BadRequest
		err = fmt.Errorf("Неверный ID участка '%s'", request[1])
		return
	}

	// Delete from [region]
	_, err = db.DB.Exec("DELETE FROM region WHERE id=? AND project_id=?", answer.ID, projectID)
	return
}

///////////////////////////////////////////////////////////////////////////////
// SetGUIFields - set up fields for GUI depending on param type, value list and current value
func (param *APIParam) SetGUIFields() {
	if param.ParamType == nil {
		return
	}

	if param.ParamType.ID < 0 || param.ParamType.ID >= int64(len(calc.Params)) {
		return
	}

	calcParam := calc.Params[param.ParamType.ID]

	// If possible value list was filtered to 0 values, disable param ort make it invisible
	if len(param.ValueList) == 0 && len(calcParam.Values) > 1 {
		param.Disable()
	} else {
		param.Enable()
	}

	// Set up control
	if len(calcParam.Values) <= 1 {
		param.Control = CTNumericEditBox // or CTEditBox
	} else if calcParam.Values[0].Name == calc.BoolParamName {
		param.Control = CTCheckBox
		if len(param.ValueList) == 1 {
			param.Disable()
		}
	} else if calcParam.Values[0].Name == calc.ColorParamName {
		param.Control = CTChooseColorBtn
	} else {
		param.Control = CTComboBox
	}
}

///////////////////////////////////////////////////////////////////////////////
// Disable - disable param or make it invisible
func (param *APIParam) Disable() {
	param.Enabled = false
	param.Visible = true
}

///////////////////////////////////////////////////////////////////////////////
// Enable - enable param and make it visible
func (param *APIParam) Enable() {
	param.Enabled = true
	param.Visible = true
}
