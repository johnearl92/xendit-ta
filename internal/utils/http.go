package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/johnearl92/xendit-ta.git/internal/model/errors"
)

// ReadEntity reads request
func ReadEntity(req *http.Request, entity interface{}) errors.JSONErrors {
	body := req.Body
	defer body.Close()
	serializedJSON, err := ioutil.ReadAll(body)
	if err != nil {
		return errors.New().Add(
			"400",
			map[string]string{"pointer": "/data"},
			"Unable to read request",
			"Request body is not readable")
	}
	return FromJSON(serializedJSON, entity)
}

// WriteEntity writes response
func WriteEntity(res http.ResponseWriter, code int, entity interface{}) {
	serializedJSON, err := ToJSON(entity)
	if err != nil {
		WriteError(res, http.StatusInternalServerError, errors.New().Add(
			"500",
			map[string]string{"pointer": "/data"},
			"Response cannot be encoded as JSON",
			err.Error()))
		return
	}

	if _, wErr := res.Write(serializedJSON); wErr != nil {
		WriteError(res, http.StatusInternalServerError, errors.New().Add(
			"500",
			map[string]string{"pointer": "/data"},
			"Unable to write response",
			"Response body is not writable",
		))
		return
	}
}

// WriteServerError writes http.StatusNotFound status code to http response.
// It writes JSONErrors to http response body with configurable param.
func WriteServerError(res http.ResponseWriter, param string, title string, desc string) {
	WriteError(res, http.StatusNotFound, errors.New().Add(
		"500",
		map[string]string{"pointer": param},
		title,
		desc))
}

// WriteError writes errors
func WriteError(res http.ResponseWriter, code int, jsonErr errors.JSONErrors) {
	res.WriteHeader(code)
	serializedJSON, err := ToJSON(jsonErr)
	if err != nil {
		log.Error("unable to serialize error response")
		return
	}
	if _, wErr := res.Write(serializedJSON); wErr != nil {
		log.WithError(wErr).Error("unable to write error response")
	}
}

// SendPostRequest sends POST request.
func SendPostRequest(data interface{}, url string) error {
	return SendRequest(data, url, http.MethodPost)
}

// SendDeleteRequest sends DELETE request.
func SendDeleteRequest(url string) error {
	return SendRequest(nil, url, http.MethodDelete)
}

// SendRequest sends HTTP request.
func SendRequest(data interface{}, url string, method string) error {
	log.Debugf("HTTP %s Request URL: %s", method, url)

	httpRequestBody, jsonErr := ToJSON(data)
	if jsonErr != nil {
		log.Error(jsonErr)
		return fmt.Errorf("Failed to create HTTP %s request body", method)
	}

	log.Debugf("HTTP %s Request body: %s", method, string(httpRequestBody))

	httpRequest, err := http.NewRequest(method, url, bytes.NewBuffer(httpRequestBody))
	if err != nil {
		log.Error(err.Error())
		return fmt.Errorf("Failed to create HTTP %s request", method)
	}

	httpRequest.Header.Add("Content-Type", "application/json")

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		log.Error(err.Error())
		return fmt.Errorf("Failed to send HTTP %s request", method)
	}

	defer httpResponse.Body.Close()

	httpResponseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Error(err.Error())
		return fmt.Errorf("Failed to read HTTP %s response body", method)
	}

	log.Debugf("HTTP %s Response Body: %s", method, string(httpResponseBody))

	if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		err = fmt.Errorf("HTTP %s Request Failure: status_code=%d body=%s", method, httpResponse.StatusCode, string(httpResponseBody))
		log.Error(err.Error())
		return err
	}

	return nil
}
