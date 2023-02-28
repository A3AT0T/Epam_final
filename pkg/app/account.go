package app

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"Epam_final/pkg/db/models"
	"Epam_final/pkg/front"
)

// CreateAccountHandler create new account
func (a *Api) CreateAccountHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("get id from path: %w", err))
		return
	}

	row := &models.Account{
		Acc:    genAccNumber(),
		UserID: id,
	}
	if err := a.accountRepo.Create(row); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't create account: %w", err))
		return
	}
	c.JSON(http.StatusOK, makeAccountRes(row))

	log := &models.Log{
		UserID:  row.UserID,
		Massage: "create account",
		Date:    time.Time{},
	}
	err = a.logRepo.Create(log)
	if err != nil {
		fmt.Errorf("create log: %w", err)
	}
}

// ListAccountHandler  show all accounts
func (a *Api) ListAccountHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("can't get list: %w", err))
		return
	}

	rows, err := a.accountRepo.List(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't show account list: %w", err))
		return
	}
	resp := make([]front.AccountRes, 0, len(rows))
	for _, r := range rows {
		resp = append(resp, *makeAccountRes(&r))
	}
	c.JSON(http.StatusOK, resp)
}

func makeAccountRes(row *models.Account) *front.AccountRes {
	res := &front.AccountRes{
		ID:        row.ID,
		Acc:       row.Acc,
		UserID:    row.UserID,
		AccStatus: row.AccStatus,
		Amount:    row.Amount,
	}
	res.Cards = make([]front.CardRes, 0, len(row.Cards))
	for _, card := range row.Cards {
		res.Cards = append(res.Cards, front.CardRes{
			ID:     card.ID,
			CardID: card.CardID,
			AccID:  card.AccID,
		})
	}
	res.UserRequest = make([]front.UserRequestRes, 0, len(row.UserRequest))
	for _, userRequest := range row.UserRequest {
		res.UserRequest = append(res.UserRequest, front.UserRequestRes{
			ID:     userRequest.ID,
			AccID:  userRequest.AccID,
			Status: userRequest.Status,
		})
	}
	res.Payments = make([]front.PaymentsRes, 0, len(row.Payments))
	for _, pay := range row.Payments {
		res.Payments = append(res.Payments, front.PaymentsRes{
			ID:     pay.ID,
			AccID:  pay.AccID,
			Amount: pay.Amount,
			Status: pay.Status,
			Date:   pay.Date,
		})
	}

	return res
}
