package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseWriter (w http.ResponseWriter, x interface{}) {
	res, _ := json.Marshal(x)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}