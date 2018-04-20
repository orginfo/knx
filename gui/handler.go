package gui

import (
	"encoding/json"
	"fmt"
	"knx/api"
	"net/http"
)

func InitGUIMux() *http.ServeMux {
	GUIMux := http.NewServeMux()
	GUIMux.HandleFunc("/", handler)
	return GUIMux
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Data to parse with html-template
	var typedResult []api.APIProject
	data := api.Answer{Result: &typedResult}

	// Return "projects" html created with JSON answer
	defer Templates["projects"].Execute(w, &data)

	// Ask API for the whole list of projects
	resp, err := http.Get("http://localhost:8080/v0/projects")
	if err != nil {
		data.Message = fmt.Sprintf("%v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		data.Message = fmt.Sprintf("Сбой запроса: %s", resp.Status)
		return
	}

	// Decode request answer into JSON
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		data.Message = fmt.Sprintf("JSON Decode error: %s", err)
		return
	}
	resp.Body.Close()
}
