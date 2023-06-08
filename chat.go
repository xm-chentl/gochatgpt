package gochatgpt

type ResponseChat struct {
	ID      string               `json:"id"`
	Object  string               `json:"object"`
	Created int64                `json:"created"`
	Choices []ResponseChatChoice `json:"choices"`
	Usage   ResponseChatUsage
}

type ResponseChatChoice struct {
	Index        int                       `json:"index"`
	Message      ResponseChatChoiceMessage `json:"message"`
	FinishReason string                    `json:"finish_reason"`
}

type ResponseChatChoiceMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseChatUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
