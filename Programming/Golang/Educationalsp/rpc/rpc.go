package rpc

import (
	"encoding/json"
	"errors"
	"fmt"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) (string, int, error) {
	return "", 0, errors.New("Not implemented yet")
}
