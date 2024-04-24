package gopushdeer

import (
	"encoding/json"
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

type PushDeerResponse struct {
	Code    int `json:"code"`
	Content struct {
		Result []string `json:"result"`
	} `json:"content"`
}

type PushDeerResponseResult struct {
	Counts  int           `json:"counts"`
	Logs    []interface{} `json:"logs"`
	Success string        `json:"success"`
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

func (gpd *GoPushDeer) request(requestUrl string) error {
	resp, getErr := http.Get(requestUrl)
	if getErr != nil {
		return getErr
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("push deer http response: " + resp.Status)
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

	//fmt.Printf("Http Resp: %s\n", body)

	pdr := PushDeerResponse{}
	if jsonErr := json.Unmarshal(body, &pdr); jsonErr != nil {
		return jsonErr
	}
	if pdr.Code != 0 {
		return errors.New(fmt.Sprintf("push deer response code not success: %d", pdr.Code))
	}
	if len(pdr.Content.Result) == 0 {
		return errors.New("push deer response result is empty")
	}

	result := PushDeerResponseResult{}
	if jsonErr := json.Unmarshal([]byte(pdr.Content.Result[0]), &result); jsonErr != nil {
		return jsonErr
	}
	if strings.Compare("ok", result.Success) != 0 {
		return errors.New("push deer response result not ok, " + pdr.Content.Result[0])
	}

	return nil
}

func (gpd *GoPushDeer) SendText(text string, optional ...string) error {
	base, err := url.Parse(gpd.Server)
	if err != nil {
		return err
	}
	base.Path += "/message/push"
	params := url.Values{}
	params.Add("pushkey", gpd.Key)
	params.Add("text", text)
	if len(optional) > 0 {
		params.Add("desp", optional[0])
	}
	base.RawQuery = params.Encode()

	return gpd.request(base.String())
}

func (gpd *GoPushDeer) SendImage(imageUrl string, optional ...string) error {
	base, err := url.Parse(gpd.Server)
	if err != nil {
		return err
	}
	base.Path += "/message/push"
	params := url.Values{}
	params.Add("pushkey", gpd.Key)
	params.Add("text", imageUrl)
	params.Add("type", "image")
	if len(optional) > 0 {
		params.Add("desp", optional[0])
	}
	base.RawQuery = params.Encode()

	return gpd.request(base.String())
}

func (gpd *GoPushDeer) SendMarkdown(md string, optional ...string) error {
	//	https://api2.pushdeer.com/message/push?pushkey=<key>&text=标题&desp=<markdown>&type=markdown
	base, err := url.Parse(gpd.Server)
	if err != nil {
		return err
	}
	base.Path += "/message/push"
	params := url.Values{}
	params.Add("pushkey", gpd.Key)
	params.Add("text", md)
	params.Add("type", "markdown")
	if len(optional) > 0 {
		params.Add("desp", optional[0])
	}
	base.RawQuery = params.Encode()

	return gpd.request(base.String())
}
