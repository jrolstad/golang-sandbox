package main

import (
	"encoding/json"
	"fmt"
)

type people struct{
	Number int `json:"number"`
}

func main() {
	text:=GetPeopleData()
	result:=ParseJson(text)
	
	fmt.Println(result.Number)
}

func GetPeopleData() string{
	return `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`
}

func ParseJson(data string) people{
	textBytes:= []byte(data)

	people1:=people{}
	err:= json.Unmarshal(textBytes,&people1)

	if(err!=nil){
		fmt.Println(err)
	}

	return people1
}