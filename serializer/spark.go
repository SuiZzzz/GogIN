package serializer

type SparkReq struct {
	Header    *Header    `json:"header,omitempty"`
	Parameter *Parameter `json:"parameter,omitempty"`
	Payload   *Payload   `json:"payload,omitempty"`
}

type SparkResp struct {
	Header  *Header  `json:"header,omitempty"`
	Payload *Payload `json:"payload,omitempty"`
}

type Header struct {
	AppId   string `json:"app_id,omitempty"`
	Uid     string `json:"uid,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Sid     string `json:"sid,omitempty"`
	Status  int    `json:"status,omitempty"`
}
type Parameter struct {
	Chat *Chat `json:"chat,omitempty"`
}

type Payload struct {
	Message *Message `json:"message,omitempty"`
	Choices *Choices `json:"choices,omitempty"`
	Usage   *Usage   `json:"usage,omitempty"`
}

type Chat struct {
	Domain      string  `json:"domain,omitempty"`
	Temperature float32 `json:"temperature,omitempty"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	TopK        int     `json:"top_k,omitempty"`
	ChatId      string  `json:"chat_id,omitempty"`
}

type Text struct {
	// 如果想获取结合上下文的回答，需要开发者每次将历史问答信息一起传给服务端
	// 注意：text里面的所有content内容加一起的tokens需要控制在8192以内，开发者如有较长对话需求，需要适当裁剪历史信息
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
	Index   int    `json:"index,omitempty"`
}

type Message struct {
	Text *[]Text `json:"text,omitempty"`
}

type Choices struct {
	Status int     `json:"status,omitempty"`
	Seq    int     `json:"seq,omitempty"`
	Text   *[]Text `json:"text,omitempty"`
}

type Usage struct {
	QuestionTokens   int `json:"question_tokens,omitempty"`
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}
