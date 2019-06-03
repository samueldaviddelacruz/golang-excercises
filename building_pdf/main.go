package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v ,height=%v", w, h)
	pdf.AddPage()
	// Basic Text
	pdf.MoveTo(0, 0)
	pdf.SetFont("Arial", "B", 38)
	_, lineHeight := pdf.GetFontSize()
	pdf.SetTextColor(255, 0, 0)
	pdf.Text(0, lineHeight, "Hello, World")
	pdf.MoveTo(0, lineHeight*2.0)

	pdf.SetFont("times", "", 18)
	pdf.SetTextColor(100, 100, 100)
	_, lineHeight = pdf.GetFontSize()
	pdf.MultiCell(0, lineHeight*1.5, "here is a lot of some text\nalot of text text text\n", gofpdf.BorderNone, gofpdf.AlignRight, false)

	//Basic shapes

	pdf.SetFillColor(0, 255, 0)
	pdf.SetDrawColor(0, 0, 255)
	pdf.Rect(10, 100, 100, 100, "FD")
	pdf.SetFillColor(100, 200, 200)
	pdf.Polygon(
		[]gofpdf.PointType{
			{110, 250},
			{160, 300},
			{110, 350},
			{60, 300}}, "F")

	pdf.ImageOptions("images/jump.png", 275, 275, 92, 0, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")

	drawGrid(pdf)
	err := pdf.OutputFileAndClose("test.pdf")
	if err != nil {
		panic(err)
	}
}
func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x += (w / 20.0) {

		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y += (h / 20.0) {

		pdf.Line(0, y, w, y)

		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}

}
