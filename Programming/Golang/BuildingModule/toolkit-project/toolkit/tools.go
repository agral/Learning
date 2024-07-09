package toolkit

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

type Tools struct {
	MaxFileSize      int64
	AllowedFileTypes []string
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
