package goopenai

import "fmt"

type ChatAI struct {
	token string
}

func NewChatAI(token string) (*ChatAI, error) {
	if len(token) == 0 {
		return nil, fmt.Errorf("token can't be empty")
	}
	return &ChatAI{
		token: token,
	}, nil
}

func (o *ChatAI) NewConversation() (*Conversation, error) {
	return newConversation(o)
}

func (o *ChatAI) GetToken() string {
	return o.token
}

func (o *ChatAI) SetToken(token string) *ChatAI {
	o.token = token
	return o
}

func (o *ChatAI) GetAccessToken() (string, error) {
	return "", nil
}
