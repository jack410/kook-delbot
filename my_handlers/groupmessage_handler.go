package my_handlers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gookit/event"
	"github.com/kaiheila/golang-bot/api/base"
	event2 "github.com/kaiheila/golang-bot/api/base/event"
	"github.com/kaiheila/golang-bot/api/helper"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"strings"
	"time"
)

type MessageDelHandler struct {
	Token   string
	BaseUrl string
}

func (gteh *MessageDelHandler) Handle(e event.Event) error {
	log.WithField("event", fmt.Sprintf("%+v", e.Data())).Info("收到频道内的文字消息.")
	err := func() error {
		if _, ok := e.Data()[base.EventDataFrameKey]; !ok {
			return errors.New("data has no frame field")
		}
		frame := e.Data()[base.EventDataFrameKey].(*event2.FrameMap)
		data, err := sonic.Marshal(frame.Data)
		if err != nil {
			return err
		}
		msgEvent := &event2.MessageKMarkdownEvent{}
		err = sonic.Unmarshal(data, msgEvent)
		log.Infof("Received json event:%+v", msgEvent)
		if err != nil {
			return err
		}
		client := helper.NewApiHelper("/v3/message/delete", gteh.Token, gteh.BaseUrl, "", "")
		if msgEvent.Author.Bot {
			log.Info("bot message")
			return nil
		}

		// 获取当前日期
		currentDate := time.Now().Format("200601")

		// 构建文件名
		fileName := fmt.Sprintf("/root/chat/80chatlog-%s.csv", currentDate)

		// 使用 os.OpenFile 创建或打开文件
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		encoder := simplifiedchinese.GBK.NewEncoder()

		//构建文件名
		volChatLogPath := fmt.Sprintf("/root/chat/80volchatlog-%s.txt", currentDate)
		volChatLogFile, err := os.OpenFile(volChatLogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer volChatLogFile.Close()

		if hasPrefix(msgEvent.Author.Nickname) {
			currentTime := time.Now()
			formattedTime := currentTime.Format("2006-01-02 15:04:05")
			line := strings.Join([]string{formattedTime, msgEvent.AuthorId, msgEvent.Author.Nickname, msgEvent.Content}, ", ")
			_, err := volChatLogFile.WriteString(line + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}

		if msgEvent.Author.Nickname != "" {
			currentTime := time.Now()
			formattedTime := currentTime.Format("2006-01-02 15:04:05")
			w := csv.NewWriter(file)
			var record []string
			record = append(record, formattedTime)
			authorId, _ := encoder.String(msgEvent.AuthorId)
			record = append(record, authorId)
			nickname, _ := encoder.String(msgEvent.Author.Nickname)
			record = append(record, nickname)
			channelName, _ := encoder.String(msgEvent.ChannelName)
			record = append(record, channelName)
			content, _ := encoder.String(msgEvent.Content)
			record = append(record, content)
			err = w.Write(record)
			if err != nil {
				return err
			}
			w.Flush()
		}

		if ContainsPhrase(msgEvent.KMarkdown.RawContent, "https://wotlk.everlook-wow.net/recruit/") || ContainsPhrase(msgEvent.KMarkdown.RawContent, "煞笔") || ContainsPhrase(msgEvent.KMarkdown.RawContent, "你妈") {
			DeleteGroupMessage(msgEvent.MsgId, client)
			currentTime := time.Now()
			formattedTime := currentTime.Format("2006-01-02 15:04:05")
			log.Infof(formattedTime, "删除了消息：", msgEvent.KMarkdown.RawContent)
		}

		return nil
	}()
	if err != nil {
		log.WithError(err).Error("GroupTextEventHandler err")
	}

	return nil
}

func ContainsPhrase(text, phrase string) bool {
	// 将文字和词组都转换为小写，然后检查是否包含词组
	return strings.Contains(strings.ToLower(text), strings.ToLower(phrase))
}

func DeleteGroupMessage(msgId string, client *helper.ApiHelper) {
	echoData := map[string]interface{}{
		"msg_id": msgId,
	}
	//将map转化成[]byte
	deleteMsgBody, err := sonic.Marshal(echoData)
	if err != nil {
		log.Infof("将map转化成[]byte出错", err)
	}

	resp, err := client.SetBody(deleteMsgBody).Post()
	log.Info("sent post:%s", client.String())
	if err != nil {
		log.Infof("删除消息出错", err)
	}
	log.Infof("resp:%s", string(resp))
}

func SendGroupTextMessage(channel_id, content string, client *helper.ApiHelper) {
	echoData := map[string]string{
		"channel_id": channel_id,
		"content":    content,
	}
	echoDataByte, err := sonic.Marshal(echoData)
	if err != nil {
		log.Infof("echoData 序列化出错", err)
	}
	resp, err := client.SetBody(echoDataByte).Post()
	log.Info("sent post:%s", client.String())
	if err != nil {
		log.Infof("发送echoDataByte出错", err)
	}
	log.Infof("resp:%s", string(resp))

}

func SendGroupCardessage(cardMessageContent, channel_id string, client *helper.ApiHelper) {
	echoData := map[string]interface{}{
		"type":       10,
		"channel_id": channel_id,
		"content":    cardMessageContent,
	}
	//将map转化成[]byte
	echoDataByte, err := sonic.Marshal(echoData)
	if err != nil {
		log.Infof("将map转化成[]byte出错", err)
	}

	resp, err := client.SetBody(echoDataByte).Post()
	log.Info("sent post:%s", client.String())
	if err != nil {
		log.Infof("发送echoDataByte出错", err)
	}
	log.Infof("resp:%s", string(resp))
}

func hasPrefix(s string) bool {
	return strings.HasPrefix(s, "CM |") || strings.HasPrefix(s, "VOL |")
}
