package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"Epam_final/pkg/db/models"
	"Epam_final/pkg/front"
)

//ListPaymentsHandler show all payments
func (a *Api) ListPaymentsHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("get id from path: %s", err))
		return
	}
	rows, err := a.paymentRepo.List(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't show payment list: %w", err))
		return
	}
	resp := make([]front.PaymentsRes, len(rows))
	for _, r := range rows {
		resp = append(resp, *makePaymentRes(&r))
	}
	c.JSON(http.StatusOK, resp)
}

//CreatePaymentHandler create payment & update amount
func (a *Api) CreatePaymentHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("get id from path: %w", err))
		return
	}

	req := &front.PaymentReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("bind request: %w", err))
		return
	}
	if err = req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("not validate value %w", err))
		return
	}

	account, err := a.accountRepo.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't find account %w", err))
		return
	}
	if account.AccStatus {
		c.JSON(http.StatusBadRequest, fmt.Errorf("account is lock"))
		return
	}

	if account.Amount <= 0 {
		if account.Amount < req.Amount*-1 {
			c.JSON(http.StatusBadRequest, fmt.Errorf("not enough money"))
			return
		}
	}
	row := &models.Payments{
		AccID:  req.AccID,
		Amount: req.Amount,
		Status: req.Status,
	}
	err = a.paymentRepo.Create(row)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("create payment %w", err))
		return
	}
	account.Amount = account.Amount + req.Amount
	err = a.accountRepo.Update(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't update account: %w", err))
	}
	c.JSON(http.StatusOK, makePaymentRes(row))
}

func makePaymentRes(row *models.Payments) *front.PaymentsRes {
	return &front.PaymentsRes{
		AccID:  row.AccID,
		Amount: row.Amount,
		Status: row.Status,
	}
}
