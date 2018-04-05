package api

import (
	"fmt"
	"strconv"
	"strings"
)

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
