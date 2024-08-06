package handler

import (
	"consume_api/domain"
	"consume_api/utils"
	"encoding/json"
	"net/http"
)

func ConsumeRest(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		utils.NewResponse(http.StatusText(resp.StatusCode), resp.StatusCode, nil).Write(w)
		return
	}

	defer resp.Body.Close()

	var todo []domain.Todo

	json.NewDecoder(resp.Body).Decode(&todo)

	utils.PrintToConsole(todo)

	utils.NewResponse(http.StatusText(resp.StatusCode), resp.StatusCode, todo).Write(w)
}