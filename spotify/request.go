package spotify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Request(token, method, url string, payload interface{}, query map[string]string, response interface{}) error {
	if token == "" {
		return errors.New("token may not be empty") // nolint: goerr113
	}

	body, err := marshalPayload(payload)
	if err != nil {
		return err
	}

	// nolint: noctx
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return clientErr(err, url, nil)
	}

	for k, v := range query {
		r.URL.Query().Add(k, v)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	c := http.Client{
		Timeout: time.Second,
	}

	resp, err := c.Do(r)
	if err != nil {
		return clientErr(err, url, nil)
	}
	defer resp.Body.Close()

	var success = false
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		success = true
	}

	if resp.StatusCode != http.StatusNoContent {
		responseBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return clientErr(err, url, resp.StatusCode)
		}

		err = json.Unmarshal(responseBytes, &response)
		if err != nil {
			return clientErr(err, url, resp.StatusCode)
		}
	}

	if !success {
		if errResp, ok := response.(error); ok {
			return clientErr(errResp.Error(), url, resp.StatusCode)
		}

		return clientErr("unexpected status code", url, resp.StatusCode)
	}

	return nil
}

func clientErr(err interface{}, url string, statusCode interface{}) error {
	// nolint: goerr113
	return fmt.Errorf(
		"spotify api client error calling url=%s status=%v error=%v",
		url,
		statusCode,
		err,
	)
}

func marshalPayload(payload interface{}) (*bytes.Buffer, error) {
	var body []byte

	if payload != nil {
		var err error

		body, err = json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("error marshallig payload: %w", err)
		}
	}

	return bytes.NewBuffer(body), nil
}
