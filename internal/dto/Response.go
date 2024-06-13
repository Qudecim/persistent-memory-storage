package dto

type Response struct {
	Id              string   `json:"i"`
	Value           string   `json:"v"`
	Items           []string `json:"t"`
	Increment_value int64    `json:"c"`
	Error           int      `json:"e"`
}

func NewResponse(id string, value string, error int) Response {
	return Response{Id: id, Value: value, Error: error, Items: []string{}}
}

func NewResponseList(id string, items []string, error int) Response {
	return Response{Id: id, Value: "", Items: items, Error: error}
}

func NewResponseIncrement(id string, value int64, error int) Response {
	return Response{Id: id, Increment_value: value, Value: "", Error: error, Items: []string{}}
}
