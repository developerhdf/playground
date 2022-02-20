package model

import (
	"encoding/json"
)

type Person struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	Gender    string `json:"gender"`
	Homeworld string `json:"homeworld"`
}

func (p Person) String() string {
	personString := "{}"
	if jsonPerson, err := json.Marshal(p); err == nil {
		personString = string(jsonPerson)
	}
	return personString
}
