package rest

import (
	"encoding/json"
	"net/http"
)

type AnswerIncorrect struct {
	Error string `json:"error"`
}

type AnswerСorrect struct {
	Result interface{} `json:"result"`
}

func JsonResponse(w http.ResponseWriter, status int, structure interface{}) {
	response, err := json.Marshal(structure)
	if err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Не удалось выполнить Marshal: " + err.Error(),
		}

		JsonResponse(w, 500, answerIncorrect)

		return
	}

	w.WriteHeader(status)
	w.Write(response)
}
