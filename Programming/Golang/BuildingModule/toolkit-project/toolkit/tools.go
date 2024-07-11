package toolkit

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

type Tools struct {
	MaxFileSize        int64
	AllowedFileTypes   []string
	MaxJsonSize        int
	AllowUnknownFields bool
}

func (t *Tools) RandomString(length int) string {
	s, runeSource := make([]rune, length), []rune(randomStringSource)
	for i := range s {
		prime, _ := rand.Prime(rand.Reader, len(runeSource))
		x, mod := prime.Uint64(), uint64(len(runeSource))
		s[i] = runeSource[x%mod]
	}

	return string(s)
}

// Stores information about an uploaded file.
type UploadedFile struct {
	NewFileName      string
	OriginalFileName string
	FileSize         int64
}

func (t *Tools) UploadOneFile(r *http.Request, uploadDir string, rename ...bool) (*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}

	files, err := t.UploadFiles(r, uploadDir, renameFile)
	if err != nil {
		return nil, err
	}

	return files[0], nil
}

func (t *Tools) UploadFiles(r *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}

	var uploadedFiles []*UploadedFile
	if t.MaxFileSize == 0 {
		t.MaxFileSize = 1024 * 1024 * 1024 // 1 GiB
	}

	if err := t.CreateDirIfNotExist(uploadDir); err != nil {
		return nil, err
	}

	err := r.ParseMultipartForm(t.MaxFileSize)
	if err != nil {
		return nil, errors.New("The file is too big")
	}

	for _, fHeaders := range r.MultipartForm.File {
		for _, header := range fHeaders {
			uploadedFiles, err = func(uploadedFiles []*UploadedFile) ([]*UploadedFile, error) {
				var uploadedFile UploadedFile
				infile, err := header.Open()
				if err != nil {
					return nil, err
				}
				defer infile.Close()

				buff := make([]byte, 512)
				if _, err := infile.Read(buff); err != nil {
					return nil, err
				}

				isAllowed := false
				fileType := http.DetectContentType(buff)

				if len(t.AllowedFileTypes) > 0 {
					for _, x := range t.AllowedFileTypes {
						if strings.EqualFold(fileType, x) {
							isAllowed = true
						}
					}
				} else {
					// When allowlist is empty, all file types are considered allowed.
					isAllowed = true
				}

				if !isAllowed {
					return nil, errors.New("The uploaded file type is not permitted")
				}

				_, err = infile.Seek(0, 0)
				if err != nil {
					return nil, err
				}

				uploadedFile.OriginalFileName = header.Filename
				if renameFile {
					uploadedFile.NewFileName = fmt.Sprintf("%s%s", t.RandomString(25),
						filepath.Ext(header.Filename))
				} else {
					uploadedFile.NewFileName = header.Filename
				}

				var outfile *os.File
				defer outfile.Close()
				if outfile, err = os.Create(filepath.Join(uploadDir, uploadedFile.NewFileName)); err != nil {
					return nil, err
				} else {
					fileSize, err := io.Copy(outfile, infile)
					if err != nil {
						return nil, err
					}
					uploadedFile.FileSize = fileSize
				}

				uploadedFiles = append(uploadedFiles, &uploadedFile)
				return uploadedFiles, nil
			}(uploadedFiles)

			if err != nil {
				return uploadedFiles, err
			}
		}
	}

	return uploadedFiles, nil
}

// Creates a directory with all necessary parent directories, if it does not exist.
// If the directory already exists, succeeds.
func (t *Tools) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

// Creates a web-safe slug from a given string
func (t *Tools) Slugify(s string) (string, error) {
	if s == "" {
		return "", errors.New("Empty string not permitted")
	}

	var re = regexp.MustCompile(`[^a-z\d]+`)
	slug := strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")
	if len(slug) == 0 {
		return "", errors.New("After removing disallowed characters, slug is empty")
	}

	return slug, nil
}

// Downloads a file.
// Tries to force the browser to avoid displaying it in the browser window by setting content-disposition.
func (t *Tools) DownloadStaticFile(w http.ResponseWriter, r *http.Request, p, file, displayName string) {
	filePath := path.Join(p, file)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", displayName))
	http.ServeFile(w, r, filePath)
}

type JsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Tries to read the body of a request and convert it from JSON to a go data variable
func (t *Tools) ReadJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024 // 1 MiB
	if t.MaxJsonSize != 0 {
		maxBytes = t.MaxJsonSize
	}

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	decoder := json.NewDecoder(r.Body)
	if !t.AllowUnknownFields {
		decoder.DisallowUnknownFields()
	}

	if err := decoder.Decode(data); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("Body contains badly formed JSON at character %d", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("Body contains badly formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("Body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("Body contains incorrect JSON type at character %d", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("Body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("Body contains unknown key %s", fieldName)

		case err.Error() == "http: request body too large":
			return fmt.Errorf("Body must not be larger than %d bytes", maxBytes)

		case errors.As(err, &invalidUnmarshalError):
			return fmt.Errorf("Error unmarshaling JSON: %s", err.Error())

		default:
			return err
		}
	}

	// Check if more than one JSON file is present
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return errors.New("Body must contain only one JSON value")
	}

	return nil
}

// Takes a response status code and arbitrary data, marshals it into a JSON and writes it to the specified client.
func (t *Tools) WriteJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}
