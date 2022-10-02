package mail

import (
	"com.caiflower/commons/thirdpart/internal/common"
	"com.caiflower/commons/thirdpart/internal/entity"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime"
	"net/smtp"
	"strings"
	"time"
)

type Client struct {
	client *smtp.Client
}

func (c *Client) SendMsg(m *entity.EmailByteMessage) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = &common.ThirdError{Msg: "发送邮件失败"}
			return
		}
	}()

	var (
		client     = c.client
		head       = ""
		body       = ""
		attachment = ""
	)
	if e = client.Mail(m.From); e != nil {
		return
	}

	// 1.RCPT，告诉收件邮箱即将发送邮件
	for _, v := range m.To {
		if e := client.Rcpt(v); e != nil {
			log.Printf(v + " not exit")
			return e
		}
	}

	// 附件标签
	boundary := "GoBoundary"

	// 2.编辑邮件头
	write, e := client.Data()
	headers := make(map[string]string)
	headers["SUBJECT"] = m.Title
	headers["FROM"] = m.From
	headers["TO"] = strings.Join(m.To, ";")
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "multipart/mixed;boundary=" + boundary
	headers["DATE"] = time.Now().String()
	for k, v := range headers {
		head += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	write.Write([]byte(head))

	// 3.发送文件主体
	body = "\r\n--" + boundary + "\r\n"
	body += m.ContentType + "\r\n"
	body += "\r\n" + string(m.Content) + "\r\n"
	write.Write([]byte(body))

	// 4. 发送附件
	if m.Attachment.WithFile {
		attachment = "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Disposition:attachment\r\n"
		attachment += "Content-Type:" + m.Attachment.ContentType + ";name=\"" + mime.BEncoding.Encode("UTF-8", m.Attachment.Name) + "\"\r\n"
		write.Write([]byte(attachment))
		writeFile(write, m.Attachment.Content)
	}
	write.Write([]byte("\r\n--" + boundary + "--"))

	// 5. 关闭结束
	if e = write.Close(); e != nil {
		return
	}
	return client.Quit()
}

func writeFile(buffer io.WriteCloser, content []byte) {
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(content)))
	base64.StdEncoding.Encode(payload, content)
	buffer.Write([]byte("\r\n"))
	for index, line := 0, len(payload); index < line; index++ {
		buffer.Write([]byte{payload[index]})
		if (index+1)%76 == 0 {
			buffer.Write([]byte("\r\n"))
		}
	}
}

func (c *Client) Close() {
	c.client.Close()
}
