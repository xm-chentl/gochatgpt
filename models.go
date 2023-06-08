package gochatgpt

type Model string

func (m Model) String() string {
	return string(m)
}

const (
	GPT35Turbo Model = "gpt-3.5-turbo"
)

type ResponseModelPermissionItem struct {
	ID                string      `json:"id"`
	Object            string      `json:"object"`
	Created           int64       `json:"created"`
	AllowCreateEngine bool        `json:"allow_create_engine"`
	AllowSampling     bool        `json:"allow_sampling"`
	AllowLogprobs     bool        `json:"allow_logprobs"`
	AllowView         bool        `json:"allow_view"`
	AllowFineTuning   bool        `json:"allow_fine_trning"`
	Organization      string      `json:"organization"`
	Group             interface{} `json:"group"`
	IsBlocking        bool        `json:"is_blocking"`
}

type ResponseModel struct {
	ID         string                        `json:"id"`
	Object     string                        `json:"object"`
	Created    int64                         `json:"created"`
	OwnedBy    string                        `json:"owned_by"`
	Permission []ResponseModelPermissionItem `json:"permission"`
	Root       string                        `json:"roor"`
	Parent     string                        `json:"parent"`
}

type ResponseModelList struct {
	Data   []ResponseModel `json:"data"`
	Object string          `json:"object"`
}
