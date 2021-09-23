package api

import (
	"encoding/json"
	"log"

	"github.com/nndergunov/RTGC-Project/api/v1"
)

func encode(r v1.Response) ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return data, nil
}

func decode(msg []byte) (v1.Request, error) {
	var req v1.Request

	err := json.Unmarshal(msg, &req)
	if err != nil {
		log.Println(err)

		return req, err
	}

	return req, nil
}

func statusEncoder(s v1.Status) ([]byte, error) {
	data, err := json.Marshal(s)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return data, nil
}
