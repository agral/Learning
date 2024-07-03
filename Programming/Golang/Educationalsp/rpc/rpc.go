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

// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, separatorFound := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !separatorFound {
		return 0, nil, nil // note: no error - this is fine, the LSP is just waiting for more data.
	}

	// Parse Content-Length: <number>
	contentLengthBytes := header[len("Content-Length: "):]
	contentLengthValue, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLengthValue {
		// The declared number of bytes have not yet arrived from source. That's okay, no error.
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLengthValue

	return totalLength, data[:totalLength], nil
}
