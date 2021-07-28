package utils

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, code int, res interface{}) {
	json_data, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(json_data))
}

func CompareString(req string, res string) (similar bool) {
	similar = true
	if req != "" {
		if req != res {
			similar = false
			return similar
		}
	}
	return similar
}

func CompareInt(req int, res int) (similar bool) {
	similar = true
	if req != 0 {
		if req != res {
			similar = false
			return similar
		}
	}
	return similar
}
