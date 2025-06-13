package ysgo

type ClientOption func(*YSClient)

// WithApiBaseUrl Set the api Base Url
func WithApiBaseUrl(apiBaseUrl string) ClientOption {
	return func(c *YSClient) {
		c.apiBaseUrl = apiBaseUrl
	}
}
