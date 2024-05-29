package dto

type Request struct {
	Id     string `json:"i"`
	Method string `json:"m"`
	Key    string `json:"k"`
	Value  string `json:"v"`
}

func (r Request) GetMethod() string {
	return r.Method
}

func (r Request) GetKey() string {
	return r.Key
}

func (r Request) GetValue() string {
	return r.Value
}
