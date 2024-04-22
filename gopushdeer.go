package gopushdeer

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type GoPushDeer struct {
	Server string
	Key    string
}

func NewGoPushDeer(key string) (*GoPushDeer, error) {
	if strings.TrimSpace(key) == "" {
		return nil, errors.New("key must not be nil or empty")
	}

	return &GoPushDeer{
		Server: "https://api2.pushdeer.com",
		Key:    key,
	}, nil
}

func NewGoPushDeerServer(server, key string) (*GoPushDeer, error) {
	if strings.TrimSpace(key) == "" {
		return nil, errors.New("key must not be nil or empty")
	}
	if server == "" {
		server = "https://api2.pushdeer.com"
	}

	return &GoPushDeer{
		Server: server,
		Key:    key,
	}, nil
}

func (gpd *GoPushDeer) sendText(text string) error {
	encodedText := url.QueryEscape(text)
	requestUrl := fmt.Sprintf("%s/message/push?pushkey=%s&text=%s", gpd.Server, gpd.Key, encodedText)
	resp, getErr := http.Get(requestUrl)
	if getErr != nil {
		return getErr
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Resp io close error, ", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Http Resp: %s\n", body)
	return nil
}
