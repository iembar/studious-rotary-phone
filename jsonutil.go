package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadJSON(filename string) (interface{}, error) {
	var data interface{}
	raw, err := ioutil.ReadFile(filename)
	json.Unmarshal(raw, &data)
	return data, err
}

func CreateJSON(content interface{}, filename string) (string, error) {
	toWrite, _ := json.Marshal(content)
	err := ioutil.WriteFile(filename, toWrite, 0644)

	if err != nil {
		fmt.Println("Error in creating JSON", err)
	}
	return "File created :" + filename, err
}
