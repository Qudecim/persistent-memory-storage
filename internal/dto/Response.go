package dto

type Response struct {
	Id    string   `json:"i"`
	Value string   `json:"v"`
	Items []string `json:"t"`
	Error int      `json:"e"`
}

func NewResponse(id string, value string, error int) Response {
	return Response{Id: id, Value: value, Error: error, Items: []string{}}
}

func NewResponseList(id string, items []string, error int) Response {
	return Response{Id: id, Value: "", Items: items, Error: error}
}
