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
