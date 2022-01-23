package models

type Personalidade struct {
	Id       int    `json:id`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades = []Personalidade{
	{Id: 1, Nome: "Nagib Shaib", Historia: "Seila 1"},
	{Id: 2, Nome: "Chico de Paula", Historia: "Seila"},
}
