package main

import (
	"io"
	"log"
	"net/http"
)

type machine struct {
	InstanceID string
	IPAddress string
}

func instanceID() (string, error) {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		return "", err
	}

	body := make([]byte, resp.ContentLength)
	if n, err := resp.Body.Read(body); err != nil {
		if err != io.EOF {
			return "", err
		}
		if int64(n) != resp.ContentLength {
			return "", err
		}
	}
	resp.Body.Close()

	log.Println(body, resp.ContentLength)
	return string(body), nil
}


func getMachine() (*machine, error) {
	inst, err := instanceID()
	if err != nil {
		return nil, err
	}

	return &machine{inst, ""}, nil
}