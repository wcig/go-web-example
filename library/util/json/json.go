package json

import (
	"encoding/json"
	"fmt"
)

func PrintJson(v interface{}) {
	if b, err := json.Marshal(v); err == nil {
		fmt.Println(string(b))
	}
}

func PrintJsonPretty(v interface{}) {
	if b, err := json.MarshalIndent(v, "", "  "); err == nil {
		fmt.Println(string(b))
	}
}
