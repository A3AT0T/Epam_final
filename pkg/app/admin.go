package app

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"Epam_final/pkg/db/models"
	"github.com/gin-gonic/gin"
)

// BlockUserHandler for admin to block user
func (a *Api) BlockUserHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("get id from path: %w", err))
		return
	}

	row := &models.User{
		ID:     id,
		Status: "blocked",
	}
	if err := a.userRepo.Update(row); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't find user: %w", err))
		return
	}
	row, err = a.userRepo.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't get user: %w", err))
		return
	}
	c.JSON(http.StatusOK, makeUserRes(row))

	log := &models.Log{
		UserID:  row.ID,
		Massage: "user has been blocked",
		Date:    time.Time{},
	}
	err = a.logRepo.Create(log)
	if err != nil {
		fmt.Errorf("create log: %w", err)
	}

}

// UnblockUserHandler for admin to unblock user
func (a *Api) UnblockUserHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("get id from path: %w", err))
		return
	}

	row := &models.User{
		ID:     id,
		Status: "active",
	}
	if err := a.userRepo.Update(row); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't update user: %w", err))
		return
	}
	row, err = a.userRepo.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't get user: %w", err))
		return
	}
	c.JSON(http.StatusOK, makeUserRes(row))

	log := &models.Log{
		UserID:  row.ID,
		Massage: "user has been unblocked",
		Date:    time.Time{},
	}
	err = a.logRepo.Create(log)
	if err != nil {
		fmt.Errorf("create log: %w", err)
	}
}

// UserRequestApproveHandler approves user request
func (a *Api) UserRequestApproveHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("get id from path: %w", err))
		return
	}

	r, err := a.userRequestRepo.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("get user request: %w", err))
		return
	}

	if r.Status == true {
		c.JSON(http.StatusBadRequest, fmt.Errorf("request already approved"))
		return
	}
	row := &models.UserRequest{
		ID:     id,
		Status: true,
	}
	if err = a.userRequestRepo.Update(row); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't update user: %w", err))
		return
	}
	row, err = a.userRequestRepo.Get(id)
	if err != nil {

		c.JSON(http.StatusInternalServerError, fmt.Errorf("can't get user: %w", err))
		return
	}

	acc := &models.Account{
		ID:        r.AccID,
		AccStatus: false,
	}

	if err = a.accountRepo.Update(acc); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("account update failed: %w", err))
		return
	}

	c.JSON(http.StatusOK, makeUserRequestRes(row))

	logs := &models.Log{
		UserID:  row.Account.UserID,
		Massage: "user has created request",
		Date:    time.Time{},
	}

	err = a.logRepo.Create(logs)
	if err != nil {
		fmt.Errorf("create log: %w", err)
	}
}
