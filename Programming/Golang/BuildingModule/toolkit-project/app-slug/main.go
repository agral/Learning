package main

import (
	"log"
	"toolkit"
)

func main() {
	toSlug := "now is THE!!? time 123"
	var tools toolkit.Tools
	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}
	log.Println(slugified)
}
