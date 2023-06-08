package gochatgpt

import (
	"errors"
)

type Config struct {
	API          string
	Organization string
	APIKey       string
}

type IClient interface {
	Models() (ResponseModelList, error)
	Model(modelID string) (ResponseModel, error)
	Chat(msg string) (ResponseChat, error)
	GetImage() IGPTImages
}

type client struct {
	cfg Config
	req *request
}

func (c client) Models() (res ResponseModelList, err error) {
	res = ResponseModelList{}
	err = c.req.Get("/models", &res)
	return
}

func (c client) Model(modelID string) (res ResponseModel, err error) {
	res = ResponseModel{}
	err = c.req.Get("/models/"+modelID, &res)
	return
}

func (c client) Chat(msg string) (res ResponseChat, err error) {
	bodyArgs := map[string]interface{}{
		"model": GPT35Turbo,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": msg,
			},
		},
	}
	res = ResponseChat{}
	err = c.req.Post("/chat/completions", bodyArgs, &res)

	return
}

func (c client) GetImage() (img IGPTImages) {
	img = newGptImage(c.req)
	return
}

func NewClient(cfg Config) (c IClient, err error) {
	if cfg.API == "" {
		cfg.API = "https://api.openai.com/v1"
	}
	if cfg.APIKey == "" {
		err = errors.New("api key is empty")
		return
	}
	if cfg.Organization == "" {
		err = errors.New("organization is empty")
		return
	}
	c = &client{
		cfg: cfg,
		req: newRequest(cfg),
	}

	return
}
