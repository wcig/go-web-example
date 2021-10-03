package json

import (
	"encoding/json"
	"fmt"
)

func ToJson(v interface{}) string {
	if b, err := json.Marshal(v); err == nil {
		return string(b)
	}
	return ""
}

func ToJsonPretty(v interface{}) string {
	if b, err := json.MarshalIndent(v, "", "  "); err == nil {
		return string(b)
	}
	return ""
}

func PrintJson(v interface{}) {
	fmt.Println(ToJson(v))
}

func PrintJsonPretty(v interface{}) {
	fmt.Println(ToJsonPretty(v))
}
