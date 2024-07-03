package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type BaseMessage struct {
	Method string `json:"method"`
}

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, separatorFound := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !separatorFound {
		return "", nil, errors.New("Message did not contain mandatory \\r\\n\\r\\n separator")
	}

	// Parse Content-Length: <number>
	contentLengthBytes := header[len("Content-Length: "):]
	contentLengthValue, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}

	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLengthValue], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLengthValue], nil
}
