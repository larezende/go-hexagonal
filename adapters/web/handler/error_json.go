package handler

import "encoding/json"

func jsonError(msg string) []byte {
	jsonMsg := struct {
		Message string `json:"message"`
	}{
		msg,
	}
	r, err := json.Marshal(jsonMsg)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
