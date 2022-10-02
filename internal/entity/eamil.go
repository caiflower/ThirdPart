package entity

type EmailStringMessage struct {
	From       string     `json:"from" binding:"required"`
	To         []string   `json:"to" binding:"required"`
	Title      string     `json:"title" binding:"required"`
	Content    string     `json:"content" binding:"required"`
	Attachment Attachment `json:"attachment"`
}

type EmailByteMessage struct {
	From        string     `json:"from" binding:"required"`
	To          []string   `json:"to" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	ContentType string     `json:"content_type" binding:"required" default:"Content-Type:text/plain;charset=utf-8"`
	Content     []byte     `json:"content" binding:"required"`
	Attachment  Attachment `json:"attachment"`
}

type Attachment struct {
	WithFile    bool   `json:"with_file"`
	Name        string `json:"name"`
	ContentType string `json:"content_type"`
	Content     []byte `json:"content"`
}
