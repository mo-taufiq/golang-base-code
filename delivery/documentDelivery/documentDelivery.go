package documentDelivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"taufiq.code/golang-base-code/helper/document/pdf"
	"taufiq.code/golang-base-code/useCase"
)

type IDocumentDelivery interface {
	GetInvoicePDF(c *gin.Context)
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
