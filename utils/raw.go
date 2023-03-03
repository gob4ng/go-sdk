package utils

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
)

func GetRawData(reader io.Reader) string {
	jsonData, err := ioutil.ReadAll(reader)
	if err != nil {
		return err.Error()
	}

	return string(jsonData)
}

func JsonToString(jsonStruct interface{}) string {
	byteJson, err := json.Marshal(jsonStruct)
	if err != nil {
		return err.Error()
	}
	return string(byteJson)
}

func JsonStringToStruct(strJson string) (*interface{}, *error) {

	var structJson interface{}
	if err := json.Unmarshal([]byte(strJson), structJson); err != nil {
		return nil, &err
	}

	return &structJson, nil
}

func XmlStringToStruct(strXml string) (*interface{}, *error) {

	var structJson interface{}
	if err := xml.Unmarshal([]byte(strXml), structJson); err != nil {
		return nil, &err
	}

	return &structJson, nil
}

func XmlToString(jsonStruct interface{}) string {
	byteXml, err := xml.Marshal(jsonStruct)
	if err != nil {
		return err.Error()
	}
	return string(byteXml)
}
