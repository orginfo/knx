package db

import (
	"database/sql"
	"fmt"
)

type DBParamValue struct {
	Value     float64            // If Value = -1 and Value list is empty, parameter is not available
	ValueList map[float64]string // If value list is empty and Value is not -1 the parameter could have any value
}

type DBPartNomenclatureValue struct {
	ID     *int64
	IDList []*int64
}

///////////////////////////////////////////////////////////////////////////////
// GetParamPartValues - return params dependent on the given param values, return nomenclature for parts, dependent on the given param values
// key of nomenclature value is an ID of part type
// Input map is local param values entered by user to replace values from DB
//
func GetParamPartValues(regionID int64, localParams map[int64]float64, localParts map[int64]int64) (resParams map[int64]DBParamValue, resParts map[int64]DBPartNomenclatureValue, err error) {
	resParams = make(map[int64]DBParamValue)
	resParts = make(map[int64]DBPartNomenclatureValue)

	// Collect all the parameter DB values + local values
	// Order by priority
	type DBParam struct {
		Prio        int
		ParamTypeID int64
		Value       float64
	}
	var regionParams []DBParam

	// Get all the parameters of this region from DB
	var rows *sql.Rows
	rows, err = DB.Query(`SELECT t.prio, p.tparam_id, p.value FROM param p
		INNER JOIN tparam t ON t.id = p.tparam_id
		WHERE region_id=? ORDER BY t.prio`, regionID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p DBParam
		err = rows.Scan(&p.Prio, &p.ParamTypeID, &p.Value)
		if err != nil {
			return
		}

		// Replace DB param value by local value
		localValue, ok := localParams[p.ParamTypeID]
		if ok {
			p.Value = localValue
		}

		regionParams = append(regionParams, p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	rows.Close()

	// Get possible values for all the parameters Prio 1 and higher
	for i, param := range regionParams {
		// Select all the possible values of this param type
		// Join possible values with set of dependencies, if there are any
		sqlFormat := `SELECT v.value, v.name FROM tparamvalue v %s WHERE v.tparam_id=?`
		sqlJoin := ""

		// Join results with dependencies
		for mI, mParam := range regionParams {

			// Main parameters only with priority less than current one
			if mI >= i {
				break
			}

			if mParam.Prio >= param.Prio {
				break
			}

			// Get count of dependencies
			var row *sql.Row
			row = DB.QueryRow(`SELECT count(*) FROM cn_tparamvalue_tparamvalue
				WHERE tparam_id=? AND value=? AND dependent_tparam_id=?`, mParam.ParamTypeID, mParam.Value, param.ParamTypeID)
			var countOfDepValues int
			err = row.Scan(&countOfDepValues)
			if err != nil {
				return
			}

			if countOfDepValues == 0 {
				continue
			}

			// Join with dependent values
			sqlJoin = fmt.Sprintf(`%[1]s INNER JOIN cn_tparamvalue_tparamvalue d%[2]d
				ON d%[2]d.dependent_value=v.value AND d%[2]d.dependent_tparam_id=v.tparam_id
				AND d%[2]d.tparam_id=%[3]d AND d%[2]d.value=%[4]f`, sqlJoin, mI, mParam.ParamTypeID, mParam.Value)
		}

		sqlSelect := fmt.Sprintf(sqlFormat, sqlJoin)

		rows, err = DB.Query(sqlSelect, param.ParamTypeID)
		if err != nil {
			return
		}
		var v DBParamValue
		v.ValueList = make(map[float64]string)
		for rows.Next() {
			var valueFloat float64
			var valueName string
			err = rows.Scan(&valueFloat, &valueName)
			if err != nil {
				return
			}
			v.ValueList[valueFloat] = valueName
		}
		err = rows.Err()
		if err != nil {
			return
		}
		rows.Close()

		// Find current param value in the list of possible values
		_, ok := v.ValueList[param.Value]
		if ok {
			v.Value = param.Value
		} else {
			// Get the first value from the list
			var k float64
			var found bool
			for k = range v.ValueList {
				found = true
				break
			}
			if found {
				v.Value = k
			} else { // If no values, set -1 in case of the dependences, don't change the value, if there were no dependencies
				if sqlJoin == "" {
					v.Value = param.Value
				} else {
					v.Value = -1
				}
			}
			param.Value = v.Value
		}

		resParams[param.ParamTypeID] = v
	}

	// Get all the part with set nomenclature from DB
	// Replace given part with local nomenclature
	rows, err = DB.Query(`SELECT p.tpart_id, p.nomenclature_id
		FROM part p INNER JOIN component c ON c.id = p.component_id
		WHERE c.region_id=?`, regionID)
	if err != nil {
		return
	}
	for rows.Next() {
		var partValue DBPartNomenclatureValue
		var tpartID int64
		err = rows.Scan(&tpartID, &partValue.ID)
		if err != nil {
			return
		}

		// Replace DB param value by local value
		localValue, ok := localParts[tpartID]
		if ok {
			if partValue.ID == nil {
				partValue.ID = new(int64)
			}
			*partValue.ID = localValue
		}

		resParts[tpartID] = partValue
	}
	err = rows.Err()
	if err != nil {
		return
	}
	rows.Close()

	// Select list of available nomenclature for each part
	for tpartID, partValue := range resParts {
		sqlFormat := `SELECT cn.nomenclature_id FROM cn_tpart_nomenclature %s cn WHERE cn.tpart_id=? ORDER BY cn.nomenclature_id`
		sqlJoin := ""

		// Join results with dependencies
		for mI, mParam := range regionParams {
			// Get count of dependencies
			var row *sql.Row
			row = DB.QueryRow(`SELECT count(*) FROM cn_tparamvalue_nomenclature n
				INNER JOIN cn_tparam_tpart p ON p.tparam_id=n.tparam_id
				WHERE n.tparam_id=? AND n.value=? AND p.tpart_id=?`, mParam.ParamTypeID, mParam.Value, tpartID)
			var countOfDepValues int
			err = row.Scan(&countOfDepValues)
			if err != nil {
				return
			}

			if countOfDepValues == 0 {
				continue
			}

			// Join with dependent values
			sqlJoin = fmt.Sprintf(`%[1]s INNER JOIN cn_tparamvalue_nomenclature d%[2]d
				ON d%[2]d.nomenclature_id=cn.nomenclature_id
				AND d%[2]d.tparam_id=%[3]d AND d%[2]d.value=%[4]f`, sqlJoin, mI, mParam.ParamTypeID, mParam.Value)
		}

		sqlSelect := fmt.Sprintf(sqlFormat, sqlJoin)
		rows, err = DB.Query(sqlSelect, tpartID)
		if err != nil {
			return
		}
		for rows.Next() {
			var nomenclatureID *int64
			err = rows.Scan(&nomenclatureID)
			if err != nil {
				return
			}
			partValue.IDList = append(partValue.IDList, nomenclatureID)
		}
		err = rows.Err()
		if err != nil {
			return
		}
		rows.Close()

		// Change value of part nomenclature if the current value not in the list
		var found bool
		for _, v := range partValue.IDList {
			if v == nil && partValue.ID == nil {
				found = true
				break
			}
			if v != nil && partValue.ID != nil && *v == *partValue.ID {
				found = true
				break
			}
		}

		if !found {
			if len(partValue.IDList) > 0 {
				if partValue.IDList[0] == nil {
					partValue.ID = nil
				} else {
					if partValue.ID == nil {
						partValue.ID = new(int64)
					}
					*partValue.ID = *partValue.IDList[0]
				}
			} else {
				partValue.ID = nil
			}
		}

		resParts[tpartID] = partValue
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
// WriteParamPartValues - write into db values of parameters and parts for the given region
//
func WriteParamPartValues(regionID int64, params map[int64]DBParamValue, parts map[int64]DBPartNomenclatureValue) (err error) {
	// Begin transaction
	var tx *sql.Tx
	tx, err = DB.Begin()
	if err != nil {
		return
	}

	// Commit or rollback transaction at the end
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Param values
	for tparamID, paramValue := range params {
		_, err = tx.Exec("UPDATE param SET value=? WHERE region_id=? AND tparam_id=?", paramValue.Value, regionID, tparamID)
		if err != nil {
			return
		}
	}
	// Part values
	for tpartID, partValue := range parts {
		_, err = tx.Exec("UPDATE part SET nomenclature_id=? WHERE tpart_id=? AND component_id IN (SELECT id FROM component WHERE region_id=?)", partValue.ID, tpartID, regionID)
		if err != nil {
			return
		}
	}

	return
}
