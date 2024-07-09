package toolkit

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"mime/multipart"
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
	{
		name: "allowed no rename", allowedTypes: []string{"image/jpeg", "image/png"}, renameFile: false, errorExpected: false,
	},
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
