package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "name"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	// Transforma um struct em json
	cachorroJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	// Converter em json
	fmt.Println(bytes.NewBuffer(cachorroJson))

	// Transformando o map em json
	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2Json, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cachorro2Json))

	// Transformar um json em struct
	cachorro3Json := `{"Nome":"Fifi","Raca":"Pastor Alemão","Idade":5}`

	var dog cachorro
	if erro := json.Unmarshal([]byte(cachorro3Json), &dog); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(dog)

	cachorro4Json := `{"Nome":"Bella","Raca":"Poodle","Idade":1}`

	dog2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro4Json), &dog2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(dog2)
}
