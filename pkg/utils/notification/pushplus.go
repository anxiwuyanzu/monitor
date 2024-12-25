package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const PushPlusAPI = "http://www.pushplus.plus/send"

// PushPlusConfig PushPlus通知配置
type PushPlusConfig struct {
	Token string `json:"token"` // 用户令牌
}

// PushPlusMessage PushPlus消息结构
type PushPlusMessage struct {
	Token    string `json:"token"`              // 用户令牌
	Title    string `json:"title"`              // 消息标题
	Content  string `json:"content"`            // 具体消息内容
	Template string `json:"template,omitempty"` // 发送模板，可选项：html/json/markdown/txt，默认为html
	Topic    string `json:"topic,omitempty"`    // 群组编码，不填仅发送给自己
	Channel  string `json:"channel,omitempty"`  // 发送渠道，可选项：wechat/webhook/mail/sms，默认为wechat
}

// PushPlusNotifier PushPlus通知实现
type PushPlusNotifier struct {
	config PushPlusConfig
	client *http.Client
}

// NewPushPlusNotifier 创建PushPlus通知实例
func NewPushPlusNotifier(config PushPlusConfig) *PushPlusNotifier {
	return &PushPlusNotifier{
		config: config,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// Send 发送通知
func (n *PushPlusNotifier) Send(title, content string) error {
	message := &PushPlusMessage{
		Token:    n.config.Token,
		Title:    title,
		Content:  content,
		Template: "markdown",
	}

	return n.SendMessage(message)
}

// SendMessage 发送消息
func (n *PushPlusNotifier) SendMessage(message *PushPlusMessage) error {
	// 构建请求体
	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("marshal message failed: %v", err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", PushPlusAPI, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create request failed: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := n.client.Do(req)
	if err != nil {
		return fmt.Errorf("send request failed: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	return nil
}
