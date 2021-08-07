package utils

import (
	"github.com/alvanrahimli/dots-cli/dlog"
	"io"
	"net/http"
	"strings"
)

// HttpPost issues HTTP POST request to given url, returns Body, StatusCode, error
func HttpPost(url string, headers map[string]string, body *strings.Reader) ([]byte, int, error) {
	client := &http.Client{}
	req, reqErr := http.NewRequest("POST", url, body)
	if reqErr != nil {
		dlog.Err(reqErr.Error())
		return nil, 0, reqErr
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	dlog.Debug("Sending request url: '%s', headers: '%s', body: *", url, headers)

	response, err := client.Do(req)
	dlog.Info("Sent POST request to '%s'", url)
	if err != nil {
		dlog.Err(err.Error())
		return nil, 0, err
	}
	//goland:noinspection ALL
	defer response.Body.Close()

	bodyBytes, bodyErr := io.ReadAll(response.Body)
	if bodyErr != nil {
		dlog.Err(bodyErr.Error())
		return nil, 0, bodyErr
	}

	dlog.Info("Received StatusCode: %d and Body: %s", response.StatusCode, string(bodyBytes))
	return bodyBytes, response.StatusCode, nil
}
