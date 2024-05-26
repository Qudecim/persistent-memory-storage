package transport

import (
	"encoding/json"
	"log"
	"qudecim/db/internal/app"
	"qudecim/db/internal/dto"
)

func handle(app *app.App, message []byte) ([]byte, bool) {

	jsonData := []byte(message)

	var request dto.Request
	err := json.Unmarshal(jsonData, &request)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	var response dto.Response
	var sendAnswer bool

	switch request.Method {
	case "g": // get
		data, _ := app.Get(&request)
		response = dto.Response{Id: request.Id, Value: data}
		sendAnswer = true
	case "s": // set
		app.Set(&request)
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	return responseJson, sendAnswer
}
