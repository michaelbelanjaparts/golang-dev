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
