package chat

import "encoding/json"

type SimpleText struct {
	Text string `json:"text"`
}

func SimpleTextMessage(text string) string {
	content := &SimpleText{
		Text: text,
	}

	data, _ := json.Marshal(content)

	return string(data)
}
