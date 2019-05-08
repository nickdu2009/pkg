package xjson

import "encoding/json"

func MustToJsonString(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func SafeMarshal(v interface{}) []byte {
	buf, _ := json.Marshal(v)
	return buf
}