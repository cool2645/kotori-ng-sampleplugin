package handler

import (
	"net/http"
	"github.com/yanzay/log"
	"encoding/json"
)

func responseJson(w http.ResponseWriter, data map[string]interface{}, httpStatusCode int) {
	resJson, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		http.Error(w, "Error occurred encoding response.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write(resJson)
	return
}

func Pong(w http.ResponseWriter, req *http.Request) {
	res := map[string]interface{}{
		"code":   http.StatusOK,
		"result": true,
		"msg":    "OK From Sample Plugin",
	}
	responseJson(w, res, http.StatusOK)
	return
}
