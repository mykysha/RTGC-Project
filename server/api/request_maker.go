package api

import (
	"encoding/json"
	"fmt"

	v1 "github.com/nndergunov/RTGC-Project/server/api/v1"
)

func encode(r v1.Response) ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("encode: %w", err)
	}

	return data, nil
}

func decode(msg []byte) (v1.Request, error) {
	var req v1.Request

	err := json.Unmarshal(msg, &req)
	if err != nil {
		return req, fmt.Errorf("decode: %w", err)
	}

	return req, nil
}

func statusEncoder(s v1.State) ([]byte, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("statusEncoder: %w", err)
	}

	return data, nil
}
