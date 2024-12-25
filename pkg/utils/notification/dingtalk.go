package notification

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// DingTalkConfig 钉钉通知配置
type DingTalkConfig struct {
	WebhookURL string `json:"webhookUrl"` // webhook地址
	Secret     string `json:"secret"`     // 签名密钥
}

// DingTalkMessage 钉钉消息结构
type DingTalkMessage struct {
	MsgType  string               `json:"msgtype"`
	Markdown *DingTalkMarkdownMsg `json:"markdown,omitempty"`
	Text     *DingTalkTextMsg     `json:"text,omitempty"`
	At       *DingTalkAtInfo      `json:"at,omitempty"`
}

// DingTalkMarkdownMsg Markdown消息
type DingTalkMarkdownMsg struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// DingTalkTextMsg 文本消息
type DingTalkTextMsg struct {
	Content string `json:"content"`
}

// DingTalkAtInfo @信息
type DingTalkAtInfo struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	AtUserIds []string `json:"atUserIds,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

// DingTalkNotifier 钉钉通知实现
type DingTalkNotifier struct {
	config DingTalkConfig
	client *http.Client
}

// NewDingTalkNotifier 创建钉钉通知实例
func NewDingTalkNotifier(config DingTalkConfig) *DingTalkNotifier {
	return &DingTalkNotifier{
		config: config,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// Send 发送通知
func (n *DingTalkNotifier) Send(title, content string) error {
	message := &DingTalkMessage{
		MsgType: "markdown",
		Markdown: &DingTalkMarkdownMsg{
			Title: title,
			Text:  content,
		},
	}

	return n.SendMessage(message)
}

// SendMessage 发送消息
func (n *DingTalkNotifier) SendMessage(message *DingTalkMessage) error {
	// 生成签名
	timestamp := time.Now().UnixMilli()
	sign := n.sign(timestamp)

	// 构建请求URL
	reqURL := fmt.Sprintf("%s&timestamp=%d&sign=%s",
		n.config.WebhookURL,
		timestamp,
		url.QueryEscape(sign),
	)

	// 构建请求体
	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("marshal message failed: %v", err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", reqURL, bytes.NewReader(body))
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

// sign 生成签名
func (n *DingTalkNotifier) sign(timestamp int64) string {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, n.config.Secret)
	mac := hmac.New(sha256.New, []byte(n.config.Secret))
	mac.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
