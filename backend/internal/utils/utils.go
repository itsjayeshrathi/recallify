package utils

import (
	"encoding/json"
	"net/http"
)

type Envelop map[string]any

func WriteJSON(w http.ResponseWriter, r *http.Request, status int, data Envelop) error {
	js, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	w.Header().Set("Conent-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

