package controllers

import (
	"fazztrack/demo/lib"
	"fazztrack/demo/models"

	// "fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FormTransaction struct {
	EventId         int `json:"eventId" form:"eventId" db:"event_id"`
	SectionId       []int `json:"sectionId" form:"sectionId[]" db:"section_id"`
	TicketQty       []int `json:"ticketQty" form:"ticketQty[]" db:"ticket_qty"`
	PaymentMethodId int `json:"paymentMethodId" form:"paymentMethodId" db:"payment_method_id"`
}

func CreateTransaction(ctx *gin.Context) {
	var form FormTransaction
	ctx.Bind(&form)
	// fmt.Println(ctx.GetInt("userId"))
	user := ctx.GetInt("userId")

	trx := models.CreateTransaction(models.Transaction{
		UserId: user,
		PaymentMethodId: form.PaymentMethodId,
		EventId: form.EventId,
	})

	for i := range form.SectionId {
		models.CreateTransactionDetail(models.TransactionDetail{
			SectionId: form.SectionId[i],
			TickectQty: form.TicketQty[i],
			TransactionId: trx.Id,
		})
	}

	data := models.DetailTransaction(trx.Id)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Success to create event",
		Results: data,
	})
}

func FindAllTransactionByUserId(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    datatransactions, err := models.FindAllTransactionOnByUserId(id)
    // fmt.Println(datatransactions)

    if err != nil {
        ctx.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "Transaction Not Found",
        })
        return
    }

    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "Transaction Found",
        Results: datatransactions,
    })
}