package util

import (
	"encoding/json"
	"fmt"
	"log"
)

func Map2json() []byte {
	data := map[string]any{
		"name":    "dudu",
		"age":     18,
		"good at": []string{"Chinese", "Math"},
		"belongings": map[string]int{
			"house": 1,
			"car":   2,
		},
		"is pussy": false,
	}
	//fmt.Println(data["name"], data["age"])
	result, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	//fmt.Printf("%T", result)
	//fmt.Println(result)
	//fmt.Println(string(result))
	return result
}

func DecodeJsonWithMap() {
	data := Map2json()
	m := make(map[string]any)
	err := json.Unmarshal(data, &m)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("key:%v,\t\tv:%T\n", k, v)
	}
}

func DecodeJsonWithStruct() {
	data := Map2json()
	type Person struct {
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Belongings struct {
			Car   int `json:"car"`
			House int `json:"house"`
		} `json:"belongings"`
		GoodAt  []string `json:"good at"`
		IsPussy bool     `json:"is pussy"`
	}
	type Person2 struct {
		Belongings struct {
			House int `json:"house"`
		} `json:"belongings"`
	}
	//var p1 Person
	var p2 Person2
	err := json.Unmarshal(data, &p2)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(p2)
	fmt.Println(p2.Belongings.House)
}

func DecodeUnknownStructureJsonData() {
	data := Map2json()
	var result interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(result.(map[string]interface{})["belongings"])
}
