package gochatgpt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"
)

type request struct {
	cfg Config
}

func (r request) Get(url string, res interface{}) (err error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodGet)
	req.SetRequestURI(r.cfg.API + url)
	req.Header.Add("Authorization", "Bearer "+r.cfg.APIKey)
	req.Header.Add("OpenAI-Organization", r.cfg.Organization)
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	if err = fasthttp.Do(req, resp); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &res)

	return
}

func (r request) Post(url string, body interface{}, res interface{}) (err error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodPost)
	req.SetRequestURI(r.cfg.API + url)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+r.cfg.APIKey)
	if body != nil {
		var v []byte
		v, err = json.Marshal(body)
		if err != nil {
			return
		}
		fmt.Println(string(v))
		req.SetBody(v)
	}
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	if err = fasthttp.Do(req, resp); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &res)

	return
}

func newRequest(cfg Config) *request {
	return &request{
		cfg: cfg,
	}
}
