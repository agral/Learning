package main

import (
	"bufio"
	"educationalsp/analysis"
	"educationalsp/lsp"
	"educationalsp/rpc"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	logger := getLogger("/tmp/educationalsp.txt")
	logger.Println("Educationalsp is starting")

	state := analysis.NewState()
	writer := os.Stdout

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Error: %s", err)
			continue
		}
		handleMessage(logger, writer, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Printf("Received message with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error: Could not parse this: %s", err)
		}
		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, msg)

		logger.Print("Sent the reply")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error in textDocument/didOpen: %s", err)
			return
		}
		logger.Printf("Opened: %s", request.Params.TextDocument.Uri)
		state.OpenDocument(request.Params.TextDocument.Uri, request.Params.TextDocument.Text)

	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error in textDocument/didChange: %s", err)
			return
		}
		logger.Printf("Changed: %s", request.Params.TextDocument.Uri)
		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.Uri, change.Text)
		}

	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Error in textDocument/hover: %s", err)
			return
		}

		response := state.Hover(request.ID, request.Params.TextDocument.Uri, request.Params.Position)
		writeResponse(writer, response)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Bad log file given")
	}
	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
