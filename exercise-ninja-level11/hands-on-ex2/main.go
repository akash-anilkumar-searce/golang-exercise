package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSONprog(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}


func toJSONprog(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("error is present in toJSONprog function: %v", err)
	}
	return bs, nil
}