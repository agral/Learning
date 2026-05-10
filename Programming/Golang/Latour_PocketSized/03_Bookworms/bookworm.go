package main

import "os"

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return nil, nil
}
