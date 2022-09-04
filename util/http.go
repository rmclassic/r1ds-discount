package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func HttpSendRequest(url string, method string, body interface{}, headers []string, response interface{}) error {

	jbody, _ := json.Marshal(&body)
	req, _ := http.NewRequest(method, url, strings.NewReader(string(jbody)))
	req.Header["Content-Type"] = []string{"application/json"}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode/100 != 2 {
		return errors.New("failed request to " + url)
	}

	rbody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if response != nil {
		if err = json.Unmarshal(rbody, response); err != nil {
			return err
		}
	}

	return nil
}
