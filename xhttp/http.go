package xhttp

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func Request(req *http.Request, timeout time.Duration) (body []byte, err error) {
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request status code not equal to 200")
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"method":     req.Method,
		"url":        req.URL,
		"reqHeader":  req.Header,
		"respHeader": resp.Header,
		"respBody":   string(body),
	}).Tracef("http request")
	return body, nil
}

func RequestJson(req *http.Request, timeout time.Duration, v interface{}) error {
	body, err := Request(req, timeout)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
