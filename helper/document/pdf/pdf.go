package pdf

import (
	"fmt"
	"os"
	"runtime"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	gologger "github.com/mo-taufiq/go-logger"
)

func CreatePDFFromTemplateToFile(fileHTMLTemplatePath, outputFilePDFPath string) {
	gologger.Info("Start create PDF file")
	gologger.Info(fmt.Sprintf("path input tempate file HTML: %s", fileHTMLTemplatePath))
	gologger.Info(fmt.Sprintf("path output file PDF: %s", outputFilePDFPath))
	setBinWkhtmltopdfPath()

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	pdfg.Cover.EnableLocalFileAccess.Set(true)
	if err != nil {
		gologger.Error(err.Error())
	}

	f, err := os.Open(fileHTMLTemplatePath)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		gologger.Error(err.Error())
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		gologger.Error(err.Error())
	}

	err = pdfg.WriteFile(outputFilePDFPath)
	if err != nil {
		gologger.Error(err.Error())
	}

	gologger.Info("Finish create PDF file")
}

func CreatePDFFromTemplateToBytes(fileHTMLTemplatePath string) []byte {
	gologger.Info("Start create PDF file")
	gologger.Info(fmt.Sprintf("path input tempate file HTML: %s", fileHTMLTemplatePath))
	setBinWkhtmltopdfPath()

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	pdfg.Cover.EnableLocalFileAccess.Set(true)
	if err != nil {
		gologger.Error(err.Error())
	}

	f, err := os.Open(fileHTMLTemplatePath)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		gologger.Error(err.Error())
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		gologger.Error(err.Error())
	}

	gologger.Info("Finish create PDF file")
	return pdfg.Bytes()
}

func CreatePDFFromTemplate(fileHTMLTemplatePath string) *wkhtmltopdf.PDFGenerator {
	gologger.Info("Start create PDF file")
	gologger.Info(fmt.Sprintf("path input tempate file HTML: %s", fileHTMLTemplatePath))
	setBinWkhtmltopdfPath()

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	pdfg.Cover.EnableLocalFileAccess.Set(true)
	if err != nil {
		gologger.Error(err.Error())
	}

	f, err := os.Open(fileHTMLTemplatePath)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		gologger.Error(err.Error())
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		gologger.Error(err.Error())
	}

	gologger.Info("Finish create PDF file")
	return pdfg
}

func setBinWkhtmltopdfPath() {
	operatingSystem := runtime.GOOS
	switch operatingSystem {
	case "windows":
		gologger.Info("Windows operating system")
	case "darwin":
		gologger.Info("MAC operating system")
		wkhtmltopdf.SetPath("./bin/wkhtmltopdf/darwin-wkhtmltopdf")
	case "linux":
		gologger.Info("Linux operating system")
	default:
		fmt.Printf("%s.\n", operatingSystem)
	}
}
