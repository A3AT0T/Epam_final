package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"Epam_final/pkg/db/models"
	"Epam_final/pkg/front"
)

//CreateCardHandler create new card
func (a *Api) CreateCardHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("get id from path: %w", err))
		return
	}
	row := &models.Card{
		CardID: genCardsNumber(),
		AccID:  id,
	}
	if err := a.cardRepo.Create(row); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't create card: %w", err))
		return
	}
	c.JSON(http.StatusOK, makeCardRes(row))
}

//GetCardHandler get info about card
func (a *Api) GetCardHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("get id from path: %w", err))
		return
	}
	row, err := a.cardRepo.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't get card info: %w", err))
		return
	}
	c.JSON(http.StatusOK, makeCardRes(row))
}

func makeCardRes(row *models.Card) *front.CardRes {
	return &front.CardRes{
		ID:     row.ID,
		CardID: row.CardID,
		AccID:  row.AccID,
	}
}
