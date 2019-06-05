package main

import (
	"flag"
	"fmt"
	"image/color"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type PDFOption func(*gofpdf.Fpdf)

func FillColor(c color.RGBA) PDFOption {
	return func(pdf *gofpdf.Fpdf) {
		r, g, b := rgb(c)
		pdf.SetFillColor(r, g, b)
	}
}
func rgb(c color.RGBA) (int, int, int) {
	alpha := float64(c.A) / 255.0
	alphaWhite := int(255 * (1.0 - alpha))
	r := int(float64(c.R)*alpha) + alphaWhite
	g := int(float64(c.G)*alpha) + alphaWhite
	b := int(float64(c.B)*alpha) + alphaWhite
	return r, g, b
}

type PDF struct {
	fpdf *gofpdf.Fpdf
	x, y float64
}

func (p *PDF) Move(xDelta, yDelta float64) {
	p.x, p.y = p.x+xDelta, p.y+yDelta
	p.fpdf.MoveTo(p.x, p.y)
}
func (p *PDF) MoveAbs(x, y float64) {
	p.x, p.y = x, y
	p.fpdf.MoveTo(p.x, p.y)
}

func (p *PDF) Text(text string) {
	p.fpdf.Text(p.x, p.y, text)
}

func (p *PDF) Polygon(pts []gofpdf.PointType, opts ...PDFOption) {
	for _, opt := range opts {
		opt(p.fpdf)
	}
	p.fpdf.Polygon(pts, "F")
}

func main() {

	name := flag.String("name", "", "the name of the person who completed the course")
	flag.Parse()
	fpdf := gofpdf.New(gofpdf.OrientationLandscape, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")

	w, h := fpdf.GetPageSize()

	fmt.Printf("width=%v ,height=%v", w, h)
	fpdf.AddPage()
	pdf := PDF{
		fpdf: fpdf,
		x:    0,
		y:    0,
	}
	primary := color.RGBA{103, 60, 79, 255}
	secondary := color.RGBA{103, 60, 79, 220}

	// Top and Bottom Graphics
	pdf.Polygon([]gofpdf.PointType{
		{0, 0},
		{0, h / 9.0},
		{w - (w / 6), 0},
	}, FillColor(secondary))

	pdf.Polygon([]gofpdf.PointType{
		{0, 0},
		{w, 0},
		{w, h / 9.0},
	}, FillColor(primary))

	pdf.Polygon([]gofpdf.PointType{
		{w, h},
		{w, h - h/8.0},
		{w / 6, h},
	}, FillColor(secondary))

	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - h/8.0},
		{w - (w / 6), h},
	}, FillColor(primary))

	fpdf.SetFont("times", "B", 50)
	fpdf.SetTextColor(50, 50, 50)
	pdf.MoveAbs(0, 100)
	_, lineHt := fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, "Certificate of Completion", gofpdf.AlignCenter)

	pdf.Move(0, lineHt*2.0)

	fpdf.SetFont("arial", "", 28)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, "This certificate is awarded to", gofpdf.AlignCenter)

	pdf.Move(0, lineHt*2.0)
	fpdf.SetFont("times", "B", 42)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, *name, gofpdf.AlignCenter)

	pdf.Move(0, lineHt*1.75)
	fpdf.SetFont("arial", "", 22)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt*1.5, "For successfully completing all twenty programming exercises in the Gophercises programming course for budding Gophers (Go developers)", gofpdf.AlignCenter)

	pdf.Move(0, lineHt*4.5)
	fpdf.ImageOptions("images/dr_flag.png", w/2.0-50.0, pdf.y, 100.0, 0, false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")

	pdf.Move(0, 65.0)

	fpdf.SetFillColor(100, 100, 100)

	fpdf.Rect(60.0, pdf.y, 240.0, 1.0, "F")
	fpdf.Rect(490.0, pdf.y, 240.0, 1.0, "F")
	fpdf.SetFont("arial", "", 12)
	pdf.Move(0, lineHt/1.5)
	fpdf.SetTextColor(100, 100, 100)
	pdf.MoveAbs(60.0+105.0, pdf.y)
	pdf.Text("Date")
	pdf.MoveAbs(490+100.0, pdf.y)
	pdf.Text("Instructor")

	pdf.MoveAbs(60.0, pdf.y-lineHt/1.5)
	fpdf.SetFont("times", "", 22)
	_, lineHt = fpdf.GetFontSize()
	pdf.Move(0, -lineHt)
	fpdf.SetTextColor(50, 50, 50)
	year, month, day := time.Now().Date()
	dateStr := fmt.Sprintf("%d/%d/%d", month, day, year)
	fpdf.CellFormat(240.0, lineHt, dateStr, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignCenter, false, 0, "")
	pdf.MoveAbs(490, pdf.y)
	fpdf.CellFormat(240.0, lineHt, "Jonathan Calhoun", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignCenter, false, 0, "")
	//drawGrid(pdf)
	err := fpdf.OutputFileAndClose("cert.pdf")
	if err != nil {
		panic(err)
	}
}

/*
func trailerLine(pdf *gofpdf.Fpdf, x, y float64, label string, amount int) (float64, float64) {
	origX := x
	w, _ := pdf.GetPageSize()
	pdf.SetFont("times", "", 14)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(180, 180, 180)
	//x, y = w/1.75, y+lineHt*2.25
	pdf.MoveTo(x, y)
	pdf.CellFormat(80.0, lineHt, label, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")

	x = w - xIndent - 2.0 - 119.5
	pdf.SetTextColor(50, 50, 50)
	pdf.MoveTo(x, y)
	pdf.CellFormat(119.5, lineHt, toUSD(amount), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	y = y + lineHt*1.5
	return origX, y
}

func toUSD(cents int) string {
	centsStr := fmt.Sprintf("%d", cents%100)
	if len(centsStr) < 2 {
		centsStr = "0" + centsStr
	}

	return fmt.Sprintf("$%d.%s", cents/100, centsStr)
}
func lineItem(pdf *gofpdf.Fpdf, x, y float64, lineItem LineItem) (float64, float64) {
	origX := x
	w, _ := pdf.GetPageSize()
	pdf.SetFont("times", "", 14)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(50, 50, 50)
	pdf.MoveTo(x, y)

	x, y = xIndent-2.0, y+lineHt*0.75
	pdf.MoveTo(x, y)
	pdf.MultiCell(w/2.65+1.5, lineHt, lineItem.UnitName, gofpdf.BorderNone, gofpdf.AlignLeft, false)
	tmp := pdf.SplitLines([]byte(lineItem.UnitName), w/2.65+1.5)
	maxY := y + float64(len(tmp)-1)*lineHt

	x = x + w/2.65 + 1.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(100, lineHt, toUSD(lineItem.PricePerUnit), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = x + 100.0
	pdf.MoveTo(x, y)
	pdf.CellFormat(80, lineHt, fmt.Sprintf("%d", lineItem.UnitsPurchased), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = w - xIndent - 2.0 - 119.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(119.5, lineHt, toUSD(lineItem.PricePerUnit*lineItem.UnitsPurchased), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")

	if maxY > y {
		y = maxY
	}

	y = y + lineHt*1.75
	pdf.SetDrawColor(180, 180, 180)

	pdf.Line(xIndent-10, y, w-xIndent+10, y)
	return origX, y
}

func summaryBlock(pdf *gofpdf.Fpdf, x, y float64, title string, data ...string) (float64, float64) {
	pdf.SetFont("times", "", 14)
	pdf.SetTextColor(180, 180, 180)
	_, lineHt := pdf.GetFontSize()
	y = y + lineHt
	//x, y = xIndent, bannerHt+lineHt*2.5
	pdf.Text(x, y, title)

	y = y + lineHt*1.5
	pdf.SetTextColor(50, 50, 50)
	for _, str := range data {
		y = y + lineHt*1.25
		pdf.Text(x, y, str)

	}
	return x, y
}
func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x += (w / 20.0) {
		pdf.SetTextColor(200, 200, 200)
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y += (h / 20.0) {
		if y < bannerHt*0.9 {
			pdf.SetTextColor(200, 200, 200)
		} else {
			pdf.SetTextColor(80, 80, 80)
		}

		pdf.Line(0, y, w, y)

		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}

}
*/
