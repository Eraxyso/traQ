package router

import (
	"net/http"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/traPtitech/traQ/model"
)

type MessageForResponce struct {
	MessageId       string
	UserId          string
	ParentChannelId string
	Content         string
	Datetime        string
	//Pin bool
	//StampList /*stampのオブジェクト*/
}

type postMessage struct {
	Text string `json:"text"`
}

func GetMessageByIdHandler(c echo.Context) error {
	if _, err := getUserId(c); err != nil {
		return err
	}

	id := c.Param("messageId") // TODO: idの検証
	raw, err := model.GetMessage(id)
	if err != nil {
		errorMessageResponse(c, http.StatusNotFound, "Message is not found")
		return fmt.Errorf("model.Getmessage returns an error : %v", err)
	}
	res := formatMessgae(raw)
	return c.JSON(http.StatusOK, res)
}

func GetMessagesByChannelIdHandler(c echo.Context) error {
	return nil
}

func PostMessageHandler(c echo.Context) error {
	return nil
}

func PutMessageByIdHandler(c echo.Context) error {
	return nil
}

func DeleteMessageByIdHandler(c echo.Context) error {
	
	return nil
}

// 実質user認証みたいなことに使っている
func getUserId(c echo.Context) (string, error) {
	sess, err := session.Get("sessions", c)
	if err != nil {
		errorMessageResponse(c, http.StatusInternalServerError, "Failed to get a session")
		return "", fmt.Errorf("Failed to get a session: %v", err)
	}
	
	var userId string
	if sess.Values["userId"] != nil {
		userId = sess.Values["userId"].(string)
	} else {
		errorMessageResponse(c, http.StatusForbidden, "Your userId doesn't exist")
		return "", fmt.Errorf("This session doesn't have a userId")
	}
	return userId, nil
}

func formatMessgae(raw *model.Messages) *MessageForResponce {
	res := new(MessageForResponce)
	res.MessageId = raw.Id
	res.UserId = raw.UserId
	res.ParentChannelId = raw.ChannelId
	res.Content = raw.Text
	res.Datetime = raw.CreatedAt
	//TODO: res.pin,res.stampListの取得
	return res
}