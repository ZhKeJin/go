package main


import (
"encoding/json"
"fmt"
)

func main() {
	txt := `{"a":1,"b":2,"c":[{"name":"1","group":"2"},{"name":"3","group":"4"}]}`
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(txt), &m); err != nil {
		panic(err)
	}
	for _, v := range m["c"].([]interface{}) {
		fmt.Println(v)
	}
}
