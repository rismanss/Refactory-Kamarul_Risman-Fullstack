package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const text_1 = "Mammals"
const text_2 = "Bruiser build"

func charCounter(str string) {
	setData := make(map[string]interface{})
	for _, c := range strings.Split(str, "") {
		if c != " " {
			setData[strings.ToLower(c)] = strings.Repeat("*", strings.Count(strings.ToLower(str), strings.ToLower(c)))
		}
	}
	dataJson, _ := json.Marshal(setData)
	fmt.Println(string(dataJson))
}

func main() {
	charCounter(text_1)
	charCounter(text_2)
}
