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
	Items []ResponseItem `json:"i"`
}

type ResponseItem struct {
	Id    string `json:"i"`
	Value string `json:"v"`
	Error int    `json:"e"`
}

func handle(message []byte) []byte {

	jsonData := []byte(message)

	var request Request
	err := json.Unmarshal(jsonData, &request)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	var response Response

	switch request.Method {
	case "g": // get
		data, _ := db.Get(request.Key)
		responseItem := ResponseItem{Id: request.Id, Value: data}
		r := []ResponseItem{responseItem}
		response = Response{Items: r}
	case "s": // set
		db.Set(request.Key, request.Value)
		responseItem := ResponseItem{Id: request.Id}
		r := []ResponseItem{responseItem}
		response = Response{Items: r}
	default:
		responseItem := ResponseItem{Id: request.Id, Error: 1}
		r := []ResponseItem{responseItem}
		response = Response{Items: r}

	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	return responseJson
}
