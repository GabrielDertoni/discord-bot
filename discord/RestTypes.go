package discord

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
)

type MessageBody struct {
	Content         string             `json:"content,omitempty"`
	Embed           *Embed             `json:"embed,omitempty"`
	TTS             bool               `json:"tts,omitempty"`
	PayloadJSON     string             `json:"payload_json,omitempty"`
	AllowedMentions AllowedMentionType `json:"allowed_mentions,omitempty"`
	Files           []*File            `json:"-"`
}

type File struct {
	Name        string
	ContentType string
	Reader      io.Reader
}

type RestError struct {
	Message string `json:"message" required:"true"`
	Code    int    `json:"code" required:"true"`
}

type ModifyCurrentUserBody struct {
	Username string    `json:"username,omitempty"`
	Avatar   ImageData `json:"avatar,omitempty"`
}

type ImageData string

type ImageFileExt string

const (
	ImageFileExtPNG ImageFileExt = "image/png"
	ImageFileExtJPG              = "image/jpg"
	ImageFileExtGIF              = "image/gif"
)

func FileToImageData(filepath string) (data ImageData, err error) {
	rawData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	var ext ImageFileExt = map[string]ImageFileExt{
		"png": ImageFileExtPNG,
		"jpg": ImageFileExtJPG,
		"gif": ImageFileExtGIF,
	}[filepath[len(filepath)-3:]]

	data = BytesToImageData(rawData, ext)
	return
}

func BytesToImageData(rawData []byte, fileExt ImageFileExt) ImageData {
	base64Raw := make([]byte, base64.StdEncoding.EncodedLen(len(rawData)))
	base64.StdEncoding.Encode(base64Raw, rawData)
	return ImageData(fmt.Sprintf("data:%s;base64,%s", fileExt, base64Raw))
}

func RequiredUnmarshal(data []byte, v interface{}) (err error) {
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}

	fields := reflect.ValueOf(v).Elem()
	for i := 0; i < fields.NumField(); i++ {
		requiredTag := fields.Type().Field(i).Tag.Get("required")
		if requiredTag == "true" && fields.Field(i).IsZero() {
			return errors.New("required field is missing")
		}
	}
	return nil
}
