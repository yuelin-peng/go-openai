package goopenai_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConversationAsk(t *testing.T) {
	//t.Skip()

	chatAI := createDefaultOpenAI(t)
	c, err := chatAI.NewConversation()
	assert.Nil(t, err)
	assert.NotNil(t, c)

	respChan, err := c.Ask(context.Background(), "hello")
	fmt.Println("finish ask", err)
	assert.Nil(t, err)
	assert.NotNil(t, respChan)
	if err == nil {
		for resp := range respChan {
			fmt.Println("resp=", resp)
		}
	}
}
