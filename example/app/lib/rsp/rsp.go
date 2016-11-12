package rsp

import (
	"encoding/json"
)

func Json(v interface{}) string {
	result, _ := json.Marshal(v)

	return string(result)
}
