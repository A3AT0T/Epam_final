package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"Epam_final/pkg/db/models"
	"Epam_final/pkg/front"
)

// CreateUserRequestHandler create user request
func (a *Api) CreateUserRequestHandler(c *gin.Context) {
	req := &front.UserRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("bind request: %w", err))
		return
	}
	row := &models.UserRequest{
		AccID: req.AccID,
	}
	err := a.userRequestRepo.Create(row)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("create user request %w", err))
	}
	c.JSON(http.StatusOK, makeUserRequestRes(row))
}

func makeUserRequestRes(row *models.UserRequest) *front.UserRequestRes {
	return &front.UserRequestRes{
		ID:     row.ID,
		AccID:  row.AccID,
		Status: row.Status,
		Date:   row.Date,
	}
}
