package transport

import (
	"encoding/json"
	"fmt"
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
		response = dto.NewResponse(request.Id, data, 0)
		sendAnswer = true
	case "s": // set
		app.Set(&request)
	case "u": // pull
		data, _ := app.Pull(&request)
		response = dto.NewResponseList(request.Id, data, 0)
		fmt.Println(response)
		sendAnswer = true
	case "p": // push
		app.Push(&request)
		response = dto.NewResponse(request.Id, "", 1)
		sendAnswer = true
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	return responseJson, sendAnswer
}
