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
