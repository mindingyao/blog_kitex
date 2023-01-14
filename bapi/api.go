package bapi

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type API struct {
	URL string
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

func (a *API) httpPost(ctx context.Context, path string, from string) ([]byte, error) {

	resp, err := http.Post(path, "application/x-www-form-urlencoded",
		strings.NewReader(from))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func (a *API) GetTagList(ctx context.Context, name string, page, pageSize int64) ([]byte, error) {
	FormBody := fmt.Sprintf("name=%s&page=%d&page_size=%d", name, page, pageSize)
	body, err := a.httpPost(ctx, "http://127.0.0.1:8888/v1/tag/query/", FormBody)
	if err != nil {
		return nil, err
	}

	return body, nil
}
