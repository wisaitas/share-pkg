package caller

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"gitlab.com/wisaitas1/trade-store-share-pkg/external"
)

type Caller interface {
	CallHttp(ctx context.Context, method external.Method, url string, headers []map[string]string, request any, response any) error
}

type caller struct {
}

func NewCaller() Caller {
	return &caller{}
}

func (c *caller) CallHttp(ctx context.Context, method external.Method, url string, headers []map[string]string, request any, response any) (err error) {
	httpHeader := http.Header{}
	for _, header := range headers {
		for key, value := range header {
			httpHeader.Set(key, value)
		}
	}

	bodyBuffer := bytes.NewBuffer(nil)
	if request != nil {
		bodyJson, err := json.Marshal(request)
		if err != nil {
			return err
		}
		bodyBuffer = bytes.NewBuffer(bodyJson)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, string(method), url, bodyBuffer)
	if err != nil {
		return err
	}
	httpRequest.Header = httpHeader

	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(responseBody, &response); err != nil {
		return err
	}

	return nil
}
