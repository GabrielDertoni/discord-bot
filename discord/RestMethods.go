package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"strings"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func CreateMessage(message *MessageBody, channelID string, token string) (sent *Message, err error) {
	apiURL := BaseURL.Concat(URL(fmt.Sprintf("channels/%s/messages", channelID)))
	var respBody []byte
	if len(message.Files) == 0 {
		respBody, err = simplePOST(apiURL, message, map[string]string{
			"Authorization": "Bot " + token,
			"Content-Type":  "application/json",
		})
	} else {
		postBody := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(postBody)

		payload, err := json.Marshal(message)
		header := make(textproto.MIMEHeader)
		header.Set("Content-Disposition", `form-data; name="payload_json"`)
		header.Set("Content-Type", "application/json")

		part, err := bodyWriter.CreatePart(header)
		if err != nil {
			return nil, err
		}

		if _, err = part.Write(payload); err != nil {
			return nil, err
		}

		for i, file := range message.Files {
			header := make(textproto.MIMEHeader)
			header.Set("Content-Disposition",
				fmt.Sprintf(`form-data; name="file%d"; filename="%s"`, i, quoteEscaper.Replace(file.Name)))
			contentType := file.ContentType
			if contentType == "" {
				contentType = "application/octet-stream"
			}
			header.Set("Content-Type", contentType)
			part, err = bodyWriter.CreatePart(header)
			if _, err = io.Copy(part, file.Reader); err != nil {
				return nil, err
			}
		}
		err = bodyWriter.Close()
		if err != nil {
			return nil, err
		}
		respBody, err = rawPOST(apiURL, postBody.Bytes(), map[string]string{
			"Authorization": "Bot " + token,
			"Content-Type":  bodyWriter.FormDataContentType(),
		})
	}
	if err != nil {
		return
	}
	errMessage := &RestError{}
	if RequiredUnmarshal(respBody, errMessage) == nil {
		err = errors.New(errMessage.Message)
	}
	sent = &Message{}
	err = json.Unmarshal(respBody, sent)
	return
}

func ModifyCurrentUser(modification *ModifyCurrentUserBody, token string) (err error) {
	apiURL := BaseURL.Concat(URL("users/@me"))
	var respBody []byte
	respBody, err = simplePATCH(apiURL, modification, map[string]string{
		"Authorization": "Bot " + token,
		"Content-Type":  "application/json",
	})
	errMessage := &RestError{}
	if RequiredUnmarshal(respBody, errMessage) == nil {
		err = errors.New(errMessage.Message)
	}
	return
}
