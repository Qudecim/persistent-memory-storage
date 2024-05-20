package socket

import (
	"encoding/json"
	"log"
	"qudecim/db/db"
)

type Request struct {
	Id     string `json:"i"`
	Method string `json:"m"`
	Key    string `json:"k"`
	Value  string `json:"v"`
}

type Response struct {
	Id    string `json:"i"`
	Value string `json:"v"`
	Error int    `json:"e"`
}

func handle(message []byte) ([]byte, bool) {

	jsonData := []byte(message)

	var request Request
	err := json.Unmarshal(jsonData, &request)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	var response Response
	var sendAnswer bool

	switch request.Method {
	case "g": // get
		data, _ := db.Get(request.Key)
		response = Response{Id: request.Id, Value: data}
		sendAnswer = true
	case "s": // set
		db.Set(request.Key, request.Value)
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	return responseJson, sendAnswer
}
