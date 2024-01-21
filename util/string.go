package util

import "encoding/json"

func JsonMarshal(v any) string {
	vByte, _ := json.Marshal(v)
	return string(vByte)
}
