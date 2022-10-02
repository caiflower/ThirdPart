package test

import (
	"bytes"
	"com.caiflower/commons/thirdpart/internal/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	FORM = "from@163.com"
	TO   = "to@163.com"
)

func TestHttp() {
	msg := new(entity.EmailByteMessage)
	msg.From = FORM
	msg.To = []string{TO}
	msg.ContentType = "Content-Type:text/plain;charset=utf-8"
	msg.Title = "测试邮件"
	msg.Content = []byte("你好")
	attachment := entity.Attachment{WithFile: true, Name: "测试附件.xlsx", ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}
	content, err := ioutil.ReadFile("/Users/lijinlong/Desktop/李金龙试用期转正工作总结.xlsx")
	if err != nil {
		log.Fatal("read file fail")
	}
	attachment.Content = content
	msg.Attachment = attachment

	body, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("json convert fail")
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", "http://localhost:8080/email/byte", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
