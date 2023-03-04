package goopenai

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/launchdarkly/eventsource"
)

type Client struct {
}

func NewClient() (*Client, error) {
	return &Client{}, nil
}

const (
	sessionURL      = "https://chat.openai.com/api/auth/session"
	sessionTokenKey = "__Secure-next-auth.session-token"
	userAgent       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"

	conversationURL = "https://chat.openai.com/backend-api/conversation"
)

type SessionResult struct {
	Error       string `json:"error"`
	Expires     string `json:"expires"`
	AccessToken string `json:"accessToken"`
}

func (o *Client) GetAccessToken(sessionToken string) (string, time.Time, error) {
	req, err := http.NewRequest("GET", sessionURL, nil)
	if err != nil {
		return "", time.Now(), fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Cookie", fmt.Sprintf(sessionTokenKey+"=%s", sessionToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", time.Now(), fmt.Errorf("failed to perform request: %v", err)
	}
	defer res.Body.Close()

	var result SessionResult
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return "", time.Now(), fmt.Errorf("failed to decode response: %v", err)
	}

	accessToken := result.AccessToken
	if accessToken == "" {
		return "", time.Now(), errors.New("unauthorized")
	}

	if result.Error != "" {
		if result.Error == "RefreshAccessTokenError" {
			return "", time.Now(), errors.New("Session token has expired")
		}

		return "", time.Now(), errors.New(result.Error)
	}

	expiryTime, err := time.Parse(time.RFC3339, result.Expires)
	if err != nil {
		return "", time.Now(), fmt.Errorf("failed to parse expiry time: %v", err)
	}

	return accessToken, expiryTime, nil
}

type MessageResponse struct {
	ConversationId string `json:"conversation_id"`
	Error          string `json:"error"`
	Message        struct {
		ID      string `json:"id"`
		Content struct {
			Parts []string `json:"parts"`
		} `json:"content"`
	} `json:"message"`
}

type ChatResponse struct {
	Message        string
	ConversationID string
	LastMessageID  string
}

// return new conversation id and response channel
func (o *Client) AskForConversation(question, conversationID, lastMessageID, accessToken string) (chan ChatResponse, error) {
	messages, err := json.Marshal([]string{question})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to encode question: %v", err))
	}

	if lastMessageID == "" {
		lastMessageID = uuid.NewString()
	}

	var conversationIdString string
	if conversationID != "" {
		conversationIdString = fmt.Sprintf(`, "conversation_id": "%s"`, conversationID)
	}

	// if conversation id is empty, don't send it
	body := fmt.Sprintf(`{
        "action": "next",
        "messages": [
            {
                "id": "%s",
                "role": "user",
                "content": {
                    "content_type": "text",
                    "parts": %s
                }
            }
        ],
        "model": "text-davinci-002-render",
		"parent_message_id": "%s"%s
    }`, uuid.NewString(), string(messages), lastMessageID, conversationIdString)

	req, err := http.NewRequest("POST", conversationURL, strings.NewReader(body))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create request: %v", err))
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to connect to openai: %v", err))
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("failed to connect to openai: %v", resp.Status))
	}
	respChan := make(chan ChatResponse)
	go func() {
		defer resp.Body.Close()
		decoder := eventsource.NewDecoder(resp.Body)
		defer close(respChan)

		for {
			event, err := decoder.Decode()
			if err != nil {
				log.Println(errors.New(fmt.Sprintf("failed to decode event: %v", err)))
				break
			}
			if event.Data() == "[DONE]" || event.Data() == "" {
				break
			}

			var res MessageResponse
			err = json.Unmarshal([]byte(event.Data()), &res)
			if err != nil {
				log.Printf("Couldn't unmarshal message response: %v", err)
				continue
			}

			if len(res.Message.Content.Parts) > 0 {
				respChan <- ChatResponse{
					Message:        res.Message.Content.Parts[0],
					ConversationID: res.ConversationId,
					LastMessageID:  res.Message.ID,
				}
			}
		}
	}()

	return respChan, nil
}
