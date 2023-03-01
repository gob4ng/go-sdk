package utils

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"time"
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

func XmlToString(jsonStruct interface{}) string {
	byteXml, err := xml.Marshal(jsonStruct)
	if err != nil {
		return err.Error()
	}
	return string(byteXml)
}

func GetUnixTimestamp() int64 {
	return time.Now().UnixNano()
}
