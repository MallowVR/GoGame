package main

import (
	"encoding/json"
	"os"
)

func ReadJsonFile(_in any, _fileName string) bool {
	file, err := os.OpenFile(_fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	fileContent, err := os.ReadFile(_fileName)
	if err != nil {
		panic(err)
	}

	file.Close()

	json.Unmarshal(fileContent, _in)
	return true
}

func WriteJsonFile(_in any, _fileName string) {
	file, err := os.OpenFile(_fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	json, err := json.MarshalIndent(_in, "", "  ")
	if err != nil {
		panic(err)
	}

	if file != nil {
		file.Write(json)
	}

	file.Sync()

	file.Close()
}
