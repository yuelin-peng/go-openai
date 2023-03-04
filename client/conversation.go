package goopenai

import (
	"fmt"
)

type Conversation struct {
	chatAI *ChatAI
}

const (
	CommonSessionName = "common_session"
	url               = "https://chat.openai.com/backend-api/conversation"
)

func newConversation(chatAI *ChatAI) (*Conversation, error) {
	if chatAI == nil {
		return nil, fmt.Errorf("create turbo session failed, open ai can't be nil")
	}
	return &Conversation{
		chatAI: chatAI,
	}, nil
}

func (s *Conversation) Ask(question string) (string, error) {
	return "", nil
}
