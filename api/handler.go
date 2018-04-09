package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func InitAPIMux() *http.ServeMux {
	APIMux := http.NewServeMux()
	APIMux.HandleFunc("/", handler)
	return APIMux
}

func handler(w http.ResponseWriter, r *http.Request) {
	var answer Answer

	// Show answer struct as a result of any request to API at the end, whatever the request or result is
	defer func() {
		// CORS for angular debugging
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "content-type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Max-Age", "5")

		// TODO: Для реальной работы использовать компактный вывод: json.Marshal
		//data, err := json.Marshal(answer)
		data, err := json.MarshalIndent(answer, " ", "   ")
		if err != nil {
			log.Fatalf("Сбой маршалинга JSON: %v\n", err)
		}
		_, err = w.Write(data)
		if err != nil {
			log.Fatalf("Ошибка при выводе JSON: %v\n", err)
		}
	}()

	// Получаем путь из URL
	path := r.URL.EscapedPath()

	// URL имеет вид: /v0/region_types/...
	// Поэтому request должен содержать минимум 2 элемента:
	// - request[0] содержит пустую строку;
	// - request[1] содержит версию программы.
	request := strings.Split(path, "/")

	if len(request) < 2 {
		answer.Code = BadRequest
		answer.Message = "URL должен содержать версию API.\n"
		return
	}

	// TODO: обрабатывать версию API. Временное решение.
	if version := request[1]; version != "v0" && version != "V0" {
		answer.Code = BadRequest
		answer.Message = fmt.Sprintf("Версия API '%s' не поддерживается.\n", request[1])
		return
	}

	// Игнорируем версию API.
	request = request[2:]

	if len(request) == 0 {
		answer.Code = BadRequest
		answer.Message = fmt.Sprintf("Запрос '%s' не предусмотрен. Обратитесь к документации по API KNX v0.", r.URL.EscapedPath())
		return
	}

	// Формируем ключ для поиска функции в мапе. Каждое второе слово в request - это ID. Чтобы сформировать шаблон, заменяем четные слова на <id>.
	var key bytes.Buffer
	// Запоминаем идентификаторы запроса.
	var IDs []string
	for i, v := range request {
		if len(v) == 0 {
			continue
		}
		if i%2 == 0 {
			key.WriteString(v)
		} else {
			key.WriteString("<id>")
			IDs = append(IDs, v)
		}
	}

	// Получаем параметры запроса
	params := r.URL.Query()

	// Получаем метод. Приводим к нижнему регистру.
	method := strings.ToLower(r.Method)

	//TODO: временное решение для тестирования API. В дальнейшем method будет определяться только через r.Method
	if mparam, ok := params["method"]; ok {
		method = mparam[0]
	}

	// По первому слову request определяем какая функция должна выполняться
	f, ok := F[key.String()]
	if !ok {
		answer.Code = BadRequest
		answer.Message = fmt.Sprintf("Запрос '%s' не предусмотрен. Обратитесь к документации по API KNX v0.", r.URL.EscapedPath())
		return
	}

	switch method {
	case "get":
		answer = f.Get(request, params)
	case "put":
		answer = f.Put(request, params)
	case "post":
		answer = f.Post(request, params)
	case "delete":
		answer = f.Del(request, params)
	default:
		answer.Code = BadRequest
		answer.Message = fmt.Sprintf("Неизвестная команда: '%s %s'.\n", method, path)
	}
}
