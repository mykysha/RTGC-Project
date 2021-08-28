package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func encode(w http.ResponseWriter, s Status) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	io.WriteString(w, string(data))
	return nil
}
