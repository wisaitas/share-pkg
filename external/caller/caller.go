package caller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wisaitas/shared-pkg/response"
)

type Caller interface {
	CallHttp(ctx context.Context, method string, url string, headers []map[string]string, request any, response *response.ApiResponse[any]) error
}

type caller struct {
}

func NewCaller() Caller {
	return &caller{}
}

// CallHttpWithResponse - Pass by reference, return only error
func (c *caller) CallHttp(ctx context.Context, method string, url string, headers []map[string]string, req any, resp *response.ApiResponse[any]) error {
	httpHeader := http.Header{}
	for _, header := range headers {
		for key, value := range header {
			httpHeader.Set(key, value)
		}
	}

	bodyBuffer := bytes.NewBuffer(nil)
	if req != nil {
		bodyJson, err := json.Marshal(req)
		if err != nil {
			return fmt.Errorf("[Share Package Caller] : %w", err)
		}
		bodyBuffer = bytes.NewBuffer(bodyJson)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, string(method), url, bodyBuffer)
	if err != nil {
		return fmt.Errorf("[Share Package Caller] : %w", err)
	}
	httpRequest.Header = httpHeader

	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return fmt.Errorf("[Share Package Caller] : %w", err)
	}
	defer httpResponse.Body.Close()

	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return fmt.Errorf("[Share Package Caller] : %w", err)
	}

	if err := json.Unmarshal(responseBody, resp); err != nil {
		return fmt.Errorf("[Share Package Caller] failed to unmarshal response: %w", err)
	}

	resp.HttpStatusCode = httpResponse.StatusCode
	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {
		return &ServiceError{
			StatusCode: httpResponse.StatusCode,
			Response:   resp,
		}
	}

	return nil
}
