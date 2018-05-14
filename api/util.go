package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// putRequest выполняет "PUT" запрос.
func putRequest(request string, data io.Reader) (id int64, err error) {
	client := &http.Client{}

	// Создаем запрос
	req, error := http.NewRequest(http.MethodPut, request, data)
	if error != nil {
		err = error
		return
	}

	// Выполняем запрос
	resp, error := client.Do(req)
	if error != nil {
		err = fmt.Errorf("%s. Ошибка при выполнении запроса: %s.", error.Error(), request)
		return
	}

	defer resp.Body.Close()

	// Проверяем статус выполненного запроса
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Сбой запроса: %s. Стасус: %s.", request, resp.Status)
		return
	}

	//  Декодируем результат запроса, чтобы узнать, не случилась ли ошибка
	var answer Answer
	if err = json.NewDecoder(resp.Body).Decode(&answer); err != nil {
		return
	}

	id = answer.ID
	return
}

// ImportNomenclature импортирует данные номенклатуры из указанной директории.
// Обрабатываются только csv-файлы. Заполняются таблицы nomenclature и tnomenclature.

func ImportNomenclature(filePath string) (err error) {
	// Открываем файл
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return
	}

	// Отложенный вызов закрытия файла.
	defer file.Close()

	// Формат файла:
	// 1. Тип номенклатуры
	// 2. Поля
	// 3. Значения полей

	reader := csv.NewReader(file)
	reader.Comma = ';'

	var nomenclatureType string
	var fields []string

	var nomenclatureID int64
	lineCount := 0
	for {
		record, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			err = error
			if perr, ok := err.(*csv.ParseError); ok && perr.Err != csv.ErrFieldCount {
				err = perr
				return
			}
		}

		if lineCount == 0 {
			// В первой строке содержится тип номенклатуры
			// Проверяем количество параметров в первой строке файла
			if count := len(record); count > 1 {
				err = fmt.Errorf("Файл '%s' в первой строке должен содержать единственный параметр - тип номенклатуры. Введено количество параметров: %d. Введенные параметры: '%s'.", filePath, count, record)
				return
			}
			// Получаем тип номенклатуры
			nomenclatureType = record[0]

			// Формируем запрос для вставки типа номенклатуры
			request := fmt.Sprintf("http://localhost:8080/v0/nomenclature_types?name=%s", url.QueryEscape(nomenclatureType))

			nomenclatureID, err = putRequest(request, nil)
			fmt.Println(nomenclatureID)
			if err != nil {
				return
			}
		} else if lineCount == 1 {
			// Вторая строка файла содержит названия полей таблицы
			fields = record
			if len(record) == 0 {
				err = fmt.Errorf("Файл '%s' во второй строке должен содержать список полей. Список полей пуст.", filePath)
				return
			}
		} else {
			// Считываем данные для полей

			countData := len(record)
			countFields := len(fields)
			// Импортируемые файлы содержат поле price. Но таблица nomenclature не содержит это поле.
			// Поле price - последнее поле в файле. Игнорируем это поле для вставки в таблицу номенклатур.
			// Поэтому количество столбцов в файле не должно быть меньше 2. Иначе, считаем, что в файле нет данных.
			if /*countData == 0 || */ countData < 2 {
				err = fmt.Errorf("Файл '%s' не содержит данных.", filePath)
				return
			} else if countData != countFields {
				err = fmt.Errorf("Файл '%s': Количество параметров в строке (%d) не совпадает с количеством полей (%d). Ошибка в строке: %d. Строка: %s.", filePath, countData, countFields, lineCount+1, record)
				return
			}

			values := url.Values{}

			for i, field := range fields {
				if field == "price" {
					// TODO: Игнорируем поле price. Это временное решение, т.к. сейчас заполяняем только 2 таблицы: tnomenclature и nomenclature, которые не содержат этого поля.
					// TODO: Реализовать заполнение таблицы price, при котором будет использоваться поле price.
					continue
				}

				if len(values) == 0 {
					values.Set(field, record[i])
				} else {
					values.Add(field, record[i])
				}
			}

			request := fmt.Sprintf("http://localhost:8080/v0/nomenclature_types/%d/nomenclature?%s", nomenclatureID, values.Encode())
			fmt.Println(request)

			_, err = putRequest(request, nil)
			if err != nil {
				return
			}
		}

		lineCount++
	}

	return
}

///////////////////////////////////////////////////////////////////////////////
// Request: GET /import_nomenclature
// GetImportNomenclature импортирует данные номенклатуры из директории, указанной в параметре 'path' запроса
// В директории должны храниться csv-файлы для импорта.
func GetImportNomenclature(request []string, params map[string][]string) (answer Answer) {
	var err error
	var report string
	defer answer.make(&err, &report)

	if params == nil || len(params) == 0 {
		answer.Code = BadRequest
		err = fmt.Errorf("Не задан параметр 'path' для запроса '%s'", request[0])
		return
	}

	paramName := "path"
	value, ok := params[paramName]
	if !ok {
		answer.Code = BadRequest
		err = fmt.Errorf("Не задан параметр '%s' для запроса '%s'", paramName, request[0])
		return
	}

	path := value[0]
	if len(path) == 0 {
		err = fmt.Errorf("Не задан путь для параметра '%s' для запроса '%s'", paramName, request[0])
		return
	}

	var filesInfo os.FileInfo
	if filesInfo, err = os.Stat(path); os.IsNotExist(err) {
		err = fmt.Errorf("Не задан путь для параметра '%s' для запроса '%s'", paramName, request[0])
		return
	}

	if !filesInfo.IsDir() {
		err = fmt.Errorf("'%s' не является директорией. Укажите директорию для параметра '%s' запроса '%s'", filesInfo.Name(), paramName, request[0])
		return
	}

	var files []os.FileInfo
	files, err = ioutil.ReadDir(path)
	if err != nil {
		return
	}

	// Обрабатываются только .csv файлы.
	var ignoredFiles []string     // Файлы с другим расширением. Добавляем для статистики.
	var unprocessedFiles []string // Файлы, обработка которых закончилась с ошибкой
	var processedFiles []string   // Успешно обработанные файлы
	var errMsg []string
	for _, file := range files {
		if fileName := file.Name(); !strings.Contains(fileName, ".csv") {
			// Обрабатываем только .csv файлы
			ignoredFiles = append(ignoredFiles, fileName)
			continue
		}

		// os.FileInfo содержит только имя файла. Для обращению к файлу требуется полный путь + имя файла.
		// TODO: Заглушка. Ожидается, что в браузере путь к папке закачнивается '\'.
		// Т.е. этот вариант работает: работает http://api.localhost:8080/v0/import_nomenclature?path=D:\Dev\Projects\Knx\nomenclature\
		// А этот не работает (т.к. путь должен заканчиваться символом '\'): http://api.localhost:8080/v0/import_nomenclature?path=D:\Dev\Projects\Knx\nomenclature
		// TODO: в unix системах и windows используется прямой и обратный слэши соответсвенно. Предусмотреть в дальнейшем, чтобы работали оба варианта и на разных ОС.
		fileName := path + file.Name()
		err = ImportNomenclature(fileName)
		if err == nil {
			processedFiles = append(processedFiles, file.Name())
		} else {
			errMsg = append(errMsg, err.Error())
			unprocessedFiles = append(unprocessedFiles, file.Name())
		}
	}

	report = fmt.Sprintf("Успешно обработанные файлы: %v\n Файлы, обработка которых закончилась с ошибкой: %v\nИгнорируемые файлы: %v\nОшибки: %v", processedFiles, unprocessedFiles, ignoredFiles, errMsg)
	answer.Result = report

	return
}
