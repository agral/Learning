package lsp

type TextDocumentItem struct {
	Uri        string `json:"uri"`
	LanguageId string `json:"languageId"`

	// The version number of this TextDocumentItem.
	// It will be increased after every change, including undo&redo.
	Version int    `json:"version"`
	Text    string `json:"text"`
}

type TextDocumentIdentifier struct {
	Uri string `json:"uri"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}
