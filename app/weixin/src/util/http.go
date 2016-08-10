package util

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	. "global"
)

const (
	HTTP_GET    = "get"
	HTTP_POST   = "post"
	HTTP_PUT    = "put"
	HTTP_DELETE = "delete"
)

func CallHttp(method string, reqUrl string, params url.Values) ([]byte, error) {
	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	var (
		resp *http.Response
		err  error
	)
	switch method {
	case HTTP_POST:
		resp, err = httpClient.PostForm(reqUrl, params)
	case HTTP_PUT:
		req, err := http.NewRequest("PUT", reqUrl, strings.NewReader(params.Encode()))
		if err != nil {
			// log
			break
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err = httpClient.Do(req)
	case HTTP_DELETE:
		req, err := http.NewRequest("DELETE", reqUrl+"?"+params.Encode(), nil)
		if err != nil {
			// log
			break
		}
		resp, err = httpClient.Do(req)
	case HTTP_GET:
		resp, err = httpClient.Get(reqUrl + "?" + params.Encode())
	default:
		// log("not consist method")
		return nil, HttpMethodNot
	}

	if err != nil {
		// logger.Errorf("url:%q, params:%v, error: %v", reqUrl, params, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		// logger.Errorf("url:%q, params:%v, status: %s(%d)", reqUrl, params, resp.Status, resp.StatusCode)
		return nil, errors.New("status code is not ok")
	}

	resultBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if strings.ToUpper(method) == "POST" {
			// logger.Errorf("url:%q, params:%v, result: %s, error: %v", reqUrl, params, resultBuf, err)
		} else {

			// logger.Errorf("url:%q, result: %s, error: %v", reqUrl, resultBuf, err)
		}
		return nil, err

	}

	return resultBuf, nil
}
