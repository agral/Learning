package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		// Oh yes, typographic points.
		PageSize: gopdf.Rect{W: 595.28, H: 841.89}, // A4, portrait
	})
	pdf.AddPage()
	err := pdf.AddTTFFont("NotoSans", "../assets/ttf/NotoSans.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("NotoSans", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetGrayFill(0.0)
	pdf.Cell(nil, "Zażółć Gęślą Jaźń")
	pdf.WritePdf("hello.pdf")
}
