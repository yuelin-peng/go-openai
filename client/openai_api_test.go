package goopenai_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	goopenai "github.com/yuelin-pengk/go-openai/client"
)

func TestGetAccessToken(t *testing.T) {
	t.Skip()

	c, err := goopenai.NewClient()
	assert.Nil(t, err)
	assert.NotNil(t, c)

	accessToken, expiryTime, err := c.GetAccessToken(getDefaultToken())
	assert.Nil(t, err)
	assert.Greater(t, len(accessToken), 0)
	assert.Greater(t, expiryTime.Unix(), int64(0))
}

func TestAskForConversation(t *testing.T) {
	t.Skip()

	c, err := goopenai.NewClient()
	assert.Nil(t, err)
	assert.NotNil(t, c)

	accessToken, expiryTime, err := c.GetAccessToken(getDefaultToken())
	assert.Nil(t, err)
	assert.Greater(t, len(accessToken), 0)
	assert.Greater(t, expiryTime.Unix(), int64(0))

	respChan, err := c.AskForConversation("hello", "", "", accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, respChan)
}
