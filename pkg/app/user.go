package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"Epam_final/pkg/db/models"
	"Epam_final/pkg/front"
)

// GetUserHandler gets user
func (a *Api) GetUserHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("get id from path: %s", err))
		return
	}

	row, err := a.userRepo.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("get user: %s", err))
		return
	}
	c.JSON(http.StatusOK, makeUserRes(row))
}

func makeUserRes(row *models.User) *front.UserRes {
	res := &front.UserRes{
		Name:    row.Name,
		Surname: row.Surname,
		Email:   row.Email,
		IsAdmin: row.IsAdmin,
	}
	res.Accounts = make([]front.AccountRes, 0, len(row.Accounts))
	for _, acc := range row.Accounts {
		res.Accounts = append(res.Accounts, front.AccountRes{
			ID:        acc.ID,
			Acc:       acc.Acc,
			UserID:    acc.UserID,
			AccStatus: acc.AccStatus,
			Amount:    acc.Amount,
		})

	}
	res.Logs = make([]front.LogRes, 0, len(row.Logs))
	for _, log := range row.Logs {
		res.Logs = append(res.Logs, front.LogRes{
			ID:      log.ID,
			UserID:  log.UserID,
			Massage: log.Massage,
			Date:    log.Date,
		})
	}

	return res
}
