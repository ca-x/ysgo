package ysgo

type YSClient struct {
	Endpoint       string `json:"endpoint"`
	ManagementPass string `json:"-"`
}

func NewClient(endpoint string, managementPass string) *YSClient {
	return &YSClient{
		Endpoint:       endpoint,
		ManagementPass: managementPass,
	}
}
