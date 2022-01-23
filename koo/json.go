package koo

import (
	"encoding/json"
	"fmt"
)

func PrintData(d interface{}) {
	x := map[string]interface{}{"a": 1, "b": 2}
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}
