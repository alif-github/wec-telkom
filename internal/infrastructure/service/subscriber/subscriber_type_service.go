package service

import (
	"bytes"
	domain_subscriber "content-management/domain/subscriber"
	"content-management/helper"
	"content-management/helper/exceptions"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type SubscriberTypeService struct {
	handlerClient *http.Client
}

func NewSubscriberTypeService(handler *http.Client) *SubscriberTypeService {
	return &SubscriberTypeService{handler}
}

func (p SubscriberTypeService) GetSubscriberType(msisdn string) (string, error) {
	var res domain_subscriber.ResponseSubscriber
	reqBody, err := json.Marshal(&domain_subscriber.RequestSubscriber{
		Initialize: false,
		Msisdn:     msisdn,
	})

	if err != nil {
		return "", err
	}

	urlStr := os.Getenv("POIN_SUBSCRIBE_URL")
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(reqBody))

	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:      "SERVICE|SUBSCRIBER|GETSUBSCRIBERTYPE|NEWREQUEST",
			StatusCode: req.Response.StatusCode,
			Method:     "POST",
			Request:    string(reqBody),
			URL:        urlStr,
			Response:   err,
		}, "error")

		return "", exceptions.ErrBadRequest
	}

	// req.Header.Set("x-msisdn", msisdn)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := p.handlerClient.Do(req)

	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:      "SERVICE|SUBSCRIBER|GETSUBSCRIBERTYPE|CLIENTDO",
			StatusCode: http.StatusInternalServerError,
			Method:     "POST",
			Request:    string(reqBody),
			URL:        urlStr,
			Message:    err.Error(),
		}, "error")

		return "", exceptions.ErrBadRequest
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:      "SERVICE|SUBSCRIBER|GETSUBSCRIBERTYPE|READALL",
			StatusCode: http.StatusInternalServerError,
			Method:     "POST",
			Request:    msisdn,
			URL:        urlStr,
			Response:   err,
		}, "error")

		return "", exceptions.ErrBadRequest
	}

	io.Copy(ioutil.Discard, resp.Body) // <= NOTE
	defer resp.Body.Close()

	defer helper.CreateLog(&helper.Log{
		Event:      "poin_subscribe",
		StatusCode: resp.StatusCode,
		Method:     resp.Request.Method,
		Request:    string(reqBody),
		URL:        urlStr,
		Response:   string(body),
	}, "info")

	if resp.StatusCode == 200 {
		json.Unmarshal([]byte(body), &res)
	} else if resp.StatusCode == 500 {
		return "", exceptions.ErrInternalServerError
	} else if resp.StatusCode == 502 {
		return "", exceptions.ErrInternalServerError
	} else if resp.StatusCode == 504 {
		return "", exceptions.ErrTimeOut
	} else {
		return "", exceptions.ErrBadRequest
	}

	return res.SubscriberType, nil
}
