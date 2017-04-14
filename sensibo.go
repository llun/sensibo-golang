package sensibo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Sensibo struct {
	Key string

	client *http.Client
}

func NewSensibo(key string) *Sensibo {
	return &Sensibo{
		Key: key,
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (s *Sensibo) get(resource string, values url.Values) ([]byte, error) {
	resp, err := s.client.Get(s.resourceUrl(resource, values))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *Sensibo) resourceUrl(resource string, values url.Values) string {
	values["apiKey"] = []string{s.Key}
	fullUrl := &url.URL{
		Scheme:   "https",
		Host:     "home.sensibo.com",
		Path:     fmt.Sprintf("/api/v2/%v", resource),
		RawQuery: values.Encode(),
	}
	return fullUrl.String()
}
