package rpc_test

import (
	"educationalsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s, actual: %s", expected, actual)
	}
}

func TestDecodeMessage(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, contentLength, err := rpc.DecodeMessage([]byte(incomingMessage))
	if err != nil {
		t.Fatal(err)
	}
	if contentLength != 15 {
		t.Fatalf("Wrong Content-Length: expected 15, got: %d", contentLength)
	}
	if method != "hi" {
		t.Fatalf("Wrong method: expected \"hi\", got \"%s\"", method)
	}
}
