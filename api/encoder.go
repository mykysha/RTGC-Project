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

	_, err = io.WriteString(w, string(data))
	if err != nil {
		return err
	}

	return nil
}
