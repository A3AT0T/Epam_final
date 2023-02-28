package app

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"Epam_final/pkg/db/models"
	"Epam_final/pkg/front"
)

// CreateUserHandler Creates user
func (a *Api) CreateUserHandler(c *gin.Context) {

	req := &front.UserReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("bind request: %w", err))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("not validate value %w", err))
		return
	}
	passhash := fmt.Sprintf("%x", sha256.Sum256([]byte(req.Pass)))

	row := &models.User{
		Name:    req.Name,
		Surname: req.Surname,
		Email:   req.Email,
		Pass:    passhash,
	}
	err := a.userRepo.Create(row)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("create user %w", err))
	}
	c.JSON(http.StatusOK, makeUserRes(row))

	log := &models.Log{
		UserID:  row.ID,
		Massage: "user has been created",
		Date:    time.Time{},
	}
	err = a.logRepo.Create(log)
	if err != nil {
		fmt.Errorf("create log: %w", err)
	}
}

// GetUserHandler gets user
func (a *Api) GetUserHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("get id from path: %s", err))
		return
	}
	log.Printf("get id from path %v", err)
	row, err := a.userRepo.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("get user: %s", err))
		return
	}
	c.JSON(http.StatusOK, makeUserRes(row))
}

func (a *Api) DeleteUserHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("delete user: %s", err))
		return
	}
	err = a.userRepo.Delete(id)
	if err != nil {
		c.JSON(http.StatusOK, err)
	}
}

func makeUserRes(row *models.User) *front.UserRes {
	res := &front.UserRes{
		ID:      row.ID,
		Name:    row.Name,
		Surname: row.Surname,
		Email:   row.Email,
		IsAdmin: row.IsAdmin,
		Status:  row.Status,
	}
	res.Accounts = make([]front.AccountRes, 0, len(row.Accounts))
	for _, acc := range row.Accounts {
		res.Accounts = append(res.Accounts, *makeAccountRes(&acc))

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
