package tinodeclient

import (
	"encoding/base64"
	"fmt"
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type TinodeClient struct {
	wsConn    *websocket.Conn
	serverURL string
	token     string
	mu        sync.Mutex
}

func NewTinodeClient(serverURL string) *TinodeClient {
	return &TinodeClient{
		serverURL: serverURL,
	}
}

func (c *TinodeClient) Connect() error {
	u := url.URL{Scheme: "ws", Host: c.serverURL, Path: "/v0/channels"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect to Tinode server: %w", err)
	}

	c.wsConn = conn
	return nil
}

func (c *TinodeClient) ReadMessage() (map[string]interface{}, error) {
	var msg map[string]interface{}
	if err := c.wsConn.ReadJSON(&msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (c *TinodeClient) Close() error {
	return c.wsConn.Close()
}

// Implement the RoomsClient interface methods

// CreateTopic creates a new topic (chat room) in Tinode
func (c *TinodeClient) CreateTopic(topicName string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	msgID := "sub-" + topicName

	subMsg := map[string]interface{}{
		"id":    msgID,
		"topic": "new",
		"cmd":   "sub",
		"params": map[string]interface{}{
			"desc": map[string]interface{}{
				"public": map[string]interface{}{
					"fn": topicName,
				},
			},
		},
	}

	if err := c.wsConn.WriteJSON(subMsg); err != nil {
		return "", fmt.Errorf("failed to send create topic message: %w", err)
	}

	// Read the response
	var resp map[string]interface{}
	if err := c.wsConn.ReadJSON(&resp); err != nil {
		return "", fmt.Errorf("failed to read create topic response: %w", err)
	}

	// Check for success
	if ctrl, ok := resp["ctrl"].(map[string]interface{}); ok {
		if code, ok := ctrl["code"].(float64); ok && code == 200 {
			if params, ok := ctrl["params"].(map[string]interface{}); ok {
				if topic, ok := params["topic"].(string); ok {
					return topic, nil
				}
			}
		} else {
			return "", fmt.Errorf("create topic failed: %v", ctrl)
		}
	}

	return "", fmt.Errorf("unexpected response: %v", resp)
}

// AddUserToTopic adds a user to a topic (chat room)
func (c *TinodeClient) AddUserToTopic(topicName, userEmail string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	msgID := "set-" + topicName + "-" + userEmail

	setMsg := map[string]interface{}{
		"id":    msgID,
		"topic": topicName,
		"cmd":   "set",
		"params": map[string]interface{}{
			"sub": map[string]interface{}{
				"user": userEmail,
				"mode": "JRWPA", // Adjust permissions as needed
			},
		},
	}

	if err := c.wsConn.WriteJSON(setMsg); err != nil {
		return fmt.Errorf("failed to send add user to topic message: %w", err)
	}

	// Read the response
	var resp map[string]interface{}
	if err := c.wsConn.ReadJSON(&resp); err != nil {
		return fmt.Errorf("failed to read add user to topic response: %w", err)
	}

	// Check for success
	if ctrl, ok := resp["ctrl"].(map[string]interface{}); ok {
		if code, ok := ctrl["code"].(float64); ok && code == 200 {
			return nil
		} else {
			return fmt.Errorf("add user to topic failed: %v", ctrl)
		}
	}

	return fmt.Errorf("unexpected response: %v", resp)
}

// RemoveUserFromTopic removes a user from a topic (chat room)
func (c *TinodeClient) RemoveUserFromTopic(topicName, userEmail string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	msgID := "del-" + topicName + "-" + userEmail

	delMsg := map[string]interface{}{
		"id":    msgID,
		"topic": topicName,
		"cmd":   "del",
		"params": map[string]interface{}{
			"what": "sub",
			"user": userEmail,
		},
	}

	if err := c.wsConn.WriteJSON(delMsg); err != nil {
		return fmt.Errorf("failed to send remove user from topic message: %w", err)
	}

	// Read the response
	var resp map[string]interface{}
	if err := c.wsConn.ReadJSON(&resp); err != nil {
		return fmt.Errorf("failed to read remove user from topic response: %w", err)
	}

	// Check for success
	if ctrl, ok := resp["ctrl"].(map[string]interface{}); ok {
		if code, ok := ctrl["code"].(float64); ok && code == 200 {
			return nil
		} else {
			return fmt.Errorf("remove user from topic failed: %v", ctrl)
		}
	}

	return fmt.Errorf("unexpected response: %v", resp)
}

func (c *TinodeClient) SendMessage(topicName, messageContent string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	msgID := "pub-" + topicName

	pubMsg := map[string]interface{}{
		"id":      msgID,
		"topic":   topicName,
		"cmd":     "pub",
		"content": messageContent,
	}

	if err := c.wsConn.WriteJSON(pubMsg); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	// Read the response
	var resp map[string]interface{}
	if err := c.wsConn.ReadJSON(&resp); err != nil {
		return fmt.Errorf("failed to read send message response: %w", err)
	}

	// Check for success
	if ctrl, ok := resp["ctrl"].(map[string]interface{}); ok {
		if code, ok := ctrl["code"].(float64); ok && code == 200 {
			return nil
		} else {
			return fmt.Errorf("send message failed: %v", ctrl)
		}
	}

	return fmt.Errorf("unexpected response: %v", resp)
}

func (c *TinodeClient) Subscribe(topicName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	msgID := "sub-" + topicName

	subMsg := map[string]interface{}{
		"id":    msgID,
		"topic": topicName,
		"cmd":   "sub",
	}

	if err := c.wsConn.WriteJSON(subMsg); err != nil {
		return fmt.Errorf("failed to send subscribe message: %w", err)
	}

	// Read the response
	var resp map[string]interface{}
	if err := c.wsConn.ReadJSON(&resp); err != nil {
		return fmt.Errorf("failed to read subscribe response: %w", err)
	}

	// Check for success
	if ctrl, ok := resp["ctrl"].(map[string]interface{}); ok {
		if code, ok := ctrl["code"].(float64); ok && code == 200 {
			return nil
		} else {
			return fmt.Errorf("subscribe failed: %v", ctrl)
		}
	}

	return fmt.Errorf("unexpected response: %v", resp)
}

func (c *TinodeClient) GetMessages(topicName string) ([]entity.Message, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	msgID := "get-" + topicName

	getMsg := map[string]interface{}{
		"id":    msgID,
		"topic": topicName,
		"cmd":   "get",
		"params": map[string]interface{}{
			"what": "data",
		},
	}

	if err := c.wsConn.WriteJSON(getMsg); err != nil {
		return nil, fmt.Errorf("failed to send get messages request: %w", err)
	}

	// Read the response
	var messages []entity.Message
	for {
		var resp map[string]interface{}
		if err := c.wsConn.ReadJSON(&resp); err != nil {
			return nil, fmt.Errorf("failed to read get messages response: %w", err)
		}

		if data, ok := resp["data"].([]interface{}); ok {
			// Parse the messages
			for _, item := range data {
				if msgData, ok := item.(map[string]interface{}); ok {
					ts, _ := time.Parse(time.RFC3339, msgData["ts"].(string))
					msg := entity.Message{
						ChatRoomID: topicName,
						Author:     msgData["from"].(string),
						Content:    msgData["content"].(string),
						Timestamp:  ts,
					}
					messages = append(messages, msg)
				}
			}
			break
		} else if ctrl, ok := resp["ctrl"].(map[string]interface{}); ok {
			if code, ok := ctrl["code"].(float64); ok && code != 200 {
				return nil, fmt.Errorf("get messages failed: %v", ctrl)
			}
		}
	}

	return messages, nil
}

func (c *TinodeClient) Register(email, password, fullName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	msgID := "acc-" + email

	accMsg := map[string]interface{}{
		"id":    msgID,
		"cmd":   "acc",
		"topic": "new",
		"params": map[string]interface{}{
			"user": map[string]interface{}{
				"email": email,
				"fn":    fullName,
			},
			"cred": []interface{}{
				map[string]interface{}{
					"meth": "email",
					"val":  email,
				},
			},
			"secret": password,
		},
	}

	if err := c.wsConn.WriteJSON(accMsg); err != nil {
		return fmt.Errorf("failed to send register message: %w", err)
	}

	// Read the response
	var resp map[string]interface{}
	if err := c.wsConn.ReadJSON(&resp); err != nil {
		return fmt.Errorf("failed to read register response: %w", err)
	}

	// Check for success
	if ctrl, ok := resp["ctrl"].(map[string]interface{}); ok {
		if code, ok := ctrl["code"].(float64); ok && code == 200 {
			return nil
		} else {
			return fmt.Errorf("register failed: %v", ctrl)
		}
	}

	return fmt.Errorf("unexpected response: %v", resp)
}

func (c *TinodeClient) Login(email, password string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	msgID := "login-" + email

	loginMsg := map[string]interface{}{
		"id":  msgID,
		"cmd": "login",
		"params": map[string]interface{}{
			"scheme": "basic",
			"secret": base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", email, password))),
		},
	}

	if err := c.wsConn.WriteJSON(loginMsg); err != nil {
		return fmt.Errorf("failed to send login message: %w", err)
	}

	// Read the response
	var resp map[string]interface{}
	if err := c.wsConn.ReadJSON(&resp); err != nil {
		return fmt.Errorf("failed to read login response: %w", err)
	}

	// Check for success and store the authentication token
	if ctrl, ok := resp["ctrl"].(map[string]interface{}); ok {
		if code, ok := ctrl["code"].(float64); ok && code == 200 {
			if params, ok := ctrl["params"].(map[string]interface{}); ok {
				if token, ok := params["token"].(string); ok {
					c.token = token
					return nil
				}
			}
		} else {
			return fmt.Errorf("login failed: %v", ctrl)
		}
	}

	return fmt.Errorf("unexpected response: %v", resp)
}
