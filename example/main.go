package main

import (
	"encoding/json"
	"fmt"
	"github.com/ekofedriyanto/typeform-webhook-parser/parser"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	pathFile := path.Join(pwd, "/example/typeform.json")
	jsonFile, err := os.Open(pathFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		_ = jsonFile.Close()
	}()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	var typeformData parser.TypeFormData
	err = json.Unmarshal(byteValue, &typeformData)

	if err != nil {
		fmt.Println(err)
		return
	}

	typeformParser := parser.NewParser(&typeformData)
	mapData, err := typeformParser.ToMap()

	if err != nil {
		fmt.Println(err)
		return
	}

	marshal, _ := json.MarshalIndent(mapData, "", "    ")
	if err != nil {
		return
	}

	fmt.Println(string(marshal))

}
