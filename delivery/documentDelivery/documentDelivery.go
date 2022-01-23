package documentDelivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gologger "github.com/mo-taufiq/go-logger"
	"gorm.io/gorm"
	"taufiq.code/golang-base-code/helper/document/excel"
	"taufiq.code/golang-base-code/helper/document/pdf"
	"taufiq.code/golang-base-code/useCase"
)

type IDocumentDelivery interface {
	GetInvoicePDF(c *gin.Context)
	GetInvoiceExcel(c *gin.Context)
}

type DocumentDelivery struct {
	UseCase useCase.UseCase
}

func NewDocumentDelivery(db *gorm.DB) IDocumentDelivery {
	return &DocumentDelivery{
		UseCase: *useCase.NewUseCase(db),
	}
}

func (d *DocumentDelivery) GetInvoicePDF(c *gin.Context) {
	filePDF := pdf.CreatePDFFromTemplate("./assets/templates/template-invoice-1.html")

	c.Header("Content-Disposition", "attachment; filename=invoice.pdf")
	c.Data(http.StatusOK, "application/octet-stream", filePDF.Bytes())
}

func (d *DocumentDelivery) GetInvoiceExcel(c *gin.Context) {
	columnsTitle := []string{"No", "Name", "Description", "Unit Price", "Quantity", "Total"}
	data := [][]interface{}{
		{
			"1",
			"Website Design",
			"Creating a recognizable design solution based on the company's existing visual identity",
			"$40.00",
			"30",
			"$1,200.00",
		},
		{
			"2",
			"Website Development",
			"Developing a Content Management System-based Website",
			"$40.00",
			"80",
			"$3,200.00",
		},
		{
			"3",
			"Search Engines Optimization",
			"Optimize the site for search engines (SEO)",
			"$40.00",
			"20",
			"$800.00",
		},
	}
	fileExcel := excel.CreateFileExcel(columnsTitle, data)
	fileExcelByte, err := fileExcel.WriteToBuffer()
	if err != nil {
		gologger.Error(err.Error())
	}

	c.Header("Content-Disposition", "attachment; filename=invoice.xlsx")
	c.Data(http.StatusOK, "application/octet-stream", fileExcelByte.Bytes())
}
