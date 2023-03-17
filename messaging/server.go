package server

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var MessageMap *map[string]map[string]string

func NewJsonMessageConfig(path string, fileName string) *error {

	wd, err := os.Getwd()
	if err != nil {
		return &err
	}

	file, err := os.Open(filepath.Join(wd, path, fileName))
	if err != nil {
		return &err
	}
	defer file.Close()

	byteString, err := ioutil.ReadAll(file)
	if err != nil {
		return &err
	}

	if err := json.Unmarshal(byteString, &MessageMap); err != nil {
		return &err
	}

	return nil
}
