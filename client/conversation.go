package goopenai

import (
	"context"
	"fmt"
	"strings"
)

type Conversation struct {
	chatAI         *ChatAI
	client         *Client
	conversationID string
	lastMessageID  string
}

func newConversation(chatAI *ChatAI) (*Conversation, error) {
	if chatAI == nil {
		return nil, fmt.Errorf("create turbo session failed, open ai can't be nil")
	}
	c, err := NewClient()
	if err != nil {
		return nil, err
	}
	return &Conversation{
		chatAI: chatAI,
		client: c,
	}, nil
}

func (c *Conversation) Ask(ctx context.Context, question string) (chan string, error) {
	accessToken, err := c.chatAI.GetAccessToken()
	if err != nil {
		return nil, err
	}
	respChan, err := c.client.AskForConversation(question, c.conversationID, c.lastMessageID, accessToken)
	if err != nil {
		return nil, err
	}
	result := make(chan string)
	go func() {
		defer close(result)
		for resp := range respChan {
			result <- resp.Message
			c.conversationID = resp.ConversationID
			c.lastMessageID = resp.LastMessageID
		}
	}()
	return result, nil
}

func (c *Conversation) AskAndAnswer(ctx context.Context, question string) (string, error) {
	respChan, err := c.Ask(ctx, question)
	if err != nil {
		return "", err
	}
	result := new(strings.Builder)
	for resp := range respChan {
		result.WriteString(resp)
	}
	return result.String(), nil
}
