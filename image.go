package gochatgpt

type ImageSize string

func (i ImageSize) String() string {
	return string(i)
}

type ImageResponseFormat string

func (i ImageResponseFormat) String() string {
	return string(i)
}

const (
	X256  ImageSize = "256x256"
	X512  ImageSize = "512x512"
	X1024 ImageSize = "1024x1024"
)

type RequestCreateImage struct {
	Prompt     string              `json:"prompt"`
	Number     int                 `json:"n"`
	Size       ImageSize           `json:"size"`
	RespFormat ImageResponseFormat `json:"response_format"`
	User       string              `json:"user"`
}

type ResponseImage struct {
	Created int64       `json:"created"`
	Data    interface{} `json:"data"`
}

type IGPTImages interface {
	Create(req RequestCreateImage) (res ResponseImage, err error)
}

type gptImage struct {
	req *request
}

func (g gptImage) Create(req RequestCreateImage) (res ResponseImage, err error) {
	if req.Number == 0 {
		req.Number = 1
	}
	if req.RespFormat == "" {
		req.RespFormat = "url"
	}
	if string(req.Size) == "" {
		req.Size = X256
	}

	res = ResponseImage{}
	err = g.req.Post("/images/generations", req, &res)
	return
}

func newGptImage(req *request) IGPTImages {
	return &gptImage{
		req: req,
	}
}
