package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	APP_KEY    = "terry"
	APP_SECRET = "blog-service"
)

type AccessToken struct {
	Token string `json:"token"`
}
type API struct {
	URL string
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	body, err := a.httpGet(ctx, fmt.Sprintf(
		"%s?name=%s&token=%s",
		"api/v1/tags",
		name,
		token,
	))
	log.Println("token", token)
	if err != nil {

		return nil, err
	}
	return body, nil
}

// 所有API請求都要帶上AccessToken
func (a *API) getAccessToken(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", a.URL+"/auth", nil)

	if err != nil {
		return "", err
	}

	//在header中加入app_key&app_secret
	req.Header.Add("app_key", APP_KEY)
	req.Header.Add("app_secret", APP_SECRET)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	//驗證 /auth 回傳token
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var accessToken AccessToken
	_ = json.Unmarshal(body, &accessToken)

	return accessToken.Token, nil

}

// API SDK統一的http get的請求方法
func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil

}
