package main

import (
	"encoding/json"
	"fmt"
)

// Name inside struct is aliase
// json is the json format
type userInfo struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	Address  []string `json:"address,omitempty"`
}

func main() {
	fmt.Println("Welcome to json video")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	info := []userInfo{
		{"Nan OO", "noh@dev", "abc123", []string{"Sanchaung", "Hlaing"}},
		{"Khin", "khin@dev", "abc1234", []string{"8 Mile"}},
		{"Ingyin", "ingyin@dev", "abce123", nil},
	}

	//package this data as JSON data
	finalData, err := json.MarshalIndent(info, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalData)
}

func DecodeJson() {
	jsonDataByte := []byte(`
	{
		"name": "Nan OO",
		"email": "noh@dev",
		"address": [
				"Sanchaung",
				"Hlaing"
		]
    }
	`)
	fmt.Println("Json Byte data")
	fmt.Println(jsonDataByte)

	// locInfo is struct
	var locInfo userInfo
	checkValid := json.Valid(jsonDataByte)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataByte, &locInfo)
		fmt.Printf("%#v\n", locInfo)
		fmt.Printf("%#v\n", locInfo.Address)
	} else {
		fmt.Println("Json was not valid")
	}

	// keyValueData is map (like associative array)
	var keyValueData map[string]interface{}
	json.Unmarshal(jsonDataByte, &keyValueData)
	fmt.Printf("%#v\n", keyValueData)

	for k, v := range keyValueData {
		fmt.Printf("Key is %v and value is %v\n", k, v)
	}
}
