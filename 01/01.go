package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
)

type Product struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    string `json:"price"`
}

func main() {
	_, b, _, _ := runtime.Caller(0)
	a := func() string {
		if runtime.GOOS != "windows" {
			return "data.csv"
		} else {
			return path.Dir(b) + "/data.csv"
		}
	}()
	csvFile, err := os.Open(a)

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Product
	var allRecords []Product

	for _, each := range csvData {
		oneRecord.Name = each[0]
		oneRecord.Category = each[1]
		oneRecord.Price = each[2]
		allRecords = append(allRecords, oneRecord)
	}

	jsondata, err := json.Marshal(allRecords)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsondata))
}
