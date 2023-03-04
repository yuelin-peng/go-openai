package goopenai

import (
	"fmt"

	"github.com/yuelin-pengk/go-openai/util"
)

type ChatAI struct {
	token          string
	accessTokenMap util.ExpiryMap
	client         *Client
}

func NewChatAI(token string) (*ChatAI, error) {
	if len(token) == 0 {
		return nil, fmt.Errorf("token can't be empty")
	}
	client, err := NewClient()
	if err != nil {
		return nil, err
	}
	return &ChatAI{
		token:          token,
		accessTokenMap: util.NewExpiryMap(),
		client:         client,
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
	if accessToken, ok := o.accessTokenMap.Get(o.token); !ok || len(accessToken) <= 0 {
		accessToken, expiryTime, err := o.client.GetAccessToken(o.token)
		if err != nil {
			return "", err
		}
		o.accessTokenMap.Set(o.token, accessToken, expiryTime)
		return accessToken, nil
	} else {
		return accessToken, nil
	}
}
