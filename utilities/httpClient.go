package utilities

import (
	"github.com/quic-go/quic-go/http3"
	"io"
	"net/http"
	"time"
	"fmt"
)

type HttpClientConfig struct {
    Timeout   time.Duration 
}

type CustomHttpClient struct {
    Client *http.Client
}

func NewCustomHttpClient(config *http.Client) *CustomHttpClient {
	httpClient := &http.Client{
		Timeout:   config.Timeout,
		Transport: &http3.RoundTripper{}, 
	}
	
	return &CustomHttpClient{
		Client: httpClient,
	}
}

func (c *CustomHttpClient) Get(url string) ([]byte, error) {
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err 
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {

			fmt.Printf("Error closing response body: %v\n", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err 
	}

	return body, nil 
}