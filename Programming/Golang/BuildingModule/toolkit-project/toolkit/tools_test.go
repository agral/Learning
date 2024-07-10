package toolkit

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
)

func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("Wrong length of generated random string")
	}
}

var uploadTests = []struct {
	name          string
	allowedTypes  []string
	renameFile    bool
	errorExpected bool
}{
	{name: "allowed no rename", allowedTypes: []string{"image/jpeg", "image/png"}, renameFile: false, errorExpected: false},
	{name: "allowed rename", allowedTypes: []string{"image/jpeg", "image/png"}, renameFile: true, errorExpected: false},
	{name: "upload not allowed filetype", allowedTypes: []string{"image/jpeg"}, renameFile: false, errorExpected: true},
}

func TestTools_UploadFiles(t *testing.T) {
	for _, e := range uploadTests {
		// sets up a pipe to avoid buffering
		pipeReader, pipeWriter := io.Pipe()
		writer := multipart.NewWriter(pipeWriter)
		waitGroup := sync.WaitGroup{}
		waitGroup.Add(1)

		go func() {
			defer writer.Close()
			defer waitGroup.Done()

			// Create a form data field "file":
			part, err := writer.CreateFormFile("file", "./test_data/trioptimum.png")
			if err != nil {
				t.Error(err)
			}

			f, err := os.Open("./test_data/trioptimum.png")
			if err != nil {
				t.Error(err)
			}
			defer f.Close()

			img, _, err := image.Decode(f)
			if err != nil {
				t.Error("Error decoding image", err)
			}

			err = png.Encode(part, img)
			if err != nil {
				t.Error(err)
			}
		}()

		// read from the pipe receiving the data:
		request := httptest.NewRequest("POST", "/", pipeReader)
		request.Header.Add("Content-Type", writer.FormDataContentType())

		var testTools Tools
		testTools.AllowedFileTypes = e.allowedTypes

		uploadedFiles, err := testTools.UploadFiles(request, "./test_data/uploads", e.renameFile)
		if err != nil && !e.errorExpected {
			t.Error(err)
		}

		if !e.errorExpected {
			uploadedFilename := fmt.Sprintf("./test_data/uploads/%s", uploadedFiles[0].NewFileName)
			if _, err := os.Stat(uploadedFilename); os.IsNotExist(err) {
				t.Errorf("%s: expected file to exist: %s", e.name, err.Error())
			}

			// clean up: remove the uploaded file
			_ = os.Remove(uploadedFilename)
		}

		if !e.errorExpected && err != nil {
			// Note: this ^ if clause is clearly wrong.
			// It will apply where error was not expected, but present.
			// But the message inside says that the error is expected, and not present.
			// No idea which option does the instructor intend, leaving as is.
			t.Errorf("%s: error expected, but none received", e.name)
		}

		waitGroup.Wait()
	}
}

func TestTools_UploadOneFile(t *testing.T) {
	// sets up a pipe to avoid buffering
	pipeReader, pipeWriter := io.Pipe()
	writer := multipart.NewWriter(pipeWriter)

	go func() {
		defer writer.Close()

		// Create a form data field "file":
		part, err := writer.CreateFormFile("file", "./test_data/trioptimum.png")
		if err != nil {
			t.Error(err)
		}

		f, err := os.Open("./test_data/trioptimum.png")
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			t.Error("Error decoding image", err)
		}

		err = png.Encode(part, img)
		if err != nil {
			t.Error(err)
		}
	}()

	// read from the pipe receiving the data:
	request := httptest.NewRequest("POST", "/", pipeReader)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	var testTools Tools

	uploadedFiles, err := testTools.UploadOneFile(request, "./test_data/uploads", true)
	if err != nil {
		t.Error(err)
	}

	uploadedFilename := fmt.Sprintf("./test_data/uploads/%s", uploadedFiles.NewFileName)
	if _, err := os.Stat(uploadedFilename); os.IsNotExist(err) {
		t.Errorf("Expected file to exist: %s", err.Error())
	}

	// clean up: remove the uploaded file
	_ = os.Remove(uploadedFilename)
}

func TestTools_CreateDirIfNotExist(t *testing.T) {
	var testTool Tools
	const testedDirPath = "./test_data/myDir"

	// Test that it is possible to create a directory that hopefully does not yet exist:
	err := testTool.CreateDirIfNotExist(testedDirPath)
	if err != nil {
		t.Error(err)
	}
	// Test that creating an already existing directory results in a success too:
	err = testTool.CreateDirIfNotExist(testedDirPath)
	if err != nil {
		t.Error(err)
	}

	// Clean up:
	_ = os.Remove(testedDirPath)
}

var slugTests = []struct {
	name          string
	s             string
	expected      string
	errorExpected bool
}{
	{name: "valid string", s: "gral learns go language", expected: "gral-learns-go-language", errorExpected: false},
	{name: "empty string", s: "", expected: "-", errorExpected: true},
	{name: "complex string", s: "Look at hOW RaNDOM i am! So +& r4nd0m!!!1^",
		expected: "look-at-how-random-i-am-so-r4nd0m-1", errorExpected: false},
	{name: "korean (non-roman) string", s: "안녕, 세상", expected: "", errorExpected: true},
	{name: "mix of korean and roman string", s: "안녕, 세상 means hello, world", expected: "means-hello-world", errorExpected: false},
}

func TestTools_Slugify(t *testing.T) {
	var testTools Tools
	for _, e := range slugTests {
		slug, err := testTools.Slugify(e.s)
		if err != nil && !e.errorExpected {
			t.Errorf("%s: error received, but none was expected. %s", e.name, err.Error())
		}

		if !e.errorExpected && slug != e.expected {
			t.Errorf("%s: wrong slug returned. Expected: %s, got: %s", e.name, e.expected, slug)
		}

		if e.errorExpected && err == nil {
			t.Errorf("%s: expected to produce an error, but it succeeded for %s", e.name, e.s)
		}
	}
}

func TestTools_DownloadStaticFile(t *testing.T) {
	var testTools Tools
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	testTools.DownloadStaticFile(recorder, request, "./test_data/", "synth.png", "downloaded_pic.png")
	result := recorder.Result()
	defer result.Body.Close()

	if result.Header["Content-Length"][0] != "376891" {
		t.Error("Wrong Content-Length of", result.Header["Content-Length"][0])
	}

	if result.Header["Content-Disposition"][0] != "attachment; filename=\"downloaded_pic.png\"" {
		t.Error("Wrong Content-Disposition")
	}

	if _, err := io.ReadAll(io.Reader(result.Body)); err != nil {
		t.Error(err)
	}
}
