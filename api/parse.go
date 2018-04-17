package api

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type RequestParamType int

const (
	Nil RequestParamType = iota
	Int
	Float
	String
	IntArray
	StringArray
	IntStringMap
	FloatStringMap
)

type RequestParamValue struct {
	Type           RequestParamType
	IntValue       int
	FloatValue     float32
	StringValue    string
	IntArray       []int
	StringArray    []string
	IntStringMap   map[int]string
	FloatStringMap map[float32]string
}

type RequestParam struct {
	Optional bool
	Type     RequestParamType
	Value    RequestParamValue
}

type RequestParams map[string]RequestParam

// parseListOfInt конвертирует строку с числами в список чисел.
// Числа в строке представлены через ','. Строка не должна содержать пробелов.
func parseListOfInt(value string) (values []int, err error) {
	if len(value) == 0 {
		return
	}

	words := strings.Split(value, ",")
	for _, word := range words {
		v, err := strconv.Atoi(word)
		if err != nil {
			return values, err
		}

		values = append(values, v)
	}
	return
}

// parseValueName - разбивает строку <Value>(<Name>),<Value>(<Name>),... на отдельные
// значения и соответствующие им имена. Значения и имена хранятся в строковом формате.
// Имя может отсутствовать, тогда оно хранится в мапе в виде пустой строки.
func parseValueName(str string) (values map[string]string, err error) {
	if len(str) == 0 {
		return
	}

	values = make(map[string]string)

	params := strings.Split(str, ",")
	for _, param := range params {
		// ibegin, iend содержат начальные и конечные индексы сперва для <name>, а затем и для <value>
		ibegin := strings.Index(param, "(")
		iend := strings.Index(param, ")")
		if ibegin != -1 && iend == -1 || ibegin == -1 && iend != -1 || ibegin > iend {
			err = fmt.Errorf("Неверно задан параметр: %s.\n", param)
			return
		}

		// Копируем имя, если оно задано.
		var name string
		if ibegin != -1 {
			name = param[ibegin+1 : iend]
			iend = ibegin
		} else {
			iend = len(param)
		}

		value := param[0:iend]
		if len(value) == 0 {
			err = fmt.Errorf("Неверно задан параметр: %s. Значение не может быть пустым.\n", param)
			return
		}

		values[value] = name
	}
	return
}

// parseFloatName
func parseFloatName(str string) (result map[float32]string, err error) {
	var values map[string]string
	values, err = parseValueName(str)
	if err != nil {
		return
	}

	result = make(map[float32]string)

	for v, name := range values {
		var value float64
		value, err = strconv.ParseFloat(v, 32)
		if err != nil {
			return
		}
		// Преобразовываем ключ к типу float32
		result[float32(value)] = name
	}
	return
}

// parseIntName
func parseIntName(str string) (result map[int]string, err error) {
	var values map[string]string
	values, err = parseValueName(str)
	if err != nil {
		return
	}

	result = make(map[int]string)

	for v, name := range values {
		// Преобразовываем ключ к типу int
		var value int
		value, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		result[value] = name
	}
	return
}

// Value - return value of RequestParamValue dependent on its type
func (v *RequestParamValue) Value() interface{} {
	switch v.Type {
	case Int:
		return v.IntValue
	case Float:
		return v.FloatValue
	case String:
		return v.StringValue
	case IntArray:
		return v.IntArray
	case StringArray:
		return v.StringArray
	case IntStringMap:
		return v.IntStringMap
	case FloatStringMap:
		return v.FloatStringMap
	default:
		return nil
	}
}

// Parse - parse string value into the RequestParamValue
func (rp *RequestParamValue) Parse(s string) (err error) {
	switch rp.Type {
	case Int:
		rp.IntValue, err = strconv.Atoi(s)
	case Float:
		{
			var value float64
			value, err = strconv.ParseFloat(s, 32)
			rp.FloatValue = float32(value)
		}
	case String:
		rp.StringValue = s
	case IntArray:
		rp.IntArray, err = parseListOfInt(s)
	case StringArray:
		rp.StringArray = strings.Split(s, ",")
	case IntStringMap:
		rp.IntStringMap, err = parseIntName(s)
	case FloatStringMap:
		rp.FloatStringMap, err = parseFloatName(s)
	default:
		err = fmt.Errorf("Неизвестный тип параметра '%d'", rp.Type)
	}
	return
}

// Parse - convert standard HTTP-Rerquest params into convinient RequestParams and return an error
func (rps *RequestParams) Parse(params map[string][]string) error {
	for name, rp := range *rps {
		p, ok := params[name]
		if !ok && !rp.Optional {
			return fmt.Errorf("Не задан обязательный параметр запроса '%s'", name)
		}

		if !ok {
			rp.Value.Type = Nil
			(*rps)[name] = rp
			continue
		}

		rp.Value.Type = rp.Type
		err := rp.Value.Parse(p[0])
		(*rps)[name] = rp
		if err != nil {
			err = fmt.Errorf("Значение параметра '%s' не удается распознать (%v)", name, err)
			return err
		}
	}
	return nil
}

// MakeSQLInsert - construct SQL-INSERT request and return text of SQL and list of values as parmeters for the SQL
func (rps *RequestParams) MakeSQLInsert(tableName string, fields []string) (sqlText string, sqlParams []interface{}) {
	var fieldDesc bytes.Buffer
	var paramDesc bytes.Buffer

	for _, f := range fields {
		rp, ok := (*rps)[f]
		if !ok || !rp.Exists() {
			continue
		}

		fmt.Fprintf(&fieldDesc, "%s,", f)
		paramDesc.WriteString("?,")

		sqlParams = append(sqlParams, rp.GetValue())
	}

	// Remove last comma
	if fieldDesc.Len() == 0 {
		return
	}

	fieldDesc.Truncate(fieldDesc.Len() - 1)
	paramDesc.Truncate(paramDesc.Len() - 1)
	sqlText = fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", tableName, fieldDesc.String(), paramDesc.String())
	return
}

// MakeSQLUpdate - construct SQL-UPDATE request and return text of SQL and list of values as parmeters for the SQL
func (rps *RequestParams) MakeSQLUpdate(tableName string, fields []string, id int) (sqlText string, sqlParams []interface{}) {
	var fieldDesc bytes.Buffer

	for _, f := range fields {
		rp, ok := (*rps)[f]
		if !ok || !rp.Exists() {
			continue
		}

		fmt.Fprintf(&fieldDesc, "%s=?,", f)
		sqlParams = append(sqlParams, rp.GetValue())
	}

	if fieldDesc.Len() == 0 {
		return
	}

	// Remove last comma
	fieldDesc.Truncate(fieldDesc.Len() - 1)
	sqlParams = append(sqlParams, id)
	sqlText = fmt.Sprintf("UPDATE %s SET %s WHERE id=?", tableName, fieldDesc.String())
	return
}

// Exists - checks if the given parameter set up by user or not
func (p *RequestParam) Exists() bool {
	return p.Value.Type != Nil
}

// GetValue - return value of the parameter in native form
func (p *RequestParam) GetValue() interface{} {
	return p.Value.Value()
}
